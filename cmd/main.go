package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryutaKimu/kakebo/internal/controller"
	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	repository "github.com/ryutaKimu/kakebo/internal/infra/postgre/user"
	"github.com/ryutaKimu/kakebo/internal/router"
	service "github.com/ryutaKimu/kakebo/internal/service/user"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}

	pg := postgres.NewPostgres()
	userRepo := repository.NewUserRepository(pg.DB)
	userService := service.NewUserService(pg, userRepo)
	userController := controller.NewUserController(userService)

	router := router.NewRouter(*userController)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server startup error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	pg.Close()
	log.Println("Starting shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
