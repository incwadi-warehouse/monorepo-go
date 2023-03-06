package manager

import (
	"encoding/json"
	"errors"
	"strings"

	merge "github.com/RaveNoX/go-jsonmerge"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Config struct {
	JsonSchema   []byte
	JsonDefaults []byte
	JsonData     []byte

	Schema   *jsonschema.Schema
	Defaults interface{}
	Data     interface{}
}

func LoadFromString(schema, defaults, file []byte) (*Config, error) {
	c := &Config{JsonSchema: schema, JsonDefaults: defaults, JsonData: file}

	if err := c.loadSchema(); err != nil {
		return nil, err
	}
	if err := c.loadDefaults(); err != nil {
		return nil, err
	}
	if err := c.loadData(); err != nil {
		return nil, err
	}

	if err := c.ValidateSchema(); err != nil {
		return nil, errors.New("INVALID VALUES")
	}

	return c, nil
}

func (c *Config) ValidateSchema() error {
	return c.Schema.Validate(c.Data)
}

func (c *Config) Get(key string) interface{} {
	k, l, err := c.findLastKey(key)
	if err != nil {
		return nil
	}

	return k.(map[string]interface{})[l]
}

func (c *Config) Add(key string, value interface{}) {
	k, l, err := c.findLastKey(key)
	if err != nil {
		return
	}

	k.(map[string]interface{})[l] = value
}

func (c *Config) Rm(key string) {
	k, l, err := c.findLastKey(key)
	if err != nil {
		return
	}

	delete(k.(map[string]interface{}), l)
}

func (c *Config) Merge() (interface{}, error) {
	data, info := merge.Merge(c.Defaults, c.Data)
	if len(info.Errors) != 0 {
		return nil, errors.New("ERROR MERGING DEFAULT VALUES")
	}

	return data, nil
}

func (c *Config) loadSchema() error {
	s, err := jsonschema.CompileString("schema.json", string(c.JsonSchema))
	if err != nil {
		return err
	}

	c.Schema = s

	return nil
}

func (c *Config) loadDefaults() error {
	return json.Unmarshal(c.JsonDefaults, &c.Defaults)
}

func (c *Config) loadData() error {
	return json.Unmarshal(c.JsonData, &c.Data)
}

func (c *Config) addKey(key string, value map[string]interface{}) {
	k, l, err := c.findLastKey(key)
	if err != nil {
		return
	}

	k.(map[string]interface{})[l] = value
}

func (c *Config) findLastKey(name string) (interface{}, string, error) {
	var pos string
	s := strings.Split(name, ".")
	key := c.Data

	if len(s) >= 10 {
		return nil, "", errors.New("NESTING TOO DEEP")
	}

	for k, v := range s {
		if k < len(s)-1 {
			if k == 0 {
				pos = v
			} else {
				pos += "." + v
			}

			if key.(map[string]interface{})[v] == nil {
				c.addKey(pos, map[string]interface{}{})
				return c.findLastKey(name)
			} else {
				key = key.(map[string]interface{})[v]
			}
		}
	}

	return key, s[len(s)-1], nil
}
