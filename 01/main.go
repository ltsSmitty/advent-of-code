package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	counter := 0
	for _, v := range data {
		// this will be a string like pqr3stu8vwx
		// need to get 38 and then add that to the counter
		splitIntoCharacters := SplitString(v)
		intCharacters := MakeIntSlice(splitIntoCharacters)
		value := 10*intCharacters[0] + intCharacters[len(intCharacters)-1]
		counter += value
	}

	return counter
}

var numberStrings = [][]string {{"zero","0"},{"one","1"},{"two","2"},{"three","3"},{"four","4"},{"five","5"},{"six","6"},{"seven","7"},{"eight","8"},{"nine","9"}}


func Part2(data []string) int {
	counter := 0

	for _, v := range data {
		value:= GetAnswerNumberFromString(v)
		counter += value		
	}

	return counter

}

func main() {
	startTime := time.Now()
	// data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	data2:= LoadInput("input.txt")
	fmt.Printf("Part 2: %v\n", Part2(data2))
	fmt.Println(time.Since(startTime))

}

func MakeInt (s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}

func MakeIntSlice (s []string) []int {
	var i []int
	for _, v := range s {
		x, err := MakeInt(v)
		if err == nil {
			i = append(i, x)
		}
	}
	return i
}

func SplitString (s string) []string {
	return strings.Split(s, "")
}

type NumberIndex struct {
	firstIndex int
	lastIndex int
}

func GetNumberIndex (s string, numString string) NumberIndex {
	return NumberIndex{strings.Index(s, numString), strings.LastIndex(s, numString)}
}

func GetEitherNumberIndex (s string, numString1 string, numString2 string) NumberIndex {
	firstNumberIndex1 := float64(strings.Index(s, numString1))
	firstNumberIndex2 :=float64(strings.Index(s, numString2))
	var lowestIndex float64
	switch {
		case firstNumberIndex1 == -1:
			lowestIndex = firstNumberIndex2
		case firstNumberIndex2 == -1:
			lowestIndex = firstNumberIndex1
		default:
			lowestIndex = math.Min(firstNumberIndex1,firstNumberIndex2)
	}

	lastNumberIndex1 :=float64(strings.LastIndex(s, numString1))
	lastNumberIndex2 :=float64(strings.LastIndex(s, numString2))


	return NumberIndex{int(lowestIndex), int(math.Max(lastNumberIndex1,lastNumberIndex2)) }
}

func GetAnswerNumberFromString (s string) int {
	firstNumber := 0
	firstNumberIndex := 100
	lastNumber := 0
	lastNumberIndex := 0
	var err error

	log.Println(s)

	for i, numberString := range numberStrings {
		numberIndex := GetEitherNumberIndex(s, numberString[0],numberString[1])
		if numberIndex.firstIndex != -1 {
			if numberIndex.firstIndex <= firstNumberIndex {
				firstNumber, err = strconv.Atoi(numberStrings[i][1])
				firstNumberIndex = numberIndex.firstIndex
			}
		}
		if numberIndex.lastIndex != -1 {
			if numberIndex.lastIndex >= lastNumberIndex {
				lastNumber, err = strconv.Atoi(numberStrings[i][1])
				lastNumberIndex = numberIndex.lastIndex
			}
		}
	}

	if err != nil {
		log.Println(err)
	}

	value := 10*firstNumber + lastNumber
	log.Println(value)
	return value
}