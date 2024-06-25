package app

import (
	"context"
	"log"
	"sync"
	"technical-test/internal/delegation/domain"
	"technical-test/internal/delegation/repository"
	"technical-test/internal/delegation/services"
	database "technical-test/pkg/storage"
	tezos "technical-test/pkg/tzkt"
	"time"
)

// App represents the application with its dependencies.
type App struct {
	Repo repository.Repository
	svc  services.Service
}

// NewApp creates a new App instance.
func NewApp(db database.Database, tezosClient tezos.Client) *App {
	repo := repository.NewRepository(db.DB())
	svc := services.NewService(tezosClient, repo)
	return &App{
		Repo: repo,
		svc:  svc,
	}
}

// StartPolling starts the polling process for fetching delegations.
func (a *App) StartPolling(ctx context.Context) {
	lastTimestamp, err := a.Repo.GetLastTimestamp()
	if err != nil {
		log.Fatal("Error fetching last timestamp:", err)
	}

	if lastTimestamp.IsZero() {
		log.Println("Database is empty. Fetching all existing delegations...")
		a.fetchAllDelegations()
	} else {
		log.Println("Database contains delegations. Starting regular polling...")
	}

	// Create a ticker to fetch every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Starting to fetch new delegations...")
			a.fetchNewDelegations()
		case <-ctx.Done():
			log.Println("Stopping polling due to context cancellation")
			return
		}
	}
}

// fetchAllDelegations fetches all existing delegations from the Tezos API and saves them to the repository.
func (a *App) fetchAllDelegations() {
	delegations, err := a.svc.FetchDelegations()
	if err != nil {
		log.Println("Error fetching delegations:", err)
		return
	}

	a.saveDelegations(delegations)
}

// fetchNewDelegations fetches new delegations from the Tezos API that are newer than the last stored timestamp.
func (a *App) fetchNewDelegations() {
	// Check if the database is empty by fetching the last timestamp
	lastTimestamp, err := a.Repo.GetLastTimestamp()
	if err != nil {
		log.Println("Error fetching last timestamp:", err)
		return
	}

	var delegations []domain.Delegation
	if lastTimestamp.IsZero() {
		// If the database is empty, fetch all delegations
		delegations, err = a.svc.FetchDelegations()
	} else {
		// Fetch delegations newer than the last timestamp
		delegations, err = a.svc.FetchDelegationsSince(lastTimestamp)
	}

	if err != nil {
		log.Println("Error fetching delegations:", err)
		return
	}

	a.saveDelegations(delegations)
}

// saveDelegations saves delegations to the repository.
func (a *App) saveDelegations(delegations []domain.Delegation) {
	if len(delegations) == 0 {
		return
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 100) // Limit the number of concurrent goroutines semaphore pattern

	for _, d := range delegations {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(delegation domain.Delegation) {
			defer wg.Done()
			defer func() { <-semaphore }()

			if err := a.Repo.Save(delegation); err != nil {
				log.Println("Error saving delegation:", err)
			} else {
				log.Printf("Delegation saved: %+v\n", delegation.TxId)
			}
		}(d)
	}
	wg.Wait()
	log.Println("Delegations fetched and saved successfully")
}
