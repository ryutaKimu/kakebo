package top

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ryutaKimu/kakebo/internal/infra/postgre/dbutil"
	"github.com/ryutaKimu/kakebo/internal/model"
	repository "github.com/ryutaKimu/kakebo/internal/repository/top"
)

var _ repository.TopRepository = (*TopRepository)(nil)

type TopRepository struct {
	db   *sql.DB
	goqu goqu.DialectWrapper
}

func NewTopRepository(db *sql.DB) *TopRepository {
	return &TopRepository{
		db:   db,
		goqu: goqu.Dialect("postgres"),
	}
}

func (r *TopRepository) GetIncome(ctx context.Context, userId int) (*model.FixedIncome, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)
	query, args, err := r.goqu.
		From("fixed_incomes").
		Select("id", "user_id", "name", "amount", "pay_day", "created_at").
		Where(goqu.C("user_id").Eq(userId)).
		Order(goqu.I("created_at").Desc()).
		Limit(1).
		ToSQL()
	if err != nil {
		return nil, err
	}

	row := exec.QueryRowContext(ctx, query, args...)
	var income model.FixedIncome
	if err := row.Scan(
		&income.ID,
		&income.UserID,
		&income.Name,
		&income.Amount,
		&income.PayDay,
		&income.CreatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan income: %w", err)
	}

	return &income, nil

}
func (r *TopRepository) GetTotalCost(ctx context.Context, userId int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)

	query, args, err := r.goqu.
		From("fixed_costs").
		Select(goqu.SUM("amount").As("total_amount")).
		Where(goqu.C("user_id").Eq(userId)).
		ToSQL()
	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to scan total cost: %w", err)
	}

	return total, nil
}
