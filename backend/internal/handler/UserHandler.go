package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
	return c.JSON(http.StatusOK, "Login  route")
}
