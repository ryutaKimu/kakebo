package adjustment

import (
	"context"
	"database/sql"
	"time"

	"github.com/ryutaKimu/kakebo/internal/infra/postgre/base"
	"github.com/ryutaKimu/kakebo/internal/repository/adjustment"
)

var _ adjustment.AdjustmentRepository = (*AdjustmentRepository)(nil)

type AdjustmentRepository struct {
	*base.BaseRepository
}

func NewAdjustmentRepository(db *sql.DB) *AdjustmentRepository {
	return &AdjustmentRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}

func (r *AdjustmentRepository) GetSumIncomeAdjustment(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.GetMonthlySum(ctx, "income_adjustments", "adjustment_date", userId, now)
}
