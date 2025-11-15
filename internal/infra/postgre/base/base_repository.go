package base

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/ryutaKimu/kakebo/internal/infra/postgre/dbutil"
)

type BaseRepository struct {
	Db   *sql.DB
	Goqu *goqu.Database
}

func NewBaseRepository(db *sql.DB) *BaseRepository {
	return &BaseRepository{
		Db:   db,
		Goqu: goqu.New("postgres", db),
	}
}

func (b *BaseRepository) GetMonthlySum(
	ctx context.Context,
	table string,
	dateColumn string,
	userID int,
	now time.Time,
) (float64, error) {
	// 月初・月末
	year, month, _ := now.Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 1, 0)

	exec := dbutil.GetDBExecutor(ctx, b.Db)

	query, args, err := b.Goqu.
		From(table).
		Select(goqu.COALESCE(goqu.SUM("amount"), 0)).
		Where(
			goqu.C("user_id").Eq(userID),
			goqu.I(dateColumn).Gte(start),
			goqu.I(dateColumn).Lt(end),
		).ToSQL()

	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}
