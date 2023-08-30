package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/utils"
)

func (h *Handler) createSegment(ctx echo.Context) error {
	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {

		utils.Logger.Error("incorrect segment data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect segment data"))
	}

	utils.Logger.Debug("got segment to create", zap.Any("segment", segment))

	id, err := h.services.CreateSegment(segment)
	if err != nil {

		utils.Logger.Error("segment creation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("segment created")
	utils.Logger.Debug("segment created with id", zap.Int("id", id))

	return responseOk(ctx, ResponseId{Id: id})
}

func (h *Handler) deleteSegment(ctx echo.Context) error {

	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {

		utils.Logger.Error("incorrect segment data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect segment data"))
	}

	utils.Logger.Debug("got segment to delete", zap.String("name", segment.Name))

	if err := h.services.DeleteSegment(segment.Name); err != nil {

		utils.Logger.Error("segment deletion error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("segment deleted")

	return responseOk(ctx, ResponseMessage{Message: "success"})
}
