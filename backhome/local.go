package backhome

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fernandogiovanini/backhome/utils"
)

// Local represents the local repository
// The base path is the directory where the files are copied
type Local struct {
	BasePath string
}

// SafeCopy represents a safe copy of the local repository
// The path is the directory where the files are copied
type SafeCopy struct {
	Path         string
	OriginalPath string
}

func NewLocal(basePath string) (*Local, error) {
	if basePath == "" {
		return nil, errors.New("base path cannot be empty")
	}

	basePath, err := utils.ResolvePath(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve base path %s: %w", basePath, err)
	}

	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base path %s: %w", basePath, err)
	}

	return &Local{BasePath: basePath}, nil
}

func (local *Local) NewSafeCopy() (*SafeCopy, error) {
	destinationPath := strings.Join([]string{local.BasePath, "backhome"}, ".")

	if err := utils.CopyDir(destinationPath, local.BasePath); err != nil {
		return nil, fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", local.BasePath, destinationPath, err)
	}

	return &SafeCopy{
		Path:         destinationPath,
		OriginalPath: local.BasePath,
	}, nil
}

func (safeCopy *SafeCopy) RestoreTo() error {
	return errors.New("RestoreTo NOT IMPLEMENTED YET")
}
