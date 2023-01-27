package web

import (
	"os"
)

type Config struct {
	Value string `json:"value"`
}

var dataDir = "data/"

var databaseName = ""

func setDatabaseName(d string) {
    databaseName = d
}

func getSchemaUrl() string {
	return os.Getenv("FILE_PATH") + dataDir + databaseName +".schema.json"
}

func getDefaultsUrl() string {
	return os.Getenv("FILE_PATH") + dataDir + databaseName +".defaults.json"
}

func getDatabaseUrl() string {
	return os.Getenv("FILE_PATH") + dataDir + databaseName +".json"
}

func readFile(file, defaults string) []byte {
    data, err := os.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}
