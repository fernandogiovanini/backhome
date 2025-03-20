package app

import (
	"fmt"
	"strconv"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/printer"
)

func (a *App) Add(files ...string) {
	for _, file := range files {
		logger.Info("adding file %s", file)
		if err := config.AddFile(file); err != nil {
			printer.Error("Failed to add file %s:\n%v", file, err)
			logger.Fatalf("failed to add files to config: %v", err)
		}
	}

	if err := config.Save(); err != nil {
		printer.Error("Failed to add files to config:\n%v", err)
		logger.Fatalf("failed to add files to config: %v", err)
	}

	// TODO: Pluralize
	fmt.Printf(strconv.Itoa(len(files))+" files(s) added to %s", config.GetConfigFilePath())
}
