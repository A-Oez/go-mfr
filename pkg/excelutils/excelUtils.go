package excelutils

import (
	"log"
	"strings"

	pReader "MFRCli/pkg"

	"github.com/xuri/excelize/v2"
)

func GetTNumbers(filePath string) []string {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetName := pReader.GetProperty("tNumberSheet")
	rows, err := file.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	var columnValues []string

	for _, row := range rows {
		columnValues = append(columnValues, row[0])
	}

	return columnValues
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
