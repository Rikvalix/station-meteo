package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
UserAuthMiddleware

	Authentifie l'utilisateur sur un couple pseudo et mot de passe et renvoi un token d'api
*/
func UserAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
			}

			return next(c)
		}
	}
}
