package router

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"station_meteo_api/internal/handler"
	"station_meteo_api/internal/repository"
	"station_meteo_api/internal/service"
)

func InitRoutes(e *echo.Echo, db *mongo.Database) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Station Meteo API")
	})

	measureRepository := repository.NewMeasureRepository(db)
	measurementService := service.NewMeasurementService(measureRepository)
	measurementHandler := handler.NewMeasurementHandler(measurementService)

	// Routes liés aux mesures
	e.POST("/api/v1/measurements", measurementHandler.CreateMeasure)

	e.GET("/api/v1/measurements", measurementHandler.GetAllMeasures)

	e.GET("/api/v1/measurements/:id", measurementHandler.GetMeasureById)

	e.PUT("/api/v1/measurements/latest", measurementHandler.GetLatestMeasure)

	// Routes liés aux statistiques
	e.GET("/api/v1/stats/daily", func(c echo.Context) error {
		return c.String(http.StatusOK, "Daily statistics")
	})

	e.GET("/api/v1/stats/weekly", func(c echo.Context) error {
		return c.String(http.StatusOK, "Weekly statistics")
	})

	e.GET("/api/v1/stats/range", func(c echo.Context) error {
		return c.String(http.StatusOK, "Statistics for a range of dates from"+c.QueryParam("from")+" to "+c.QueryParam("to"))
	})

	// Routes liés aux systèmes
	e.GET("/api/v1/meta/info", func(c echo.Context) error {
		return c.String(http.StatusOK, "System information")
	})

	e.GET("/api/v1/system/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
}
