package interfaces

import (
	"context"
	"time"
)

type TopService interface {
	GetMonthlyPageSummary(ctx context.Context, userId int, now time.Time) (float64, float64, float64, float64, error)
}
