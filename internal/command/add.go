package command

import (
	"fmt"

	"github.com/fernandogiovanini/backhome/internal/app"
	"github.com/fernandogiovanini/backhome/internal/utils"
	"github.com/spf13/cobra"
)

func buildAddCommand(newApp func() (*app.App, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "add <file> <file> ...",
		Short: "Set files to be copied to the local repository",
		Long: "\nSet files to be copied to the local repository.\n\n" +
			"The files will be copied to the local repository when you run the copy command.\n\n" +
			"backhome copy --local path/to/local \n\n" +
			"To add a file with spaces in the name, use quotes. For example: backhome add 'my file.txt'",
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			app, err := newApp()
			if err != nil {
				return err
			}

			if err := app.Add(utils.Unique(args)...); err != nil {
				return fmt.Errorf("failed to add files to config: %w", err)
			}

			return nil
		},
	}
}
