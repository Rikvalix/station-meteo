package main

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"station_meteo_api/internal/config"
	"station_meteo_api/internal/form"
	"station_meteo_api/internal/router"
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

	router.InitRoutes(e, db.Database("station_meteo"))

	e.Validator = &form.MeasurementFormValidator{
		Validator: validator.New(),
	}
	port := os.Getenv("SERVER_PORT")
	logger.Infof("Serveur lancé sur http://localhost%s", port)
	if err := e.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
