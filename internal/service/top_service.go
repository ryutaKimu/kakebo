package service

import (
	"context"
	"log"
	"time"

	"github.com/ryutaKimu/kakebo/internal/model"
	"github.com/ryutaKimu/kakebo/internal/repository/adjustment"
	"github.com/ryutaKimu/kakebo/internal/repository/cost"
	"github.com/ryutaKimu/kakebo/internal/repository/income"
	"github.com/ryutaKimu/kakebo/internal/repository/saving"
	"github.com/ryutaKimu/kakebo/internal/repository/want"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
	"golang.org/x/sync/errgroup"
)

type TopServiceImpl struct {
	IncomeRepo     income.IncomeRepository
	CostRepo       cost.CostRepository
	AdjustmentRepo adjustment.AdjustmentRepository
	SavingRepo     saving.SavingRepository
	WantRepo       want.WantRepository
}

func NewTopService(
	incomeRepo income.IncomeRepository,
	costRepo cost.CostRepository,
	adjustmentRepo adjustment.AdjustmentRepository,
	savingRepo saving.SavingRepository,
	wantRepo want.WantRepository,
) interfaces.TopService {
	return &TopServiceImpl{
		IncomeRepo:     incomeRepo,
		CostRepo:       costRepo,
		AdjustmentRepo: adjustmentRepo,
		SavingRepo:     savingRepo,
		WantRepo:       wantRepo,
	}
}

func (s *TopServiceImpl) GetMonthlyPageSummary(
	ctx context.Context, userID int, now time.Time,
) (float64, float64, float64, float64, error) {

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
		amount, err := s.IncomeRepo.GetSumFixedIncome(ctx, userID, now)
		if err != nil {
			return err
		}
		fixedIncomeAmount = amount
		return nil
	})

	g.Go(func() error {
		amount, err := s.IncomeRepo.GetSumSubIncome(ctx, userID, now)
		if err != nil {
			return err
		}
		subIncomeAmount = amount
		return nil
	})

	g.Go(func() error {
		amount, err := s.AdjustmentRepo.GetSumIncomeAdjustment(ctx, userID, now)
		if err != nil {
			return err
		}
		adjustmentAmount = amount
		return nil
	})

	g.Go(func() error {
		amount, err := s.CostRepo.GetSumCost(ctx, userID, now)
		if err != nil {
			return err
		}
		totalCost = amount
		return nil
	})

	g.Go(func() error {
		amount, err := s.SavingRepo.GetSumSaving(ctx, userID, now)
		if err != nil {
			return err
		}
		saving = amount
		return nil
	})

	g.Go(func() error {
		amount, err := s.WantRepo.GetWantAmount(ctx, userID)
		if err != nil {
			return err
		}
		want = amount
		return nil
	})

	if err := g.Wait(); err != nil {
		return 0, 0, 0, 0, err
	}

	totalIncome := fixedIncomeAmount + subIncomeAmount + adjustmentAmount
	amountDistance := want - saving
	if amountDistance <= 0 {
		log.Printf("🎉 目標達成！ 貯金額が目標 %.2f を超えました。", want)
		return totalIncome, totalCost, saving, 0, nil
	}

	return totalIncome, totalCost, saving, amountDistance, nil
}

func (s *TopServiceImpl) GetLatestWant(ctx context.Context, userId int) (*model.Want, error) {
	return s.WantRepo.FetchLatestWant(ctx, userId)
}
