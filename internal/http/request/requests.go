package request

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	httpUtil "MFRCli/pkg/httpUtils"
)

func GetServiceRequestAndStepData(tNumber string) string {
	encodedTNumber := url.QueryEscape(tNumber)
	apiUrl := fmt.Sprintf("https://portal.mobilefieldreport.com/odata/ServiceRequests?$filter=ExternalId%%20eq%%20'%s'&$expand=Appointments,Steps%%0A", encodedTNumber)

	statusCode, jsonBody := httpUtil.HttpGetRequest(apiUrl)

	if statusCode != 200 {
		log.Fatal("http statuscode: " + strconv.Itoa(statusCode) + "| t-number: " + tNumber)
	}

	return jsonBody
}
