package app

import (
	"fmt"
)

func (a *App) Init() error {
	fmt.Fprintf(a.Writer, "Initializing local repository... \n")

	if err := a.ConfigStorage.MakeLocalRepository(); err != nil {
		return fmt.Errorf("failed to setup local repository: %w", err)
	}

	if err := a.ConfigStorage.CreateConfigFile(); err != nil {
		return fmt.Errorf("failed to set up config file: %w", err)
	}

	localPath, _ := a.ConfigStorage.GetConfig().GetLocalPath()
	message := "\n" +
		"Local repository initialized at %s\n" +
		"Run 'backhome help' for more commands"
	fmt.Fprintf(a.Writer, message, localPath)

	return nil
}
