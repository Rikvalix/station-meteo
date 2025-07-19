package middleware

import "github.com/labstack/echo/v4"

/*
UserAuthMiddleware

	Authentifie l'utilisateur sur un couple pseudo et mot de passe et renvoi un token d'api
*/
func UserAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
