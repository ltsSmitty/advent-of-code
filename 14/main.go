package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	rolledData := data
	for i, row := range data {
		log.Printf("Row %v: %v\n", i, row)
		for j, col := range row {
			if string(col) == "O" {
				// fmt.Printf("Found a rock at %v,%v\n", j,i)
				newPos := RollRock(rolledData, Coord{j,i})
				// fmt.Printf("New pos: %v\n", newPos)
				rolledData[i] = rolledData[i][:j] + "." + rolledData[i][j+1:]
				rolledData[newPos.y] = rolledData[newPos.y][:newPos.x] + "O" + rolledData[newPos.y][newPos.x+1:]
				weight := len(data) -newPos.y
				counter+= weight
			}
		}	
	}

	for _, line := range rolledData {
		fmt.Println(line)
	}
	return counter
}

type RowsMap map[int][]int
type ColumnsMap map[int][]int


func Part2(data []string) int {
	rowsMap := make(RowsMap)
	columnsMap := make(ColumnsMap)

	// fill the maps wherever there is a non . character
	for i, row := range data {
		for j, col := range row {
			if string(col) != "." {
				rowsMap[i] = append(rowsMap[i], j)
				columnsMap[j] = append(columnsMap[j], i)
			}
		}
	}
	// log.Printf("Rows map: %v\n", rowsMap)
	// log.Printf("Columns map: %v\n", columnsMap)

	rockPositions := FindRockPositions(data)
	rockRowsMap := make(RowsMap)
	rockColumnsMap := make(ColumnsMap)
	for _, rock := range rockPositions {
		rockRowsMap[rock.y] = append(rockRowsMap[rock.y], rock.x)
		rockColumnsMap[rock.x] = append(rockColumnsMap[rock.x], rock.y)
	}
	// log.Printf("Rock rows map: %v\n", rockRowsMap)
	// log.Printf("Rock columns map: %v\n", rockColumnsMap)

	newCoords, newMap := SlideRocksVertical(rockPositions, rockColumnsMap, columnsMap, "north")

	log.Printf("New coords: %v\nNew map: %v", newCoords,newMap)

	return 999
}

func SlideRocksVertical(rockPositions []Coord, rockColumnsMap ColumnsMap, columns ColumnsMap ,direction string) ([]Coord, ColumnsMap) {
	if direction == "north" {
		// start with the rock with the lowest y, and move it up
		// check the column map for the next y value which is less than it
		// if that exists, set the new y value to one more that that
		// otherwise set the new y value to 0

		// todo if this is too slow, maybe can sort less frequently

		// make sure rockColumnsMap is sorted
		for _, column := range rockColumnsMap {
			sort.Ints(column)
		}

		// make sure columns is sorted
		for _, column := range columns {  // e.g. [2,4,7]
			sort.Ints(column)
		}

		// newRockColumnsMap := make(ColumnsMap)
		// newColumnMap := make(ColumnsMap)

		// go through one column of rockColumnsMap at a time
		for i, rockColumn := range rockColumnsMap { // e.g. [4,7]
			newRockCol := rockColumn
			// newCol := []int{}
			log.Printf("\tAll blockage in column %d: %v\n", i, columns[i])
			if (len(rockColumn) == 0) {
				log.Printf("No rocks in column %v\n", i)
			}
			// go through each rock in the column
			// ! this one is the specific north part
			for _, rock := range rockColumn { // e.g 4
				log.Printf("Rocks in column %v: %v. Next rock: %v\n", i,rockColumn, rock)

				// go through each blockage in the column and see if its above the rock
				// if it is, stop the rock one above it?
				Subloop:
				for j :=0; j<len(columns[i]); j++ { // e.g. 2, 4, 7
					log.Printf("j: %v",j)
					r := columns[i][j] // e.g. 2
					newValue := rock
					log.Printf("Blocker at %v, current rock to slide (r) at %v\n", r, rock)

					if (j ==0 && r == newValue) { // this is the first obstruction in the row, so set it to 0
						log.Printf("This rock is the first obstruction in the row, so rolling the rock to 0\n")
						newValue = 0
						newRockCol = append(append(rockColumn[:j],newValue), rockColumn[j+1:]...)
						log.Printf("New rock column: %v\n", newRockCol)
						break Subloop
					}

					if r < rock {
						newValue = r+1
						log.Printf("%v < %v\n", r, rock)
						log.Printf("Sliding up from %v to %v\n", rock, newValue)
					} else if (r >= rock) {
						log.Printf("r >= rock: %v > %v, done sliding\n", r, rock)
						// need to append it in at the right spot
						log.Printf("Trying to slide a rock forward")
						log.Printf("rockColumn[:j]: %v", rockColumn[:j])
						log.Printf("newValue: %v", newValue)
						log.Printf("rockColumn[j:]: %v", rockColumn[j:])
						newRockCol = append(append(rockColumn[:j],newValue), rockColumn[j:]...)
						log.Printf("New rock column: %v\n", newRockCol)
						continue Subloop
					}
				}					
			}
		}
			






	} 
	// else if direction == "south" {
	// }
	// for each column, find the lowest rock
	// for each rock, find the lowest floor
	// move the rock to the lowest floor
	// repeat until no more rocks can be moved
	// return the new data
	return rockPositions, rockColumnsMap
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	// fmt.Printf("Part 2: %v\n", Part2(data))
	// fmt.Println(time.Since(startTime))

}

type Coord struct {
	x int
	y int
}

func RollRock (data []string, startingPos Coord) Coord {
	startingChar := string(data[startingPos.y][startingPos.x])
	if startingChar != "O" {
		log.Printf("There's no rock at %v (actually %v)\n", startingPos,startingChar)
		return startingPos
	}
	newPos := startingPos
	for i:=startingPos.y-1; i>=0; i-- {
		// log.Printf("Searching for a floor at %v,%v. Printing current state of map:\n", startingPos.x, i)
		// for _, line := range data {
		// 	fmt.Println(line)
		// }
		v := string(data[i][startingPos.x])
		// log.Printf("Found %v\n", v)		
		if v == "." {
			newPos.y = i
		} else {
			return newPos
		}
	}
	return newPos
}

func FindRockPositions (data []string) []Coord {
	rocks := []Coord{}
	for i, row := range data {
		for j, col := range row {
			if string(col) == "O" {
				rocks = append(rocks, Coord{j,i})
			}
		}
	}
	return rocks
}
