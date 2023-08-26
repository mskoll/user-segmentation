package handler

import (
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"userSegmentation/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/create", h.createUser1).Methods("POST")
	return router
}

func (h *Handler) Route(e *echo.Echo) {

	user := e.Group("/user")
	user.POST("/create", h.createUser)
	e.GET("/user/:id", h.getUserById)

}
