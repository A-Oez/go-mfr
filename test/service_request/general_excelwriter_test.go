package service_request_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/A-Oez/go-mfr/internal/service/excel_handler"
	json_parser "github.com/A-Oez/go-mfr/internal/service/json_parser"
)

type MockHttpGetByTNumber struct{}

func (h *MockHttpGetByTNumber) GetByTNumber(tNumber string) string {
	return getJsonFileByID("1")
}

func init() {
	json_parser.HttpGetService = &MockHttpGetByTNumber{}
}

func TestExcelWriter(t *testing.T) {
	var excelHandler = &excel_handler.SREQGeneral{}
	excelPath := filepath.Join("..", "..", "test_files", "template_mfr_export.xlsx")
	SREQGeneral, err := excelHandler.GetExcelModel("_")

	if err == nil {
		excelHandler.WriteExcel(excelPath, SREQGeneral)
	} else {
		log.Fatal(err)
	}
}

func getJsonFileByID(jsonID string) string{
	jsonFile := fmt.Sprintf("../test_files/json_%s.json", jsonID)
	filePath := filepath.Join("..", "..", "test_files", jsonFile)

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}