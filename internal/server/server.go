package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		return fmt.Errorf("TOKEN is missing")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(authenticator(token))

	e.PATCH("/switch", patchSwitch)

	return e.Start(fmt.Sprintf(":%s", port))
}
