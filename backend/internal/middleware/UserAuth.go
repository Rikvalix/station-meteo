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

/*
UserAuthMiddleware

	Authentifie l'utilisateur sur un couple pseudo et mot de passe et renvoi un token d'api
*/
func UserAuthMiddleware(db *mongo.Database) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
			}

			bearer := strings.TrimPrefix(authHeader, "Bearer ")
			if bearer == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
			}

			// Vérifier si l'utilisateur existe
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			repo := repository.NewUserRepository(db)
			user, err := repo.GetUserByAuthToken(ctx, bearer)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API Key")
			}

			// Set l'utilisateur courant
			c.Set("user", user)
			return next(c)
		}
	}
}
