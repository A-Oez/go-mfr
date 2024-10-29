package http

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	httpUtils "github.com/A-Oez/go-mfr/pkg/http_utils"
)

type HttpHandler struct{}

func (h *HttpHandler) GetByTNumber(tNumber string) string {
	encodedTNumber := url.QueryEscape(tNumber)
	apiUrl := fmt.Sprintf("https://portal.mobilefieldreport.com/odata/ServiceRequests?$filter=ExternalId%%20eq%%20'%s'&$expand=Appointments,Steps%%0A", encodedTNumber)

	statusCode, jsonBody := httpUtils.HttpGetRequest(apiUrl)

	if statusCode != 200 {
		log.Fatal("http statuscode: " + strconv.Itoa(statusCode) + "| t-number: " + tNumber)
	}

	return jsonBody
}
