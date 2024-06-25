package repository

import (
	"database/sql"
	"technical-test/internal/delegation/domain"
	"time"
)

// Repository defines the methods that any data storage provider needs to implement to get and store delegations.
type Repository interface {
	Save(delegation domain.Delegation) error
	FindAll(year string) ([]domain.Delegation, error)
	GetLastTimestamp() (time.Time, error)
}

type repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Save stores a delegation in the database.
func (r *repository) Save(delegation domain.Delegation) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO delegations (tx_id, timestamp, amount, delegator, level, block) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(delegation.TxId, delegation.Timestamp, delegation.Amount, delegation.Sender.Address, delegation.Level, delegation.Block)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// FindAll retrieves all delegations from the database, optionally filtering by year.
func (r *repository) FindAll(year string) ([]domain.Delegation, error) {
	query := "SELECT timestamp, amount, delegator, level, block FROM delegations"
	args := []interface{}{}
	if year != "" {
		query += " WHERE YEAR(timestamp) = ? ORDER BY timestamp DESC"
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
		if err := rows.Scan(&d.Timestamp, &d.Amount, &d.Sender.Address, &d.Level, &d.Block); err != nil {
			return nil, err
		}
		delegations = append(delegations, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return delegations, nil
}

// GetLastTimestamp retrieves the most recent timestamp from the delegations table.
func (r *repository) GetLastTimestamp() (time.Time, error) {
	var lastTimestamp time.Time
	query := "SELECT MAX(timestamp) FROM delegations"
	err := r.db.QueryRow(query).Scan(&lastTimestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return time.Time{}, nil // Return zero time if no records found
		}
		return time.Time{}, err
	}
	return lastTimestamp, nil
}
