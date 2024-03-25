package httputils

import (
	pReader "MFRCli/pkg"
	"io"
	"log"
	"net/http"
)

func HttpGetRequest(url string) (statusCode int, body string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	req.Header.Add("Authorization", pReader.GetProperty("auth"))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return resp.StatusCode, bodyString
	}

	return resp.StatusCode, ""
}
