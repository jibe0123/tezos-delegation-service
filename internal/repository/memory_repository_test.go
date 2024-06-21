package repository_test

import (
	"technical-test/internal/domain"
	"technical-test/internal/repository"
	"testing"
	"time"
)

func TestMemoryRepository_Save(t *testing.T) {
	repo := repository.NewMemoryRepository()

	delegation := domain.Delegation{
		Delegator: "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL",
		Timestamp: time.Now(),
		Amount:    1000,
		Level:     12345,
	}

	err := repo.Save(delegation)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	delegations, err := repo.FindAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(delegations) != 1 {
		t.Fatalf("expected 1 delegation, got %d", len(delegations))
	}
}

func TestMemoryRepository_FindByYear(t *testing.T) {
	repo := repository.NewMemoryRepository()

	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)

	repo.Save(domain.Delegation{
		Delegator: "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL",
		Timestamp: now,
		Amount:    1000,
		Level:     12345,
	})
	repo.Save(domain.Delegation{
		Delegator: "KT1JejNYjmQYh8yw95u5kfQDRuxJcaUPjUnf",
		Timestamp: lastYear,
		Amount:    2000,
		Level:     67890,
	})

	delegations, err := repo.FindByYear(now.Year())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(delegations) != 1 {
		t.Fatalf("expected 1 delegation, got %d", len(delegations))
	}

	if delegations[0].Timestamp.Year() != now.Year() {
		t.Fatalf("expected year %d, got %d", now.Year(), delegations[0].Timestamp.Year())
	}
}
