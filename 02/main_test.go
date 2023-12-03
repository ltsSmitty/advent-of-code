package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", LoadInput("input_test.txt")[0])
}

func TestGetGameId(t *testing.T) {
	assert.Equal(t, 1, GetGameId("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
	assert.Equal(t, 50, GetGameId("Game 50: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
}

func TestGetNumColors(t *testing.T) {
	assert.Equal(t, 23, GetNumColors("3 blue, 4 red, 20 blue").Blue)
	assert.Equal(t, 2, GetNumColors("1 red, 2 green, 6 blue").Green)
	assert.Equal(t, 2, GetNumColors("2 green").Green)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 8, Part1(LoadInput("input_test.txt")))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2286, Part2(LoadInput("input_test.txt")))
}