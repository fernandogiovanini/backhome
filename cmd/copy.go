package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fernandogiovanini/backhome/config"
	"github.com/fernandogiovanini/backhome/logger"
	"github.com/fernandogiovanini/backhome/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(copyCommand)
}

type Target struct {
	originalPath string
	resolvedPath string
}

var copyCommand = &cobra.Command{
	Use:   "copy",
	Short: "Copy target files",
	Long:  "Copy target files to local repository, replacing any changes on destination",
	Run: func(cmd *cobra.Command, args []string) {
		var resolvedPaths, err = resolvePaths(config.Configuration.Targets)
		if err != nil {
			logger.Fatalf("failed to resolve target paths: %s", err)
		}
		if err := validateTargets(resolvedPaths); err != nil {
			logger.Fatalf("failed to validate target paths: %s", err)
		}
		if err := copyFilesToLocal(resolvedPaths, config.Configuration.Local); err != nil {
			logger.Fatalf("failed to copy files: %s", err)
		}
	},
}

func validateTargets(targets []Target) error {
	for i := range targets {
		target := targets[i]
		logger.Debug("validating target", logger.Args("target", target.resolvedPath))
		file, err := os.Open(target.resolvedPath)
		if err != nil {
			logger.Err(err)
			return fmt.Errorf("failed to open target %s: %s", target.resolvedPath, err)
		}
		defer file.Close()

		if err != nil {
			logger.Err(err)
			return fmt.Errorf("target %s is not a readable file: %s", target.resolvedPath, err)
		}
	}
	return nil
}

func resolvePaths(fileNames []string) ([]Target, error) {
	targets := []Target{}
	for i := range fileNames {
		resolvedPath, err := utils.ResolvePath(fileNames[i])
		if err != nil {
			return nil, fmt.Errorf("failed to resolve path of target %s: %s", fileNames[i], err)
		}
		targets = append(targets, Target{
			originalPath: fileNames[i],
			resolvedPath: resolvedPath,
		})
	}
	return targets, nil
}

func copyFilesToLocal(targets []Target, localDestination string) error {
	localDestination, err := utils.ResolvePath(localDestination)
	if err != nil {
		return fmt.Errorf("failed to resolve local destination path %s: %s", localDestination, err)
	}
	if err := ensureValidLocalDestination(localDestination); err != nil {
		return fmt.Errorf("failed to ensure local destination path %s exist and is writeable: %s", localDestination, err)
	}
	for i := range targets {
		if err := targets[i].copyFile(localDestination); err != nil {
			return fmt.Errorf("failed to copy file %s to local destination %s: ", targets[i], localDestination, err)
		}
	}
	return nil
}

func (target Target) copyFile(localDestination string) error {
	logger.Debug("copying file", logger.Args("target", target.resolvedPath), logger.Args("destination", localDestination))

	logger.Debug("open target file", logger.Args("target", target.resolvedPath))
	srcFile, err := os.Open(target.resolvedPath)
	if err != nil {
		return fmt.Errorf("failed to open target file %s: %s", target.resolvedPath, err)
	}
	defer srcFile.Close()

	// Check source file stats
	logger.Debug("check target file stats", logger.Args("target", target.resolvedPath))
	srcFileStats, err := os.Stat(target.resolvedPath)
	if err != nil {
		return fmt.Errorf("failed to verify source file %s stats: %s", target.resolvedPath, err)
	}

	// Check source dir status (to copy mode to destination)
	logger.Debug("check target directory", logger.Args("target", target.resolvedPath))
	srcPathStats, err := os.Stat(filepath.Dir(target.resolvedPath))
	if err != nil {
		return fmt.Errorf("failed to verify source directory %s stats: %s", target.resolvedPath, err)
	}

	// Get the destination absolute file name and path
	dstFilename := filepath.Join(localDestination, target.resolvedPath)
	dstPathName := filepath.Dir(dstFilename)

	// Create destination path
	logger.Debug("create destination path if it is missing", logger.Args("destination", dstPathName))
	dstPath, err := os.Open(dstPathName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(dstPathName, srcPathStats.Mode())
			if err != nil {
				return fmt.Errorf("failed to create destination path %s: %s", dstPathName, err)
			}
		} else {
			return fmt.Errorf("failed to open destination path %s: %s", dstPathName, err)
		}
	}
	defer dstPath.Close()

	logger.Debug("create destination file", logger.Args("destination", dstFilename))
	dstFile, err := os.OpenFile(dstFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, srcFileStats.Mode())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("failed to create destination file %s: %s", dstFilename, err)
		}
		return fmt.Errorf("failed to open destination file %s: %s", dstFilename, err)
	}
	defer dstFile.Close()

	logger.Debug("copy target file", logger.Args("target", dstFilename))
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy %s to %s: %s", target.resolvedPath, dstFilename, err)
	}

	return nil
}

func ensureValidLocalDestination(localDestination string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory for validation: %s", err)
	}

	if localDestination == home {
		return fmt.Errorf("local destination is invalid: %s", localDestination)
	}

	info, err := os.Stat(localDestination)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(localDestination, os.ModeDir); err != nil {
				return fmt.Errorf("failed to create local destination %s: %s", localDestination, err)
			}
		} else {
			return fmt.Errorf("failed to open local destination %s: %s", localDestination, err)
		}
	}

	if !info.IsDir() {
		return fmt.Errorf("focal destination %s is not a valid directory", localDestination)
	}

	return nil
}
