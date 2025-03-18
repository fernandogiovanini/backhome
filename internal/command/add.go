package subcmd

import (
	"fmt"
	"strconv"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/logger"
	"github.com/fernandogiovanini/backhome/internal/printer"
	"github.com/fernandogiovanini/backhome/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add <file> <file> ...",
	Short: "Set files to be copied to the local repository",
	Long: "\nSet files to be copied to the local repository.\n\n" +
		"The files will be copied to the local repository when you run the copy command.\n\n" +
		"backhome copy --local path/to/local \n\n" +
		"To add a file with spaces in the name, use quotes. For example: backhome add 'my file.txt'",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		args = utils.Unique(args)
		for _, file := range args {
			logger.Info("adding file %s", file)
			if err := config.AddFile(file); err != nil {
				printer.Error("Failed to add file %s:\n%v", file, err)
				logger.Fatalf("failed to add files to config: %v", err)
			}
		}

		if err := config.Save(); err != nil {
			printer.Error("Failed to add files to config:\n%v", err)
			logger.Fatalf("failed to add files to config: %v", err)
		}

		fmt.Printf(strconv.Itoa(len(args))+" files(s) added to %s", config.GetConfigFilePath())
	},
}
