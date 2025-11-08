package top

import (
	"context"
)

type TopRepository interface {
	GetSumFixedIncome(ctx context.Context, userId int) (float64, error)
	GetSumSubIncome(ctx context.Context, userId int) (float64, error)
	GetSumIncomeAdjustment(ctx context.Context, userId int) (float64, error)
	GetSumCost(ctx context.Context, userId int) (float64, error)
}
