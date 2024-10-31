package excel

import (
	"errors"
	"strings"
	"time"

	"fmt"
	"log"

	excelModel "github.com/A-Oez/go-mfr/internal/model/excel"
	jsonModel "github.com/A-Oez/go-mfr/internal/model/json"
	"github.com/pterm/pterm"

	pReader "github.com/A-Oez/go-mfr/pkg"
	excelUtils "github.com/A-Oez/go-mfr/pkg/excel_utils"
	"github.com/xuri/excelize/v2"
)

type JsonParser interface {
	ParseSREQResponse(tNumber string) (jsonModel.SREQ, []jsonModel.StepDataField)
}

type ServiceRequest struct{
	JsonParser JsonParser
}

func (sreq *ServiceRequest) GetExcelModel(tNumber string) (excelModel.SREQ, error) {
	var model excelModel.SREQ
	serviceRequests, stepDataField := sreq.JsonParser.ParseSREQResponse(tNumber)

	if stepDataField == nil {
		return model, errors.New("keine checklisten hinterlegt")
	}

	model.TNummer = tNumber
	err := assignSReqDataToExcel(serviceRequests, &model)
	if err != nil {
		return model, err
	}

	//method overwrites excel columns from step data (checklists)  
	assignStepDataToExcel(stepDataField, &model)

	return model, nil
}

func assignSReqDataToExcel(serviceRequests jsonModel.SREQ, SREQGeneral *excelModel.SREQ) error {
	//date + kw
	if len(serviceRequests.Value[0].Appointments) > 0 {
		timeObj, _ := time.Parse(time.RFC3339, serviceRequests.Value[0].Appointments[0].EndDateTime)
		formattedDate := timeObj.Format("01.01.2000")
		_, week := timeObj.ISOWeek()
		SREQGeneral.Datum = formattedDate
		SREQGeneral.KW = week
	} else {
		return errors.New("kein datum hinterlegt")
	}

	//address
	splittedAddress := strings.Split(serviceRequests.Value[0].Name, "_")
	if len(splittedAddress) >= 6 {
		SREQGeneral.Straße = assignIfExists(splittedAddress, 2)
		SREQGeneral.Hausnummer = assignIfExists(splittedAddress, 3)
		SREQGeneral.Stadt = assignIfExists(splittedAddress, 4)
		SREQGeneral.Ort = assignIfExists(splittedAddress, 5)
	} else if len(splittedAddress) >= 5 {
		SREQGeneral.Straße = assignIfExists(splittedAddress, 2)
		SREQGeneral.Hausnummer = assignIfExists(splittedAddress, 3)
		SREQGeneral.Stadt = assignIfExists(splittedAddress, 4)
	} else {
		SREQGeneral.Stadt = serviceRequests.Value[0].Name
	}			

	//direct assignments
	SREQGeneral.Auftragsname = serviceRequests.Value[0].Name
	SREQGeneral.Description = serviceRequests.Value[0].Description

	return nil
}

func assignStepDataToExcel(stepDataField []jsonModel.StepDataField, SREQGeneral *excelModel.SREQ){
	for _, stepData := range stepDataField {
		switch stepData.Name {
		case "Verband & Röhrchen Farbe NVT? (Foto)":
			SREQGeneral.Rohrfarbe = stepData.Result
		case "ONT Seriennummer?":
			SREQGeneral.OntSerialNummer1 = stepData.Result
		case "1. ONT Seriennummer?":
			SREQGeneral.OntSerialNummer1 = stepData.Result
		case "2. ONT Seriennummer?":
			SREQGeneral.OntSerialNummer2 = stepData.Result
		case "3. ONT Seriennummer?":
			SREQGeneral.OntSerialNummer3 = stepData.Result
		case "4. ONT Seriennummer?":
			SREQGeneral.OntSerialNummer4 = stepData.Result
		case "1. ONT KDnr?":
			SREQGeneral.Vertragsnummer1 = stepData.Result
		case "2. ONT KDnr?":
			SREQGeneral.Vertragsnummer2 = stepData.Result
		case "3. ONT KDnr?":
			SREQGeneral.Vertragsnummer3 = stepData.Result
		case "4. ONT KDnr?":
			SREQGeneral.Vertragsnummer4 = stepData.Result
		case "Art des Microkabels?":
			SREQGeneral.Kabel = stepData.Result
		case "KVZ Nummer?":
			SREQGeneral.KVZH = stepData.Result
		case "Meterzahl  Anfang":
			SREQGeneral.KabelStart = stepData.Result
		case "Meterzahl Ende":
			SREQGeneral.KabelEnde = stepData.Result
		case "Wie viele ONTs?":
			SREQGeneral.NumberOfAssembledONTs = stepData.Result
		case "Art des verbauten AP":
			SREQGeneral.WE = strings.ReplaceAll(stepData.Result, "WE", "")
		case "LED rot oder grün?":
			SREQGeneral.ONTStatus = stepData.Result
		case "Bemerkungen?":
			if stepData.Result != "" {
				SREQGeneral.Bemerkungen += " | " + stepData.Result
			}
		}

	}
}

func (sreq *ServiceRequest) WriteExcel(filePath string, excelModel excelModel.SREQ) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sheetName := pReader.GetProperty("service_request_export")

	row := excelUtils.FindNextEmptyRow(file, sheetName)

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
		"K": excelModel.Vertragsnummer1,
		"L": excelModel.Vertragsnummer2,
		"M": excelModel.Vertragsnummer3,
		"N": excelModel.Vertragsnummer4,
		"O": excelModel.OntSerialNummer1,
		"P": excelModel.OntSerialNummer2,
		"Q": excelModel.OntSerialNummer3,
		"R": excelModel.OntSerialNummer4,
		"S": excelModel.KVZH,
		"T": excelModel.Kabel,
		"U": excelModel.KabelStart,
		"V": excelModel.KabelEnde,
		"W": excelModel.GezogenesKabel,
		"X": excelModel.AplMontageStatus,
		"Y": excelModel.Bemerkungen,
		"Z": excelModel.NumberOfAssembledONTs,
		"AA": excelModel.WE,
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

	serviceRequestExport := pReader.GetProperty("service_request_export")
	pterm.Info.Printf("* %s %s\n", serviceRequestExport, excelModel.TNummer)
}

func assignIfExists(slice []string, index int) string {
	if index < len(slice) {
		return slice[index]
	}
	return ""
}
