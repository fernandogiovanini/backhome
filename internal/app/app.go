package app

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/spf13/viper"
)

type App struct {
	config config.IConfig
	output io.Writer
}

func New(command string) (*App, error) {

	config, err := config.InitConfig()
	// return pointer of App if config.InitConfig() returns no error or if
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
		config: config,
		output: os.Stdout,
	}, nil
}

func (a *App) Error(message string, args ...any) {
	fmt.Fprintf(a.output, "\nERROR! %s\n", fmt.Sprintf(message, args...))
}
