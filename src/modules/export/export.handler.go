package export

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ExportXLSHandler(c echo.Context) error {
	var exportRequest ExportRequest
	err := c.Bind(&exportRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	buffer, err := ExportXLS(exportRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Set response headers for file download
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=data.xlsx")
	c.Response().Header().Set("Content-Length", strconv.Itoa(len(buffer)))
	c.Response().Header().Set("Content-Transfer-Encoding", "binary")
	c.Response().Header().Set("Cache-Control", "no-cache")

	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer)
}
