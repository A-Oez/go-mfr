package service_request_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/A-Oez/go-mfr/internal/service"
	"github.com/A-Oez/go-mfr/pkg"
)

type MockHttpGetByTNumber struct{}

func (h *MockHttpGetByTNumber) GetByTNumber(tNumber string) string {
	return getJsonFileByID("1")
}

func TestExcelWriter(t *testing.T) {
	excelPath := pkg.GetProperty("excel_path")
	err := service.HandleServiceRequestExport(excelPath, &MockHttpGetByTNumber{})
	if err != nil{
		log.Fatal(err)
	} 
}

func getJsonFileByID(jsonID string) string{
	jsonFile := fmt.Sprintf("/test_files/json_%s.json", jsonID)
	filePath := filepath.Join("..", "..", "test_files", jsonFile)

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}