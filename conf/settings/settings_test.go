package settings

import (
	"os"
	"testing"
)

func TestLoadFromUrl(t *testing.T) {
	data, err := LoadFromUrl("./example.schema.json", "./example.json")
	if err != nil {
		t.Fatal(err)
	}

	if err := data.Validate(); err != nil {
		t.Fatal(err)
	}

	if data.Value == nil {
		t.Fatal("Data is nil")
	}
}

func TestLoadFromString(t *testing.T) {
	s, err := os.ReadFile("./example.schema.json")
	if err != nil {
		t.Fatal(err)
	}

	v, err := os.ReadFile("./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data, err := LoadFromString(s, v)
	if err != nil {
		t.Fatal(err)
	}

	if err := data.Validate(); err != nil {
		t.Fatal(err)
	}

	if data.Value == nil {
		t.Fatal("Data is nil")
	}
}

func TestGet(t *testing.T) {
	data, err := LoadFromUrl("./example.schema.json", "./example.json")
	if err != nil {
		t.Fatal(err)
	}

	v := data.Get("app.key")

	if v != "value" {
		t.Fatal("Value equals not 'value'")
	}
}

func TestAdd(t *testing.T) {
	data, err := LoadFromUrl("./example.schema.json", "./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data.Add("app.key", '1')

	if data.Value == nil {
		t.Fatal("Data is nil")
	}

	if data.Value.(map[string]interface{})["app"].(map[string]interface{})["key"] != '1' {
		t.Fatal("Value not set")
	}
}

func TestRm(t *testing.T) {
	data, err := LoadFromUrl("./example.schema.json", "./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data.Add("app.key", '1')
	data.Rm("app.key")

	if data.Value == nil {
		t.Fatal("Data is nil")
	}

	if data.Value.(map[string]interface{})["app"].(map[string]interface{})["key"] == '1' {
		t.Fatal("Value not removed")
	}
}
