package command

import (
	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/spf13/cobra"
)

func buildSyncCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Push changes to remote reposiitory",
		Long:  "Push files in local folder to remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			app := app.New()
			if err := app.Sync(); err != nil {
				printer.Error("%v", err)
			}
		},
	}
}
