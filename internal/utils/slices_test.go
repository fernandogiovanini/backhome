package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueWithRepeatedStrings(t *testing.T) {
	input := []string{"a", "b", "a"}

	output := Unique(input)
	expected := []string{"a", "b"}

	assert.ElementsMatch(t, output, expected, "they should be equal")
}

func TestUniqueWithUniqueStrings(t *testing.T) {
	input := []string{"a", "b"}

	output := Unique(input)

	assert.ElementsMatch(t, output, input, "they should be equal")
}
