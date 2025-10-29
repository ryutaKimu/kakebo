package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/kakebo/internal/controller"
)

func NewRouter(userController controller.UserController) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/signup", userController.CreateUser)
	r.Post("/login", userController.Login)

	return r
}
