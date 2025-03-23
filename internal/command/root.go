package command

import (
	"fmt"

	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/spf13/cobra"
)

func buildRootCommand(newApp func(string) (*app.App, error)) *cobra.Command {
	cobra.OnInitialize(func() {
		logger.InitLogger()
	})

	rootCommand := &cobra.Command{
		Use:          "backhome",
		Short:        "CLI tool to backup files to a git repository",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return fmt.Errorf("failed to show help message U+1F4A3 : %w", err)
			}
			return nil
		},
	}

	// bind flags
	rootCommand.PersistentFlags().StringVar(&config.LocalPath, "local", config.DefaultLocal(), "Local repository")
	rootCommand.PersistentFlags().StringVar(&logger.LogLevelStr, "logLevel", "ERROR", "INFO, DEBUG, ERROR, SILENCE")

	// add subcommands
	rootCommand.AddCommand(buildAddCommand(newApp))
	rootCommand.AddCommand(buildCopyCommand(newApp))
	rootCommand.AddCommand(buildInitCommand(newApp))

	return rootCommand
}

func Execute() error {
	rootCommand := buildRootCommand(app.New)
	return rootCommand.Execute()
}
