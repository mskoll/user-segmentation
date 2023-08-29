package handler

import (
	"github.com/labstack/echo/v4"
	"userSegmentation/internal/service"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Route(e *echo.Echo) {

	user := e.Group("/user")
	user.POST("/", h.createUser)
	user.GET("/:id", h.userById)
	user.POST("/segment", h.addDelSegment)
	user.POST("/operations", h.operations)

	segment := e.Group("/segment")
	segment.POST("/", h.createSegment)
	segment.DELETE("/", h.deleteSegment)

}
