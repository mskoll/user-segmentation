package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	logd "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"userSegmentation/config"
	"userSegmentation/internal/handler"
	"userSegmentation/internal/lib/logger"
	"userSegmentation/internal/repo"
	"userSegmentation/internal/service"
)

func Run() {

	logd.Print("config initializing")

	cfg, err := config.New()
	if err != nil {
		logd.Fatal("reading config err", zap.String("error", err.Error()))
	}

	log := logger.CreateLogger()
	defer log.Sync()

	log.Info("starting app")

	db, err := repo.Init(&cfg.DB)
	if err != nil {
		log.Fatal("database connection err", zap.String("error", err.Error()))
	}

	log.Info("database connected")

	repos := repo.New(db)
	services := service.New(repos)
	handlers := handler.New(services, log)

	e := echo.New()

	handlers.Route(e)

	log.Info("server starting")
	go func() {
		if err = e.Start(":" + cfg.HTTP.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("app shutting down")

	if err = e.Shutdown(ctx); err != nil {
		log.Error("server shutting down err", zap.String("error", err.Error()))
	}

	if err = db.Close(); err != nil {
		log.Error("database connection close err", zap.String("error", err.Error()))
	}

}
