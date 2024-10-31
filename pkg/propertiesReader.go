package pkg

import (
	"encoding/json"
	"log"
	"os"
)

func GetProperty(key string) string {
	data, err := os.ReadFile(os.Getenv("GO_MFR_PATH"))
	if err != nil {
		log.Fatalf("properties datei konnte nicht geöffnet werden, GO_MFR_PATH systemvariable prüfen\n error_msg: %v", err)
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	value, _ := jsonData[key].(string)

	return value
}
