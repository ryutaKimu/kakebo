package adjustment

import (
	"context"
	"time"
)

type AdjustmentRepository interface {
	GetSumIncomeAdjustment(ctx context.Context, userId int, now time.Time) (float64, error)
}
