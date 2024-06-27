package service_request_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/A-Oez/MFRCli/internal/service/excel_handler"
	json_parser "github.com/A-Oez/MFRCli/internal/service/json_parser"
)

type MockHttpGetByTNumber struct{}

func (h *MockHttpGetByTNumber) GetByTNumber(tNumber string) string {
	return getJsonFileByID("1")
}

func init() {
	json_parser.HttpGetService = &MockHttpGetByTNumber{}
}

func StartExcelWriter(t *testing.T) {
	var excelHandler = &excel_handler.SREQGeneral{}
	excelPath := "../test_files/template_mfr_export.xlsx"
	SREQGeneral, err := excelHandler.GetExcelModel("_")

	if err == nil {
		excelHandler.WriteExcel(excelPath, SREQGeneral)
	} else {
		log.Fatal(err)
	}
}

func getJsonFileByID(jsonID string) string{
	path := fmt.Sprintf("../test_files/json_%s.json", jsonID)

	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}