package command

import (
	"fmt"
	"os"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/spf13/cobra"
)

func buildRootCommand() *cobra.Command {
	cobra.OnInitialize(func() {
		logger.InitLogger()

		// TODO: Refactor config loading logic do not load config
		// if command is init so we can create the config file
		if len(os.Args) > 1 && os.Args[1] != "init" {
			config.InitConfig()
		}
	})

	rootCommand := &cobra.Command{
		Use:   "backhome",
		Short: "CLI tool to backup files to a git repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				printer.Error("Failed to show help message U+1F4A3 : %v", err)
				os.Exit(1)
			}
		},
	}

	// TODO: Check if this is the right way of passing flags to config file
	rootCommand.PersistentFlags().StringVar(&config.LocalPath, "local", config.DefaultLocal(), "Local repository")
	rootCommand.PersistentFlags().StringVar(&logger.LogLevelStr, "logLevel", "ERROR", "INFO, DEBUG, ERROR, SILENCE")

	// add subcommands
	rootCommand.AddCommand(buildAddCommand())
	rootCommand.AddCommand(buildCopyCommand())
	rootCommand.AddCommand(buildInitCommand())
	rootCommand.AddCommand(buildSyncCommand())

	return rootCommand
}

func Execute() {
	rootCommand := buildRootCommand()
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(0)
}
