package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 1, Part1(LoadInput("input_test.txt")))
}

func TestRunPart1 (t *testing.T) {
	assert.Equal(t, 288, Part1(LoadInput("input_test.txt")))
}

func TestRunPart2 (t *testing.T) {
	assert.Equal(t, 71503, Part1(LoadInput("input_test.txt")))
}
