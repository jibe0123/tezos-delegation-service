package tzkt

import (
	"encoding/json"
	"net/http"
	"time"
)

// Delegation represents a delegation operation from the Tzkt API
type Delegation struct {
	Sender struct {
		Address string `json:"address"`
	} `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
	Amount    int64     `json:"amount"`
	Level     int64     `json:"level"`
}

// Client provides methods to interact with the Tzkt API
type Client struct {
	baseURL string
}

// NewClient creates a new Tzkt API client
func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

// GetDelegations fetches delegations from the Tzkt API
func (c *Client) GetDelegations() ([]Delegation, error) {
	resp, err := http.Get(c.baseURL + "/v1/operations/delegations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var delegations []Delegation
	err = json.NewDecoder(resp.Body).Decode(&delegations)
	if err != nil {
		return nil, err
	}

	return delegations, nil
}
