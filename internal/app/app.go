package app

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"userSegmentation/internal/handler"
	"userSegmentation/internal/repo"
	"userSegmentation/internal/service"
	"userSegmentation/pkg/logger"
)

func Run() {

	// config
	if err := initConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}
	// logger
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
	e.Logger.Fatal(e.StartServer(s))
	//serv := new(server.Server)
	//go func() {
	//	if err := serv.Run(viper.GetString("port"), handlers); err != nil {
	//		log.Printf("Server error: %s", err.Error())
	//	}
	//}()
	log.Printf("Server started\n")

	//
	//stop := make(chan os.Signal, 1)
	//
	//signal.Notify(stop, os.Interrupt)
	//<-stop

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := serv.Shutdown(ctx); err != nil {
	//	log.Printf("Server shutdown error: %s", err.Error())
	//}
	//
	//if err := db.Close(); err != nil {
	//	log.Printf("DB connection close error: %s", err.Error())
	//}
}

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
