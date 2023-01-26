package branch

import (
	"os"
)

type Config struct {
	Value string `json:"value"`
}

type BaseConfig struct {
	Schema string `json:"schema/$databaseName.schema.json"`
}

var databaseName = ""

func setDatabaseName(d string) {
    databaseName = d
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

func readFile(file, defaults string) []byte {
    data, err := os.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}
