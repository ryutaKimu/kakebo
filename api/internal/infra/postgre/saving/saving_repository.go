package saving

import (
	"context"
	"database/sql"
	"time"

	"github.com/ryutaKimu/kakebo/api/internal/infra/postgre/base"
	repo "github.com/ryutaKimu/kakebo/api/internal/repository/saving"
)

var _ repo.SavingRepository = (*SavingRepository)(nil)

type SavingRepository struct {
	*base.BaseRepository
}

func NewSavingRepository(db *sql.DB) *SavingRepository {
	return &SavingRepository{
		BaseRepository: base.NewBaseRepository(db),
	}
}

func (r *SavingRepository) GetSumSaving(ctx context.Context, userId int, now time.Time) (float64, error) {
	return r.GetMonthlySum(ctx, "savings", "saved_at", userId, now)
}
