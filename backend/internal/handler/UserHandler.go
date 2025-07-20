package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Login authentification de l'utilisateur
func (userHandler *UserHandler) Login(c echo.Context) error {
	dataForm := new(form.LoginForm)
	if err := c.Bind(dataForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dataForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	auth, err := userHandler.service.Login(dataForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, auth)
}
