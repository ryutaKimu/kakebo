package interfaces

import (
	"context"

	"github.com/ryutaKimu/kakebo/internal/model"
)

type TopService interface {
	GetIncome(ctx context.Context, userId int) (*model.FixedIncome, error)
	GetTotalCost(ctx context.Context, userId int) (float64, error)
}
