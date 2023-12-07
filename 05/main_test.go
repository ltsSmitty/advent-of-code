package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestTranslateInstructions(t *testing.T) {
instructions := []Instruction{{50, 98, 2},{52, 50, 48}}
	assert.Equal(t, 81, TranslateInstructions(79,instructions))
	assert.Equal(t, 14, TranslateInstructions(14,instructions))
	assert.Equal(t, 57, TranslateInstructions(55,instructions))
	assert.Equal(t, 13, TranslateInstructions(13,instructions))
}

func TestGetSeedNumbers(t *testing.T) {
	assert.Equal(t, []int{79, 14, 55, 13}, GetSeedNumbers("seeds: 79 14 55 13"))
}

func TestLoadSections(t *testing.T) {
	assert.Equal(t, 1, LoadSections(LoadInput("input_test.txt")))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 35, Part1(LoadInput("input_test.txt")))
}

func TestGetSeedRanges(t *testing.T) {
	assert.Equal(t, []Range{{79,92},{55,67}}, GetSeedRanges("seeds: 79 14 55 13"))
}

func TestSortRanges(t *testing.T) {
	assert.Equal(t, []Range{{55,67},{79,92}}, SortRanges([]Range{{79,92},{55,67}}))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 46, Part2(LoadInput("input_test.txt")))
}