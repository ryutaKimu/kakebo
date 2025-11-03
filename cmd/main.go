package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryutaKimu/kakebo/internal/app"
	appLog "github.com/ryutaKimu/kakebo/internal/pkg/log"
	"go.uber.org/zap"
)

func main() {
	if err := appLog.InitLogger(); err != nil {
		log.Fatalf("Logger initialization failed: %v", err)
	}
	defer appLog.Sync()
	defer appLog.Close()

	a, err := app.NewApp()
	if err != nil {
		appLog.Fatal("App initialization failed", zap.Error(err))
	}

	go func() {
		if err := a.Start(); err != nil && err != http.ErrServerClosed {
			appLog.Error("Server startup error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLog.Info("Starting shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.Shutdown(ctx)

	appLog.Info("Server stopped gracefully")
}
