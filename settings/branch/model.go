package branch

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

var schema = "schema/branch.schema.json"
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

func getDatabaseUrl() string {
	return os.Getenv("FILE_PATH") + database
}

func writeBaseConfig() error {
	d, err := json.Marshal(BaseConfig{schema})
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, d, "", "\t"); err != nil {
		return err
	}

	if err := storage.Write(getDatabaseUrl(), out.Bytes()); err != nil {
		return err
	}

	return nil
}
