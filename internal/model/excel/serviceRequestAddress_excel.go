package model_excel

import (
	"fmt"
	"log"

	_ "github.com/A-Oez/MFRCli/internal/interfaces"
	pReader "github.com/A-Oez/MFRCli/pkg"
	excelutils "github.com/A-Oez/MFRCli/pkg/excelutils"
	"github.com/xuri/excelize/v2"
)

type ServiceRequestsAddressExcel struct {
	TNummer      string
	Auftragsname string
	Email        string
	Telefon      string
}

func (sw *ServiceRequestsAddressExcel) WriteExcel(filePath string, model interface{}) {
	if excelModel, ok := model.(ServiceRequestsAddressExcel); ok {
		file, err := excelize.OpenFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		sheetName := pReader.GetProperty("serviceRequestAddress")

		row := excelutils.FindNextEmptyRow(file, sheetName)

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
}
