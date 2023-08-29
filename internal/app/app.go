package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
	"userSegmentation/internal/handler"
	"userSegmentation/internal/lib/logger"
	"userSegmentation/internal/repo"
	"userSegmentation/internal/service"
)

func Run() {
	log := logger.CreateLogger()
	defer log.Sync()
	if err := initConfig(); err != nil {
		log.Fatal("config init err", zap.String("error", err.Error()))
	}

	log.Info("starting app")

	db, err := repo.Init(repo.Conf{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatal("database connection error", zap.String("error", err.Error()))
	}

	log.Info("database connected")

	repos := repo.New(db)
	services := service.New(repos)
	handlers := handler.New(services, log)

	e := echo.New()

	handlers.Route(e)

	log.Info("server starting", zap.String("port", ":8000"))
	go func() {
		if err = e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("app shutting down")

	if err = e.Shutdown(ctx); err != nil {
		log.Fatal("server shutting down err", zap.String("error", err.Error()))
	}

}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
