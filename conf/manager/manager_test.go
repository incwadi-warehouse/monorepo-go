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

	v := data.Get("app.string")

	if v != "value 1" {
		t.Fatal("Value equals not 'value'")
	}

    v2 := data.Get("app.bool")

	if v2 != true {
		t.Fatal("Value equals not 'true'")
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

	v := data.Get("app.string")

	if v != "value 1" {
		t.Fatal("Value equals not 'value'")
	}

	v2 := data.Get("app.merge")


	if v2 != true {
		t.Fatal("Value equals not 'true'")
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

	data.Add("app.string2", "string 2")

	if data.Data == nil {
		t.Fatal("Data is nil")
	}

	if data.Data.(map[string]interface{})["app"].(map[string]interface{})["string2"] != "string 2" {
		t.Fatal("Value not set")
	}

    data.Add("app.bool2", true)

	if data.Data == nil {
		t.Fatal("Data is nil")
	}

	if data.Data.(map[string]interface{})["app"].(map[string]interface{})["bool2"] != true {
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

	data.Add("app.rm", 1)

    if data.Data.(map[string]interface{})["app"].(map[string]interface{})["rm"] != 1 {
		t.Fatal("Value not added")
	}

	data.Rm("app.rm")

	if data.Data == nil {
		t.Fatal("Data is nil")
	}

	if data.Data.(map[string]interface{})["app"].(map[string]interface{})["rm"] == 1 {
		t.Fatal("Value not removed")
	}
}
