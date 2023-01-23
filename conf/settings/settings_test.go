package settings

import (
	"os"
	"testing"
)

func TestLoadFromString(t *testing.T) {
	s, err := os.ReadFile("./example.schema.json")
	if err != nil {
		t.Fatal(err)
	}

	d, err := os.ReadFile("./example.defaults.json")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.ReadFile("./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data, err := LoadFromString(s, d, f)
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
	s, err := os.ReadFile("./example.schema.json")
	if err != nil {
		t.Fatal(err)
	}

	d, err := os.ReadFile("./example.defaults.json")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.ReadFile("./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data, err := LoadFromString(s, d, f)
	if err != nil {
		t.Fatal(err)
	}

	v := data.Get("app.key")

	if v != "value" {
		t.Fatal("Value equals not 'value'")
	}

	v2 := data.Get("app.key2")

	if v2 != "value" {
		t.Fatal("Value equals not 'value'")
	}
}

func TestAdd(t *testing.T) {
	s, err := os.ReadFile("./example.schema.json")
	if err != nil {
		t.Fatal(err)
	}

	d, err := os.ReadFile("./example.defaults.json")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.ReadFile("./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data, err := LoadFromString(s, d, f)
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
	s, err := os.ReadFile("./example.schema.json")
	if err != nil {
		t.Fatal(err)
	}

	d, err := os.ReadFile("./example.defaults.json")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.ReadFile("./example.json")
	if err != nil {
		t.Fatal(err)
	}

	data, err := LoadFromString(s, d, f)
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
