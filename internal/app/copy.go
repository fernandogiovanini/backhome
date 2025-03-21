package app

import (
	"fmt"

	"github.com/fernandogiovanini/backhome/internal/backhome"
	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/printer"
)

func (a *App) Copy() error {
	localPath, err := config.GetLocalPath()
	if err != nil {
		return fmt.Errorf("failed to get local path %s: %w", localPath, err)
	}

	logger.Info("copying files on %s to %s", config.GetFilenames(), localPath)
	fmt.Printf("Copying files to %s\n\n", localPath)

	local, err := backhome.NewLocal(localPath)
	if err != nil {
		return fmt.Errorf("failed to open local repository %s: %v", localPath, err)
	}

	files, err := backhome.NewFileList(config.GetFilenames())
	if err != nil {
		return fmt.Errorf("failed to get the list of files: %v", err)
	}

	if len(files.Files) == 0 {
		fmt.Println("No files to copy")
		return nil
	}

	err = copyFiles(files, local)
	if err != nil {
		return fmt.Errorf("failed to copy files: %w", err)
	}

	fmt.Println("\nDone")

	return nil
}

func copyFiles(files *backhome.FileList, local *backhome.Local) error {
	safeCopy, err := local.NewSafeCopy()
	if err != nil {
		return fmt.Errorf("failed to create safe copy: %v", err)
	}

	if err := files.CopyTo(local); err != nil {
		if err := backhome.RestoreSafeCopy(safeCopy); err != nil {
			logger.Err("failed to restore safe copy: %v", err)
		}
		return fmt.Errorf("failed to copy file to %s:\n%v", local.GetPath(), err)
	}

	// failing to delete the safe copy is not a fatal error
	if err := safeCopy.Delete(); err != nil {
		// TODO: It feels like this printer shouldn't be here
		// Maybe just call sdtout.Print?
		printer.Error("Failed to delete safe copy. Check log file at %s for more informarion", logger.GetLogFile())
		logger.Err("failed to delete safe copy: %v", err)
	}

	return nil
}
