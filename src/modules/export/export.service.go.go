package export

import (
	"errors"
	"fmt"

	"github.com/indraprasetya154/go-export-xls/src/constant"
	"github.com/indraprasetya154/go-export-xls/src/helper"
	"github.com/xuri/excelize/v2"
)

func ExportXLS(payload ExportRequest) ([]byte, error) {
	// Create a new Excel file
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			// Optional: Log the error
		}
	}()

	// Create a new sheet
	sheetIndex, err := f.NewSheet(constant.ExcelDefaultSheetName)
	if err != nil {
		return nil, err
	}

	// Get headers from payload
	headers := payload.Header
	if len(headers) == 0 {
		return nil, errors.New("no headers found in payload")
	}

	// Get content from payload
	content := payload.Content
	if len(content) == 0 {
		return nil, errors.New("no content found in payload")
	}

	// Set headers in Excel
	for i, header := range headers {
		colLetter := helper.GenerateColumn(i + 1) // +1 because Excel columns start at 1
		cellRef := fmt.Sprintf("%s1", colLetter)
		if err := f.SetCellValue(constant.ExcelDefaultSheetName, cellRef, header.Value); err != nil {
			return nil, err
		}
	}

	// Write content rows
	for rowIndex, row := range content {
		for colIndex, header := range headers {
			colLetter := helper.GenerateColumn(colIndex + 1)
			cellRef := fmt.Sprintf("%s%d", colLetter, rowIndex+2) // +2 because row 1 is header

			// Retrieve the value by header.Key
			if val, ok := row[header.Key]; ok {
				if err := f.SetCellValue(constant.ExcelDefaultSheetName, cellRef, val); err != nil {
					return nil, err
				}
			} else {
				// Optionally: leave blank or write "N/A"
				if err := f.SetCellValue(constant.ExcelDefaultSheetName, cellRef, ""); err != nil {
					return nil, err
				}
			}
		}
	}

	f.SetActiveSheet(sheetIndex)

	// Save to buffer instead of file
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
