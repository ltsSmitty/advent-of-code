package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

// one line
// func LoadInput(fileName string) []string {
// 	input := []string{}

// 	file, err := os.ReadFile(fileName)
// 	check(err)
// 	for _, data := range string(file) {
// 		input = append(input, string(data))
// 	}
// 	return input
// }

// multiple lines
func LoadInput(fileName string) []string {
	input := []string{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		block := scanner.Text()
		input = append(input, block)

	}
	return input
}

func Part1(data []string) int {
	counter :=0;
	for _, line := range data {
		nums:= []int{}
		numStrings := strings.Split(line, " ")
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			check(err)
			nums = append(nums, num)
		}
		counter+= FindNthDifferent(nums)
	}
	return counter
}

func Part2(data []string) int {
	counter :=0;
	for _, line := range data {
		nums:= []int{}
		numStrings := strings.Split(line, " ")
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			check(err)
			nums = append(nums, num)
		}
		counter+= FindNthDifferent2(nums)
	}
	return counter
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}

func FindNthDifferent(data []int) int {
	diffs := []int{}
	for i := 0; i < len(data)-1; i++ {
		nextDiff := data[i+1] - data[i]
		diffs = append(diffs, nextDiff)
	}
	if allIntsEqual(diffs) {
		return data[len(data)-1] + diffs[0]
	}	else {
		return FindNthDifferent(diffs)+data[len(data)-1]
	}
}

func FindNthDifferent2(data []int) int {
	diffs := []int{}
	for i := 0; i < len(data)-1; i++ {
		nextDiff := data[i+1] - data[i]
		diffs = append(diffs, nextDiff)
	}
	if allIntsEqual(diffs) {
		log.Printf("Found it: %v\n", data[0]-diffs[0])
		return data[0] - diffs[0]
	}	else {
		return data[0]-FindNthDifferent2(diffs)
	}
}

func allIntsEqual (data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] != data[i+1] {
			return false
		}
	}
	return true
}