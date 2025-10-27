package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/kakebo/internal/controller"
	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	repository "github.com/ryutaKimu/kakebo/internal/infra/postgre/user"
	service "github.com/ryutaKimu/kakebo/internal/service/user"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	pg := postgres.NewPostgres()

	userRepo := repository.NewUserRepository(pg.DB)

	userService := service.NewUserService(pg, userRepo)

	userController := controller.NewUserController(userService)

	r.Post("/signup", userController.CreateUser)

	return r
}
