package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/lib/errTypes"
)

func (h *Handler) createUser(ctx echo.Context) error {

	var user entity.User

	if err := ctx.Bind(&user); err != nil {

		h.log.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(errTypes.ErrBadRequest, "incorrect user data"))
	}

	h.log.Debug("got user to create", zap.String("username", user.Username))

	id, err := h.services.CreateUser(user)
	if err != nil {

		h.log.Error("user creation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	h.log.Info("user created")
	h.log.Debug("user created with id", zap.Int("id", id))

	type response struct {
		Id int `json:"id"`
	}

	return ctx.JSON(http.StatusOK, response{Id: id})
}

func (h *Handler) userById(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {

		h.log.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(errTypes.ErrBadRequest, "bad request"))
	}

	h.log.Debug("got user id to select", zap.Int("id", id))

	user, err := h.services.UserById(id)
	if err != nil {

		h.log.Error("user selection error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	h.log.Info("user selected")
	h.log.Debug("user selected", zap.Any("user", user))

	return ctx.JSON(http.StatusOK, user)
}

func (h *Handler) addDelSegment(ctx echo.Context) error {

	var segments entity.AddDelSegments

	if err := ctx.Bind(&segments); err != nil {

		h.log.Error("incorrect user data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(errTypes.ErrBadRequest, "bad request"))
	}

	h.log.Debug("got data to update user segments", zap.Any("data", segments))

	err := h.services.AddDeleteSegment(segments)
	if err != nil {

		h.log.Error("update user segments error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	h.log.Info("user segments updated")

	type response struct {
		Message string `json:"message"`
	}

	return ctx.JSON(http.StatusOK, response{Message: "success"})
}

func (h *Handler) operations(ctx echo.Context) error {

	var userOperations entity.UserOperations

	if err := ctx.Bind(&userOperations); err != nil {

		h.log.Error("incorrect input data", zap.String("error", err.Error()))

		return responseErr(errors.Wrap(errTypes.ErrBadRequest, "incorrect input data"))
	}

	h.log.Debug("got data to receive a report", zap.Any("data", userOperations))

	res, err := h.services.Operations(userOperations)
	if err != nil {

		h.log.Error("operation report generation error", zap.String("error", err.Error()))

		return responseErr(err)
	}

	h.log.Info("operation report generated")

	return ctx.JSON(http.StatusOK, res)
}
