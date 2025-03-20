package command

import (
	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/spf13/cobra"
)

func buildInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize backhome",
		Long:  "Initialize backhome local repository",
		Run: func(cmd *cobra.Command, args []string) {
			app := app.New()
			if err := app.Init(); err != nil {
				printer.Error("Failed to initialize backhome local: %v", err)
			}
		},
	}
}
