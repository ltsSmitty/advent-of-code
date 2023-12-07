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

func TestCalculateWaysToRace(t *testing.T) {
	assert.Equal(t,[]int{1,2,3} , CalculateWaysToRace(Race{7, 9}))
}

func TestRunPart1 (t *testing.T) {
	assert.Equal(t, 288, Part1(LoadInput("input_test.txt")))
}