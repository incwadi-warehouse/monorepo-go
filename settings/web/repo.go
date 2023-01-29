package web

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/conf/manager"
	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

func readEmbeddedFile(file, defaults string) []byte {
	data, err := fs.ReadFile(file)
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
