package web

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/conf-api/schema"
	"github.com/incwadi-warehouse/monorepo-go/conf-api/storage"
	"github.com/incwadi-warehouse/monorepo-go/conf/manager"
)

func readEmbeddedFile(file, defaults string) []byte {
	data, err := schema.Fs.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}

func readFile(file, defaults string) []byte {
	data, err := os.ReadFile(file)
	if err != nil {
		data = []byte(defaults)
	}

	return data
}

func loadData() (*manager.Config, error) {
	return manager.LoadFromString(
		readEmbeddedFile(getSchemaUrl(), "{}"),
		readEmbeddedFile(getDefaultsUrl(), "{}"),
		readFile(getDatabaseUrl(), "{}"),
	)
}

func loadDataAndMerge() (*manager.Config, error) {
	s, err := loadData()
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

func writeData(data interface{}) error {
    var out bytes.Buffer

	v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Indent(&out, v, "", "\t"); err != nil {
		return err
	}

	if err := storage.Write(getDatabaseUrl(), out.Bytes()); err != nil {
		return err
	}

	return nil
}
