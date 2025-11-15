package income

import (
	"context"
	"database/sql"
	"time"

	"github.com/ryutaKimu/kakebo/internal/infra/postgre/base"
	repo "github.com/ryutaKimu/kakebo/internal/repository/income"
)

var _ repo.IncomeRepository = (*IncomeRepository)(nil)

type IncomeRepository struct {
	*base.BaseRepository
}

func NewIncomeRepository(db *sql.DB) *IncomeRepository {
	return &IncomeRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}

func (r *IncomeRepository) GetSumFixedIncome(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.GetMonthlySum(ctx, "fixed_incomes", "payment_date", userId, now)
}

func (r *IncomeRepository) GetSumSubIncome(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.GetMonthlySum(ctx, "sub_incomes", "payment_date", userId, now)
}
