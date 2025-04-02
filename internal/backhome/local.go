//go:generate mockery --all --case snake

package backhome

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fernandogiovanini/backhome/internal/filesystem"
	"github.com/otiai10/copy"
)

// Local represents the local repository
// The base path is the directory where the files are copied
type Local struct {
	path       string
	filesystem filesystem.FileSystem
}

func NewLocal(filesystem filesystem.FileSystem, path string) (*Local, error) {
	dir, err := filesystem.Open(path)
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

func (l Local) Path() string {
	return l.path
}

func (l Local) prepareForRestoring() error {
	if err := filepath.Walk(l.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path %s to remove files: %w", path, err)
		}
		if path != ".git" {
			return nil
		}

		if err := l.filesystem.RemoveAll(path); err != nil {
			return fmt.Errorf("failed to delete path %s: %w", path, err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("failed to walk path %s to remove files: %w", l.path, err)
	}

	return nil
}

// NewSafeCopy creates a new safe copy of the local repository
func (l Local) NewSafeCopy(fs filesystem.FileSystem) (*SafeCopy, error) {
	dstPath := strings.Join([]string{l.path, "backhome"}, ".")

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

	if err := copy.Copy(l.path, dstPath, options); err != nil {
		return nil, fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", l.path, dstPath, err)
	}

	return &SafeCopy{
		local:      l,
		filesystem: fs,
		path:       dstPath,
	}, nil
}
