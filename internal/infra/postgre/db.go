package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

type txKey struct{}

var TxContextKey txKey = txKey{}

type Postgres struct {
	DB *sql.DB
}

func NewPostgres() *Postgres {

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DATABASE")

	// 開発環境では SSL を無効にする
	sslmode := os.Getenv("POSTGRES_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	encodedUser := url.QueryEscape(user)
	encodedPass := url.QueryEscape(password)

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		encodedUser, encodedPass, host, port, dbname, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return &Postgres{DB: db}
}

func (p *Postgres) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if ctx.Value(TxContextKey) != nil {
		return fn(ctx)
	}

	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, TxContextKey, tx)

	if err := fn(ctx); err != nil {
		if e := tx.Rollback(); e != nil {
			log.Printf("transaction rollback failed: %v", e)
		}
		return err
	}

	return tx.Commit()
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}
