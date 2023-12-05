package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestGetNumberIndices(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestExtractVals(t *testing.T) {
	assert.Equal(t, []Val{{58, 7, 0}}, ExtractVals([]string{".....+.58."}))
}

func TestGetIndiciesToCheckForSymbol(t *testing.T) {
	assert.Equal(t, []Coord{{6,0},{9,0}}, GetIndiciesToCheckForSymbol(Val{58, 7, 1}))
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 4361, Part1(LoadInput("input_test.txt")))
}

func TestEmptyStringAtoi(t *testing.T) {
	num, _ := strconv.Atoi("")
	assert.Equal(t, 0, num)
}

func TestPart2(t *testing.T) {	
	assert.Equal(t, 467835, Part2(LoadInput("input_test.txt")))
}

func TestGetValFromDigit (t *testing.T) {
	data:=[]string{"467..114..","...*......","..35..633."}
	assert.Equal(t, Val{35,2,2}, GetValFromDigit(StringCoord{"3",2,2},&data))
}