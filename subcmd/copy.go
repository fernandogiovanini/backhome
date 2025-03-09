package subcmd

import (
	"github.com/fernandogiovanini/backhome/backhome"
	"github.com/fernandogiovanini/backhome/config"
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/spf13/cobra"
)

var (
	safeRun bool

	copyCommand = &cobra.Command{
		Use:   "copy",
		Short: "Copy target files",
		Long:  "Copy target files to local repository, replacing any changes on destination",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("copying target files...")

			local, err := backhome.NewLocal(config.Configuration.Local)
			if err != nil {
				logger.Fatalf("failed to create local repository %s: %v", config.Configuration.Local, err)
			}

			safeCopy := &backhome.SafeCopy{}
			if safeRun {
				safeCopy, err = local.NewSafeCopy()
				if err != nil {
					logger.Fatalf("failed to create safe copy: %v", err)
				}
			}

			items, err := backhome.NewItemList(config.Configuration.BackupItems)
			if err != nil {
				logger.Fatalf("failed to resolve target paths: %v", err)
			}

			if err := items.CopyTo(local); err == nil {
				safeCopy.Delete()
				logger.Info("done.")
			} else {
				logger.Err("failed to copy files: %v", err)
				if safeRun {
					if err := backhome.RestoreSafeCopy(safeCopy); err != nil {
						logger.Err("failed to restore safe copy: %w", err)
					}
				}
				logger.Fatalf("failed to copy file to %s: %v", local.BasePath, err)
			}
		},
	}
)

func init() {
	copyCommand.Flags().BoolVar(&safeRun, "safe", true, "Create a safe copy of the local repository files before copying news files to it. Default to true")
	rootCommand.AddCommand(copyCommand)
}
