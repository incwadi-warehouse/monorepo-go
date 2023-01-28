package web

import (
	"os"

	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

type Config struct {
	Value string `json:"value"`
}

var databaseName = ""

func setDatabaseName(d string) error {
    if err := storage.Exists(os.Getenv("DATA_DIR") + d +".schema.json"); err != nil {
        return err
    }

    databaseName = d

    return nil
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
