package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Database defines the interface for database operations.
type Database interface {
	Begin() (*sql.Tx, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

// InitDB initializes the database and creates the necessary tables.
func InitDB(filepath string) (Database, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	createTableQuery := `CREATE TABLE IF NOT EXISTS delegations (
		timestamp TEXT,
		amount TEXT,
		delegator TEXT,
		level TEXT,
		UNIQUE(timestamp, delegator, level)
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	log.Println("Database table created successfully")
	return db, nil
}
