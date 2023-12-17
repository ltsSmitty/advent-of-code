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
	counter :=0
	words := strings.Split(data[0],",")

	for _, word := range words {
	v := ProcessWord(word)
	log.Printf("v: %v", v)
	counter+=v
	}
	return counter
}

type Operation string

const (
	ADD Operation = "add"
	REMOVE Operation = "remove"
)

type Instruction struct {
	box int
	op Operation
	lens Lens
}

type Lens struct {
	label string
	value int
}

func Part2(data []string) int {
	boxes := make(map[byte][]Lens)
	words := strings.Split(data[0],",")
	for _, word := range words {
		instruction := ProcessWord2(word)
		box := instruction.box
		lens := instruction.lens

		if instruction.op == REMOVE {
			// go to the box and remove the lens with the same label if there is one
			lensesInBox := boxes[byte(box)]
			for i, l := range lensesInBox {
				if l.label == lens.label {
					// remove the lens with the same value from lens
					boxes[byte(box)] = append(boxes[byte(box)][:i], boxes[byte(box)][i+1:]...)
				}
			}
		} else if instruction.op == ADD {
			// see if a lens with the same label exists
			// if it does, update the value
			// if it doesn't, add it
			lensesInBox := boxes[byte(box)]
			found := false
			for i, l := range lensesInBox {
				if l.label == lens.label {
					// update the value
					boxes[byte(box)][i].value = lens.value
					found = true
				}
			}
			if !found {
				boxes[byte(box)] = append(boxes[byte(box)], lens)
			}
		}
	}
	log.Printf("Final boxes")
	for k, v := range boxes {
		log.Printf("box %v: %v", k, v)
	}

	// calculate the score 
	counter := 0
	for boxNumber, lenses := range boxes {
		for indexInBox, lens := range lenses {
			
			boxVal := int(boxNumber) +1
			slotVal := indexInBox +1
			focalLength := lens.value
			product := boxVal * slotVal * focalLength
			log.Printf("boxVal: %v, slotVal: %v, focalLength: %v, product %v", boxVal, slotVal, focalLength, product)
			
			counter += product
		}
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

func ProcessCharByte(b byte) int {
	// log.Printf("b: %v", b)
	v := b * 17
	return int(v)
}

func ProcessWord (word string) int {
	log.Printf("word: %v", word)
	v := 0
	for _, b := range word {
		v+=int(b)
		log.Printf("v: %v (%v)", v,b)
		val := ProcessCharByte(byte(v))
		log.Printf("val: %v", val)

		v=val
	}
	return int(byte(v))
}

func ProcessWord2 (word string) Instruction {
	label := word[0:2]

	v := 0
	for _, b := range label {
		v+=int(b)
		val := ProcessCharByte(byte(v))
		v=val
	}

	command := word[2:]
	value := -1
	if len(command) ==2 {
		// log.Printf("command: %v", command[1])
		value,_ = strconv.Atoi(string(command[1]))
	} 
	
	op := Operation("")
	if string(command[0]) == "=" {
		op = ADD
	} else {
		op = REMOVE
	}

	i := int(byte(v))

	return Instruction{i, op, Lens{label, value}}
}