package app

import (
	"errors"

	"github.com/fernandogiovanini/backhome/internal/config"
	"github.com/spf13/viper"
)

type App struct {
	config *config.Config
}

func New(command string) (*App, error) {

	config, err := config.InitConfig()
	// return add if no error of if command is init and
	// error is ConfigFileNotFoundError (because init will create the file)
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
	}, nil
}
