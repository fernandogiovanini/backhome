package app

import (
	"github.com/fernandogiovanini/backhome/internal/config"
)

type App struct {
	config *config.Config
}

func New() (*App, error) {
	config, err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	return &App{
		config: config,
	}, nil
}
