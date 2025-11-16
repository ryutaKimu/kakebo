// internal/router/router.go
package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/kakebo/api/internal/controller"
	"github.com/ryutaKimu/kakebo/api/internal/middleware"
)

func NewRouter(userController *controller.UserController, topController *controller.TopController) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/signup", userController.CreateUser)
	r.Post("/login", userController.Login)

	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/profile", userController.GetProfile)
		r.Get("/top", topController.GetTop)
	})

	return r
}
