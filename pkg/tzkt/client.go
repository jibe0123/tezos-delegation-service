package tezos

import (
	"encoding/json"
	"net/http"
	"technical-test/internal/delegation/domain"
)

// Client defines the interface for fetching delegations from the Tezos API.
type Client interface {
	GetDelegations() ([]domain.Delegation, error)
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
	resp, err := http.Get(c.baseURL + "operations/delegations")
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
