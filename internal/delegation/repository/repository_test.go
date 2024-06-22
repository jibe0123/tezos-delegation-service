package repository

import (
	"technical-test/internal/delegation/domain"
	database "technical-test/pkg/sqlite"
	"testing"
)

func TestSaveAndFindAll(t *testing.T) {
	db, err := database.InitDB(":memory:")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	delegation := domain.Delegation{
		Timestamp: "2022-05-05T06:29:14Z",
		Amount:    125896,
		Sender: struct {
			Address string `json:"address"`
		}{
			Address: "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL",
		},
		Level: 2338084,
	}

	err = repo.Save(delegation)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	delegations, err := repo.FindAll("")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(delegations) != 1 {
		t.Fatalf("expected 1 delegation, got %d", len(delegations))
	}

	if delegations[0].Timestamp != delegation.Timestamp {
		t.Errorf("expected timestamp %s, got %s", delegation.Timestamp, delegations[0].Timestamp)
	}

	if delegations[0].Amount != delegation.Amount {
		t.Errorf("expected amount %d, got %d", delegation.Amount, delegations[0].Amount)
	}

	if delegations[0].Sender.Address != delegation.Sender.Address {
		t.Errorf("expected sender address %s, got %s", delegation.Sender.Address, delegations[0].Sender.Address)
	}

	if delegations[0].Level != delegation.Level {
		t.Errorf("expected level %d, got %d", delegation.Level, delegations[0].Level)
	}
}
