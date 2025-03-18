package subcmd

import (
	"fmt"
	"os"

	"github.com/fernandogiovanini/backhome/internal/backhome"
	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize backhome",
	Long:  "Initialize backhome local repository",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.InitLocalPath(); err != nil {
			printer.Error("Failed to set local path:\n%v", err)
			logger.Fatalf("failed to set local path: %v", err)
		}

		if config.ConfigExists() {
			if err := config.LoadConfig(); err != nil {
				printer.Error("Failed to load config file:\n%v", err)
				logger.Fatalf("failed to load config file: %v", err)
				os.Exit(1)
			}

			message := "\n" +
				"Config file found at %s\n" +
				"Run 'backhome help' for instructions\n"
			fmt.Printf(message, config.GetConfigFilePath())
			os.Exit(0)
		}

		fmt.Print("Initializing local repository... ")

		if err := setupLocalRepository(); err != nil {
			printer.Error("Failed to load config file:\n%v", err)
			logger.Fatalf("failed to setup local repository: %v", err)
		}

		if err := setupConfig(); err != nil {
			printer.Error("Failed to set up config file:\n%v", err)
			logger.Fatalf("failed to set up config file: %v", err)
		}

		if err := config.LoadConfig(); err != nil {
			printer.Error("Failed to load config file:\n%v", err)
			logger.Fatalf("failed to load config file: %v", err)
		}

		fmt.Println("Done.")

		localPath, err := config.GetLocalPath()
		if err != nil {
			printer.Error("Failed to get local path:\n%v", err)
			logger.Fatalf("failed to get local path: %v", err)
		}

		message := "\n" +
			"Local repository initialized at %s\n" +
			"Run 'backhome help' for instructions"
		fmt.Printf(message, localPath)
	},
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
