package want

import (
	"context"

	"github.com/ryutaKimu/kakebo/api/internal/model"
)

type WantRepository interface {
	GetWantAmount(ctx context.Context, userId int) (float64, error)
	FetchLatestWant(ctx context.Context, userId int) (*model.Want, error)
}
