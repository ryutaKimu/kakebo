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
	appLog "github.com/ryutaKimu/kakebo/internal/pkg/log" // ← パスはあなたのlogパッケージに合わせて
	"go.uber.org/zap"
)

func main() {
	// ① zapログ初期化
	if err := appLog.InitLogger(); err != nil {
		log.Fatalf("Logger initialization failed: %v", err)
	}
	defer appLog.Sync() // 最後にフラッシュして安全に閉じる

	// ② アプリ初期化
	a, err := app.NewApp()
	if err != nil {
		appLog.Fatal("App initialization failed", zap.Error(err))
	}

	// ③ サーバー起動
	go func() {
		if err := a.Start(); err != nil && err != http.ErrServerClosed {
			appLog.Error("Server startup error", zap.Error(err))
		}
	}()

	// ④ 終了シグナルを待つ
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLog.Info("Starting shutdown...")

	// ⑤ シャットダウン処理
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.Shutdown(ctx)

	appLog.Info("Server stopped gracefully")
}
