package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"userSegmentation/internal/entity"
)

func (h *Handler) createUser(ctx echo.Context) error {

	var user entity.User

	if err := ctx.Bind(&user); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := h.services.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}

func (h *Handler) userById(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	// todo: refactor to segmentList
	// var res entity.SegmentList
	res, err := h.services.UserById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}

func (h *Handler) addDelSegment(ctx echo.Context) error {

	var segments entity.AddDelSegments

	if err := ctx.Bind(&segments); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	// todo: refactor to addDelSegments
	// AddDelSegment(ctx, segments)
	err := h.services.AddDeleteSegment(ctx.Request().Context(), segments.Id, segments.ToAdd, segments.ToDel)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "OK")
}
