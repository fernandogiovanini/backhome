package command

import (
	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/spf13/cobra"
)

func buildCopyCommand(newApp func() (*app.App, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "copy",
		Short: "Copy files",
		Long:  "Copy files to local repository, replacing files on destination",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}

			if err := app.Copy(); err != nil {
				printer.Error("Failed to copy files: %v", err)
			}
			return nil
		},
	}
}
