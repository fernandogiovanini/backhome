package app

import (
	"bytes"
	"testing"

	cfgmock "github.com/fernandogiovanini/backhome/internal/config/mocks"
	fsmock "github.com/fernandogiovanini/backhome/internal/filesystem/mocks"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	filesystem := &fsmock.FileSystem{}
	configStorage := &cfgmock.ConfigStorage{}
	buffer := &bytes.Buffer{}
	app := &App{
		configStorage: configStorage,
		filesystem:    filesystem,
		writer:        buffer,
	}

	app.Error("string %s", "value")

	assert.Equal(t, "\nERROR! string value\n", buffer.String())
}
