package server

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"station_meteo_api/internal/config"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	logger := config.NewLogger()

	e := echo.New()
	
	port := os.Getenv("SERVER_PORT")
	logger.Infof("Serveur lancé sur http://localhost%s", port)
	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Erreur au démarrage du serveur : %v", err)
	}
}
