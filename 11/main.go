package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	columnsExpanded := []string{}
	columnsExpanded = append(columnsExpanded, expanded...)

	for i :=  len(expanded[0])-1; i >= 0; i-- {
		if ColumnIsEmpty(expanded, i) {
			// fmt.Printf("empty column %d\n", i)
			columnsExpanded = InsertRuneAtIndex(columnsExpanded, i, '.')		
		}
	}
	
	starMap := MakeStarMap(columnsExpanded)
	log.Printf("%v", starMap)

	totalDistance:=0
	counter := 1
	for star1 :=0; star1 < len(starMap); star1++ {
		log.Printf("star %v", star1)		
		for star2 := star1+1; star2 < len(starMap); star2++ {
			log.Printf("star1: %v, star2: %v", starMap[star1], starMap[star2])
			distance := math.Abs(float64(starMap[star1].X - starMap[star2].X)) + math.Abs(float64(starMap[star1].Y - starMap[star2].Y))
			log.Printf("distance: %v", distance)
			totalDistance += int(distance)
			counter++
			log.Printf("Loops %v", counter)
		}		
	}
	return totalDistance
}

func Part2(data []string) int {
	starMap := MakeStarMap(data)
	log.Printf("%v", starMap)

	totalDistance:=0
	DISTANCE_MULPLIER := 2

	for star1 :=0; star1 < len(starMap); star1++ {
		log.Printf("star %v", star1)		
		for star2 := star1+1; star2 < len(starMap); star2++ {
			currentLocation := Coord{starMap[star2].X, starMap[star2].Y}
			distanceMoved := 0

			for currentX := starMap[star2].X; startingX != starMap[star1].X; startingX++ {
				// move in the x direction
				if starMap[star1].X < starMap[star2].X {
					// move right
					for currentLocation.X < starMap[star1].X {
						currentLocation.X++
						if ColumnIsEmpty(data, currentLocation.X) {
							distanceMoved += 1*DISTANCE_MULPLIER
						} else {
							distanceMoved += 1
						}
					}
				} else {
					// move left
					for currentLocation.X > starMap[star1].X {
						currentLocation.X--
						if ColumnIsEmpty(data, currentLocation.X) {
							distanceMoved += 1*DISTANCE_MULPLIER
						} else {
							distanceMoved += 1
						}
					}
				}
			} 
			for startingY := starMap[star2].Y; startingY != starMap[star1].Y; startingY++ {
				// move in the y direction
				if starMap[star1].Y < starMap[star2].Y {
					// move down
					for currentLocation.Y < starMap[star1].Y {
						currentLocation.Y++
						if RowIsEmpty(data[currentLocation.Y]) {
							distanceMoved += 1*DISTANCE_MULPLIER
						} else {
							distanceMoved += 1
						}
					}
				} else {
					// move up
					for currentLocation.Y > starMap[star1].Y {
						currentLocation.Y--
						if RowIsEmpty(data[currentLocation.Y]) {
							distanceMoved += 1*DISTANCE_MULPLIER
						} else {
							distanceMoved += 1
						}
					}
				}
			} else {
				// same star
				distanceMoved = 0
			}
			

			// log.Printf("star1: %v, star2: %v", starMap[star1], starMap[star2])
			// distance := math.Abs(float64(starMap[star1].X - starMap[star2].X)) + math.Abs(float64(starMap[star1].Y - starMap[star2].Y))
			// log.Printf("distance: %v", distance)
			// totalDistance += int(distance)
			// counter++
			// log.Printf("Loops %v", counter)
		}		
	}
	return totalDistance


	return 0
}

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

type Coord struct {
	X int
	Y int
}


func MakeStarMap (data []string) []Coord {
	starMap := []Coord{}
	for y, row := range data {
		for x, char := range row {
			if char == '#' {
				starMap = append(starMap, Coord{x, y})
			}
		}
	}
	return starMap
}

/*

	start from the original map
	step toward the next star
		start by moving in an x position
		check if the new x is in an empty column. 
			if so, distanceMoved +=1e6
			otherwise distanceMoved +=1	
*/