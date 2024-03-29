package excel_handler

import (
	"errors"
	"strings"
	"time"

	"fmt"
	"log"

	excelModel "github.com/A-Oez/MFRCli/internal/model/excel_model"
	jsonModel "github.com/A-Oez/MFRCli/internal/model/json_model"
	jsonParser "github.com/A-Oez/MFRCli/internal/service/json_parser"

	_ "github.com/A-Oez/MFRCli/internal/interfaces"
	pReader "github.com/A-Oez/MFRCli/pkg"
	excelUtils "github.com/A-Oez/MFRCli/pkg/excel_utils"
	"github.com/xuri/excelize/v2"
)

type SREQGeneral struct{}

func (sreq SREQGeneral) WriteExcel(filePath string, model interface{}) {
	if excelModel, ok := model.(excelModel.SREQGeneral); ok {
		file, err := excelize.OpenFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		sheetName := pReader.GetProperty("serviceRequestExport")

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

func (e *SREQGeneral) GetExcelModel(tNumber string) (interface{}, error) {
	var SREQGeneral excelModel.SREQGeneral
	serviceRequests, stepDataField := jsonParser.ParseSREQResponse(tNumber)

	if stepDataField == nil {
		return SREQGeneral, errors.New("-- ERROR | Keine Checklisten hinterlegt")
	}

	assignStepDataToExcel(stepDataField, &SREQGeneral)
	err := assignSReqDataToExcel(serviceRequests, &SREQGeneral)

	if err != nil {
		return SREQGeneral, errors.New("-- ERROR | " + err.Error())
	}
	SREQGeneral.TNummer = tNumber

	return SREQGeneral, nil
}

func assignSReqDataToExcel(serviceRequests jsonModel.ServiceRequestResponse, SREQGeneral *excelModel.SREQGeneral) error {
	//date + kw
	if len(serviceRequests.Value[0].Appointments) > 0 {
		timeObj, _ := time.Parse(time.RFC3339, serviceRequests.Value[0].Appointments[0].EndDateTime)
		formattedDate := timeObj.Format("02.01.2006")
		_, week := timeObj.ISOWeek()
		SREQGeneral.Datum = formattedDate
		SREQGeneral.KW = week
	} else {
		return errors.New("Kein Termin hinterlegt")
	}

	//address
	spllittedAddress := strings.Split(serviceRequests.Value[0].Name, "_")
	if len(spllittedAddress) == 6 {
		SREQGeneral.Straße = spllittedAddress[2]
		SREQGeneral.Hausnummer = spllittedAddress[3]
		SREQGeneral.Stadt = spllittedAddress[4]
		SREQGeneral.Ort = spllittedAddress[5]
	} else if len(spllittedAddress) > 4 && len(spllittedAddress) < 6 {
		SREQGeneral.Straße = spllittedAddress[2]
		SREQGeneral.Hausnummer = spllittedAddress[3]
		SREQGeneral.Stadt = spllittedAddress[4]
	} else {
		SREQGeneral.Stadt = serviceRequests.Value[0].Name
	}

	//customer
	splittedDescrByCustomer := strings.Split(serviceRequests.Value[0].Description, "|")
	for customerIndex := range splittedDescrByCustomer {
		splittedCustomer := strings.Split(splittedDescrByCustomer[customerIndex], ";")
		if len(splittedCustomer) != 4 {
			SREQGeneral.Vertragsnehmer = serviceRequests.Value[0].Description
			break
		} else {
			SREQGeneral.Vertragsnummer += splittedCustomer[0] + "\r\n"
			SREQGeneral.Vertragsnehmer += splittedCustomer[1] + "\r\n"
		}
	}

	//direct assignments
	SREQGeneral.Auftragsname = serviceRequests.Value[0].Name
	SREQGeneral.Description = serviceRequests.Value[0].Description

	return nil
}

func assignStepDataToExcel(stepDataField []jsonModel.StepDataField, SREQGeneral *excelModel.SREQGeneral) {

	for _, stepData := range stepDataField {
		switch stepData.Name {
		case "Verband & Röhrchen Farbe NVT? (Foto)":
			SREQGeneral.Rohrfarbe = stepData.Result
		case "ONT Seriennummer?":
			SREQGeneral.OntSerialNummer = stepData.Result
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
