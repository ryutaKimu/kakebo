package saving

import (
	"context"
	"time"
)

type SavingRepository interface {
	GetSumSaving(ctx context.Context, userId int, now time.Time) (float64, error)
}
