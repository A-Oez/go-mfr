package pkg

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func GetProperty(key string) string {
	absoluteFilePath := filepath.Join("config", "properties.json")

	data, err := os.ReadFile(absoluteFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	value, _ := jsonData[key].(string)

	return value
}
