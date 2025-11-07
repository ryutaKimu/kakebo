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
	return r.getSumAmount(ctx, "fixed_incomes", "payment_month", userId, month)
}

func (r *TopRepository) GetSumSubIncome(ctx context.Context, userId int, month int) (float64, error) {
	return r.getSumAmount(ctx, "sub_incomes", "payment_month", userId, month)
}

func (r *TopRepository) GetSumIncomeAdjustment(ctx context.Context, userId int, month int) (float64, error) {
	return r.getSumAmount(ctx, "income_adjustments", "adjustment_month", userId, month)
}

func (r *TopRepository) GetSumCost(ctx context.Context, userId int, month int) (float64, error) {
	return r.getSumAmount(ctx, "fixed_costs", "payment_month", userId, month)
}

func (r *TopRepository) getSumAmount(ctx context.Context, tableName string, monthColumnName string, userId int, month int) (float64, error) {
	exec := dbutil.GetDBExecutor(ctx, r.db)
	query, args, err := r.goqu.
		From(tableName).
		Select(goqu.COALESCE(goqu.SUM("amount"), 0)).
		Where(
			goqu.C("user_id").Eq(userId),
			goqu.L(fmt.Sprintf("%s = ?", monthColumnName), month),
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
