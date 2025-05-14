package app

import (
	"log"

	"github.com/labstack/echo/v4"
)

type (
	App struct {
		Echo *echo.Echo
	}
)

func InitApp() *App {
	err := InitConfig()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	e := echo.New()
	return &App{Echo: e}
}
