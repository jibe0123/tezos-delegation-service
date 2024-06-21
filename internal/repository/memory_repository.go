package repository

import (
	"sync"
	"technical-test/internal/domain"
)

// MemoryRepository stores delegations in memory with thread-safety
type MemoryRepository struct {
	delegations []domain.Delegation
	mutex       sync.Mutex
}

// NewMemoryRepository creates a new instance of MemoryRepository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		delegations: []domain.Delegation{},
	}
}

// Save adds a new delegation to the in-memory store
func (r *MemoryRepository) Save(delegation domain.Delegation) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.delegations = append(r.delegations, delegation)
	return nil
}

// FindAll retrieves all delegations
func (r *MemoryRepository) FindAll() ([]domain.Delegation, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.delegations, nil
}

// FindByYear retrieves delegations for a specific year
func (r *MemoryRepository) FindByYear(year int) ([]domain.Delegation, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var result []domain.Delegation
	for _, delegation := range r.delegations {
		if delegation.Timestamp.Year() == year {
			result = append(result, delegation)
		}
	}
	return result, nil
}
