package app

import (
	"fmt"
	"strconv"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
)

func (a *App) Add(files ...string) error {
	for _, file := range files {
		logger.Info("adding file %s", file)
		if err := config.AddFile(file); err != nil {
			return fmt.Errorf("failed to add files to config: %w", err)
		}
	}

	if err := config.Save(); err != nil {
		return fmt.Errorf("failed to add files to config: %w", err)
	}

	// TODO: Pluralize
	fmt.Printf(strconv.Itoa(len(files))+" files(s) added to %s", config.GetConfigFilePath())

	return nil
}
