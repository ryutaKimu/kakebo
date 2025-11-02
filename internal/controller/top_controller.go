package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ryutaKimu/kakebo/internal/common"
	"github.com/ryutaKimu/kakebo/internal/model"
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

	fixedIncome, err := s.service.GetIncome(r.Context(), userId)
	if err != nil {
		log.Printf("[ERROR] failed to get income: %v", err)
		http.Error(w, "failed to get income", http.StatusInternalServerError)
		return
	}

	totalCost, err := s.service.GetTotalCost(r.Context(), userId)

	if err != nil {
		log.Printf("[ERROR] failed to get Total cost: %v", err)
		http.Error(w, "failed to get cost", http.StatusInternalServerError)
		return
	}

	response := struct {
		Income *model.FixedIncome `json:"income"`
		Cost   float64            `json:"cost"`
	}{
		Income: fixedIncome,
		Cost:   totalCost,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
