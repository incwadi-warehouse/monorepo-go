package web

import (
	"os"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Config struct {
	Value interface{} `json:"value"`
}

var schemaName string

var databaseId string

func setSchemaName(d string) {
	schemaName = d
}

func setDatabaseId(d string) {
	databaseId = d
}

func getSchemaUrl() string {
	return "data/" + schemaName + ".schema.json"
}

func getDefaultsUrl() string {
	return "data/" + schemaName + ".defaults.json"
}

func getDatabaseUrl() string {
	return os.Getenv("DATA_DIR") + schemaName + "-" + databaseId + ".json"
}
