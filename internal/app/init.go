package app

import (
	"fmt"
)

func (a *App) Init() error {
	fmt.Fprintf(a.writer, "Initializing local repository... \n")

	if err := a.configStorage.MakeLocalRepository(); err != nil {
		return fmt.Errorf("failed to setup local repository: %w", err)
	}

	if err := a.configStorage.CreateConfigFile(); err != nil {
		return fmt.Errorf("failed to set up config file: %w", err)
	}

	localPath, _ := a.configStorage.GetConfig().GetLocalPath()
	message := "\n" +
		"Local repository initialized at %s\n" +
		"Run 'backhome help' for more commands"
	fmt.Fprintf(a.writer, message, localPath)

	return nil
}
