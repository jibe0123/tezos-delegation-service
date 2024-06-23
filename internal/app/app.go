package app

import (
	"context"
	"log"
	"sync"
	"technical-test/internal/delegation/domain"
	"technical-test/internal/delegation/repository"
	"technical-test/internal/delegation/services"
	database "technical-test/pkg/sqlite"
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
	repo := repository.NewRepository(db)
	svc := services.NewService(tezosClient, repo)
	return &App{
		Repo: repo,
		svc:  svc,
	}
}

// StartPolling starts the polling process for fetching delegations.
func (a *App) StartPolling(ctx context.Context) {
	log.Println("Starting to fetch delegations...")
	a.fetchDelegations()

	// Create a ticker to fetch every hour
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Starting to fetch delegations...")
			a.fetchDelegations()
		case <-ctx.Done():
			log.Println("Stopping polling due to context cancellation")
			return
		}
	}
}

// fetchDelegations fetches new delegations from the Tezos API and saves them to the repository.
func (a *App) fetchDelegations() {
	delegations, err := a.svc.FetchDelegations()
	if err != nil {
		log.Println("Error fetching delegations:", err)
		return
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 2) // Limit the number of concurrent goroutines semaphore pattern

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
