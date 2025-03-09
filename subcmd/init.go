package subcmd

import (
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize backhome",
	Long:  "Initialize backhome local repository",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("initializing local repository...")

		// local, err := backhome.NewLocal(config.Configuration.Local)
		// if err != nil {
		// 	logger.Fatalf("failed to create local repository %s: %v", config.Configuration.Local, err)
		// }

		// // if err := local.Init(); err != nil {
		// // 	logger.Fatalf("failed to initialize local repository: %v", err)
		// // }

		logger.Info("done.")
	},
}

func init() {
	rootCommand.AddCommand(initCommand)
}
