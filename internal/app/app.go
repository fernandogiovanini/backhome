package app

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/fernandogiovanini/backhome/internal/filesystem"
	"github.com/spf13/viper"
)

type App struct {
	configStorage config.ConfigStorage
	filesystem    filesystem.FileSystem
	writer        io.Writer
}

func New(command string) (*App, error) {
	filesystem := filesystem.NewFileSystem()
	configStorage, err := config.NewConfigStorage(config.LocalPath, config.DefaultConfigFilename, filesystem, viper.New())

	// return pointer of App if config.NewConfigStorage() returns no error or if
	// command is init and error is ConfigFileNotFoundError (because init will create the file)
	if err != nil {
		if command != "init" {
			return nil, err
		}
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, err
		}
	}

	return &App{
		configStorage: configStorage,
		filesystem:    filesystem,
		writer:        os.Stdout,
	}, nil
}

func (a *App) Error(message string, args ...any) {
	fmt.Fprintf(a.writer, "\nERROR! %s\n", fmt.Sprintf(message, args...))
}
