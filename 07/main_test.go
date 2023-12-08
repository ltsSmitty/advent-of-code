package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 6441, Part1(LoadInput("input_test.txt")))
}

func TestCompareHands(t *testing.T) {
	v, _ := CompareHands(Hand{"TTTT6", 540, 6}, Hand{"TTJTT", 450, 6})
	assert.Equal(t, 0, v)
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 5905, Part2(LoadInput("input_test.txt")))
}