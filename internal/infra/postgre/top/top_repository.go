package top

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

func (r *TopRepository) GetSumFixedIncome(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.getSumAmount(ctx, "fixed_incomes", "payment_date", userId, now)
}

func (r *TopRepository) GetSumSubIncome(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.getSumAmount(ctx, "sub_incomes", "payment_date", userId, now)
}

func (r *TopRepository) GetSumIncomeAdjustment(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.getSumAmount(ctx, "income_adjustments", "adjustment_date", userId, now)
}

func (r *TopRepository) GetSumCost(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.getSumAmount(ctx, "fixed_costs", "payment_date", userId, now)
}

func (r *TopRepository) GetSumSaving(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.getSumAmount(ctx, "savings", "saved_at", userId, now)
}

func (r *TopRepository) GetWant(ctx context.Context, userId int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)
	query, args, err := r.goqu.
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
			return 0, err
		}
		return 0, fmt.Errorf("failed to scan: %w", err)
	}
	return amount, nil
}

func (r *TopRepository) getSumAmount(ctx context.Context, tableName string, columnName string, userId int, now time.Time) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)

	year, month, _ := now.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	query, args, err := r.goqu.
		From(tableName).
		Select(goqu.COALESCE(goqu.SUM("amount"), 0)).
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.I(columnName).Gte(startOfMonth),
			goqu.I(columnName).Lt(endOfMonth),
		).
		ToSQL()
	if err != nil {
		return 0, err
	}

	var total float64
	row := exec.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil {
		return 0, fmt.Errorf("failed to scan total from %s: %w", tableName, err)
	}

	return total, nil
}
