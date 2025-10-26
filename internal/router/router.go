package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/ryutaKimu/kakebo/internal/controller"
	repository "github.com/ryutaKimu/kakebo/internal/infra/postgre/user"
	service "github.com/ryutaKimu/kakebo/internal/service/user"
)

func NewRouter(pg *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	userRepo := repository.NewUserRepository(pg)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	r.Post("/signup", userController.CreateUser)

	return r
}
