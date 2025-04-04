package app

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	buffer := &bytes.Buffer{}
	app := &App{
		Writer: buffer,
	}

	app.Error("string %s", "value")

	assert.Equal(t, "\nERROR! string value\n", buffer.String())
}
