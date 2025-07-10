package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/model"
	"station_meteo_api/internal/service"
	"strconv"
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
	dataForm := new(form.MeasurementForm)
	if err := c.Bind(dataForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(dataForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	station := c.Get("station").(*model.StationModel)
	measurement, err := h.service.CreateMeasurement(dataForm, station)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create measurements",
		})
	}
	return c.JSON(http.StatusOK, measurement)
}

// GetAllMeasures Récupération de toutes les mesures
func (h *MeasurementHandler) GetAllMeasures(c echo.Context) error {
	limit := c.QueryParam("limit")
	// Conversion int
	limitConvert, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Limit must be an integer")
	}
	measurements, err := h.service.GetAllMeasurements(limitConvert)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch measurements",
		})
	}

	return c.JSON(http.StatusOK, measurements)
}

// GetMeasureById Récupération d'une mesure par son ID
func (h *MeasurementHandler) GetMeasureById(c echo.Context) error {
	id := c.Param("id")
	measure, err := h.service.GetMeasureById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch measurements",
		})
	}
	return c.JSON(http.StatusOK, measure)

}

// GetLatestMeasure Récupération de la dernière mesure
func (h *MeasurementHandler) GetLatestMeasure(c echo.Context) error {
	station := c.Get("station").(*model.StationModel)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello " + station.Name,
	})
}
