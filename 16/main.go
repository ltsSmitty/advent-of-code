package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
	walkHistory := make(map[Position][]Direction)
	startingDirection := E
	startingPosition := Position{x: -1, y: 0}
	beam := Beam{
		direction: startingDirection, 
		position: startingPosition, 
		history: []WalkHistory{{
			direction: startingDirection,
			position: startingPosition}}}

	beams := []Beam{beam}

	for i:=0; (i<10000 && len(beams) != 0); i++ {
		newBeams := []Beam{}
		log.Printf("current independent beams: %d", len(beams))

		for i, beam := range beams {
			if v, ok := walkHistory[beam.position]; !ok {
				walkHistory[beam.position] = []Direction{beam.direction}
			} else {

				if slices.Contains(v, beam.direction) {
					log.Printf("Beam %d has been here (at %v) before, and in this direction (%v)",i,beam.position,beam.direction)			
					if (len(beams) > i) {
						beams = append(beams[:i], beams[i+1:]...)
					} else {
						beams = beams[:i]
					}
				} else {
					log.Printf("Beam %d has been here (at %v) before, but not in this direction (%v)",i,beam.position,beam.direction)
					walkHistory[beam.position] = append(walkHistory[beam.position], beam.direction)
				}
			}
		}

		for _, beam := range beams {
			beamsFromThisLoop := WalkBeam(beam, data)
			log.Printf("beamsFromThisLoop %d:",i )
			for i, beam := range beamsFromThisLoop {
				log.Printf("Beam %d: %+v, %v\n\n",i, beam.position,beam.direction)
			}


			newBeams = append(newBeams, beamsFromThisLoop...)
		}

		if len(newBeams) == 0 {
			log.Printf("No more beams to walk")
			// return 999
		}
		
		beams = newBeams
	}
	
	log.Printf("Ending with %d beams", len(beams))

	// print out the graph

	graph := make([][]string, len(data))

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			if v, ok := walkHistory[Position{x: j, y: i}]; ok {
				if len(v) == 1 {
					graph[i] = append(graph[i], string(v[0]))
				} else {
					graph[i] = append(graph[i], fmt.Sprintf("%d", len(v)))
				}
			} else {
				graph[i] = append(graph[i], ".")
			}
		}
	}

	for _, row := range graph {
		fmt.Println(row)
	}

	// for v := range walkHistory {
		// log.Printf("walkHistory: %v", v)
	// }
	return len(walkHistory)-1
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

type Direction string

const (
	N Direction = "N"
	S Direction = "S"
	E Direction = "E"
	W Direction = "W"
)

type Position struct {
	x int
	y int
}

type WalkHistory struct {
	position Position
	direction Direction
}

type Beam struct {
	direction Direction
	position  Position
	history   []WalkHistory
}


func FindNextPosition (beam Beam) Position {
	nextLocation := Position{}
	switch beam.direction {
	case N:
		nextLocation = Position{x: beam.position.x, y: beam.position.y - 1}
	case S:
		nextLocation = Position{x: beam.position.x, y: beam.position.y + 1}
	case E:
		nextLocation = Position{x: beam.position.x + 1, y: beam.position.y}
	case W:
		nextLocation = Position{x: beam.position.x - 1, y: beam.position.y}
	}
	return nextLocation
}


func WalkBeam (beam Beam, data []string) []Beam {
	nextLocation := FindNextPosition(beam)
	log.Printf("Original location: %v, nextLocation: %+v", beam.position, nextLocation)

	// check if the nextLocation is valid in data
	// if valid, see what kind of item is in it
	if nextLocation.x < 0 || nextLocation.y < 0 || nextLocation.x >= len(data[0]) || nextLocation.y >= len(data) {
		log.Printf("beam leaving the grid")
		return []Beam{}
	}

	nextVal := string(data[nextLocation.y][nextLocation.x])
	nextDirection := Direction("")

	log.Printf("nextVal: %v", nextVal)
	switch nextVal {
		case "/": {
			switch beam.direction {
				case N: nextDirection = E
				case E: nextDirection = N
				case S: nextDirection = W
				case W: nextDirection = S
			}
			return []Beam{{direction: nextDirection, position: nextLocation, history: append(beam.history, WalkHistory{position: nextLocation, direction: nextDirection})}}
		}
		case `\`: {
			switch beam.direction {
				case N: nextDirection = W
				case E: nextDirection = S
				case S: nextDirection = E
				case W: nextDirection = N
			}
			return []Beam{{direction: nextDirection, position: nextLocation, history: append(beam.history, WalkHistory{position: nextLocation, direction: nextDirection})}}
		}
		case "|":{
			if beam.direction == N || beam.direction == S {
				return []Beam{{direction: beam.direction, position: nextLocation, history: append(beam.history, WalkHistory{position: nextLocation, direction: beam.direction})}}
			} else {
				log.Printf("splitting beam")
				wh1 := WalkHistory{position: nextLocation, direction: N}
				wh2 := WalkHistory{position: nextLocation, direction: S}

				newWh1 := make([]WalkHistory, len(beam.history))
				newWh2 := make([]WalkHistory, len(beam.history))
				copy(newWh1, beam.history)
				copy(newWh2, beam.history)

				newWh1 = append(newWh1, wh1)
				newWh2 = append(newWh2, wh2)

				b2 := Beam{direction: S, position: nextLocation, history: newWh1}
				b1 := Beam{direction: N, position: nextLocation, history: newWh2}
				newBeams := []Beam{b1,b2}

				return newBeams
			}
		}
		case "-":{
			if beam.direction == E || beam.direction == W {
				return []Beam{{direction: beam.direction, position: nextLocation, history: append(beam.history, WalkHistory{position: nextLocation, direction: beam.direction})}}
				} else {
					// split into two beams, with one going east and one going west
					log.Printf("splitting beam")
					wh1 := WalkHistory{position: nextLocation, direction: E}
					wh2 := WalkHistory{position: nextLocation, direction: W}

					newWh1 := make([]WalkHistory, len(beam.history))
					newWh2 := make([]WalkHistory, len(beam.history))
					copy(newWh1, beam.history)
					copy(newWh2, beam.history)

					newWh1 = append(newWh1, wh1)
					newWh2 = append(newWh2, wh2)

					b2 := Beam{direction: W, position: nextLocation, history: newWh1}
					b1 := Beam{direction: E, position: nextLocation, history: newWh2}
					newBeams := []Beam{b1,b2}

						
					return newBeams
				}
		}
		case ".":{
			return []Beam{{direction: beam.direction, position: nextLocation, history: append(beam.history, WalkHistory{position: nextLocation, direction: beam.direction})}}
		}
	}
	log.Panic("hmm returning nil value")
	return []Beam{}
}