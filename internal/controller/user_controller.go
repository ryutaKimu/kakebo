package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/ryutaKimu/kakebo/internal/common"
	"github.com/ryutaKimu/kakebo/internal/request"
	"github.com/ryutaKimu/kakebo/internal/service/interfaces"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(s interfaces.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input request.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.CreateUser(r.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, interfaces.ErrUserAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		log.Printf("failed to create user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var input request.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "リクエストボディが不正です", http.StatusBadRequest)
		return
	}

	if err := input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signed, err := c.service.Login(r.Context(), input.Email, input.Password)
	if err != nil {
		log.Printf("failed to login: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if signed == "" {
		http.Error(w, "メールアドレスまたはパスワードが正しくありません", http.StatusUnauthorized)
		return
	}

	var response struct {
		Token string `json:"token"`
	}
	response.Token = signed

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId, err := common.GetCurrentUserID(r.Context())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := c.service.GetProfile(r.Context(), userId)
	if err != nil {
		log.Printf("failed to get profile: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
