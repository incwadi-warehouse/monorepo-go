package storage

import (
	"os"
	"testing"
)

func TestFilesystemStorage_Save(t *testing.T) {
	file, err := os.CreateTemp("", "export-*.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	fs := &filesystemStorage{name: file.Name()}

	testData := []byte("{\"test\":\"data\"}")

	err = fs.save(testData)
	if err != nil {
		t.Errorf("Save() error = %v; want nil", err)
	}

	data, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("Failed to read data from file: %v", err)
	}

	if string(data) != string(testData) {
		t.Errorf("Saved data does not match original data. Got: %s, Want: %s", data, testData)
	}
}
