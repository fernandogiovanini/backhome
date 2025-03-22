package backhome

import (
	"errors"
	"fmt"
	"os"

	"github.com/otiai10/copy"
)

// SafeCopy represents a safe copy of the local repository
// The path is the directory where the files are copied
type SafeCopy struct {
	path  string
	local Local
}

func (safeCopy SafeCopy) Delete() error {
	if safeCopy.path == "" {
		return errors.New("safe copy path is empty")
	}

	if err := os.RemoveAll(safeCopy.path); err != nil {
		return fmt.Errorf("failed to delete safe copy %s: %w", safeCopy.path, err)
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

func (safeCopy SafeCopy) Restore() error {
	if safeCopy.path == "" {
		return errors.New("safe copy path is empty")
	}

	if err := safeCopy.local.prepareForRestoring(); err != nil {
		return fmt.Errorf("failed to clean up safe copy %s: %w", safeCopy.path, err)
	}

	options := copy.Options{
		Sync:          true,
		PreserveTimes: true,
		PreserveOwner: true,
	}

	if err := copy.Copy(safeCopy.path, safeCopy.local.GetPath(), options); err != nil {
		return fmt.Errorf("failed to copy local %s to safe copy %s dir: %w", safeCopy.path, safeCopy.local, err)
	}

	return nil
}

func (safeCopy SafeCopy) GetPath() string {
	return safeCopy.path
}
