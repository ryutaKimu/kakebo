package top

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/ryutaKimu/kakebo/internal/infra/postgre/dbutil"
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

func (r *TopRepository) GetSumFixedIncome(ctx context.Context, userId int, month int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)
	query, args, err := r.goqu.
		From("fixed_incomes").
		Select(goqu.COALESCE(goqu.SUM("amount"), 0)).
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.L("EXTRACT(MONTH FROM created_at) = ?", month),
		).
		ToSQL()
	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		return 0, fmt.Errorf("failed to scan total: %w", err)
	}

	return total, nil

}

func (r *TopRepository) GetSumSubIncome(ctx context.Context, userId int, month int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)
	query, args, err := r.goqu.
		From("sub_incomes").
		Select(goqu.COALESCE(goqu.SUM("amount"), 0)).
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.L("EXTRACT(MONTH FROM created_at) = ?", month),
		).
		ToSQL()
	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		return 0, fmt.Errorf("failed to scan total: %w", err)
	}

	return total, nil

}

func (r *TopRepository) GetSumIncomeAdjustment(ctx context.Context, userId int, month int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)

	query, args, err := r.goqu.
		From("income_adjustments").
		Select(goqu.COALESCE(goqu.SUM("amount"), 0)).As("total_amount").
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.L("EXTRACT(MONTH FROM created_at) = ?", month),
		).
		ToSQL()

	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		return 0, fmt.Errorf("failed to scan total: %w", err)
	}

	return total, nil
}

func (r *TopRepository) GetSumCost(ctx context.Context, userId int, month int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)

	query, args, err := r.goqu.
		From("fixed_costs").
		Select(goqu.COALESCE(goqu.SUM("amount"), 0).As("total_amount")).
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.L("EXTRACT(MONTH FROM created_at) = ?", month),
		).
		ToSQL()
	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		return 0, fmt.Errorf("failed to scan total cost: %w", err)
	}

	return total, nil
}
