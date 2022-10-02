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
	log.SetFlags(0)
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
	s := strings.Split(key, ".")

	return c.Value.(map[string]interface{})[s[0]].(map[string]interface{})[s[1]]
}

func (c *Config) Add(key string, value interface{}) {
	s := strings.Split(key, ".")
	c.Value.(map[string]interface{})[s[0]].(map[string]interface{})[s[1]] = value
}

func (c *Config) Rm(key string) {
	s := strings.Split(key, ".")
	delete(c.Value.(map[string]interface{})[s[0]].(map[string]interface{}), s[1])
}
