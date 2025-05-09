package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type Data struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

func generateData() []Data {
	data := make([]Data, 10)
	for i := 0; i < 10; i++ {
		data[i] = Data{
			Name:      "Item",
			Category:  "Category " + strconv.Itoa(i),
			CreatedAt: time.Now(),
		}
	}
	return data
}

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Route for generating Excel file
	e.GET("/export", exportExcel)

	// Start server
	e.Start(":8080")
}

func exportExcel(c echo.Context) error {
	// Create a new Excel file
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	}()

	// Create a new sheet
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Generate data
	data := generateData()

	// Set headers
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Category")
	f.SetCellValue("Sheet1", "C1", "Created At")

	// Populate data
	for i, item := range data {
		row := i + 2
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(row), item.Name)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(row), item.Category)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(row), item.CreatedAt.Format(time.RFC3339))
	}

	// Set active sheet of the workbook
	f.SetActiveSheet(index)

	// Save to buffer instead of file
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Set response headers for file download
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=data.xlsx")
	c.Response().Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
	c.Response().Header().Set("Content-Transfer-Encoding", "binary")
	c.Response().Header().Set("Cache-Control", "no-cache")

	// Write the file to the response
	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())
}
