package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 405, Part1(LoadInput("input_test.txt")))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 401, Part2(LoadInput("input_test.txt")))
}