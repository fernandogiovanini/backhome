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
	Config        config.Config
	ConfigManager config.ConfigManager
	ConfigStorage config.ConfigStorage
	Filesystem    filesystem.FileSystem
	Writer        io.Writer
}

func New(command string) (*App, error) {
	v := viper.New()
	fs := filesystem.NewFileSystem()

	cfgStorage, err := config.NewConfigStorage(config.LocalPath, config.DefaultConfigFilename, fs, v)
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

	cfg := cfgStorage.GetConfig()

	return &App{
		Config:        cfg,
		ConfigStorage: cfgStorage,
		ConfigManager: config.NewConfigManager(v, fs, cfg),
		Filesystem:    fs,
		Writer:        os.Stdout,
	}, nil
}

func (a *App) Error(message string, args ...any) {
	fmt.Fprintf(a.Writer, "\nERROR! %s\n", fmt.Sprintf(message, args...))
}
