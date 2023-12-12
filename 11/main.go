package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	expanded := []string{}

	// expand each row
	for _, line := range data {
		if RowIsEmpty(line) {
			fmt.Println("empty row")
			expanded = append(expanded, line)
			expanded = append(expanded, line)
		} else {
		expanded = append(expanded, line)
		}
	}

	log.Printf("Expanded rows")
	for _, line := range expanded {
		log.Printf("%v", line)
	}

	// expand each column
	// if column is empty, expand to 2 columns
	columnsExpanded := []string{}
	columnsExpanded = append(columnsExpanded, expanded...)

	for i :=  len(expanded[0])-1; i >= 0; i-- {
		if ColumnIsEmpty(expanded, i) {
			fmt.Printf("empty column %d\n", i)
			columnsExpanded = InsertRuneAtIndex(columnsExpanded, i, '.')		
		}
	}
	

	log.Printf("Expanded columns")
	for _, line := range columnsExpanded {
		log.Printf("%v", line)
	}

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

func RowIsEmpty (row string) bool {
	for _, char := range row {
		if char == '#' {
			return false
		}
	}
	return true
}

func ColumnIsEmpty (data []string, column int) bool {
	for _, row := range data {
		if row[column] == '#' {
			return false
		}
	}
	return true
}


func InsertRuneAtIndex(data []string, index int, r rune) []string {
	for i := range data {
		if index >= 0 && index < len(data[i]) {
			data[i] = data[i][:index] + string(r) + data[i][index:]
		}
	}
	return data
}
