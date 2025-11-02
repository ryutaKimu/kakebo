package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ryutaKimu/kakebo/internal/controller"
	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	"github.com/ryutaKimu/kakebo/internal/infra/postgre/top"
	userRepoPkg "github.com/ryutaKimu/kakebo/internal/infra/postgre/user"
	"github.com/ryutaKimu/kakebo/internal/router"
	"github.com/ryutaKimu/kakebo/internal/service"
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

	topRepo := top.NewTopRepository(pg.DB)
	topService := service.NewTopService(topRepo)
	topController := controller.NewTopController(topService)

	r := router.NewRouter(userController, topController)

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

func (a *App) Shutdown(ctx context.Context) error {
	serverErr := a.server.Shutdown(ctx)
	if serverErr != nil {
		log.Printf("Server shutdown error: %v", serverErr)
	}

	dbErr := a.pg.Close()
	if dbErr != nil {
		log.Printf("Database close error: %v", dbErr)
	}

	if serverErr != nil {
		return serverErr
	}
	return dbErr
}
