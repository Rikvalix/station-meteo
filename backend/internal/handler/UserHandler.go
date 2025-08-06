package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/model"
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

func (userHandler *UserHandler) Me(c echo.Context) error {
	user := c.Get("user").(*model.UserModel)
	data, err := userHandler.service.Me(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
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

// GetAllStations Renvoi toutes les stations disponibles
func (userHandler *UserHandler) GetAllStations(c echo.Context) error {
	stations, err := userHandler.service.GetAllStations()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, stations)
}

func (userHandler *UserHandler) UpdatePassword(c echo.Context) error {
	user := c.Get("user").(*model.UserModel)
	newPassword := c.FormValue("password")
	oldPassword := c.FormValue("password")
	result := userHandler.service.UpdatePassword(user, oldPassword, newPassword)
	if result != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	return c.JSON(http.StatusOK, "")
}
