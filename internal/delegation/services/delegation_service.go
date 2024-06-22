package services

import (
	"technical-test/internal/delegation/domain"
	"technical-test/internal/delegation/repository"
	tezos "technical-test/pkg/tzkt"
)

// Service defines the methods that a delegation service should implement.
type Service interface {
	FetchDelegations() ([]domain.Delegation, error)
}

type service struct {
	client tezos.Client
	repo   repository.Repository
}

// NewService creates a new Service instance.
func NewService(client tezos.Client, repo repository.Repository) Service {
	return &service{client: client, repo: repo}
}

// FetchDelegations retrieves delegations from the Tezos API.
func (s *service) FetchDelegations() ([]domain.Delegation, error) {
	return s.client.GetDelegations()
}
