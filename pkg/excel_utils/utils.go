package excel_utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"

	pReader "github.com/A-Oez/go-mfr/pkg"
)

func GetTNumbers(filePath string) ([]string, error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: \n error_msg: %w", filePath, err)
	}

	sheetName := pReader.GetProperty("tNumberSheet")
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet %s: \n error_msg: %w", sheetName, err)
	}

	var columnValues []string

	for _, row := range rows {
		columnValues = append(columnValues, row[0])
	}

	return columnValues, nil
}

func FindNextEmptyRow(file *excelize.File, sheetName string) int {
	rows, err := file.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	for i := len(rows) - 1; i >= 0; i-- {
		row := rows[i]
		isEmpty := true
		for _, cell := range row {
			if strings.TrimSpace(cell) != "" {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			return i + 1
		}
	}

	return len(rows) + 1
}
