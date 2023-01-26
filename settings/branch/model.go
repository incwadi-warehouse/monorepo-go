package branch

import (
	"os"
)

var schema = "schema/branch.schema.json"
var defaults = "schema/branch.defaults.json"
var database = "database/branch.json"

type Config struct {
	Value string `json:"value"`
}

type BaseConfig struct {
	Schema string `json:"$schema"`
}

func getSchemaUrl() string {
	return os.Getenv("FILE_PATH") + schema
}

func getDefaultsUrl() string {
	return os.Getenv("FILE_PATH") + defaults
}

func getDatabaseUrl() string {
	return os.Getenv("FILE_PATH") + database
}

func getSchema() []byte {
	schema, err := os.ReadFile(getSchemaUrl())
	if err != nil {
		schema = []byte("{}")
	}

	return schema
}

func getDefaults() []byte {
	defaults, err := os.ReadFile(getDefaultsUrl())
	if err != nil {
		defaults = []byte("{}")

	}

	return defaults
}


func getFile() []byte {
    file, err1 := os.ReadFile(getDatabaseUrl())
    if err1 != nil {
        file = []byte("{}")
    }

    return file
}
