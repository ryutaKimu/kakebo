package dbutil

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() (*sql.DB, error) {
	dbURL := os.Getenv("GOOSE_DBSTRING")
	if dbURL == "" {
		return nil, fmt.Errorf("GOOSE_DBSTRING is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect db: %w", err)
	}

	return db, nil
}
