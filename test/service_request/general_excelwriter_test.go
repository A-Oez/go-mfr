package service_request_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/A-Oez/go-mfr/internal/service"
)

type MockHttpGetByTNumber struct{}

func (h *MockHttpGetByTNumber) GetByTNumber(tNumber string) string {
	return getJsonFileByID("1")
}

func TestExcelWriter(t *testing.T) {
	excelPath := filepath.Join("..", "..", "test_files", "template_mfr_export.xlsx")
	service.HandleServiceRequestExport(excelPath, &MockHttpGetByTNumber{})
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