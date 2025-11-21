package interfaces

import (
	"context"
	"time"

	"github.com/ryutaKimu/kakebo/api/internal/model"
)

type TopService interface {
	GetMonthlyPageSummary(ctx context.Context, userId int, now time.Time) (float64, float64, float64, float64, error)
	GetLatestWants(ctx context.Context, userId int) ([]*model.Want, error)
}
