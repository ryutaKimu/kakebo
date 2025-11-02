package controller

import (
	"net/http"

	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
)

type TopController struct {
	service interfaces.TopService
}

func NewTopController(s interfaces.TopService) *TopController {
	return &TopController{service: s}
}

func (s *TopController) GetTop(w http.ResponseWriter, r *http.Request) {

}
