package json_parser

import (
	"encoding/json"
	"log"
	"strings"

	jsonModel "github.com/A-Oez/go-mfr/internal/model/json"
)

type HTTPClient interface {
	GetByTNumber(tNumber string) string
}

type SREQParser struct{
	Client HTTPClient
}

func (sreqp *SREQParser) ParseSREQResponse(tNumber string) (jsonModel.SREQ, []jsonModel.StepDataField) {
	var serviceRequests jsonModel.SREQ
	var stepDataField []jsonModel.StepDataField

	jsonBody := sreqp.Client.GetByTNumber(tNumber)

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

func parseStepData(serviceRequests jsonModel.SREQ) []jsonModel.StepDataField {
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

func relevantStepData(serviceRequests jsonModel.SREQ, index int) bool {
	return serviceRequests.Value[0].Steps[index].Name == "Sonstige Bemerkungen" ||
		serviceRequests.Value[0].Steps[index].Name == "FTTX_Montage/Einblasen NVT" ||
		serviceRequests.Value[0].Steps[index].Name == "FTTX_Montage AP"
}
