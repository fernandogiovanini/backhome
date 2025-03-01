package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fernandogiovanini/backhome/logger"
)

func ExpandHome(fileName string) string {
	if !strings.HasPrefix(fileName, "~/") {
		return fileName
	}
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, fileName[2:])
}

func ResolvePath(fileName string) (string, error) {
	fileName = ExpandHome(fileName)
	fileName, err := filepath.Abs(fileName)
	if err != nil {
		return "", fmt.Errorf("path to %s cannot be resolved", fileName)
	}
	logger.Debug("path resolved", logger.Args("file", fileName))
	return fileName, nil
}
