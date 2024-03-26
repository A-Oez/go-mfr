package parser

import (
	"MFRCli/internal/http/request"
	"MFRCli/internal/model"
	"encoding/json"
	"log"
	"strings"
	"time"
)

func GetExcelModel(tNumber string) model.ServiceRequestsExcel {
	var serviceRequestsExcel model.ServiceRequestsExcel
	serviceRequests, stepDataField := parseResponse(tNumber)
	serviceRequestsExcel.TNummer = tNumber

	if stepDataField == nil {
		return serviceRequestsExcel
	}

	assignStepDataToExcel(stepDataField, &serviceRequestsExcel)
	assignSReqDataToExcel(serviceRequests, &serviceRequestsExcel)
	return serviceRequestsExcel
}

func parseResponse(tNumber string) (model.ServiceRequests, []model.StepDataField) {
	var stepData []model.StepDataField
	var serviceRequests model.ServiceRequests

	jsonBody := request.GetServiceRequestAndStepData(tNumber)

	serviceRequests = parseServiceRequest(jsonBody)

	if len(serviceRequests.Value) > 0 {
		stepData = parseStepData(serviceRequests)
	}

	return serviceRequests, stepData
}

func parseServiceRequest(jsonString string) model.ServiceRequests {
	var serviceRequests model.ServiceRequests
	err := json.NewDecoder(strings.NewReader(jsonString)).Decode(&serviceRequests)
	if err != nil {
		log.Fatal(err)
	}

	return serviceRequests
}

func parseStepData(serviceRequests model.ServiceRequests) []model.StepDataField {
	var stepDataFieldArr []model.StepDataField
	stepArr := serviceRequests.Value[0].Steps

	for stepIndex := range stepArr {
		if !relevantStepData(serviceRequests) {
			jsonBytes := []byte(stepArr[stepIndex].Data)

			var stepData model.StepData
			err := json.Unmarshal(jsonBytes, &stepData)
			if err != nil {

				log.Fatal(err)
			}

			stepDataFieldArr = append(stepDataFieldArr, stepData.Fields...)
		}
	}

	return stepDataFieldArr
}

func relevantStepData(serviceRequests model.ServiceRequests) bool {
	return serviceRequests.Value[0].Name == "Sonstige Bemerkungen" ||
		serviceRequests.Value[0].Name == "FTTX_Montage/Einblasen NVT" ||
		serviceRequests.Value[0].Name == "FTTX_Montage AP"
}

func assignSReqDataToExcel(serviceRequests model.ServiceRequests, serviceRequestsExcel *model.ServiceRequestsExcel) {

	//date + kw
	if len(serviceRequests.Value[0].Appointments) > 0 {
		timeObj, _ := time.Parse(time.RFC3339, serviceRequests.Value[0].Appointments[0].EndDateTime)
		formattedDate := timeObj.Format("02.01.2006")
		_, week := timeObj.ISOWeek()
		serviceRequestsExcel.Datum = formattedDate
		serviceRequestsExcel.KW = week
	}

	//address
	spllittedAddress := strings.Split(serviceRequests.Value[0].Name, "_")
	if len(spllittedAddress) > 4 {
		serviceRequestsExcel.Straße = spllittedAddress[2]
		serviceRequestsExcel.Hausnummer = spllittedAddress[3]
		serviceRequestsExcel.Stadt = spllittedAddress[4]
	}

	//customer
	splittedDescrByCustomer := strings.Split(serviceRequests.Value[0].Description, "|")
	for customerIndex := range splittedDescrByCustomer {
		splittedCustomer := strings.Split(splittedDescrByCustomer[customerIndex], ";")
		if len(splittedCustomer) != 3 {
			serviceRequestsExcel.Vertragsnehmer = serviceRequests.Value[0].Description
			break
		} else {
			serviceRequestsExcel.Vertragsnummer += splittedCustomer[0] + "\r\n"
			serviceRequestsExcel.Vertragsnehmer += splittedCustomer[1] + "\r\n"
		}
	}
}

func assignStepDataToExcel(stepDataField []model.StepDataField, serviceRequestsExcel *model.ServiceRequestsExcel) {

	for _, stepData := range stepDataField {
		switch stepData.Name {
		case "Verband & Röhrchen Farbe NVT? (Foto)":
			serviceRequestsExcel.Rohrfarbe = stepData.Result
		case "ONT Seriennummer?":
			serviceRequestsExcel.OntSerialNummer = stepData.Result
		case "Art des Microkabels?":
			serviceRequestsExcel.Kabel = stepData.Result
		case "KVZ Nummer?":
			serviceRequestsExcel.KVZH = stepData.Result
		case "Meterzahl  Anfang":
			serviceRequestsExcel.KabelStart = stepData.Result
		case "Meterzahl Ende":
			serviceRequestsExcel.KabelEnde = stepData.Result
		case "Wie viele ONTs?":
			serviceRequestsExcel.NumberOfAssembledONTs = stepData.Result
		case "Art des verbauten AP":
			serviceRequestsExcel.WE = strings.ReplaceAll(stepData.Result, "WE", "")
		case "LED rot oder grün?":
			serviceRequestsExcel.ONTStatus = stepData.Result
		case "Bemerkungen?":
			serviceRequestsExcel.Bemerkungen += " | " + stepData.Result
		}

	}
}
