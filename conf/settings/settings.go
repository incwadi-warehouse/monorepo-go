package settings

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func init() {
	log.SetPrefix("settings: ")
}

type Config struct {
	FileUrl string

	SchemaString []byte
	FileString   []byte

	Schema *jsonschema.Schema
	Value  interface{}
}

func LoadFromUrl(schema, file string) (*Config, error) {
	s, err := os.ReadFile(schema)
	if err != nil {
		return nil, err
	}

	v, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	c := &Config{FileUrl: file, SchemaString: s, FileString: v}

	if err := c.parse(); err != nil {
		return nil, err
	}

	return c, nil
}

func LoadFromString(schema, file []byte) (*Config, error) {
	c := &Config{SchemaString: schema, FileString: file}

	if err := c.parse(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) parse() error {
	sch, err := jsonschema.CompileString("schema.json", string(c.SchemaString))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(c.FileString, &c.Value); err != nil {
		return err
	}

	c.Schema = sch

	if err := c.Validate(); err != nil {
		return errors.New("INVALID VALUES")
	}

	return nil
}

func (c *Config) Write() error {
	if c.FileUrl == "" {
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

	if err := os.WriteFile(c.FileUrl, out.Bytes(), 0644); err != nil {
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
	v, l := c.findLastKey(key)

	return v.(map[string]interface{})[l]
}

func (c *Config) Add(key string, value interface{}) {
	v, l := c.findLastKey(key)

	v.(map[string]interface{})[l] = value
}

func (c *Config) Rm(key string) {
	v, l := c.findLastKey(key)

	delete(v.(map[string]interface{}), l)
}

func (c *Config) findLastKey(name string) (interface{}, string) {
	s := strings.Split(name, ".")
	v := c.Value

	for key, value := range s {
		if key < len(s)-1 {
			v = v.(map[string]interface{})[value]
		}
	}

	return v, s[len(s)-1]
}
