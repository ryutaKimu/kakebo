package service

import (
	"context"

	"github.com/ryutaKimu/kakebo/internal/model"
	"github.com/ryutaKimu/kakebo/internal/repository/top"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
)

type TopServiceImpl struct {
	repo top.TopRepository
}

func NewTopService(TopRepository top.TopRepository) interfaces.TopService {
	return &TopServiceImpl{repo: TopRepository}
}

func (s *TopServiceImpl) GetIncome(ctx context.Context, userId int) (*model.FixedIncome, error) {
	return s.repo.GetIncome(ctx, userId)
}

func (s *TopServiceImpl) GetTotalCost(ctx context.Context, userId int) (float64, error) {
	return s.repo.GetTotalCost(ctx, userId)
}
