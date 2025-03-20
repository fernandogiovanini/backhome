package command

import (
	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/spf13/cobra"
)

func buildCopyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "copy",
		Short: "Copy files",
		Long:  "Copy files to local repository, replacing files on destination",
		Run: func(cmd *cobra.Command, args []string) {
			app := app.New()
			if err := app.Copy(); err != nil {
				printer.Error("Failed to copy files:\n%v", err)
				app.Fatal("Failed to copy files: %v", err)
			}
		},
	}
}
