package app

import (
	"fmt"
)

func (a *App) Init() error {
	fmt.Fprintf(a.output, "Initializing local repository... \n")

	if err := a.config.MakeLocalRepository(); err != nil {
		return fmt.Errorf("failed to setup local repository: %w", err)
	}

	if err := a.config.CreateConfigFile(); err != nil {
		return fmt.Errorf("failed to set up config file: %w", err)
	}

	localPath, _ := a.config.GetLocalPath()
	message := "\n" +
		"Local repository initialized at %s\n" +
		"Run 'backhome help' for more commands"
	fmt.Fprintf(a.output, message, localPath)

	return nil
}
