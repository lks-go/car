package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect establishes connection to PostgresSQL
// and ping it if connection is not established
// or ping fails then Connect returns error
func Connect(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprint(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
