package storage

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/conf-api/schema"
	"github.com/incwadi-warehouse/monorepo-go/conf/manager"
)

func LoadData(schemaName, databaseId string) (*manager.Config, error) {
	return manager.LoadFromString(
		readEmbeddedFile("data/" + schemaName + ".schema.json", "{}"),
		readEmbeddedFile("data/" + schemaName + ".defaults.json", "{}"),
		readFile(os.Getenv("DATA_DIR") + schemaName + "-" + databaseId + ".json", "{}"),
	)
}

func LoadDataAndMerge(schemaName, databaseId string) (*manager.Config, error) {
	s, err := LoadData(schemaName, databaseId)
	if err != nil {
		return nil, err
	}

	data, err := s.Merge()
	if err != nil {
		return nil, err
	}
	s.Data = data

	return s, nil
}

func WriteData(schemaName, databaseId string, data interface{}) error {
    var out bytes.Buffer

	v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Indent(&out, v, "", "\t"); err != nil {
		return err
	}

	if err := os.WriteFile(os.Getenv("DATA_DIR") + schemaName + "-" + databaseId + ".json", out.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func readFile(file, defaults string) []byte {
	data, err := os.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}

func readEmbeddedFile(file, defaults string) []byte {
	data, err := schema.Fs.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}
