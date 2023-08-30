package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/lib/errTypes"
)

func (h *Handler) createSegment(ctx echo.Context) error {
	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {

		h.log.Error("incorrect segment data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(errTypes.ErrBadRequest, "incorrect segment data"))
	}

	h.log.Debug("got segment to create", zap.Any("segment", segment))

	id, err := h.services.CreateSegment(segment)
	if err != nil {

		h.log.Error("segment creation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	h.log.Info("segment created")
	h.log.Debug("segment created with id", zap.Int("id", id))

	type response struct {
		Id int `json:"id"`
	}

	return ctx.JSON(http.StatusOK, response{Id: id})
}

func (h *Handler) deleteSegment(ctx echo.Context) error {

	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {

		h.log.Error("incorrect segment data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(errTypes.ErrBadRequest, "incorrect segment data"))
	}

	h.log.Debug("got segment to delete", zap.String("name", segment.Name))

	if err := h.services.DeleteSegment(segment.Name); err != nil {

		h.log.Error("segment deletion error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	h.log.Info("segment deleted")

	type response struct {
		Message string `json:"message"`
	}

	return ctx.JSON(http.StatusOK, response{Message: "success"})
}
