package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"userSegmentation/internal/entity"
)

func (h *Handler) createUser1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	resp := "userCreated"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func (h *Handler) createUser(ctx echo.Context) error {
	var user entity.User
	er := ctx.Bind(&user)

	if er != nil {
		return ctx.String(http.StatusInternalServerError, er.Error())
	}

	id, err := h.services.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, id)
}

func (h *Handler) getUserById(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	res, err := h.services.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
