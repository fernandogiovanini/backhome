package command

import (
	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/spf13/cobra"
)

func buildInitCommand(newApp func() (*app.App, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize backhome",
		Long:  "Initialize backhome local repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}
			return app.Init()
		},
	}
}
