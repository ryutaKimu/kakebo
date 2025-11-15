package income

import (
	"context"
	"time"
)

type IncomeRepository interface {
	GetSumFixedIncome(ctx context.Context, userId int, now time.Time) (float64, error)
	GetSumSubIncome(ctx context.Context, userId int, now time.Time) (float64, error)
	GetSumIncomeAdjustment(ctx context.Context, userId int, now time.Time) (float64, error)
}
