package web

import (
	"os"
)

type Config struct {
	Value string `json:"value"`
}

var databaseName = ""

func setDatabaseName(d string) {
    databaseName = d
}

func getSchemaUrl() string {
	return os.Getenv("DATA_DIR") + databaseName +".schema.json"
}

func getDefaultsUrl() string {
	return os.Getenv("DATA_DIR") + databaseName +".defaults.json"
}

func getDatabaseUrl() string {
	return os.Getenv("DATA_DIR") + databaseName +".json"
}

func readFile(file, defaults string) []byte {
    data, err := os.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}
