package content

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// GetIndex returns the content index as a JSON string.
func GetIndex() (string, error) {
	indexPath := filepath.Join(contentRoot, "index.yaml")

	entries, err := loadIndex(indexPath)
	if err != nil {
		return "", fmt.Errorf("error getting index: %w", err)
	}

	jsonData, err := json.Marshal(entries)
	if err != nil {
		return "", fmt.Errorf("error marshalling index to JSON: %w", err)
	}

	return string(jsonData), nil
}

// loadIndex reads the content index file and returns its content as []IndexEntry.
// If the file doesn't exist, it creates a default one.
func loadIndex(indexPath string) ([]IndexEntry, error) {
	if err := createIndexFileIfNotExists(indexPath); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("error reading index file: %w", err)
	}

	var entries []IndexEntry
	if err := yaml.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("error unmarshalling index file: %w", err)
	}

	return entries, nil
}

// createIndexFileIfNotExists creates the index file if it doesn't exist.
// If the file is created, it will include an example entry.
func createIndexFileIfNotExists(indexPath string) error {
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		exampleContent := []byte("- slug: example\n  date: 2024-09-24\n  summary: This is an example article.\n")
		if err := os.WriteFile(indexPath, exampleContent, 0644); err != nil {
			return fmt.Errorf("error creating index file: %w", err)
		}
	}

	return nil
}
