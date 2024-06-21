package repository

import "technical-test/internal/domain"

// DelegationRepository provides an interface for storing and retrieving delegations
type DelegationRepository interface {
	Save(delegation domain.Delegation) error
	FindAll() ([]domain.Delegation, error)
	FindByYear(year int) ([]domain.Delegation, error)
}
