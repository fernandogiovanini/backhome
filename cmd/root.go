package cmd

import (
	"fmt"
	"os"

	"github.com/fernandogiovanini/backhome/config"
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "backhome",
	Short: "CLI tool to backup text files to a git repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func init() {
	cobra.OnInitialize(logger.InitLogger, config.InitConfig)

	rootCommand.PersistentFlags().StringVar(&config.ConfigFile, "config", "", "Default to  $HOME/.config.yaml")
	rootCommand.PersistentFlags().StringVar(&logger.LogLevelStr, "logLevel", "INFO", "INFO or DEBUG. Default to  INFO")
}
