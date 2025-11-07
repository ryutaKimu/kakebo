package service

import (
	"context"
	"time"

	"github.com/ryutaKimu/kakebo/internal/repository/top"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
)

type TopServiceImpl struct {
	repo top.TopRepository
}

func NewTopService(topRepository top.TopRepository) interfaces.TopService {
	return &TopServiceImpl{repo: topRepository}
}

func (s *TopServiceImpl) GetMonthlyTotalIncome(ctx context.Context, userId int) (float64, error) {
	month := int(time.Now().Month())
	fixedIncome, err := s.repo.GetSumFixedIncome(ctx, userId, month)
	if err != nil {
		return 0, err
	}

	subIncome, err := s.repo.GetSumSubIncome(ctx, userId, month)
	if err != nil {
		return 0, err
	}

	incomeAdjusted, err := s.repo.GetSumIncomeAdjustment(ctx, userId, month)
	if err != nil {
		return 0, err
	}

	costs, err := s.repo.GetSumCost(ctx, userId, month)
	if err != nil {
		return 0, err
	}

	totalAmount := fixedIncome + subIncome

	total := (totalAmount + incomeAdjusted) - costs

	return total, nil
}

func (s *TopServiceImpl) GetSumCost(ctx context.Context, userId int) (float64, error) {
	month := int(time.Now().Month())
	return s.repo.GetSumCost(ctx, userId, month)
}
