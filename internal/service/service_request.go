package service

import (
	"fmt"
	"log"

	"github.com/A-Oez/go-mfr/internal/excel"
	jsonParser "github.com/A-Oez/go-mfr/internal/json_parser"
	pReader "github.com/A-Oez/go-mfr/pkg"
	excelUtils "github.com/A-Oez/go-mfr/pkg/excel_utils"
)


func HandleServiceRequestExport(excelPath string, client jsonParser.HTTPClient) {
	tNumbers, err := excelUtils.GetTNumbers(excelPath)
	if err != nil{
		log.Fatal(err)
	}

	parser := &jsonParser.SREQParser{
		Client:  client,
	}

	excelHandler := &excel.ServiceRequest{
		JsonParser: parser,
	}

	for i := range tNumbers {
		SREQGeneral, err := excelHandler.GetExcelModel(tNumbers[i])
		if err != nil {
			fmt.Printf("* %s %s: %s\n", pReader.GetProperty("serviceRequestExport"), tNumbers[i], err.Error())
		} 

		excelHandler.WriteExcel(excelPath, SREQGeneral)
	}

}