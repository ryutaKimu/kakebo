package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ryutaKimu/kakebo/internal/controller"
	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	userRepoPkg "github.com/ryutaKimu/kakebo/internal/infra/postgre/user"
	"github.com/ryutaKimu/kakebo/internal/router"
	userServicePkg "github.com/ryutaKimu/kakebo/internal/service/user"
)

type App struct {
	server *http.Server
	pg     *postgres.Postgres
}

func NewApp() (*App, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	pg := postgres.NewPostgres()

	userRepo := userRepoPkg.NewUserRepository(pg.DB)
	userService, err := userServicePkg.NewUserService(pg, userRepo)
	if err != nil {
		return nil, err
	}
	userController := controller.NewUserController(userService)

	r := router.NewRouter(userController)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Println("App initialized on port", port)
	return &App{
		server: srv,
		pg:     pg,
	}, nil
}

func (a *App) Start() error {
	log.Println("Starting server...")
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) {
	if err := a.server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	if err := a.pg.Close(); err != nil {
		log.Printf("Database close error: %v", err)
	}
}
