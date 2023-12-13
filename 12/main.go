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
	return 999
}

// func Part2(data []string) int {
// 	return 999
// }

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	// fmt.Printf("Part 2: %v\n", Part2(data))
	// fmt.Println(time.Since(startTime))
}

func RunRow (row string) int {
	r := strings.Split(row, " ")
	code := r[0]
	instructions := strings.Split(r[1], ",")

	numInstructions := []int{}
	for _, i := range instructions {
		num, _ := strconv.Atoi(i)
		numInstructions = append(numInstructions, num)
	}
	// log.Printf("Code: %v\n Instructions: %v", code, numInstructions)

	answer := AnalyzeString(code, numInstructions)
	log.Printf("Answer: %v", answer)
	return answer
}

func AnalyzeString (code string, instructions []int) int {
	counter :=0

	numSymbolsToProcess:=0
	for _, i := range instructions {
		numSymbolsToProcess += i
	}
	if len(code) < numSymbolsToProcess {
		log.Printf(`Impossible to finish - Not enough code to process %v symbols`, numSymbolsToProcess)
		return 0
	
	}

	if (len(code) == 0) {
		if len(instructions) == 0 {
			log.Printf(`out of code and instructions - should be a win!`)
			return 1
		} else {
			log.Printf(`out of code, but still have instructions left: %v`, instructions)
			return 0
		}
	}
	log.Printf(`Current data: %v - %v`, code, instructions)
	if string(code[0]) == "." {
		counter += AnalyzeString(code[1:], instructions)
	}
	if (string(code[0]) == "?") {
		log.Printf(`At a '?', so running with both '.' and '#'`)
		counter += AnalyzeString(string("." +code[1:]), instructions)
		counter += AnalyzeString(string("#" +code[1:]), instructions)
	}
	if (string(code[0]) == "#") {
		log.Printf(`is #`)
		if DoesStartWithNHashes(code, instructions[0]) {
			// remove the first n chars from code and the 0th instruction
			// if the next character is a ?, turn it into a .
			// otherwise keep removing #s until its ? or . 
				// and if it is a ?, turn it into a .
			log.Printf(`Code %v does start with at least %v hashes`, code, instructions[0])
			s := code[(instructions[0]):]
			log.Printf(`stripped off from the instructions, s is now %v`, s)
			// if len(code) > instructions[0] {
				Subloop:
				for i, char := range s {
					log.Printf(`Checking char %v (%v)`, i, string(char))
					switch string(char) {
						case ".":
						case "?":
							// todo this isnt quite working
							s = "." + s[i+1:]
							log.Printf("There's a ? or . at %v, so attempting to continue/break (%v)", i,s)
							break Subloop
						case "#":
							log.Printf("There's a # at %v (%v), so continuing", i,string(char))
						}
					// if string(char) != "#" {
					// 	s = "." + code[(instructions[0]+1):]
					// 	continue
					// } else {
					// 	s += string(char)
					// }
				}

			// } else {
			// 	log.Printf(`Code %v is too short to remove %v chars`, code, instructions[0])
			// 	s = code[(instructions[0]):]
			// }
			// do i need to check if the next is a ? and turn it into a . to create a new chunk?
			return counter + AnalyzeString(s, instructions[1:])
			
		} else {
			log.Printf(`Code %v doesn't start with %v hashes`, code, instructions[0])
		}		
	}
	return counter
}

func DoesStartWithNHashes(s string, n int) bool {
	log.Printf(`Trying to see if %v starts with %v hashes`, s, n)
	counter := 0
	for _, char := range s {
		if string(char) == "." {
			return false
		}
		counter++
		if counter == n {
			return true
		}
	}
	return false
}