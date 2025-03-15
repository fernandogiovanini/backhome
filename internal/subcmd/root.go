package subcmd

import (
	"fmt"
	"os"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(func() {
		logger.InitLogger()

		// do not load config if command is init so we can create
		// the config file
		cmd := initCommand.CalledAs()
		if cmd != "init" {
			config.InitConfig()
		}
	})

	rootCommand.PersistentFlags().StringVar(&config.LocalPath, "local", config.DefaultLocal(), "Local repository")
	rootCommand.PersistentFlags().StringVar(&logger.LogLevelStr, "logLevel", "ERROR", "INFO, DEBUG, ERROR, SILENCE")
}

var rootCommand = &cobra.Command{
	Use:   "backhome",
	Short: "CLI tool to backup text files (dot files) to a git repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(0)
}
