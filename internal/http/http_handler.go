package http

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	httpUtils "github.com/A-Oez/go-mfr/pkg/http_utils"
	"github.com/pterm/pterm"
)

type HttpHandler struct{}

func (h *HttpHandler) GetByTNumber(tNumber string) string {
	encodedTNumber := url.QueryEscape(tNumber)
	apiUrl := fmt.Sprintf("https://portal.mobilefieldreport.com/odata/ServiceRequests?$filter=ExternalId%%20eq%%20'%s'&$expand=Appointments,Steps%%0A", encodedTNumber)

	statusCode, jsonBody := httpUtils.HttpGetRequest(apiUrl)

	if statusCode != 200 {
		err := fmt.Errorf("http: " + strconv.Itoa(statusCode) + " | t-number: " + tNumber)
		pterm.Error.Println(err)
		os.Exit(1)
	}

	return jsonBody
}
