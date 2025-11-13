package top

import (
	"context"
	"time"
)

type TopRepository interface {
	GetSumFixedIncome(ctx context.Context, userId int, now time.Time) (float64, error)
	GetSumSubIncome(ctx context.Context, userId int, now time.Time) (float64, error)
	GetSumIncomeAdjustment(ctx context.Context, userId int, now time.Time) (float64, error)
	GetSumCost(ctx context.Context, userId int, now time.Time) (float64, error)
	GetSumSaving(ctx context.Context, userId int, now time.Time) (float64, error)
	GetWantAmount(ctx context.Context, userId int) (float64, error)
}
