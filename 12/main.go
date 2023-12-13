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

var instructionMap InstructionMap = make(map[CodeInstruction]int)

func Part1(data []string) int {
	counter := 0
	multiplier:= 1

	for _, row := range data {
		counter += RunRow(row, multiplier)
	}
	return counter
}

type CodeInstruction struct {
	code string
	instructionsFlat int
}

type InstructionMap map[CodeInstruction]int


func Part2(data []string) int {
	counter := 0
	multiplier:= 5
	for i, row := range data {
		log.Printf("Row :%d",i)
		counter += RunRow(row, multiplier)
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

func RunRow (row string, multiplier int ) int {
	r := strings.Split(row, " ")
	code := r[0]
	instructions := strings.Split(r[1], ",")

	numInstructions := []int{}
	for _, i := range instructions {
		num, _ := strconv.Atoi(i)
		numInstructions = append(numInstructions, num)
	}
	log.Printf("Original Code: %v  Instructions: %v", code, numInstructions)

	multipleCode := code
	multipleInstructions := numInstructions
	for i:=1; i<multiplier; i++ {
		multipleCode = multipleCode + "?" + code
		multipleInstructions = append(multipleInstructions, numInstructions...)
	}

	log.Printf("Multiple Code: %v  Instructions: %v", multipleCode, multipleInstructions)

	answer := AnalyzeString(multipleCode, multipleInstructions)
	log.Printf("Answer: %v", answer)
	return answer
}

func AnalyzeString (code string, instructions []int ) int {
	// log.Printf("Code: %v  Instructions: %v", code, instructions)

	if val := LoadFromMap(code, instructions); val != -1 {
		// log.Printf("Found [%v : %v] in map",code, instructions)
		return val
	}

	// numSymbolsToProcess:= len(instructions)-1
	numSymbolsToProcess:= 0
	for _, i := range instructions {
		numSymbolsToProcess += i
	}
	if len(code) < numSymbolsToProcess {
		log.Printf("Saving false - not enough symbols to process. Code: %v  Instructions: %v ", code, instructions)
		SaveToMap(code, instructions, 0)
		return 0	
	}

	if len(instructions) == 0 {
		if strings.Count(code, "#") > 0 {
			log.Printf("Saving false - out of instructions but more #s remain. Code: %v  Instructions: %v ", code, instructions)
			SaveToMap(code, instructions, 0)
			return 0
		}
		log.Printf("Saving true - no more #s left. Code: %v  Instructions: %v ", code, instructions)
		SaveToMap(code, instructions, 1)
		return 1
	}

	if (len(code) == 0) {
		if len(instructions) == 0 {
			log.Printf("Saving true- code and instructions are empty. Code: %v  Instructions: %v", code, instructions)
			SaveToMap(code, instructions, 1)
			return 1
			} else {
			log.Printf("Saving false - code empty but instructions remain. Code: %v  Instructions: %v", code, instructions)
			SaveToMap(code, instructions, 0)
			return 0
		}
	}
	if string(code[0]) == "." {
		return AnalyzeString(code[1:], instructions)
	}

	if (string(code[0]) == "?") {
		value := AnalyzeString(string(code[1:]), instructions) + AnalyzeString(string("#" +code[1:]), instructions)
		log.Printf("Saving in ? after branch. Code: %v  Instructions: %v  Value: %v", code, instructions, value)
		SaveToMap(code, instructions, value)
		return value
	}

	if (string(code[0]) == "#") {
		if DoesStartWithNHashes(code, instructions[0]) {
			s := code[(instructions[0]):]
			trimmedInstructions := instructions[1:]

			if len(s) == 0 {
				v := IsAtGoodEndpoint(s, trimmedInstructions)
				log.Printf("Saving at terminus. Code: %v  Instructions: %v  Value: %v", code, instructions, v)
				SaveToMap(code, instructions, v)
				return v
			}

			switch (string(s[0])) {
				case "?":
					s = "." + s[1:]
				case "#":
					for string(s[0]) == "#" {
						s = s[1:]
					}
					if string(s[0]) == "?" {
						s = "." + s[1:]
					}
				// default aka '.', do nothing
			}
			// do i need to save here
			val:= AnalyzeString(s, trimmedInstructions)	
			log.Printf("Saving in #. Code: %v  Instructions: %v  Value: %v", code, instructions, val)
			return val	
		} else {
			log.Printf("Saving false - code starts with # but wrong amount for instructions. Code: %v  Instructions: %v", code, instructions)
			SaveToMap(code, instructions, 0)
			return 0
		
		}
	}
	return 0
}

func DoesStartWithNHashes(s string, n int) bool {
	counter := 0
	isSuccess := false
	for _, char := range s {
		if string(char) == "." {
			return false
		}
		counter++
		if counter == n {
			isSuccess = true
			break
		}
	}
	// make sure the next character isn't a #
	if isSuccess {
		if len(s) == counter {
			return true
		}
		if string(s[counter]) == "#" {
			return false
		} else {
			return true
		}
	}
	return false
}

func IsAtGoodEndpoint (code string, instructions []int) int {
	// log.Printf(`Checking that %v is empty`,instructions)
	if len(code) == 0 && len(instructions) == 0 {
		// log.Printf(`It is!`)
		return 1
	}
	// log.Printf(`It isn't!`)
	return 0
}

func compressIntSliceIntoInt (instructions []int) int {
	// log.Printf(`Compressing %v into an int`, instructions)
	counter := 0
	for _, i := range instructions {
		counter = counter*10 + i
	}
	return counter
}

func SaveToMap(code string, instructions []int, value int) {
	num := compressIntSliceIntoInt(instructions)
	instructionMap[CodeInstruction{code, num}] = value
}

func LoadFromMap(code string, instructions []int) int {
	num := compressIntSliceIntoInt(instructions)
	ci := CodeInstruction{code, num}
	if val, ok := (instructionMap)[ci]; ok {
		return val
	}
	return -1
}