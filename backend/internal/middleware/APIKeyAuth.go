package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"station_meteo_api/internal/repository"
	"strings"
	"time"
)

func APIKeyAuthMiddleware(db *mongo.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
			}

			apiKey := strings.TrimPrefix(authHeader, "ApiKey")
			if apiKey == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "API Key is empty")
			}

			// Vérifier la clé dans mongo
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			stationRepo := repository.NewStationRepository(db)
			userRepo := repository.NewUserRepository(db)
			user, err := userRepo.GetUserByAuthToken(ctx, apiKey)
			if err == nil {
				c.Set("user", user)
				return next(c)
			} else {
				station, err := stationRepo.FindByAuthKey(ctx, apiKey)
				if err == nil {
					c.Set("station", station)
					return next(c)
				} else {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API key")
				}
			}
		}
	}
}
