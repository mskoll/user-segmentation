package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	errType "userSegmentation/internal/lib/errTypes"
	"userSegmentation/internal/service"
)

type Handler struct {
	services *service.Service
	log      *zap.Logger
}

func New(services *service.Service, log *zap.Logger) *Handler {
	return &Handler{services: services, log: log}
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

func responseErr(err error) *echo.HTTPError {

	switch errors.Cause(err) {

	case errType.ErrNotFound:
		return echo.NewHTTPError(http.StatusNotFound, err)
	case errType.ErrAlreadyExists:
		return echo.NewHTTPError(http.StatusConflict, err)
	case errType.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}
