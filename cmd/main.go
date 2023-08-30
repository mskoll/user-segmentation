package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
	"userSegmentation/internal/app"
	"userSegmentation/internal/utils"
)

func main() {

	utils.CreateLogger()
	defer utils.Logger.Sync()

	app, err := app.New()
	if err != nil {
		utils.Logger.Fatal("init app err", zap.String("error", err.Error()))
	}

	if err = app.Run(); err != nil {
		utils.Logger.Fatal("shutting down the server", zap.String("error", err.Error()))
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("app shutting down")

	if err = app.Shutdown(ctx); err != nil {
		log.Error("server shutting down err", zap.String("error", err.Error()))
	}
}
