package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/utils"
)

// @Summary 	create segment
// @Tags 		segment
// @Description	create segment
// @Accept 		json
// @Produce 	json
// @Param 		segment body entity.Segment true "segment"
// @Success 	200 {object} ResponseId
// @Failure 	400 {object} echo.HTTPError
// @Failure 	500 {object} echo.HTTPError
// @Router 		/segment/ [post]
func (h *Handler) createSegment(ctx echo.Context) error {
	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {

		utils.Logger.Error("incorrect segment data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect segment data"))
	}

	utils.Logger.Debug("got segment to create", zap.Any("segment", segment))

	id, err := h.services.CreateSegment(ctx.Request().Context(), segment)
	if err != nil {

		utils.Logger.Error("segment creation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("segment created")
	utils.Logger.Debug("segment created with id", zap.Int("id", id))

	return responseOk(ctx, ResponseId{Id: id})
}

// @Summary 	delete segment
// @Tags 		segment
// @Description	create segment
// @Accept 		json
// @Produce 	json
// @Param 		segment body entity.Segment true "segment"
// @Success 	200 {object} ResponseMessage
// @Failure 	400 {object} echo.HTTPError
// @Failure 	500 {object} echo.HTTPError
// @Router 		/segment/ [delete]
func (h *Handler) deleteSegment(ctx echo.Context) error {

	var segment entity.Segment

	if err := ctx.Bind(&segment); err != nil {

		utils.Logger.Error("incorrect segment data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect segment data"))
	}

	utils.Logger.Debug("got segment to delete", zap.String("name", segment.Name))

	if err := h.services.DeleteSegment(ctx.Request().Context(), segment.Name); err != nil {

		utils.Logger.Error("segment deletion error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("segment deleted")

	return responseOk(ctx, ResponseMessage{Message: "success"})
}
