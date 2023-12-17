package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestTransformChar(t *testing.T) {
	assert.Equal(t, 200, ProcessCharByte(byte('H')))
}

func TestProcessWord(t *testing.T) {
	assert.Equal(t, 0, ProcessWord("rn"))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 1320, Part1(LoadInput("input_test.txt")))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 144, Part2(LoadInput("input.txt")))
	// assert.Equal(t, 144, Part2(LoadInput("input_test.txt")))
}

func TestProcessWord2(t *testing.T) {
	// assert.Equal(t, 0, ProcessWord2("rn=1"))
	assert.Equal(t, 0, ProcessWord2("cm-"))
}