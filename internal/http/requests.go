package http

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	_ "github.com/A-Oez/MFRCli/internal/interfaces"
	httpUtils "github.com/A-Oez/MFRCli/pkg/http_utils"
)

type SREQHttpHandler struct{}

func (h *SREQHttpHandler) GetByTNumber(tNumber string) string {
	encodedTNumber := url.QueryEscape(tNumber)
	apiUrl := fmt.Sprintf("https://portal.mobilefieldreport.com/odata/ServiceRequests?$filter=ExternalId%%20eq%%20'%s'&$expand=Appointments,Steps%%0A", encodedTNumber)

	statusCode, jsonBody := httpUtils.HttpGetRequest(apiUrl)

	if statusCode != 200 {
		log.Fatal("http statuscode: " + strconv.Itoa(statusCode) + "| t-number: " + tNumber)
	}

	return jsonBody
}
