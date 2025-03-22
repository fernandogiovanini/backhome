package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExpandHome(path string) string {
	if !strings.HasPrefix(path, "~/") {
		return path
	}
	home, _ := os.UserHomeDir()
	expanded := filepath.Join(home, path[2:])
	return expanded
}

func ResolvePath(path string) (string, error) {
	path = ExpandHome(path)
	resolved, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("path %s cannot be resolved: %w", path, err)
	}

	return resolved, nil
}
