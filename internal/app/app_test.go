package app

import (
	"bytes"
	"testing"

	"github.com/fernandogiovanini/backhome/internal/config/mocks"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	config := &mocks.ConfigI{}
	config.On("dsa")
	buffer := &bytes.Buffer{}
	app := &App{
		config: config,
		output: buffer,
	}

	app.Error("string %s", "value")

	assert.Equal(t, "\nERROR! string value\n", buffer.String())
}
