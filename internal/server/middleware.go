package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func authenticator(token string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(authHeader) < 2 || authHeader[0] != "Bearer" {
				return c.NoContent(http.StatusUnauthorized)
			}
			if authHeader[1] != token {
				return c.NoContent(http.StatusForbidden)
			}
			return next(c)
		}
	}
}
