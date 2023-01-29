package manager

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

	if err := data.ValidateSchema(); err != nil {
		t.Fatal(err)
	}

	if data.Data == nil {
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

	v := data.Get("app.key1")

	if v != "value 1" {
		t.Fatal("Value equals not 'value'")
	}
}

func TestMerge(t *testing.T) {
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

    b, err := data.Merge()
    if err != nil {
        t.Fatal("Merge failed")
    }
    data.Data = b

	v := data.Get("app.key1")

	if v != "value 1" {
		t.Fatal("Value equals not 'value'")
	}

	v2 := data.Get("app.key2")

	if v2 != "default 2" {
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

	data.Add("app.key3", "value 3")

	if data.Data == nil {
		t.Fatal("Data is nil")
	}

	if data.Data.(map[string]interface{})["app"].(map[string]interface{})["key3"] != "value 3" {
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

	data.Add("app.key4", "value 4")

    if data.Data.(map[string]interface{})["app"].(map[string]interface{})["key4"] != "value 4" {
		t.Fatal("Value not added")
	}

	data.Rm("app.key4")

	if data.Data == nil {
		t.Fatal("Data is nil")
	}

	if data.Data.(map[string]interface{})["app"].(map[string]interface{})["key4"] == "value 4" {
		t.Fatal("Value not removed")
	}
}
