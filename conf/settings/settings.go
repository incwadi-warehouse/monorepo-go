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
	SchemaUrl string
	FileUrl   string

	Schema *jsonschema.Schema
	Value  interface{}
}

func Load(schema, file string) *Config {
	c := &Config{SchemaUrl: schema, FileUrl: file}

	if err := c.Read(); err != nil {
		log.Fatal(err)
	}

	return c
}

func (c *Config) Read() error {
	sch, err := jsonschema.Compile(c.SchemaUrl)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(c.FileUrl)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &c.Value); err != nil {
		return err
	}

	c.Schema = sch

    if err := c.Validate(); err != nil {
        return errors.New("INVALID VALUES")
    }

	return nil
}

func (c *Config) Write() error {
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

func (c *Config) Add(key string, value interface{}) {
    s := strings.Split(key, ".")
	c.Value.(map[string]interface{})[s[0]].(map[string]interface{})[s[1]] = value
}

func (c *Config) Rem(key string) {
    s := strings.Split(key, ".")
	delete(c.Value.(map[string]interface{})[s[0]].(map[string]interface{}), s[1])
}
