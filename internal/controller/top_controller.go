package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryutaKimu/kakebo/internal/common"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
)

type TopController struct {
	service interfaces.TopService
}

func NewTopController(s interfaces.TopService) *TopController {
	return &TopController{service: s}
}

func (s *TopController) GetTop(w http.ResponseWriter, r *http.Request) {
	userId, err := common.GetCurrentUserID(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	totalIncome, err := s.service.GetMonthlyTotalIncome(r.Context(), userId)
	if err != nil {
		log.Printf("[ERROR] failed to get monthly total income: %v", err)
http.Error(w, "failed to get monthly total income", http.StatusInternalServerError)
		return
	}

	costTotalAmount, err := s.service.GetMonthlyTotalCost(r.Context(), userId)

	if err != nil {
		log.Printf("[ERROR] failed to get monthly total cost: %v", err)
		http.Error(w, "failed to get cost", http.StatusInternalServerError)
		return
	}

	response := struct {
		TotalIncome     float64 `json:"total_income"`
		CostTotalAmount float64 `json:"total_cost"`
	}{
		TotalIncome:     totalIncome,
		CostTotalAmount: costTotalAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
