package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/utils"
)

// @Summary 	create user
// @Tags 		user
// @Description	create user
// @Accept 		json
// @Produce 	json
// @Param 		user body entity.User true "user"
// @Success 	200 {object} ResponseId
// @Failure 	400 {object} echo.HTTPError
// @Failure 	500 {object} echo.HTTPError
// @Router 		/user/ [post]
func (h *Handler) createUser(ctx echo.Context) error {

	var user entity.User

	if err := ctx.Bind(&user); err != nil {

		utils.Logger.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect user data"))
	}

	utils.Logger.Debug("got user to create", zap.String("username", user.Username))

	id, err := h.services.CreateUser(ctx.Request().Context(), user)
	if err != nil {

		utils.Logger.Error("user creation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("user created")
	utils.Logger.Debug("user created with id", zap.Int("id", id))

	return responseOk(ctx, ResponseId{Id: id})
}

// @Summary 	user by id
// @Tags 		user
// @Description	get user by id
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "user id"
// @Success 	200 {object} entity.SegmentList
// @Failure 	400 {object} echo.HTTPError
// @Failure 	404 {object} echo.HTTPError
// @Failure 	500 {object} echo.HTTPError
// @Router 		/user/:id [get]
func (h *Handler) userById(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		utils.Logger.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "bad request"))
	}

	utils.Logger.Debug("got user id to select", zap.Int("id", id))

	user, err := h.services.UserById(ctx.Request().Context(), id)
	if err != nil {

		utils.Logger.Error("user selection error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("user selected")
	utils.Logger.Debug("user selected", zap.Any("user", user))

	return responseOk(ctx, user)
}

// @Summary 	add and delete user segments
// @Tags 		user
// @Description	add and delete user segments
// @Accept 		json
// @Produce 	json
// @Param 		segments body entity.AddDelSegments true "segments"
// @Success 	200 {object} ResponseMessage
// @Failure 	400 {object} echo.HTTPError
// @Failure 	404 {object} echo.HTTPError
// @Failure 	409 {object} echo.HTTPError
// @Failure 	500 {object} echo.HTTPError
// @Router 		/user/segment [post]
func (h *Handler) addDelSegment(ctx echo.Context) error {

	var segments entity.AddDelSegments

	if err := ctx.Bind(&segments); err != nil {

		utils.Logger.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "bad request"))
	}

	utils.Logger.Debug("got data to update user segments", zap.Any("data", segments))

	err := h.services.AddDeleteSegment(ctx.Request().Context(), segments)
	if err != nil {

		utils.Logger.Error("update user segments error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("user segments updated")

	return responseOk(ctx, ResponseMessage{Message: "success"})
}

// @Summary 	user operation
// @Tags 		user
// @Description	report on adding and removing a user to a segment
// @Accept 		json
// @Produce 	json
// @Param 		userOperations body entity.UserOperations true "userOperations"
// @Success 	200 {array} entity.Operation
// @Failure 	400 {object} echo.HTTPError
// @Failure 	500 {object} echo.HTTPError
// @Router 		/user/operations [post]
func (h *Handler) operations(ctx echo.Context) error {

	var userOperations entity.UserOperations

	if err := ctx.Bind(&userOperations); err != nil {

		utils.Logger.Error("incorrect input data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect input data"))
	}

	utils.Logger.Debug("got data to receive a report", zap.Any("data", userOperations))

	res, err := h.services.Operations(ctx.Request().Context(), userOperations)
	if err != nil {

		utils.Logger.Error("operation report generation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("operation report generated")

	return responseOk(ctx, res)
}
