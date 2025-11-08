package interfaces

import (
	"context"
)

type TopService interface {
	GetMonthlyTotalIncome(ctx context.Context, userId int) (float64, error)
	GetMonthlyTotalCost(ctx context.Context, userId int) (float64, error)
}
