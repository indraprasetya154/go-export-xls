package router

import (
	"net/http"

	"github.com/indraprasetya154/go-export-xls/src/modules/export"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	v1 := e.Group("/v1")

	exportRoute := v1.Group("/export")
	exportRoute.POST("/xls", export.ExportXLSHandler)

	return e
}
