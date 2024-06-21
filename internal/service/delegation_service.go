package service

import (
	"technical-test/internal/domain"
	"technical-test/internal/repository"
	"time"
)

// DelegationService provides methods to manage delegations
type DelegationService struct {
	repo repository.DelegationRepository
}

// NewDelegationService creates a new instance of DelegationService
func NewDelegationService(repo repository.DelegationRepository) *DelegationService {
	return &DelegationService{repo: repo}
}

// SaveDelegation saves a new delegation
func (s *DelegationService) SaveDelegation(delegator string, timestamp time.Time, amount, level int64) error {
	delegation := domain.Delegation{
		Delegator: delegator,
		Timestamp: timestamp,
		Amount:    amount,
		Level:     level,
	}
	return s.repo.Save(delegation)
}

// GetDelegations retrieves delegations, optionally filtered by year
func (s *DelegationService) GetDelegations(year int) ([]domain.Delegation, error) {
	if year == 0 {
		return s.repo.FindAll()
	}
	return s.repo.FindByYear(year)
}
