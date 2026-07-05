package testutils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadJSON loads a JSON file from the specified path and unmarshals it into the provided interface.
// It assumes the path is relative to the calling test's directory unless it's an absolute path.
func LoadJSON(path string, v interface{}) error {
	if strings.TrimSpace(path) == "" {
		return fmt.Errorf("path is empty")
	}

	// #nosec G304
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// GetTestDataPath returns the absolute path to a file in the testdata directory.
func GetTestDataPath(filename string) string {
	return filepath.Join("testdata", filename)
}
