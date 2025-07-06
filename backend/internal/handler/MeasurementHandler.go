package handler

import (
	"github.com/labstack/echo/v4"
	"station_meteo_api/internal/service"
)

type MeasurementHandler struct {
	service *service.MeasurementService
}

func NewMeasurementHandler(service *service.MeasurementService) *MeasurementHandler {
	return &MeasurementHandler{
		service: service,
	}
}

// CreateMeasure Création d'une mesure
func (h *MeasurementHandler) CreateMeasure(c echo.Context) error {
	return c.JSON(201, "TODO")
}

// GetAllMeasures Récupération de toutes les mesures
func (h *MeasurementHandler) GetAllMeasures(c echo.Context) error {
	return c.JSON(200, "TODO")
}

// GetMeasureById Récupération d'une mesure par son ID
func (h *MeasurementHandler) GetMeasureById(c echo.Context) error {
	return c.JSON(200, "TODO")
}

// GetLatestMeasure Récupération de la dernière mesure
func (h *MeasurementHandler) GetLatestMeasure(c echo.Context) error {
	return c.JSON(200, "TODO")
}
