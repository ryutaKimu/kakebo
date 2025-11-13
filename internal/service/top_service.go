package service

import (
	"context"
	"log"
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
func (s *TopServiceImpl) GetMonthlyPageSummary(ctx context.Context, userId int, now time.Time) (float64, float64, float64, float64, error) {
	var (
		fixedIncomeAmount float64
		subIncomeAmount   float64
		adjustmentAmount  float64
		totalCost         float64
		saving            float64
		want              float64
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

	g.Go(func() error {
		sv, err := s.repo.GetSumSaving(ctx, userId, now)
		if err != nil {
			return err
		}
		saving = sv
		return nil
	})

	g.Go(func() error {
		wt, err := s.repo.GetWantAmount(ctx, userId)
		if err != nil {
			return err
		}
		want = wt
		return nil
	})

	if err := g.Wait(); err != nil {
		return 0, 0, 0, 0, err
	}

	totalIncome := fixedIncomeAmount + subIncomeAmount + adjustmentAmount
	amountDistance := want - saving
	if amountDistance <= 0 {
		log.Printf("🎉 目標達成！ 貯金額が目標 %.2f を超えました。", want)
		return 0, 0, 0, 0, nil
	}
	return totalIncome, totalCost, saving, amountDistance, nil
}
