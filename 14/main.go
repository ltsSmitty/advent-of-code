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

func RollRocksInDirection (data []string, direction string) []string {
	newData := data
	// log.Printf("Rolling rocks in direction %v\n", direction)
	if direction == "north" || direction == "west" {
		for i, row := range data {
			for j, col := range row {
				if string(col) == "O" {
					// it's not hitting this branch for some reason
					if direction == "north" {
						newPos := RollRockNorth(newData, Coord{j,i})
						newData[i] = newData[i][:j] + "." + newData[i][j+1:]
						newData[newPos.y] = newData[newPos.y][:newPos.x] + "O" + newData[newPos.y][newPos.x+1:]
					} else {
						newPos := RollRockWest(newData, Coord{j,i})
						newData[i] = newData[i][:j] + "." + newData[i][j+1:]
						newData[newPos.y] = newData[newPos.y][:newPos.x] + "O" + newData[newPos.y][newPos.x+1:]
					}
				}
			}
		} 
	} else if direction == "south" || direction == "east" {
			// start the loops at the max and work down
			for i:=len(data)-1; i>=0; i-- {
				for j:=len(data[i])-1; j>=0; j-- {
					if string(data[i][j]) == "O" {
						if direction == "south" {
							newPos := RollRockSouth(newData, Coord{j,i})
							newData[i] = newData[i][:j] + "." + newData[i][j+1:]
							newData[newPos.y] = newData[newPos.y][:newPos.x] + "O" + newData[newPos.y][newPos.x+1:]
						} else {
							newPos := RollRockEast(newData, Coord{j,i})
							newData[i] = newData[i][:j] + "." + newData[i][j+1:]
							newData[newPos.y] = newData[newPos.y][:newPos.x] + "O" + newData[newPos.y][newPos.x+1:]
						}}
					}
				}
		}

		return newData	
}

func FindWeight (data []string) int {
	weight := 0
	for i, row := range data {
		for _, col := range row {
			if string(col) == "O" {
				weight += len(data) - i
			}
		}
	}
	return weight
}
				
func Part1(data []string) int {
	counter :=0
	rolledData := data
	for i, row := range data {
		log.Printf("Row %v: %v\n", i, row)
		for j, col := range row {
			if string(col) == "O" {
				newPos := RollRockNorth(rolledData, Coord{j,i})
				rolledData[i] = rolledData[i][:j] + "." + rolledData[i][j+1:]
				rolledData[newPos.y] = rolledData[newPos.y][:newPos.x] + "O" + rolledData[newPos.y][newPos.x+1:]
				weight := len(data) -newPos.y
				counter+= weight
			}
		}	
	}

	log.Printf("Final weight of this map: %v\n", FindWeight(rolledData))

	for _, line := range rolledData {
		fmt.Println(line)
	}
	return counter
}

func Part2(data []string) int {
	rolledData := data
	rockMap := make(map[string]int)

	NUM_LOOPS := int(1e9)
	key := MakeKey(rolledData)
	cycleLength := 0
	cycleStartIndex := 0


	for i:=1; i<=NUM_LOOPS; i++ {
		log.Printf("Loop %v\n\n", i)
		// check if this key exists in the map

		if v, ok := rockMap[key]; ok {
			if (v == 1 && cycleStartIndex == 0) {
				log.Printf("Found first duplicate key at %d",i)
				cycleStartIndex = i
			}
			if (v == 2 && cycleLength == 0) {
				log.Printf("Found the second duplicate key at %d",i)
				// we've found a cycle
				cycleLength = i-cycleStartIndex
				log.Printf("Found a cycle of length %v starting at %v\n", cycleLength, cycleStartIndex)

				// instead of breaking, increase i by almost the whole of num loops minus the mod of the cycle length?

				linesLeft := NUM_LOOPS - i
				remainder := linesLeft % cycleLength
				rockMap[key] = rockMap[key] + 1
				i = NUM_LOOPS - remainder  - 1
				log.Printf("i is now %v\n", i)
				log.Printf("dont see this twice")
				continue
			}
			log.Printf("Found a duplicate key %d: %v",i,v)
			log.Printf("Weight of this map: %v\n", FindWeight(rolledData))
			rockMap[key] = rockMap[key] + 1
			} else {
				rockMap[key] = 1
			}

		rolledData = RollRocksInDirection(rolledData, "north")
		rolledData = RollRocksInDirection(rolledData, "west")
		rolledData = RollRocksInDirection(rolledData, "south")
		rolledData = RollRocksInDirection(rolledData, "east")
		key = MakeKey(rolledData)
	}
	w := FindWeight(rolledData)
	log.Printf("Final weight of this map: %v\n", w)
	return FindWeight(rolledData)
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}

type Coord struct {
	x int
	y int
}

func RollRockNorth (data []string, startingPos Coord) Coord {
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

func RollRockSouth (data []string, startingPos Coord) Coord {
	startingChar := string(data[startingPos.y][startingPos.x])
	if startingChar != "O" {
		log.Printf("There's no rock at %v (actually %v)\n", startingPos,startingChar)
		return startingPos
	}
	newPos := startingPos
	for i:=startingPos.y+1; i<len(data); i++ {
		v := string(data[i][startingPos.x])
		if v == "." {
			newPos.y = i
		} else {
			return newPos
		}
	}
	return newPos
}

func RollRockEast (data []string, startingPos Coord) Coord {
	startingChar := string(data[startingPos.y][startingPos.x])
	if startingChar != "O" {
		log.Printf("There's no rock at %v (actually %v)\n", startingPos,startingChar)
		return startingPos
	}
	newPos := startingPos
	for i:=startingPos.x+1; i<len(data[startingPos.y]); i++ {
		// log.Printf("Searching for a floor at %v,%v. Printing current state of map:\n", startingPos.x, i)
		// for _, line := range data {
		// 	fmt.Println(line)
		// }
		v := string(data[startingPos.y][i])
		// log.Printf("Found %v\n", v)		
		if v == "." {
			newPos.x = i
		} else {
			return newPos
		}
	}
	return newPos
}

func RollRockWest (data []string, startingPos Coord) Coord {
	startingChar := string(data[startingPos.y][startingPos.x])
	if startingChar != "O" {
		log.Printf("There's no rock at %v (actually %v)\n", startingPos,startingChar)
		return startingPos
	}
	newPos := startingPos
	for i:=startingPos.x-1; i>=0; i-- {
		// log.Printf("Searching for a floor at %v,%v. Printing current state of map:\n", startingPos.x, i)
		// for _, line := range data {
		// 	fmt.Println(line)
		// }
		v := string(data[startingPos.y][i])
		// log.Printf("Found %v\n", v)		
		if v == "." {
			newPos.x = i
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

func MakeKey (data []string) string {
	key := ""
	for _, line := range data {
		key += line
	}
	return key
}