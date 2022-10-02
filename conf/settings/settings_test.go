package settings

import (
	"testing"
)

func TestValidate(t *testing.T) {
	data := Load("./example.schema.json", "./example.json")
	err := data.Validate()

	if err != nil {
		t.Fatal("Validation not possible")
	}

	if data.Value == nil {
		t.Fatal("Data is nil")
	}
}

func TestAdd(t *testing.T) {
	data := Load("./example.schema.json", "./example.json")
	data.Add("app.key", '1')

	if data.Value == nil {
		t.Fatal("Data is nil")
	}

    if data.Value.(map[string]interface{})["app"].(map[string]interface{})["key"] != '1' {
        t.Fatal("Value not set")
    }
}

func TestRem(t *testing.T) {
	data := Load("./example.schema.json", "./example.json")
    data.Add("app.key", '1')
	data.Rem("app.key")

	if data.Value == nil {
		t.Fatal("Data is nil")
	}

    if data.Value.(map[string]interface{})["app"].(map[string]interface{})["key"] == '1' {
        t.Fatal("Value not removed")
    }
}
