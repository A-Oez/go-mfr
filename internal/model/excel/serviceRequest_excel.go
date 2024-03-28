package model_excel

import (
	"fmt"
	"log"

	_ "github.com/A-Oez/MFRCli/internal/interfaces"
	pReader "github.com/A-Oez/MFRCli/pkg"
	excelutils "github.com/A-Oez/MFRCli/pkg/excelutils"
	"github.com/xuri/excelize/v2"
)

type ServiceRequestsExcel struct {
	TNummer               string
	KW                    int
	Datum                 string
	ONTStatus             string
	Stadt                 string
	Ort                   string
	Straße                string
	Hausnummer            string
	Vertragsnehmer        string
	Rohrfarbe             string
	Vertragsnummer        string
	OntSerialNummer       string
	KVZH                  string
	Kabel                 string
	KabelStart            string
	KabelEnde             string
	GezogenesKabel        string
	AplMontageStatus      string
	Bemerkungen           string
	NumberOfAssembledONTs string
	WE                    string
	Description           string
	Auftragsname          string
}

func (sw ServiceRequestsExcel) WriteExcel(filePath string, model interface{}) {
	if excelModel, ok := model.(ServiceRequestsExcel); ok {
		file, err := excelize.OpenFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		sheetName := pReader.GetProperty("serviceRequestExport")

		row := excelutils.FindNextEmptyRow(file, sheetName)

		data := map[string]interface{}{
			"A": excelModel.TNummer,
			"B": excelModel.KW,
			"C": excelModel.Datum,
			"D": excelModel.ONTStatus,
			"E": excelModel.Stadt,
			"F": excelModel.Ort,
			"G": excelModel.Straße,
			"H": excelModel.Hausnummer,
			"I": excelModel.Vertragsnehmer,
			"J": excelModel.Rohrfarbe,
			"K": excelModel.Vertragsnummer,
			"L": excelModel.OntSerialNummer,
			"M": excelModel.KVZH,
			"N": excelModel.Kabel,
			"O": excelModel.KabelStart,
			"P": excelModel.KabelEnde,
			"Q": excelModel.GezogenesKabel,
			"R": excelModel.AplMontageStatus,
			"S": excelModel.Bemerkungen,
			"T": excelModel.NumberOfAssembledONTs,
			"U": excelModel.WE,
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

		fmt.Println(fmt.Sprintf("* %s %s", pReader.GetProperty("serviceRequestExport"), excelModel.TNummer))
	}
}
