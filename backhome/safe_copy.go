package backhome

import (
	"errors"
	"fmt"
	"os"

	"github.com/fernandogiovanini/backhome/logger"
	"github.com/otiai10/copy"
)

// SafeCopy represents a safe copy of the local repository
// The path is the directory where the files are copied
type SafeCopy struct {
	Path  string
	Local *Local
}

func (safeCopy *SafeCopy) Delete() error {
	if safeCopy == nil {
		return errors.New("nil pointer: SafeCopy is not initialized")
	}

	if safeCopy.Path == "" {
		return errors.New("safe copy path is empty")
	}

	if err := os.RemoveAll(safeCopy.Path); err != nil {
		return fmt.Errorf("failed to delete safe copy %s: %w", safeCopy.Path, err)
	}

	return nil
}

func RestoreSafeCopy(safeCopy *SafeCopy) error {
	if safeCopy == nil {
		return errors.New("nil pointer: SafeCopy is not initialized")
	}

	if err := safeCopy.Restore(); err != nil {
		return fmt.Errorf("failed to restore safe copy from %s to %s: %w", safeCopy.Path, safeCopy.Local, err)
	}

	if err := safeCopy.Delete(); err != nil {
		return fmt.Errorf("failed to delete safe copy at %s: %w", safeCopy.Path, err)
	}

	logger.Debug("safe copy restored from %s to %s and deleted", safeCopy.Path, safeCopy.Local)

	return nil
}

func (safeCopy *SafeCopy) Restore() error {
	if safeCopy == nil {
		return errors.New("nil pointer: SafeCopy is not initialized")
	}

	if safeCopy.Path == "" {
		return errors.New("safe copy path is empty")
	}

	if safeCopy.Local == nil {
		return errors.New("nil pointer: Local is not initialized")
	}

	if err := safeCopy.Local.prepareForRestoring(); err != nil {
		return fmt.Errorf("failed to clean up safe copy %s: %w", safeCopy.Path, err)
	}

	options := copy.Options{
		Sync:          true,
		PreserveTimes: true,
		PreserveOwner: true,
	}

	if err := copy.Copy(safeCopy.Path, safeCopy.Local.BasePath, options); err != nil {
		return fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", safeCopy.Path, safeCopy.Local, err)
	}

	logger.Debug("safe copy %s restored to %s", safeCopy.Path, safeCopy.Local)

	return nil
}
