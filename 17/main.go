package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

var globalInput = []string{}

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
	globalInput = input
	return input
}

func Part1(_ []string) int {
	
	startingX := len(globalInput[0]) - 1
	startingY := len(globalInput) - 1
	startingScore, _ := strconv.Atoi(string(globalInput[startingX][startingY]))
	startingCoord := CoordXYS{startingX, startingY, startingScore}

	log.Printf("Starting coord: %v", startingCoord)

	startingPath := Path{startingCoord}

	activePaths := []Path{startingPath}
	for i:=0; i<10; i++ {
		nextPaths := []Path{}
		for j, path := range activePaths {
			// take the next step in each direction
			// then add the new paths to the activePaths
			// then remove the old path from the activePaths
			// then sort the activePaths by score
			// then take the top 10
			log.Printf("Starting new loop :%d with path: %v", j, path[len(path)-1])
			newP := TakeNextSteps(path)
			log.Printf("new paths: %v", newP)
			nextPaths = append(nextPaths,newP...)
			log.Printf("next paths: %v", len(nextPaths))
		}
		activePaths = nextPaths
		log.Printf("Paths for the next loop: %v", len(activePaths))
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

// start at the bottom right. starting value is whatever is in that spot
// see what directions are available 
	// making sure not to visit already travelled squares
	// and make sure not on edge of grid
// do dynamic programming by going what's available (up/left/right/down)
// try branching for all 9 options (1x, 2x, and 3x for each direction)
// 



// idea #1
// each square has a map of how many steps it took to get there
	// and store the history of that path to get to that minimum?
	// and hold onto the 10 lowest scores?

// idea #2
// check if the current score is within some threshold (100? 500? of the lowest score at that point)
	// continue if it's within that threshold

type Direction int

const (
	North Direction = iota
	East 
	South
	West
)

type CoordXYS struct {
	x int
	y int
	score int
}

type PointHistory struct {
	// stores the 10 lowest Coord lengths
	bestCoords []CoordXYS
}

type Path  []CoordXYS

func TakeNextSteps(path Path) []Path {
	currentCoord := path[len(path)-1]
	if (currentCoord.x == 0 && currentCoord.y == 0) {
		// we're at the end
		log.Printf("Path completed. Score: %v", currentCoord.score)
		return []Path{path}
	}

	// need to determine which directions are available
	// then determine how many steps in that directoin are available

	availableDirections := map[Direction]bool{
		North: true,
		East:  true,
		South: true,
		West:  true,
	}
	
	if (currentCoord.x <= 0) {
		log.Printf("On left edge: %v", currentCoord)
		availableDirections[West] = false
	}
	if (currentCoord.y <= 0) {
		log.Printf("On top edge: %v", currentCoord)
		availableDirections[North] = false
	}
	if (currentCoord.x >= len(globalInput[0])-1) {
		log.Printf("On right edge: %v", currentCoord)
		availableDirections[East] = false
	}
	if (currentCoord.y >= len(globalInput)-1) {
		log.Printf("On bottom edge: %v", currentCoord)
		availableDirections[South] = false
	}

	nextPaths := []Path{}

	var northPath, eastPath, southPath, westPath = Path{}, Path{}, Path{}, Path{}
	err := fmt.Errorf("")

	if (availableDirections[North]) {
		northPath, err = GetNextPathInDirection(path, North); if (err != nil) {
			log.Printf("Error: %v", err)
			availableDirections[North] = false
	}}

	if (availableDirections[East]) {
		eastPath, err = GetNextPathInDirection(path, East); if (err != nil) {
			log.Printf("Error: %v", err)
			availableDirections[East] = false
	}}

	if (availableDirections[South]) {
		southPath, err = GetNextPathInDirection(path, South); if (err != nil) {
			log.Printf("Error: %v", err)
			availableDirections[South] = false
	}}

	if (availableDirections[West]) {
		westPath, err = GetNextPathInDirection(path, West); if (err != nil) {
			log.Printf("Error: %v", err)
			availableDirections[West] = false
	}}


	if (availableDirections[North]) {
		nextPaths = append(nextPaths, northPath)
	}
	if (availableDirections[East]) {
		nextPaths = append(nextPaths, eastPath)
	}
	if (availableDirections[South]) {
		nextPaths = append(nextPaths, southPath)
	}
	if (availableDirections[West]) {
		nextPaths = append(nextPaths, westPath)
	}

	// filter out the paths that are empty
	cleanPaths := []Path{}
	for i:=0; i<len(nextPaths); i++ {
		if (len(nextPaths[i]) == 0) {
			continue
		}
		cleanPaths = append(cleanPaths, nextPaths[i])
	}

	log.Printf("Clean paths: %v", cleanPaths)

	return cleanPaths
}

func GetNextPathInDirection(path Path, direction Direction) (Path, error) {
	// return how many steps are available in that direction
	// check that not at edge of grid
	// check that not in history
	// return between 0-3
	currentPoint := path[len(path)-1]
	newPath := Path{}

	if direction == North {
		if currentPoint.y-1 < 0 {
			return newPath, fmt.Errorf("north is out of bounds")
		}
		// check if the next point is in the history
		if IsCoordInPath(path, CoordXYS{currentPoint.x, currentPoint.y - 1, 0}) {
			return newPath, fmt.Errorf("already visited north")
		}

		// check that there aren't two others directly behind it
		if IsCoordInPath(path, CoordXYS{currentPoint.x, currentPoint.y + 1, 0}) && IsCoordInPath(path, CoordXYS{currentPoint.x, currentPoint.y + 2, 0}) {
			return newPath, fmt.Errorf("this is the third in a row: %v", path[(len(path)-3):])
		}

		newPath = path
		oldScore := path[len(path)-1].score
		log.Printf("Current point: %v", currentPoint)
		log.Printf("(x,y) for score: (%v,%v). Score: %v", currentPoint.y, currentPoint.x-1, string(globalInput[currentPoint.x-1][currentPoint.y]))
		newScore, _ := strconv.Atoi(string(globalInput[currentPoint.x-1][currentPoint.y]))
		log.Printf("New score: %v", newScore)
		newPath = append(newPath, CoordXYS{currentPoint.x, currentPoint.y - 1, oldScore + newScore})
		return newPath, nil
	}

	if direction == East {
		if currentPoint.x+1 >= len(globalInput[0]) {
			return newPath, fmt.Errorf("east is out of bounds")
		}
		// check if the next point is in the history
		if IsCoordInPath(path, CoordXYS{currentPoint.x + 1, currentPoint.y, 0}) {
			return newPath, fmt.Errorf("already visited east")
		}

		// check that there aren't two others directly behind it
		if IsCoordInPath(path, CoordXYS{currentPoint.x - 1, currentPoint.y, 0}) && IsCoordInPath(path, CoordXYS{currentPoint.x - 2, currentPoint.y, 0}) {
			return newPath, fmt.Errorf("this is the third in a row")
			// return newPath, fmt.Errorf("this is the third in a row: %v", path[(len(path)-3):])
		}

		newPath = path
		oldScore := path[len(path)-1].score
		// log.Printf("Current point: %v", currentPoint)
		// log.Printf("(x,y) for score: (%v,%v). Score: %v", currentPoint.y, currentPoint.x+1, string(globalInput[currentPoint.x+1][currentPoint.y]))
		newScore, _ := strconv.Atoi(string(globalInput[currentPoint.x+1][currentPoint.y]))
		// log.Printf("New score: %v", newScore)
		newPath = append(newPath, CoordXYS{currentPoint.x + 1, currentPoint.y, oldScore + newScore})
		return newPath, nil
	}

	if direction == South {
		if currentPoint.y+1 >= len(globalInput) {
			return newPath, fmt.Errorf("south is out of bounds")
		}
		// check if the next point is in the history
		if IsCoordInPath(path, CoordXYS{currentPoint.x, currentPoint.y + 1, 0}) {
			return newPath, fmt.Errorf("already visited south")
		}

		// check that there aren't two others directly behind it
		if IsCoordInPath(path, CoordXYS{currentPoint.x, currentPoint.y - 1, 0}) && IsCoordInPath(path, CoordXYS{currentPoint.x, currentPoint.y - 2, 0}) {
			return newPath, fmt.Errorf("this is the third in a row: %v", path[(len(path)-3):])
		}

		newPath = path
		oldScore := path[len(path)-1].score
		// log.Printf("Current point: %v", currentPoint)
		// log.Printf("(x,y) for score: (%v,%v). Score: %v", currentPoint.y+1, currentPoint.x, string(globalInput[currentPoint.x][currentPoint.y+1]))
		newScore, _ := strconv.Atoi(string(globalInput[currentPoint.x][currentPoint.y+1]))
		// log.Printf("New score: %v", newScore)
		newPath = append(newPath, CoordXYS{currentPoint.x, currentPoint.y + 1, oldScore + newScore})
		return newPath, nil
	}

	if direction == West {
		if currentPoint.x-1 < 0 {
			return newPath, fmt.Errorf("west is out of bounds")
		}
		// check if the next point is in the history
		if IsCoordInPath(path, CoordXYS{currentPoint.x - 1, currentPoint.y, 0}) {
			return newPath, fmt.Errorf("already visited west")
		}

		// check that there aren't two others directly behind it
		if IsCoordInPath(path, CoordXYS{currentPoint.x + 1, currentPoint.y, 0}) && IsCoordInPath(path, CoordXYS{currentPoint.x + 2, currentPoint.y, 0}) {
			return newPath, fmt.Errorf("this is the third in a row: %v", path[(len(path)-3):])
		}

		newPath = path
		oldScore := path[len(path)-1].score
		// log.Printf("Current point: %v", currentPoint)
		// log.Printf("(x,y) for score: (%v,%v). Score: %v", currentPoint.y, currentPoint.x-1, string(globalInput[currentPoint.x-1][currentPoint.y]))
		newScore, _ := strconv.Atoi(string(globalInput[currentPoint.x-1][currentPoint.y]))
		// log.Printf("New score: %v", newScore)
		newPath = append(newPath, CoordXYS{currentPoint.x - 1, currentPoint.y, oldScore + newScore})
		return newPath, nil
	}

	return newPath, nil
}

func IsCoordInPath(path Path, coord CoordXYS) bool {
	for _, v := range path {
		if (v.x == coord.x && v.y == coord.y) {
			return true
		}
	}
	return false
}