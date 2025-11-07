package interfaces

import (
	"context"
)

type TopService interface {
	GetMonthlyTotalIncome(ctx context.Context, userId int) (float64, error)
	GetSumCost(ctx context.Context, userId int) (float64, error)
}
