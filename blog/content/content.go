package content

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// contentRoot is the root directory for content files.
const contentRoot = "data/content"

// GetContent retrieves content from the filesystem safely.
func GetContent(path string) (string, error) {
	fullPath := filepath.Join(contentRoot, filepath.Clean("/"+path))

	if !strings.HasPrefix(fullPath, contentRoot) {
		return "", fmt.Errorf("error: path traversal attempt detected")
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("error reading content: %w", err)
	}

	return string(data), nil
}
