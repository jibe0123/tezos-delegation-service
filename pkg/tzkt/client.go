package tezos

import (
	"encoding/json"
	"net/http"
	"technical-test/internal/delegation/domain"
	"time"
)

// Client defines the interface for fetching delegations from the Tezos API.
type Client interface {
	GetDelegations() ([]domain.Delegation, error)
	GetDelegationsSince(lastTimestamp time.Time) ([]domain.Delegation, error)
}

type client struct {
	baseURL string
}

// NewClient creates a new Tezos Client instance.
func NewClient(baseURL string) Client {
	return &client{baseURL: baseURL}
}

// GetDelegations fetches delegations from the Tezos API.
func (c *client) GetDelegations() ([]domain.Delegation, error) {
	resp, err := http.Get(c.baseURL + "operations/delegations?limit=10000")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var delegations []domain.Delegation
	if err := json.NewDecoder(resp.Body).Decode(&delegations); err != nil {
		return nil, err
	}

	return delegations, nil
}

// GetDelegationsSince fetches delegations from the Tezos API that are newer than the given timestamp.
func (c *client) GetDelegationsSince(lastTimestamp time.Time) ([]domain.Delegation, error) {
	url := c.baseURL + "operations/delegations?limit=10000&timestamp.gt=" + lastTimestamp.Format(time.RFC3339)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var delegations []domain.Delegation
	if err := json.NewDecoder(resp.Body).Decode(&delegations); err != nil {
		return nil, err
	}

	return delegations, nil
}
