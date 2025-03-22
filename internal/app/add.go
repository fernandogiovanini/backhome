package app

import (
	"fmt"
	"strconv"
)

func (a *App) Add(files ...string) error {
	config := a.config
	for _, file := range files {
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
