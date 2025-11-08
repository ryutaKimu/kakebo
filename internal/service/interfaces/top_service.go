package interfaces

import (
	"context"
	"time"
)

type TopService interface {
	GetMonthlyTotalIncome(ctx context.Context, userId int, now time.Time) (float64, error)
	GetMonthlyTotalCost(ctx context.Context, userId int, now time.Time) (float64, error)
}
