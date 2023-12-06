package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	score := 0

	for i, line := range data {
		allCards := strings.Split(line, ":")[1]
		winningNumbers := strings.Split(strings.Split(allCards, "|")[0], " ")
		myNumbers := strings.Split(strings.Split(allCards, "|")[1], " ")
		matches := 0
		for _ , n := range winningNumbers {
			for _ , m := range myNumbers {
				// log.Printf("n: %v, m: %v", n, m)
				if n == m && n != ""{
					log.Printf(`line %d match: %v`, i, n)
					matches ++ 
				}
			}
		}
		if (matches != 0) {
			log.Printf("matches: %v", matches)
			score += int(math.Pow(2, float64(matches-1)))
		}
	}
	return score
}


func Part2(data []string) int {
	score := 0
	totalCards:= make(map[int]int)

	for lineNumber, line := range data {
		totalCards[lineNumber] += 1
		allCards := strings.Split(line, ":")[1]
		winningNumbers := strings.Split(strings.Split(allCards, "|")[0], " ")
		myNumbers := strings.Split(strings.Split(allCards, "|")[1], " ")
		matches := 0
		for _ , n := range winningNumbers {
			for _ , m := range myNumbers {
				// log.Printf("n: %v, m: %v", n, m)
				if n == m && n != ""{
					// log.Printf(`line %d match: %v`, i, n)
					matches ++ 
				}
			}
		}
		if (matches != 0) {
			log.Printf("matches: %v", matches)
			for i := 0; i < matches; i++ {
				totalCards[lineNumber+i+1] += totalCards[lineNumber]
			}

			// score += int(math.Pow(2, float64(matches-1)))
		}
	}
	log.Printf("totalCards: %v", totalCards)
	for _, v := range totalCards {
		score += v
	}
	return score
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}