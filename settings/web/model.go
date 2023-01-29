package web

import (
	"errors"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
	"github.com/incwadi-warehouse/monorepo-go/settings/validation"
)

type Config struct {
	Value string `json:"value"`
}

var schemaName string
var databaseId string

func setSchemaName(d string) error {
    if err := storage.Exists(os.Getenv("DATA_DIR") + d +".schema.json"); err != nil {
        return err
    }

    schemaName = d

    return nil
}

func setDatabaseId(d string) error {
    if err := storage.Exists(os.Getenv("DATA_DIR") + schemaName +".schema.json"); err != nil {
        return err
    }

    if err := validation.Validate(d, "required,settingsDatabaseId"); err != nil {
        return errors.New("VALIDATION FAILED")
    }

    databaseId = d

    return nil
}

func getSchemaUrl() string {
	return os.Getenv("DATA_DIR") + schemaName +".schema.json"
}

func getDefaultsUrl() string {
	return os.Getenv("DATA_DIR") + schemaName +".defaults.json"
}

func getDatabaseUrl() string {
	return os.Getenv("DATA_DIR") + schemaName + "-"+ databaseId +".json"
}

func readFile(file, defaults string) []byte {
    data, err := os.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}
