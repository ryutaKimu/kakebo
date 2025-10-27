package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"unicode/utf8"

	"github.com/ryutaKimu/kakebo/internal/controller/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "名前は必須です", http.StatusBadRequest)
		return
	}

	if input.Email == "" {
		http.Error(w, "メールアドレスは必須です", http.StatusBadRequest)
		return
	}

	if utf8.RuneCountInString(input.Password) < 8 {
		err := errors.New("password must be at least 8 characters")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.CreateUser(r.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
