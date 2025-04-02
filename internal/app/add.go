package app

import (
	"fmt"
	"strconv"
)

func (a *App) Add(files ...string) error {
	cs := a.configStorage
	for _, file := range files {
		if err := cs.AddFile(file); err != nil {
			return fmt.Errorf("failed to add files to config: %w", err)
		}
	}

	if err := cs.Save(); err != nil {
		return fmt.Errorf("failed to add files to config: %w", err)
	}

	// TODO: Pluralize
	cfg := cs.GetConfig()
	fmt.Printf(strconv.Itoa(len(files))+" files(s) added to %s", cfg.GetConfigFilePath())

	return nil
}
