package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 136, Part1(LoadInput("input_test.txt")))
}

func TestRollRock(t *testing.T) {
	data := LoadInput("input_test.txt")
	newLocation := RollRock(data, Coord{1,9})
	assert.Equal(t, Coord{1,5}, newLocation)
	// newLocation := RollRock(data, Coord{9,3})
	// assert.Equal(t, Coord{9,2}, newLocation)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 64, Part2(LoadInput("input_test.txt")))
}