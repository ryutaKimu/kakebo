package want

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ryutaKimu/kakebo/internal/infra/postgre/base"
	"github.com/ryutaKimu/kakebo/internal/infra/postgre/dbutil"
	"github.com/ryutaKimu/kakebo/internal/model"
	repo "github.com/ryutaKimu/kakebo/internal/repository/want"
)

var _ repo.WantRepository = (*WantRepository)(nil)

type WantRepository struct {
	*base.BaseRepository
}

func NewWantRepository(db *sql.DB) *WantRepository {
	return &WantRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}

func (r *WantRepository) GetWantAmount(ctx context.Context, userId int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.Db)
	query, args, err := r.Goqu.
		From("wants").
		Select(goqu.C("target_amount")).
		Where(goqu.C("user_id").Eq(userId)).
		Order(goqu.I("created_at").Desc()).
		Limit(1).
		ToSQL()
	if err != nil {
		return 0, err
	}
	var amount float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&amount); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to scan: %w", err)
	}
	return amount, nil
}

func (r *WantRepository) FetchLatestWant(ctx context.Context, userId int) (*model.Want, error) {
	exec := dbutil.GetDBExecutor(ctx, r.Db)
	query, args, err := r.Goqu.
		From("wants").
		Select(
			"id",
			"user_id",
			"name",
			"target_amount",
			"target_date",
			"purchased",
			"purchased_at",
			"created_at",
		).
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.C("purchased").Eq(false)).
		Order(goqu.I("created_at").Desc()).
		Limit(1).
		ToSQL()

	if err != nil {
		return nil, err
	}

	var Want model.Want
	row := exec.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&Want.ID,
		&Want.UserId,
		&Want.Name,
		&Want.TargetAmount,
		&Want.TargetDate,
		&Want.Purchased,
		&Want.PurchasedAt,
		&Want.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &Want, nil

}
