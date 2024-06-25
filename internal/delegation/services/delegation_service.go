package services

import (
	"technical-test/internal/delegation/domain"
	"technical-test/internal/delegation/repository"
	tzkt "technical-test/pkg/tzkt"
	"time"
)

// Service defines the delegation service.
type Service interface {
	FetchDelegations() ([]domain.Delegation, error)
	FetchDelegationsSince(lastTimestamp time.Time) ([]domain.Delegation, error)
}

type service struct {
	client tzkt.Client
	repo   repository.Repository
}

// NewService creates a new Service instance.
func NewService(client tzkt.Client, repo repository.Repository) Service {
	return &service{client: client, repo: repo}
}

// FetchDelegations fetches all delegations from the Tezos API.
func (s *service) FetchDelegations() ([]domain.Delegation, error) {
	return s.client.GetDelegations()
}

// FetchDelegationsSince fetches delegations from the Tezos API that are newer than the given timestamp.
func (s *service) FetchDelegationsSince(lastTimestamp time.Time) ([]domain.Delegation, error) {
	return s.client.GetDelegationsSince(lastTimestamp)
}
