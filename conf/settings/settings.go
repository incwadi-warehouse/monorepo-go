package settings

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Config struct {
	SchemaUrl   string
	DatabaseUrl string

	SchemaString   []byte
	DatabaseString []byte

	Schema *jsonschema.Schema
	Value  interface{}
}

func LoadFromUrl(schema, file string) (*Config, error) {
	s, err := os.ReadFile(schema)
	if err != nil {
		return nil, err
	}

	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	c := &Config{SchemaUrl: schema, DatabaseUrl: file, SchemaString: s, DatabaseString: f}

	if err := c.parse(); err != nil {
		return nil, err
	}

	return c, nil
}

func LoadFromString(schema, file []byte) (*Config, error) {
	c := &Config{SchemaString: schema, DatabaseString: file}

	if err := c.parse(); err != nil {
		return nil, err
	}

	return c, nil
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

func (c *Config) Write() error {
	if c.DatabaseUrl == "" {
		return errors.New("NO FILE URL GIVEN")
	}

	if err := c.Validate(); err != nil {
		return errors.New("INVALID VALUES")
	}

	d, err := json.Marshal(c.Value)
	if err != nil {
		return err
	}

	var out bytes.Buffer
	if err := json.Indent(&out, d, "", "\t"); err != nil {
		return err
	}

	if err := os.WriteFile(c.DatabaseUrl, out.Bytes(), 0644); err != nil {
		return err
	}

	return nil
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
