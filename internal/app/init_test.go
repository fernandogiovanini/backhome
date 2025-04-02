package app

import (
	"bytes"
	"errors"
	"testing"

	"github.com/fernandogiovanini/backhome/internal/config"
	cfgmock "github.com/fernandogiovanini/backhome/internal/config/mocks"
	fsmock "github.com/fernandogiovanini/backhome/internal/filesystem/mocks"
	"github.com/stretchr/testify/assert"
)

func newFileSystemMock() *fsmock.FileSystem {
	return &fsmock.FileSystem{}
}

func TestInit(t *testing.T) {
	filesystem := newFileSystemMock()
	cfg, _ := config.NewConfig("/the/local/path", config.DefaultConfigFilename)
	configStorage := cfgmock.NewConfigStorage(t)
	configStorage.
		On("MakeLocalRepository").Return(nil).
		On("CreateConfigFile").Return(nil).
		On("GetConfig").Return(cfg)
	buffer := &bytes.Buffer{}
	app := &App{
		configStorage: configStorage,
		filesystem:    filesystem,
		writer:        buffer,
	}

	result := app.Init()

	assert.Nil(t, result)
	assert.Equal(t, "Initializing local repository... \n\nLocal repository initialized at /the/local/path\nRun 'backhome help' for more commands", buffer.String())
}

func TestInitShouldFailToMakeLocalRepository(t *testing.T) {
	filesystem := newFileSystemMock()
	configStorage := cfgmock.NewConfigStorage(t)
	configStorage.On("MakeLocalRepository").Return(errors.New("failed to make local repository"))
	buffer := &bytes.Buffer{}
	app := &App{
		configStorage: configStorage,
		filesystem:    filesystem,
		writer:        buffer,
	}

	result := app.Init()

	assert.Error(t, result)
	assert.Equal(t, "failed to setup local repository: failed to make local repository", result.Error())
}

func TestInitShouldFailToCreateConfigFile(t *testing.T) {
	filesystem := newFileSystemMock()
	configStorage := cfgmock.NewConfigStorage(t)
	configStorage.
		On("MakeLocalRepository").Return(nil).
		On("CreateConfigFile").Return(errors.New("failed to create config file"))
	buffer := &bytes.Buffer{}
	app := &App{
		configStorage: configStorage,
		filesystem:    filesystem,
		writer:        buffer,
	}

	result := app.Init()

	assert.Error(t, result)
	assert.Equal(t, "failed to set up config file: failed to create config file", result.Error())
}
