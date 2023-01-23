package settings

import (
	"encoding/json"
	"errors"
	"strings"

	merge "github.com/RaveNoX/go-jsonmerge"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Config struct {
	SchemaString           []byte
	DatabaseString         []byte
	DatabaseDefaultsString []byte

	Schema   *jsonschema.Schema
	Value    interface{}
	Defaults interface{}
}

func LoadFromString(schema, defaults, file []byte) (*Config, error) {
	c := &Config{SchemaString: schema, DatabaseDefaultsString: defaults, DatabaseString: file}

	if err := c.loadSchema(); err != nil {
		return nil, err
	}
	if err := c.loadDefaults(); err != nil {
		return nil, err
	}
	if err := c.loadValue(); err != nil {
		return nil, err
	}

	if err := c.merge(); err != nil {
		return nil, err
	}

	if err := c.Validate(); err != nil {
		return nil, errors.New("INVALID VALUES")
	}

	return c, nil
}

func (c *Config) Validate() error {
	if err := c.Schema.Validate(c.Value); err != nil {
		return err
	}

	return nil
}

func (c *Config) Get(key string) interface{} {
	k, l := c.findLastKey(key)

	return k.(map[string]interface{})[l]
}

func (c *Config) Add(key string, value interface{}) {
	k, l := c.findLastKey(key)

	k.(map[string]interface{})[l] = value
}

func (c *Config) Rm(key string) {
	k, l := c.findLastKey(key)

	delete(k.(map[string]interface{}), l)
}

func (c *Config) loadSchema() error {
	s, err := jsonschema.CompileString("schema.json", string(c.SchemaString))
	if err != nil {
		return err
	}

	c.Schema = s

	return nil
}

func (c *Config) loadDefaults() error {
	if err := json.Unmarshal(c.DatabaseDefaultsString, &c.Defaults); err != nil {
		return err
	}

	return nil
}

func (c *Config) loadValue() error {
	if err := json.Unmarshal(c.DatabaseString, &c.Value); err != nil {
		return err
	}

	return nil
}

func (c *Config) merge() error {
	data, info := merge.Merge(c.Defaults, c.Value)
	if len(info.Errors) != 0 {
		return errors.New("ERROR MERGING DEFAULT VALUES")
	}

	c.Value = data

	return nil
}

func (c *Config) findLastKey(name string) (interface{}, string) {
	s := strings.Split(name, ".")
	key := c.Value

	for k, v := range s {
		if k < len(s)-1 {
			key = key.(map[string]interface{})[v]
		}
	}

	return key, s[len(s)-1]
}
