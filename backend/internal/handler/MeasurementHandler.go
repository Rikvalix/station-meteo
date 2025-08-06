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
	date := c.QueryParam("date")
	stationId := c.QueryParam("station")

	if len(stationId) == 0 {
		return c.JSON(http.StatusBadRequest, "Station ID is required")
	}
	var measurements []model.MeasureModel
	var serviceError error
	// Si date est non nul
	if len(date) > 0 {
		measurements, serviceError = h.service.GetMeasuresByDate(date)
	} else {
		// Conversion int
		limitConvert, err := strconv.Atoi(limit)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Limit must be an integer")
		}
		measurements, serviceError = h.service.GetAllMeasurements(limitConvert, stationId)
	}
	if serviceError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": serviceError.Error()})
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

// GetMeasuresByDate Récupération d'une mesure avec une date
func (h *MeasurementHandler) GetMeasuresByDate(c echo.Context) error {
	date := c.QueryParam("date")

	measure, err := h.service.GetMeasuresByDate(date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch measurements",
		})
	}
	return c.JSON(http.StatusOK, measure)
}

// GetLatestMeasure Récupération de la dernière mesure
func (h *MeasurementHandler) GetLatestMeasure(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID is required",
		})
	}
	measure, err := h.service.GetLatestMesure(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch measurement",
		})
	}
	return c.JSON(http.StatusOK, measure)
}
