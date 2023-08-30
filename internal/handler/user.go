package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strconv"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/utils"
)

func (h *Handler) createUser(ctx echo.Context) error {

	var user entity.User

	if err := ctx.Bind(&user); err != nil {

		utils.Logger.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect user data"))
	}

	utils.Logger.Debug("got user to create", zap.String("username", user.Username))

	id, err := h.services.CreateUser(user)
	if err != nil {

		utils.Logger.Error("user creation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("user created")
	utils.Logger.Debug("user created with id", zap.Int("id", id))

	type response struct {
		Id int `json:"id"`
	}

	return responseOk(ctx, ResponseId{Id: id})
}

func (h *Handler) userById(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		utils.Logger.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "bad request"))
	}

	utils.Logger.Debug("got user id to select", zap.Int("id", id))

	user, err := h.services.UserById(id)
	if err != nil {

		utils.Logger.Error("user selection error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("user selected")
	utils.Logger.Debug("user selected", zap.Any("user", user))

	return responseOk(ctx, user)
}

func (h *Handler) addDelSegment(ctx echo.Context) error {

	var segments entity.AddDelSegments

	if err := ctx.Bind(&segments); err != nil {

		utils.Logger.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "bad request"))
	}

	utils.Logger.Debug("got data to update user segments", zap.Any("data", segments))

	err := h.services.AddDeleteSegment(segments)
	if err != nil {

		utils.Logger.Error("update user segments error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("user segments updated")

	type response struct {
		Message string `json:"message"`
	}

	return responseOk(ctx, ResponseMessage{Message: "success"})
}

func (h *Handler) operations(ctx echo.Context) error {

	var userOperations entity.UserOperations

	if err := ctx.Bind(&userOperations); err != nil {

		utils.Logger.Error("incorrect input data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(utils.ErrBadRequest, "incorrect input data"))
	}

	utils.Logger.Debug("got data to receive a report", zap.Any("data", userOperations))

	res, err := h.services.Operations(userOperations)
	if err != nil {

		utils.Logger.Error("operation report generation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	utils.Logger.Info("operation report generated")

	return responseOk(ctx, res)
}
