// internal/router/server.go
package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/kakebo/internal/controller"
)

func NewRouter(userController *controller.UserController) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", userController.CreateUser)
		r.Post("/login", userController.Login)
	})

	return r
}
