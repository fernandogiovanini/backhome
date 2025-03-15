package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fernandogiovanini/backhome/internal/logger"
)

func ExpandHome(path string) string {
	if !strings.HasPrefix(path, "~/") {
		return path
	}
	home, _ := os.UserHomeDir()
	expanded := filepath.Join(home, path[2:])
	logger.Debug("path %s expanded to %s", path, expanded)
	return expanded
}

func ResolvePath(path string) (string, error) {
	path = ExpandHome(path)
	resolved, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("path %s cannot be resolved: %w", path, err)
	}
	logger.Debug("path %s resolved to %s", path, resolved)
	return resolved, nil
}
