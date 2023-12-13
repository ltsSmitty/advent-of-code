package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

// func TestRunRow(t *testing.T) {
// 	assert.Equal(t, 1, RunRow("???.### 1,1,3",1))
// 	assert.Equal(t, 4, RunRow(".??..??...?##. 1,1,3",1))
// 	assert.Equal(t, 1, RunRow("?#?#?#?#?#?#?#? 1,3,1,6",1))
// 	assert.Equal(t, 1, RunRow("????.#...#... 4,1,1",1))
// 	assert.Equal(t, 4, RunRow("????.######..#####. 1,6,5",1))
// 	assert.Equal(t, 10, RunRow("?###???????? 3,2,1",1))

// }

func TestRunRow1(t *testing.T) {
	assert.Equal(t, 2, RunRow("???.### 1,1,3",1))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 20, Part1(LoadInput("input_test.txt")))
}

func TestRunRowMultiplier(t *testing.T){
	// assert.Equal(t, 2, RunRow(".# 1",5))
	val := RunRow("??.??????????.????.??????????.????.??????????.????.??????????.????.??????????.? 1,1,3,1,1,1,1,3,1,1,1,1,3,1,1,1,1,3,1,1,1,1,3,1,1",1)
	assert.Equal(t, 1, val)
	// assert.Equal(t, 15, RunRow("????.#...#... 4,1,1",5))
}

func TestRunPart2(t *testing.T) {
	assert.Equal(t, 525151, Part2(LoadInput("input_test.txt")))
}