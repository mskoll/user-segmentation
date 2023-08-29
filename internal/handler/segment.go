package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"userSegmentation/internal/entity"
)

func (h *Handler) createSegment(ctx echo.Context) error {
	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	id, err := h.services.CreateSegment(ctx.Request().Context(), segment)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, id)
}

func (h *Handler) deleteSegment(ctx echo.Context) error {
	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := h.services.DeleteSegment(ctx.Request().Context(), segment.Name); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, "OK")
}
