package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Database defines the interface for database operations.
type Database interface {
	Begin() (*sql.Tx, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
	DB() *sql.DB
}

// mariadbImpl implements the Database interface for MariaDB.
type mariadbImpl struct {
	db *sql.DB
}

// NewMariaDB creates a new instance of the MariaDB database.
func NewMariaDB(dataSourceName string) (Database, error) {
	var db *sql.DB
	var err error

	// Try to connect to the database with retries
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err == nil {
			break
		}

		fmt.Println("Waiting for the database to be ready...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}

	if err := initializeDatabase(db); err != nil {
		return nil, err
	}

	return &mariadbImpl{db}, nil
}

// initializeDatabase initializes the database and creates the necessary tables.
func initializeDatabase(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS delegations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		tx_id VARCHAR(255) NOT NULL,
		timestamp DATETIME NOT NULL,
		amount BIGINT,
		delegator VARCHAR(255),
		block VARCHAR(255),
		level BIGINT NOT NULL,
		UNIQUE(tx_id, level, timestamp)
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	fmt.Println("Database table created successfully")
	return nil
}

func (m *mariadbImpl) Begin() (*sql.Tx, error) {
	return m.db.Begin()
}

func (m *mariadbImpl) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.db.Query(query, args...)
}

func (m *mariadbImpl) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.db.Exec(query, args...)
}

func (m *mariadbImpl) Close() error {
	return m.db.Close()
}

func (m *mariadbImpl) DB() *sql.DB {
	return m.db
}
