package storage

import (
	"log"
	"os"
)

func init() {
	log.SetPrefix("storage: ")
}

func Read(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Write(filename string, content []byte) error {
	if err := os.WriteFile(filename, content, 0644); err != nil {
		return err
	}

	return nil
}

func Exists(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return err
	}

	return nil
}
