package app

import (
	"fmt"
	"os"

	"github.com/fernandogiovanini/backhome/internal/backhome"
	"github.com/fernandogiovanini/backhome/internal/config"
)

func (a *App) Init() error {
	if err := config.InitLocalPath(); err != nil {
		return fmt.Errorf("failed to set local path: %w", err)
	}

	if config.ConfigExists() {
		if err := config.LoadConfig(); err != nil {
			return fmt.Errorf("failed to load config file: %w", err)
		}

		message := "\n" +
			"Config file found at %s\n" +
			"Run 'backhome help' for instructions\n"
		fmt.Printf(message, config.GetConfigFilePath())
		os.Exit(0)
	}

	fmt.Print("Initializing local repository... ")

	if err := setupLocalRepository(); err != nil {
		return fmt.Errorf("failed to setup local repository: %w", err)
	}

	if err := setupConfig(); err != nil {
		return fmt.Errorf("failed to set up config file: %w", err)
	}

	if err := config.LoadConfig(); err != nil {
		return fmt.Errorf("failed to load config file: %w", err)
	}

	fmt.Println("Done.")

	localPath, err := config.GetLocalPath()
	if err != nil {
		return fmt.Errorf("failed to get local path: %w", err)
	}

	message := "\n" +
		"Local repository initialized at %s\n" +
		"Run 'backhome help' for instructions"
	fmt.Printf(message, localPath)

	return nil
}

func setupLocalRepository() error {
	localPath, err := config.GetLocalPath()
	if err != nil {
		return fmt.Errorf("failed to get local path: %w", err)
	}

	if _, err := backhome.MakeLocal(localPath); err != nil {
		return fmt.Errorf("failed to create local repository %s: %w", localPath, err)
	}

	return nil
}

func setupConfig() error {
	file, err := os.OpenFile(config.GetConfigFilePath(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("failed to create config file %s: %w", config.GetConfigFilePath(), err)
	}
	defer file.Close()

	return nil
}
