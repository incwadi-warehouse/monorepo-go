package branch

import (
	"bytes"
	"encoding/json"

	"github.com/incwadi-warehouse/monorepo-go/conf/settings"
	"github.com/incwadi-warehouse/monorepo-go/settings/storage"
)

func loadData() (*settings.Config, error){
    return settings.LoadFromString(getSchema(), getDefaults(), getFile())
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

func writeBaseConfig() error {
    return writeData(BaseConfig{schema})
}
