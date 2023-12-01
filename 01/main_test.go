package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, []string{
		"1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet"}, LoadInput("input_test.txt"))
}

// create a function that splits a string into a slice of strings (e.g characters => []string{"c", "h", "a", "r", "a", "c", "t", "e", "r", "s"})
func TestSplitString (t *testing.T) {
	assert.Equal(t, []string{"1","a","b","c","2"}, SplitString("1abc2"))
	}


// create a function that takes a slice of strings and returns a slice of ints, checking which are ints and rejecting the rest
func TestMakeIntSlice (t *testing.T) {
	assert.Equal(t, []int{1, 2}, MakeIntSlice([]string{"1","a","b","c","2"}))
	} 
	
func TestGetNumberIndex (t *testing.T) {
	assert.Equal(t, NumberIndex{firstIndex:0,lastIndex:0}, GetNumberIndex("two1nine","two"))
	assert.Equal(t, NumberIndex{firstIndex:3,lastIndex:3}, GetNumberIndex("two1nine","1"))
	assert.Equal(t, NumberIndex{firstIndex:0,lastIndex:8}, GetNumberIndex("two1ninetwo","two"))
	assert.Equal(t, NumberIndex{firstIndex:4,lastIndex:4}, GetNumberIndex("two1nine","nine"))
	}

func TestGetEitherNumberIndex (t *testing.T) {
	assert.Equal(t, NumberIndex{firstIndex:0,lastIndex:0}, GetEitherNumberIndex("two1nine","two","2"))
	assert.Equal(t, NumberIndex{firstIndex:3,lastIndex:3}, GetEitherNumberIndex("two1nine","1","one"))
	assert.Equal(t, NumberIndex{firstIndex:2,lastIndex:9}, GetEitherNumberIndex("tw2o1ninetwo","two","2"))
	assert.Equal(t, NumberIndex{firstIndex:4,lastIndex:4}, GetEitherNumberIndex("two1nine","nine","9"))
	}

	func TestGetAnswerNumberFromString(t *testing.T) {
		assert.Equal(t, 19, GetAnswerNumberFromString("onetwo1nine"))
	}