package cost

import (
	"context"
	"database/sql"
	"time"

	"github.com/ryutaKimu/kakebo/internal/infra/postgre/base"
	"github.com/ryutaKimu/kakebo/internal/repository/cost"
)

var _ cost.CostRepository = (*CostRepository)(nil)

type CostRepository struct {
	*base.BaseRepository
}

func NewCostRepository(db *sql.DB) *CostRepository {
	return &CostRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}

func (r *CostRepository) GetSumCost(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.GetMonthlySum(ctx, "fixed_costs", "payment_date", userId, now)
}
