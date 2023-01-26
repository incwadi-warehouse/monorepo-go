package branch

import (
	"os"
)

var databaseName = ""

type Config struct {
	Value string `json:"value"`
}

type BaseConfig struct {
	Schema string `json:"schema/$databaseName.schema.json"`
}

func setDatabaseName(name string) {
    databaseName = name
}

func getSchemaUrl() string {
	return os.Getenv("FILE_PATH") + "schema/"+ databaseName +".schema.json"
}

func getDefaultsUrl() string {
	return os.Getenv("FILE_PATH") + "schema/"+ databaseName +".defaults.json"
}

func getDatabaseUrl() string {
	return os.Getenv("FILE_PATH") + "database/"+ databaseName +".json"
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
