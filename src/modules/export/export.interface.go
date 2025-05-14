package export

import "github.com/labstack/echo/v4"

type ExportService interface {
	ExportXLS(payload ExportRequest) ([]byte, error)
}

type ExportHandler interface {
	ExportXLS(e *echo.Echo)
}
