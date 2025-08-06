package router

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"station_meteo_api/internal/handler"
	"station_meteo_api/internal/middleware"
	"station_meteo_api/internal/repository"
	"station_meteo_api/internal/service"
)

func InitRoutes(e *echo.Echo, db *mongo.Database) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to the Station Meteo API")
	})

	userRepository := repository.NewUserRepository(db)
	stationRepository := repository.NewStationRepository(db)
	measureRepository := repository.NewMeasureRepository(db)

	measurementService := service.NewMeasurementService(measureRepository, stationRepository)
	measurementHandler := handler.NewMeasurementHandler(measurementService)

	userService := service.NewUserService(userRepository, stationRepository)
	userHandler := handler.NewUserHandler(userService)

	// Routes liés à l'utilisateur
	userGroup := e.Group("/api/v1/user")

	// Routes publiques
	userGroup.POST("/login", userHandler.Login)

	// Routes protégées
	userAuthGroup := userGroup.Group("")
	userAuthGroup.Use(middleware.APIKeyAuthMiddleware(db))
	userAuthGroup.GET("/station", userHandler.GetAllStations)
	userAuthGroup.GET("/me", userHandler.Me)
	userAuthGroup.PUT("/update-password", userHandler.UpdatePassword)
	// Routes liés aux mesures
	stationGroup := e.Group("/api/v1/station")
	stationGroup.Use(middleware.APIKeyAuthMiddleware(db))

	stationGroup.POST("/measurements", measurementHandler.CreateMeasure)
	stationGroup.GET("/measurements", measurementHandler.GetAllMeasures)
	stationGroup.GET("/measurements/:id", measurementHandler.GetMeasureById)
	stationGroup.GET("/measurements/latest/:id", measurementHandler.GetLatestMeasure)

	// Routes liés aux statistiques
	stationGroup.GET("/stats/daily", func(c echo.Context) error {
		return c.String(http.StatusOK, "Daily statistics")
	})
	stationGroup.GET("/stats/weekly", func(c echo.Context) error {
		return c.String(http.StatusOK, "Weekly statistics")
	})
	stationGroup.GET("/stats/range", func(c echo.Context) error {
		return c.String(http.StatusOK, "Statistics for a range of dates from"+c.QueryParam("from")+" to "+c.QueryParam("to"))
	})

	// Routes liés aux systèmes
	stationGroup.GET("/meta/info", func(c echo.Context) error {
		return c.String(http.StatusOK, "System information")
	})
	stationGroup.GET("/system/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})
}
