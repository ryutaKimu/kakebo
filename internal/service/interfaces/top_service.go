package interfaces

import (
	"context"
	"time"

	"github.com/ryutaKimu/kakebo/internal/model"
)

type TopService interface {
	GetMonthlyPageSummary(ctx context.Context, userId int, now time.Time) (float64, float64, float64, float64, error)
	GetLatestWant(ctx context.Context, userId int) (*model.Want, error)
}
