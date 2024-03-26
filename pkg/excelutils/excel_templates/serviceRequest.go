package excelutils_templates

import (
	"MFRCli/internal/model"
	"fmt"
	"log"

	pReader "MFRCli/pkg"

	"github.com/xuri/excelize/v2"

	excelutils "MFRCli/pkg/excelutils"
)

func WriteToExcel(filePath string, serviceRequestsExcel model.ServiceRequestsExcel) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetName := pReader.GetProperty("serviceRequestExport")

	row := excelutils.FindNextEmptyRow(file, sheetName)

	data := map[string]interface{}{
		"A": serviceRequestsExcel.TNummer,
		"B": serviceRequestsExcel.KW,
		"C": serviceRequestsExcel.Datum,
		"D": serviceRequestsExcel.ONTStatus,
		"E": serviceRequestsExcel.Stadt,
		"F": serviceRequestsExcel.Ort,
		"G": serviceRequestsExcel.Straße,
		"H": serviceRequestsExcel.Hausnummer,
		"I": serviceRequestsExcel.Vertragsnehmer,
		"J": serviceRequestsExcel.Rohrfarbe,
		"K": serviceRequestsExcel.Vertragsnummer,
		"L": serviceRequestsExcel.OntSerialNummer,
		"M": serviceRequestsExcel.KVZH,
		"N": serviceRequestsExcel.Kabel,
		"O": serviceRequestsExcel.KabelStart,
		"P": serviceRequestsExcel.KabelEnde,
		"Q": serviceRequestsExcel.GezogenesKabel,
		"R": serviceRequestsExcel.AplMontageStatus,
		"S": serviceRequestsExcel.Bemerkungen,
		"T": serviceRequestsExcel.NumberOfAssembledONTs,
		"U": serviceRequestsExcel.WE,
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

	fmt.Println(fmt.Sprintf("* EXPORT %s", serviceRequestsExcel.TNummer))
}

func WriteToAddressExcel(filePath string, serviceRequestsAddressExcel model.ServiceRequestsAddressExcel, tNumber string) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetName := "ADDRESS"

	row := excelutils.FindNextEmptyRow(file, sheetName)

	data := map[string]interface{}{
		"A": serviceRequestsAddressExcel.Auftragsname,
		"B": serviceRequestsAddressExcel.Email,
		"C": serviceRequestsAddressExcel.Telefon,
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

	fmt.Println(fmt.Sprintf("* ADDRESS %s", tNumber))
}
