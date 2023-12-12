package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadInput(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestFollowPipe(t *testing.T) {
	log.Printf(`testing |`)
	nextPipe := "|"
	res, err := FollowNextPipe(Coord{1,1},Coord{1,2}, nextPipe)
	assert.Nil(t, err)
	assert.Equal(t, Coord{1,3}, res)
}

func TestFollowL (t *testing.T) {
	log.Printf(`testing L`)
	nextPipe := "L"
	res, err := FollowNextPipe(Coord{1,1},Coord{1,2}, nextPipe)
	assert.Nil(t, err)
	assert.Equal(t, Coord{2,2}, res)
}

func TestFollowDash (t *testing.T) {
	log.Printf(`testing -`)
	nextPipe := "-"
	res, err := FollowNextPipe(Coord{1,1},Coord{2,1}, nextPipe)
	assert.Nil(t, err)
	assert.Equal(t, Coord{3,1}, res)
}


func TestFollowJ (t *testing.T) {
	log.Printf(`testing J`)
	nextPipe := "J"
	res, err := FollowNextPipe(Coord{1,1},Coord{1,2}, nextPipe)
	assert.Nil(t, err)
	assert.Equal(t, Coord{0,2}, res)
}

func TestFollow7 (t *testing.T) {
	log.Printf(`testing 7`)
	nextPipe := "7"
	res, err := FollowNextPipe(Coord{1,1},Coord{2,1}, nextPipe)
	assert.Nil(t, err)
	assert.Equal(t, Coord{2,2}, res)
}

func TestFollowF (t *testing.T) {
	log.Printf(`testing F`)
	nextPipe := "F"
	res, err := FollowNextPipe(Coord{2,1},Coord{1,1}, nextPipe)
	assert.Nil(t, err)
	assert.Equal(t, Coord{1,2}, res)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 10, Part1(LoadInput("input_test.txt")[18:]))
}