package services

import (
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.Equal(t, 1, len(delegations))
	assert.Equal(t, "2022-05-05T06:29:14Z", delegations[0].Timestamp)
	assert.Equal(t, 125896, delegations[0].Amount)
	assert.Equal(t, "tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL", delegations[0].Sender.Address)
	assert.Equal(t, 2338084, delegations[0].Level)
}
