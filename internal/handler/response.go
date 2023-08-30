package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"userSegmentation/internal/utils"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseId struct {
	Id int `json:"id"`
}

func responseOk(ctx echo.Context, data interface{}) error {

	return ctx.JSON(http.StatusOK, data)
}

func responseErr(err error) *echo.HTTPError {

	switch errors.Cause(err) {
	case utils.ErrNotFound:
		return echo.NewHTTPError(http.StatusNotFound, err)
	case utils.ErrAlreadyExists:
		return echo.NewHTTPError(http.StatusConflict, err)
	case utils.ErrBadRequest:
		return echo.NewHTTPError(http.StatusBadRequest, err)
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
}
