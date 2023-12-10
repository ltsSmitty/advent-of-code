package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestFindNthDifference(t *testing.T) {
	assert.Equal(t, 18, FindNthDifferent([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 28, FindNthDifferent([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 68, FindNthDifferent([]int{10, 13, 16, 21, 30, 45}))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 114, Part1(LoadInput("input_test.txt")))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2, Part2(LoadInput("input_test.txt")))
}