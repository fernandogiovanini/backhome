package cmd

import (
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Push changes to remote reposiitory",
	Long:  "Push target files in local folder to remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("syncing local to remote...")

		logger.Info("done.")
	},
}
