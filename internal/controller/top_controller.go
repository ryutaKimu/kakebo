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

	fixedIncome, err := s.service.GetIncome(r.Context(), userId)
	if err != nil {
		log.Printf("[ERROR] failed to get income: %v", err)
		http.Error(w, "failed to get income", http.StatusInternalServerError)
		return
	}

	fixedCost, err := s.service.GetTotalCost(r.Context(), userId)

	if err != nil {
		log.Printf("[ERROR] failed to get Total cost: %v", err)
		http.Error(w, "failed to get cost", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"income": fixedIncome,
		"cost":   fixedCost,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
