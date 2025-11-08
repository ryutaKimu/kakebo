package interfaces

import (
	"context"
	"time"
)

type TopService interface {
	GetMonthlyPageSummary(ctx context.Context, userId int, now time.Time) (totalIncome float64, totalCost float64, err error)
}
