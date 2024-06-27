package service_request_test

import (
	"testing"

	_ "github.com/A-Oez/MFRCli/internal/interfaces"
	excelHandler "github.com/A-Oez/MFRCli/internal/service/excel_handler"
	json_parser "github.com/A-Oez/MFRCli/internal/service/json_parser"
)

type DescrMockHttpGetByTNumber struct{}

func (h *DescrMockHttpGetByTNumber) GetByTNumber(tNumber string) string {
	descriptionPlaceholder := "Kundennummer:1234;NAME;EMAIL;TELEFON"

	jsonString := `{
    "odata.metadata": "https://portal.mobilefieldreport.com",
		"value": [
			{
				"Id": "123456",
				"Name": "Gf_TNG_TEST",
				"ExternalId": "T-1234",
				"InvoiceId": null,
				"ClosedAt": "2022-03-24T17:30:56Z",
				"ReleasedAt": "2022-03-22T17:20:20Z",
				"WorkDoneAt": null,
				"TargetTimeInMinutes": "60",
				"DateModified": "2024-03-28T17:31:17",
				"DateOfCreation": "2024-03-26T23:13:14",
				"DueDateRangeStart": null,
				"DueDateRangeEnd": null,
				"PortalLink": "https://portal.mobilefieldreport.com",
				"CostCenterId": "0",
				"Description": "` + descriptionPlaceholder + `",
				"State": "Closed",
				"CustomValues": [],
				"CurrentOwnerId": "12345",
				"CustomerId": "12345",
				"ParentServiceRequestId": "0",
				"Location": null,
				"Version": 12,
				"IsTemplate": false,
				"IsTemplateMobile": false,
				"CreateFromServiceRequestTemplateId": "0",
				"Type": "IsServiceRequest"
			}
		]
	}`

	return jsonString
}

func init() {
	json_parser.HttpGetService = &DescrMockHttpGetByTNumber{}
}

func TestGetExcelModel(t *testing.T) {
	var excelParser = &excelHandler.SREQAddress{}
	_, err := excelParser.GetExcelModel("test")

	//TODO: extend test with excelwriter & create gitignore with ignored folder test_files

	if err != nil {
		t.Fatalf("GetExcelModel returned an error: %v", err)
	}
}
