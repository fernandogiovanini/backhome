package app

import (
	"fmt"
	"strconv"
)

func (a *App) Add(files ...string) error {
	cm := a.ConfigManager
	for _, file := range files {
		if err := cm.AddFile(file); err != nil {
			return fmt.Errorf("failed to add files to config: %w", err)
		}
	}

	if err := cm.Save(); err != nil {
		return fmt.Errorf("failed to add files to config: %w", err)
	}

	// TODO: Pluralize
	fmt.Printf(strconv.Itoa(len(files))+" files(s) added to %s", a.Config.GetConfigFilePath())

	return nil
}
