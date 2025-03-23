package app

import (
	"bytes"
	"errors"
	"testing"

	"github.com/fernandogiovanini/backhome/internal/config/mocks"
	"github.com/stretchr/testify/assert"
)

func newConfigIMock() *mocks.ConfigI {
	return &mocks.ConfigI{}
}

func TestInit(t *testing.T) {
	config := newConfigIMock()
	config.
		On("MakeLocalRepository").
		Return(nil).
		On("CreateConfigFile").
		Return(nil).
		On("GetLocalPath").
		Return("/the/local/path", nil)
	buffer := &bytes.Buffer{}
	app := &App{
		config: config,
		output: buffer,
	}

	result := app.Init()

	assert.Nil(t, result)
	assert.Equal(t, "Initializing local repository... \n\nLocal repository initialized at /the/local/path\nRun 'backhome help' for more commands", buffer.String())
}

func TestInitShouldFailToMakeLocalRepository(t *testing.T) {
	config := newConfigIMock()
	config.
		On("MakeLocalRepository").
		Return(errors.New("failed to make local repository"))
	buffer := &bytes.Buffer{}
	app := &App{
		config: config,
		output: buffer,
	}

	result := app.Init()

	assert.Error(t, result)
	assert.Equal(t, "failed to setup local repository: failed to make local repository", result.Error())
}

func TestInitShouldFailToCreateConfigFile(t *testing.T) {
	config := newConfigIMock()
	config.
		On("MakeLocalRepository").
		Return(nil).
		On("CreateConfigFile").
		Return(errors.New("failed to create config file"))
	buffer := &bytes.Buffer{}
	app := &App{
		config: config,
		output: buffer,
	}

	result := app.Init()

	assert.Error(t, result)
	assert.Equal(t, "failed to set up config file: failed to create config file", result.Error())
}
