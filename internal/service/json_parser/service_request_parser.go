package json_parser

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/A-Oez/MFRCli/internal/http"
	"github.com/A-Oez/MFRCli/internal/interfaces"
	jsonModel "github.com/A-Oez/MFRCli/internal/model/json_model"
)

var HttpGetService interfaces.HttpGetByTNumber

func init() {
	HttpGetService = &http.SREQHttpHandler{}
}

func ParseSREQResponse(tNumber string) (jsonModel.ServiceRequestResponse, []jsonModel.StepDataField) {
	var serviceRequests jsonModel.ServiceRequestResponse
	var stepDataField []jsonModel.StepDataField

	jsonBody := HttpGetService.GetByTNumber(tNumber)

	err := json.NewDecoder(strings.NewReader(jsonBody)).Decode(&serviceRequests)
	if err != nil {
		log.Fatal(err)
	}

	if len(serviceRequests.Value) > 0 {
		stepDataField = parseStepData(serviceRequests)
		return serviceRequests, stepDataField
	}

	return serviceRequests, nil
}

func parseStepData(serviceRequests jsonModel.ServiceRequestResponse) []jsonModel.StepDataField {
	var stepDataFieldArr []jsonModel.StepDataField
	stepArr := serviceRequests.Value[0].Steps

	for stepIndex := range stepArr {
		if relevantStepData(serviceRequests, stepIndex) {
			jsonBytes := []byte(stepArr[stepIndex].Data)

			var stepData jsonModel.StepData
			err := json.Unmarshal(jsonBytes, &stepData)
			if err != nil {

				log.Fatal(err)
			}

			stepDataFieldArr = append(stepDataFieldArr, stepData.Fields...)
		}
	}

	return stepDataFieldArr
}

func relevantStepData(serviceRequests jsonModel.ServiceRequestResponse, index int) bool {
	return serviceRequests.Value[0].Steps[index].Name == "Sonstige Bemerkungen" ||
		serviceRequests.Value[0].Steps[index].Name == "FTTX_Montage/Einblasen NVT" ||
		serviceRequests.Value[0].Steps[index].Name == "FTTX_Montage AP"
}
