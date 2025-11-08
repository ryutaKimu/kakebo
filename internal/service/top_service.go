package service

import (
	"context"
	"time"

	"github.com/ryutaKimu/kakebo/internal/repository/top"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
	"golang.org/x/sync/errgroup"
)

type TopServiceImpl struct {
	repo top.TopRepository
}

func NewTopService(topRepository top.TopRepository) interfaces.TopService {
	return &TopServiceImpl{repo: topRepository}
}
func (s *TopServiceImpl) GetMonthlyPageSummary(ctx context.Context, userId int, now time.Time) (float64, float64, error) {
	var (
		fixedIncomeAmount float64
		subIncomeAmount   float64
		adjustmentAmount  float64
		totalCost         float64
	)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		fi, err := s.repo.GetSumFixedIncome(ctx, userId, now)
		if err != nil {
			return err
		}
		fixedIncomeAmount = fi
		return nil
	})

	g.Go(func() error {
		si, err := s.repo.GetSumSubIncome(ctx, userId, now)
		if err != nil {
			return err
		}
		subIncomeAmount = si
		return nil
	})

	g.Go(func() error {
		adj, err := s.repo.GetSumIncomeAdjustment(ctx, userId, now)
		if err != nil {
			return err
		}
		adjustmentAmount = adj
		return nil
	})

	g.Go(func() error {
		cost, err := s.repo.GetSumCost(ctx, userId, now)
		if err != nil {
			return err
		}
		totalCost = cost
		return nil
	})

	if err := g.Wait(); err != nil {
		return 0, 0, err
	}

	totalIncome := fixedIncomeAmount + subIncomeAmount + adjustmentAmount
	return totalIncome, totalCost, nil
}
