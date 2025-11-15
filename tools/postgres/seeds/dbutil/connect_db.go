package dbutil

import (
	"database/sql"
	"log"
	"os"
)

func Connect() (*sql.DB, error) {
	dbURL := os.Getenv("GOOSE_DBSTRING")
	if dbURL == "" {
		log.Fatal("GOOSE_DBSTRING is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	return db, nil
}
