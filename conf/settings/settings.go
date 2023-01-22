package settings

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Config struct {
	SchemaString   []byte
	DatabaseString []byte

	Schema *jsonschema.Schema
	Value  interface{}
}

func LoadFromString(schema, file []byte) (*Config, error) {
	c := &Config{SchemaString: schema, DatabaseString: file}

	if err := c.parse(); err != nil {
		return nil, err
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

func (c *Config) parse() error {
	s, err := jsonschema.CompileString("schema.json", string(c.SchemaString))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(c.DatabaseString, &c.Value); err != nil {
		return err
	}

	c.Schema = s

	if err := c.Validate(); err != nil {
		return errors.New("INVALID VALUES")
	}

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
