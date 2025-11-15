package want

import (
	"context"
)

type WantRepository interface {
	GetWantAmount(ctx context.Context, userId int) (float64, error)
}
