package subcmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fernandogiovanini/backhome/config"
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "backhome",
	Short: "CLI tool to backup text files (dot files) to a git repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

func init() {
	cobra.OnInitialize(func() {
		cmd := rootCommand.CalledAs()

		// do not load config if command is init so we can create
		// the config file
		if cmd != "init" {
			config.InitConfig()
		}
		logger.InitLogger()
	})

	rootCommand.PersistentFlags().StringVar(&config.ConfigFile, "config", "", strings.Join([]string{"Default to ", config.DefaultConfigPath(), "/.backhome.yaml"}, ""))
	rootCommand.PersistentFlags().StringVar(&logger.LogLevelStr, "logLevel", "INFO", "INFO, DEBUG, ERROR. Default to  INFO")
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
