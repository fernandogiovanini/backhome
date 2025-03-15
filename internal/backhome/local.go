package backhome

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
)

// Local represents the local repository
// The base path is the directory where the files are copied
type Local struct {
	path string
}

func NewLocal(path string) (*Local, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open local path %s: %w", path, err)
	}
	defer dir.Close()

	fileinfo, err := dir.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to read local %s stats: %w", path, err)
	}
	if !fileinfo.IsDir() {
		return nil, fmt.Errorf("local %s is not a directory", path)
	}

	return &Local{path: path}, nil
}

// MakeLocal creates a new local repository in the specified path
// if it does not exist and returns a pointer to the Local struct
func MakeLocal(path string) (*Local, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create local path %s: %w", path, err)
	}

	return NewLocal(path)
}

func (local *Local) GetPath() string {
	return local.path
}

func (local *Local) prepareForRestoring() error {
	if err := filepath.Walk(local.path, func(path string, info os.FileInfo, err error) error {
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
		return fmt.Errorf("failed to walk path %s to remove files: %w", local.path, err)
	}

	return nil
}

// NewSafeCopy creates a new safe copy of the local repository
func (local *Local) NewSafeCopy() (*SafeCopy, error) {
	destinationPath := strings.Join([]string{local.path, "backhome"}, ".")

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

	if err := copy.Copy(local.path, destinationPath, options); err != nil {
		return nil, fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", local.path, destinationPath, err)
	}

	return &SafeCopy{
		path:  destinationPath,
		local: local,
	}, nil
}
