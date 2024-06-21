package service_test

import (
	"technical-test/internal/repository"
	"technical-test/internal/service"
	"testing"
	"time"
)

func TestSaveDelegation(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewDelegationService(repo)

	err := svc.SaveDelegation("tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL", time.Now(), 1000, 12345)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	delegations, err := svc.GetDelegations(0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(delegations) != 1 {
		t.Fatalf("expected 1 delegation, got %d", len(delegations))
	}
}

func TestGetDelegationsByYear(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewDelegationService(repo)

	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)

	svc.SaveDelegation("tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL", now, 1000, 12345)
	svc.SaveDelegation("KT1JejNYjmQYh8yw95u5kfQDRuxJcaUPjUnf", lastYear, 2000, 67890)

	delegations, err := svc.GetDelegations(now.Year())
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
