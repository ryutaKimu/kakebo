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
	const passwordLengthError = "パスワードは8文字以上である必要があります"
	const nameRequiredError = "名前は必須です"
	const emailRequiredError = "メールアドレスは必須です"
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
		http.Error(w, nameRequiredError, http.StatusBadRequest)
		return
	}

	if input.Email == "" {
		http.Error(w, emailRequiredError, http.StatusBadRequest)
		return
	}

	if utf8.RuneCountInString(input.Password) < 8 {
		err := errors.New(passwordLengthError)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.CreateUser(r.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
