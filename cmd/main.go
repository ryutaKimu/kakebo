package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryutaKimu/kakebo/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("App initialization failed: %v", err)
	}

	go func() {
                if err := a.Start(); err != nil && err != http.ErrServerClosed {
                        log.Printf("Server startup error: %v", err)
                }
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Starting shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.Shutdown(ctx)

	log.Println("Server stopped gracefully")
}
