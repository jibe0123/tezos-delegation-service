package database

import (
	"database/sql"
	"fmt"
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
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
		tx_id integer NOT NULL,
		timestamp DATETIME NOT NULL,
		amount TEXT,
		delegator TEXT,
		block TEXT,
		level integer NOT NULL,
		UNIQUE(tx_id, level, timestamp)
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database table created successfully")
	return db, nil
}
