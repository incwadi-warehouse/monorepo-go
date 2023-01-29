package storage

import (
	"os"
)

func Read(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func Write(filename string, content []byte) error {
	return os.WriteFile(filename, content, 0644)
}

func Exists(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return err
	}

	return nil
}
