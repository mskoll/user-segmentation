package app

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"time"
	"userSegmentation/internal/handler"
	"userSegmentation/internal/repo"
	"userSegmentation/internal/service"
	"userSegmentation/pkg/logger"
)

func Run() {

	if err := initConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}
	logger.Init()

	// repo (pg)

	log.Info("Starting app")
	db, err := repo.Init(repo.Conf{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalf("DB-init error: %s", err.Error())
	}
	log.Info("DB connected")
	log.Info("Initializing repositories")
	repos := repo.New(db)

	// service
	services := service.New(repos)
	// handler
	handlers := handler.New(services)
	// server
	e := echo.New()

	handlers.Route(e)
	s := &http.Server{Addr: ":8000"}

	go func() {
		if err := e.StartServer(s); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
