//go:generate mockery --all --case snake

package backhome

import (
	"errors"
	"fmt"

	"github.com/fernandogiovanini/backhome/internal/filesystem"
	"github.com/otiai10/copy"
)

// SafeCopy represents a safe copy of the local repository
// The path is the directory where the files are copied
type SafeCopy struct {
	path       string
	local      Local
	filesystem filesystem.FileSystem
}

func (sc SafeCopy) Delete() error {
	if sc.path == "" {
		return errors.New("safe copy path is empty")
	}

	if err := sc.filesystem.RemoveAll(sc.path); err != nil {
		return fmt.Errorf("failed to delete safe copy %s: %w", sc.path, err)
	}

	return nil
}

func RestoreSafeCopy(safeCopy *SafeCopy) error {
	if err := safeCopy.Restore(); err != nil {
		return fmt.Errorf("failed to restore safe copy from %s to %s: %w", safeCopy.path, safeCopy.local, err)
	}

	if err := safeCopy.Delete(); err != nil {
		return fmt.Errorf("failed to delete safe copy at %s: %w", safeCopy.path, err)
	}

	return nil
}

func (sc SafeCopy) Restore() error {
	if sc.path == "" {
		return errors.New("safe copy path is empty")
	}

	if err := sc.local.prepareForRestoring(); err != nil {
		return fmt.Errorf("failed to clean up safe copy %s: %w", sc.Path(), err)
	}

	options := copy.Options{
		Sync:          true,
		PreserveTimes: true,
		PreserveOwner: true,
	}

	if err := copy.Copy(sc.path, sc.local.Path(), options); err != nil {
		return fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", sc.path, sc.local, err)
	}

	return nil
}

func (sc SafeCopy) Path() string {
	return sc.path
}
