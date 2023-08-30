package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"userSegmentation/config"
	"userSegmentation/internal/database"
	"userSegmentation/internal/handler"
	"userSegmentation/internal/repo"
	"userSegmentation/internal/service"
	"userSegmentation/internal/utils"
)

type App struct {
	db     *database.DB
	router *echo.Echo
	cfg    *config.Config
}

func New() (app *App, err error) {

	app = &App{}

	utils.Logger.Info("config initializing")

	app.cfg, err = config.New()
	if err != nil {
		return nil, errors.Wrap(err, "reading config err")
	}

	app.db, err = database.New(&app.cfg.DB)
	if err != nil {
		return nil, errors.Wrap(err, "database connection err")
	}

	log.Info("database connected")

	app.router = echo.New()

	return app, err
}

func (app *App) Run() error {

	log.Info("starting app")

	repos := repo.New(app.db.DB)
	services := service.New(repos)
	handlers := handler.New(services)

	handlers.Route(app.router)

	log.Info("server starting")

	return app.router.Start(":" + app.cfg.HTTP.Port)
}

func (app *App) Shutdown(ctx context.Context) error {
	return app.router.Shutdown(ctx)
}
