package cost

import (
	"context"
	"time"
)

type CostRepository interface {
	GetSumCost(ctx context.Context, userId int, now time.Time) (float64, error)
}
