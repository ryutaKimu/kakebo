package dbutil

import (
	"context"
	"database/sql"

	postgres "github.com/ryutaKimu/kakebo/api/internal/infra/postgre"
)

type dbExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func GetDBExecutor(ctx context.Context, db *sql.DB) dbExecutor {
	if tx, ok := ctx.Value(postgres.TxContextKey).(*sql.Tx); ok && tx != nil {
		return tx
	}
	return db
}
