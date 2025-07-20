package main

import (
	"errors"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"station_meteo_api/internal/config"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/router"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	logger := config.NewLogger()

	e := echo.New()

	// Configuration de la base de données
	db := config.ConnectDb()

	// Validator
	e.Validator = &form.MeasurementFormValidator{
		Validator: validator.New(),
	}

	router.InitRoutes(e, db.Database("station_meteo"))

	// Cors
	e.Use(middleware.CORS())
	port := os.Getenv("SERVER_PORT")
	logger.Infof("Serveur lancé sur http://localhost%s", port)
	if err := e.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
