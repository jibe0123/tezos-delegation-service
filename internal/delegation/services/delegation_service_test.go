package services

import (
	"technical-test/internal/delegation/domain"
	"technical-test/internal/delegation/repository"
	"testing"
)

// MockTezosClient is a mock implementation of the Tezos Client interface.
type MockTezosClient struct{}

// GetDelegations returns mock delegations for testing purposes.
func (m *MockTezosClient) GetDelegations() ([]domain.Delegation, error) {
	return []domain.Delegation{
		{
			Timestamp: "2022-05-05T06:29:14Z",
			Amount:    125896,
			Sender: struct {
				Address string `json:"address"`
			}{
				Address: "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL",
			},
			Level: 2338084,
		},
	}, nil
}

func TestFetchDelegations(t *testing.T) {
	mockClient := &MockTezosClient{}
	mockRepo := repository.NewRepository(nil)
	svc := NewService(mockClient, mockRepo)
	delegations, err := svc.FetchDelegations()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(delegations) != 1 {
		t.Fatalf("expected 1 delegation, got %d", len(delegations))
	}
	if delegations[0].Timestamp != "2022-05-05T06:29:14Z" {
		t.Errorf("expected timestamp '2022-05-05T06:29:14Z', got '%s'", delegations[0].Timestamp)
	}
	if delegations[0].Amount != 125896 {
		t.Errorf("expected amount 125896, got %d", delegations[0].Amount)
	}
	if delegations[0].Sender.Address != "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL" {
		t.Errorf("expected sender address 'tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL', got '%s'", delegations[0].Sender.Address)
	}
	if delegations[0].Level != 2338084 {
		t.Errorf("expected level 2338084, got %d", delegations[0].Level)
	}
}
