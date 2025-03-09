package backhome

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fernandogiovanini/backhome/utils"
	"github.com/otiai10/copy"
)

// Local represents the local repository
// The base path is the directory where the files are copied
type Local struct {
	BasePath string
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

	options := copy.Options{
		Skip: func(srcinfo os.FileInfo, src, dest string) (bool, error) {
			// skip copying the .git directory
			if srcinfo.IsDir() && srcinfo.Name() == ".git" {
				return true, nil
			}

			// skip copying the .DS_Store file
			if srcinfo.Name() == ".DS_Store" {
				return true, nil
			}

			return false, nil
		},
		OnSymlink: func(src string) copy.SymlinkAction {
			// do not follow symlinks
			return copy.Skip
		},
		Sync:          true,
		PreserveTimes: true,
		PreserveOwner: true,
	}

	if err := copy.Copy(local.BasePath, destinationPath, options); err != nil {
		return nil, fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", local.BasePath, destinationPath, err)
	}

	return &SafeCopy{
		Path:  destinationPath,
		Local: local,
	}, nil
}

func (local *Local) prepareForRestoring() error {
	if local == nil {
		return errors.New("nil pointer: Local is not initialized")
	}

	if local.BasePath == "" {
		return errors.New("base path is empty")
	}

	if err := filepath.Walk(local.BasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path %s to remove files: %w", path, err)
		}
		if path != ".git" {
			return nil
		}

		if err := os.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to delete path %s: %w", path, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to walk path %s to remove files: %w", local.BasePath, err)
	}

	return nil
}
