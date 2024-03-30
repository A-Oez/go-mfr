package excel_handler

import (
	"errors"
	"strings"

	"fmt"
	"log"

	excelModel "github.com/A-Oez/MFRCli/internal/model/excel_model"
	jsonParser "github.com/A-Oez/MFRCli/internal/service/json_parser"

	pReader "github.com/A-Oez/MFRCli/pkg"
	excelUtils "github.com/A-Oez/MFRCli/pkg/excel_utils"
	"github.com/xuri/excelize/v2"
)

type SREQAddress struct{}

func (sreq *SREQAddress) GetExcelModel(tNumber string) ([]excelModel.SREQAddress, error) {
	var SREQAddressArr []excelModel.SREQAddress
	serviceRequests, _ := jsonParser.ParseSREQResponse(tNumber)

	if len(serviceRequests.Value) == 0 {
		return SREQAddressArr, errors.New("-- ERROR | Keine Daten verfügbar")
	}

	splittedDescrByCustomer := strings.Split(serviceRequests.Value[0].Description, "|")
	for customerIndex := range splittedDescrByCustomer {
		splittedCustomer := strings.Split(splittedDescrByCustomer[customerIndex], ";")
		if len(splittedCustomer) == 4 {
			var SREQAddress excelModel.SREQAddress
			SREQAddress.TNummer = tNumber
			SREQAddress.Auftragsname = serviceRequests.Value[0].Name
			SREQAddress.Email = splittedCustomer[2]
			SREQAddress.Telefon = splittedCustomer[3]
			SREQAddressArr = append(SREQAddressArr, SREQAddress)
		} else {
			return nil, errors.New("-- ERROR | Description stimmt nicht mit Struktur überein")
		}
	}

	return SREQAddressArr, nil
}

func (sreq *SREQAddress) WriteExcel(filePath string, excelModel excelModel.SREQAddress) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetName := pReader.GetProperty("serviceRequestAddress")

	row := excelUtils.FindNextEmptyRow(file, sheetName)

	data := map[string]interface{}{
		"A": excelModel.Auftragsname,
		"B": excelModel.Email,
		"C": excelModel.Telefon,
	}

	for col, value := range data {
		cell := fmt.Sprintf("%s%d", col, row)
		if err := file.SetCellValue(sheetName, cell, value); err != nil {
			log.Fatal(err)
		}
	}

	if err := file.Save(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("* %s %s", pReader.GetProperty("serviceRequestAddress"), excelModel.TNummer))

}
