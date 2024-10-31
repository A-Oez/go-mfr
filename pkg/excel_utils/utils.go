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
		return nil, fmt.Errorf("excel-datei konnte nicht geöffnet werden %s:\n error_msg: %w", filePath, err)
	}

	sheetName := pReader.GetProperty("tnumber_sheet")
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("t-nummern konnten nicht extrahiert werden %s:\n error_msg: %w", sheetName, err)
	}

	var columnValues []string

	for _, row := range rows {
		columnValues = append(columnValues, row[0])
	}

	if(len(columnValues) == 0){
		return nil, fmt.Errorf("keine t-nummern in der excel arbeitsmappe: %s hinterlegt", sheetName)
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
