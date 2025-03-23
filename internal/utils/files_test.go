package utils

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandHome(t *testing.T) {
	input := "~/file"

	output := ExpandHome(input)

	homeDir, _ := os.UserHomeDir()
	assert.Equal(t, output, strings.Join([]string{homeDir, "file"}, string(os.PathSeparator)))
}

func TestDoNotExpandHome(t *testing.T) {
	input := "/dir/file"

	output := ExpandHome(input)

	assert.Equal(t, output, input)
}
