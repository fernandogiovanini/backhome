package command

import (
	"fmt"

	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/spf13/cobra"
)

func buildInitCommand(newApp func(string) (*app.App, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize backhome",
		Long:  "Initialize backhome local repository",
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp(cmd.CalledAs())
			if err != nil {
				return err
			}

			if err := app.Init(); err != nil {
				return fmt.Errorf("failed to init local repository and config: %w", err)
			}

			return nil
		},
	}
}
