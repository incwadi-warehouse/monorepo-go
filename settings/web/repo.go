package web

import (
	"bytes"
	"encoding/json"

	"github.com/incwadi-warehouse/monorepo-go/conf/manager"
	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

func loadData() (*manager.Config, error) {
    return manager.LoadFromString(
        readFile(getSchemaUrl(), "{}"),
        readFile(getDefaultsUrl(), "{}"),
        readFile(getDatabaseUrl(), "{}"),
    )
}

func writeData(data interface{}) error {
    v, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, v, "", "\t"); err != nil {
		return err
	}

	if err := storage.Write(getDatabaseUrl(), out.Bytes()); err != nil {
		return err
	}

    return nil
}
