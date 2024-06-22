package repository

import (
	"technical-test/internal/delegation/domain"
	database "technical-test/pkg/sqlite"
)

// Repository defines the methods that any data storage provider needs to implement to get and store delegations.
type Repository interface {
	Save(delegation domain.Delegation) error
	FindAll(year string) ([]domain.Delegation, error)
}

type repository struct {
	db database.Database
}

// NewRepository creates a new Repository instance.
func NewRepository(db database.Database) Repository {
	return &repository{db: db}
}

// Save stores a delegation in the database.
func (r *repository) Save(delegation domain.Delegation) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT OR IGNORE INTO delegations (timestamp, amount, delegator, level) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(delegation.Timestamp, delegation.Amount, delegation.Sender.Address, delegation.Level)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// FindAll retrieves all delegations from the database, optionally filtering by year.
func (r *repository) FindAll(year string) ([]domain.Delegation, error) {
	query := "SELECT timestamp, amount, delegator, level FROM delegations"
	args := []interface{}{}
	if year != "" {
		query += " WHERE strftime('%Y', timestamp) = ? ORDER BY timestamp DESC"
		args = append(args, year)
	} else {
		query += " ORDER BY timestamp DESC"
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var delegations []domain.Delegation
	for rows.Next() {
		var d domain.Delegation
		if err := rows.Scan(&d.Timestamp, &d.Amount, &d.Sender.Address, &d.Level); err != nil {
			return nil, err
		}
		delegations = append(delegations, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return delegations, nil
}
