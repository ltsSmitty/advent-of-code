package main

import (
	"bufio"
	"fmt"
	"log"
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
	RED_LIMIT := 12
	GREEN_LIMIT := 13
	BLUE_LIMIT := 14
	counter := 0
	for _, line := range data {
		success := true
		parts := strings.Split(line, ":")
		gameId := GetGameId(parts[0])
		pulls := strings.Split(parts[1], ";")
		for _, pull := range pulls {
			p := GetNumColors(pull)
			log.Printf("Game %d: %+v", gameId, p)
			
			if p.Red > RED_LIMIT || p.Blue > BLUE_LIMIT || p.Green > GREEN_LIMIT {
				success = false
				break
			}
		}
		if success {
			counter+=gameId
			log.Printf("Game %d passed", gameId)
		} else {
			log.Printf("Game %d failed", gameId)
		}
	log.Printf("counter: %d", counter)
	}
	return counter
}

func Part2(data []string) int {
	counter := 0
	for _, line := range data {
		redMin := 0
		greenMin := 0
		blueMin := 0
		parts := strings.Split(line, ":")
		gameId := GetGameId(parts[0])
		pulls := strings.Split(parts[1], ";")
		for _, pull := range pulls {
			p := GetNumColors(pull)
			log.Printf("Game %d: %+v", gameId, p)
			
			if p.Red > redMin {
				log.Printf("redMin increasing from %d to %d", redMin, p.Red)
				redMin = p.Red
			}
			if p.Green > greenMin {
				log.Printf("greenMin increasing from %d to %d", greenMin, p.Green)
				greenMin = p.Green
			}
			if p.Blue > blueMin {
				log.Printf("blueMin increasing from %d to %d", blueMin, p.Blue)
				blueMin = p.Blue
			}
		}
		product := redMin*greenMin*blueMin
		log.Printf("redMin: %d, greenMin: %d, blueMin: %d, product: %d", redMin, greenMin, blueMin, product)
		counter += (blueMin*greenMin*redMin)
	}
	return counter
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}

func GetGameId(game string) int {
	var gameId int
	fmt.Sscanf(game, "Game %d:", &gameId)
	return gameId
}

type Pull struct {
	Red int
	Blue int
	Green int
}

// e.g. "3 blue, 4 red"
func GetNumColors(pull string) Pull {
	pulls := strings.Split(pull, ",")
	red:=0; blue:=0; green:=0
	for _, p := range pulls {
		// "3 blue"
		var numColors int
		var inputColor string
		fmt.Sscanf(p, "%d %s", &numColors, &inputColor)
		switch inputColor {
			case "red":
				red += numColors
			case "blue":
				blue += numColors
			case "green":
				green += numColors
		}
	}

	// log.Printf("red: %d, blue: %d, green: %d", red, blue, green)

	return Pull{Red: red, Blue: blue, Green: green}
}