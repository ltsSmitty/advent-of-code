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

func Part1(data []string) int {
	startingCoord := FindS(data)
	pipeCoords := FollowPipe(data, startingCoord)
	cleanMap := TransformOthersToBlanks(data, pipeCoords)
	for _, row := range cleanMap {
		log.Printf("%v", row)
	}

	// do a horizontal expand
	expandedMap := []string{}
	for _, row := range cleanMap {
		newRow := ""
		for _, char := range row {
			r := HorizontalExpand(string(char))
			newRow += r
		}
		expandedMap = append(expandedMap, newRow)
	}

	for _, row := range expandedMap {
		log.Printf("%v", row)
	}

	// do a vertical expand
	expandedMap2 := []string{}
	for i:=0; i<len(expandedMap); i++ {
		newRow1 := ""
		newRow2 := ""
		for j:=0; j<len(expandedMap[i]); j++ {
			char1, char2 := VerticalExpand(string(expandedMap[i][j]))
			newRow1 += char1
			newRow2 += char2
		}
		expandedMap2 = append(expandedMap2, newRow1)
		expandedMap2 = append(expandedMap2, newRow2)
	}

	log.Printf("Fully expanded map:")
	for _, row := range expandedMap2 {
		log.Printf("%v", row)
	}
	
	split := [][]string{}
	for _, row := range expandedMap2 {
		splitRow := []string{}
		for _, char := range row {
			splitRow = append(splitRow, string(char))
		}
		split = append(split, splitRow)
	}
	
	
	// now we have a fully expanded map, we can do a flood fill
	flooded := FloodFill(split, Coord{0,0})
	log.Printf("Flooded map:")
	for _, row := range flooded {
		expandedMap2[row.Y] = expandedMap2[row.Y][:row.X] + "X" + expandedMap2[row.Y][row.X+1:]
		log.Printf("%v", expandedMap2[row.Y])
	}
	
	log.Printf("Fully expanded map:")
	for _, row := range expandedMap2 {
		log.Printf("%v", row)
	}

	// remove all the odd rows	
	originalSizedMap := []string{}
	for i:=0; i<len(expandedMap2); i++ {
		if i%2 == 0 {
			originalSizedMap = append(originalSizedMap, expandedMap2[i])
		}
	}

	// remove all the odd columns
	for i:=0; i<len(originalSizedMap); i++ {
		newRow := ""
		for j:=0; j<len(originalSizedMap[i]); j++ {
			if j%2 == 0 {
				newRow += string(originalSizedMap[i][j])
			}
		}
		originalSizedMap[i] = newRow
	}

	log.Printf("Original sized map:")
	for _, row := range originalSizedMap {
		log.Printf("%v", row)
	}

	// loop through and count remaining dots
	count := 0
	for _, row := range originalSizedMap {
		for _, char := range row {
			if string(char) == "." {
				count++
			}
		}
	}

	return count
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

// start at S, put its coords somewhere.
// walk all 4 directions
	// if the direction makes sense, walk it
		// and keep walking until either it doesn't make sense or you get back to the starting point
		// if you get back to the starting point, you have a loop
	// if the direction doesn't make sense, don't walk it
		// if you get back to the starting point, you have a loop

var MapSymbol = map[string]string{
	".": "empty",
	"|": "vertical",
	"-": "horizontal",
	"L": "left",
	"J": "right",
	"7": "up",
	"F": "down",
}

type ApproachingX string

const (
	Left  ApproachingX = "fromLeft"
	Right ApproachingX = "fromRight"
)

type ApproachingY string

const (
	Up   ApproachingY = "fromAbove"
	Down ApproachingY = "fromBelow"
)

type PipeMap map[Coord]string

type Coord struct {
	X int
	Y int
}

func FollowNextPipe (previousCoord Coord, currentCoord Coord, currentPipe string) (Coord, error) {
	approachingX := ApproachingX("")
	approachingY := ApproachingY("")
	currentX := currentCoord.X
	currentY := currentCoord.Y

	if previousCoord.X > currentCoord.X {
		approachingX = Right
	} else if previousCoord.X < currentCoord.X {
		approachingX = Left
	} else if previousCoord.Y > currentCoord.Y {
		approachingY = Down
	} else if previousCoord.Y < currentCoord.Y {
		approachingY = Up
	} else {
		return Coord{}, fmt.Errorf("both coords are the same")
	}
	
	switch currentPipe {
	case "|":
		if approachingY == "fromAbove" {
			return Coord{currentX,currentY+1}, nil
		}
		if approachingY == "fromBelow" {
			return Coord{currentX,currentY-1}, nil
		}
		return Coord{}, fmt.Errorf("invalid approach")
	case "-":
		if approachingX == "fromLeft" {
			return Coord{currentX +1,currentY}, nil			
		}
		if approachingX == "fromRight" {
			return Coord{currentX-1,currentY}, nil
		}
		return Coord{}, fmt.Errorf("invalid approach")
	case "L":
		if approachingY == "fromAbove" {
			return Coord{currentX+1,currentY}, nil
		}
		if approachingX == "fromRight" {
			return Coord{currentX,currentY-1}, nil
		}
		return Coord{}, fmt.Errorf("invalid approach")
	case "J":
		if approachingY == "fromAbove" {
			return Coord{currentX-1,currentY}, nil
		}
		if approachingX == "fromLeft" {
			return Coord{currentX,currentY-1}, nil
		}
		return Coord{}, fmt.Errorf("invalid approach")
	case "7":
		if approachingY == "fromBelow" {
			return Coord{currentX-1,currentY}, nil
		}
		if approachingX == "fromLeft" {
			return Coord{currentX,currentY+1}, nil
		}
		return Coord{}, fmt.Errorf("invalid approach")
	case "F":
		if approachingY == "fromBelow" {
			return Coord{currentX+1,currentY}, nil
		}
		if approachingX == "fromRight" {
			return Coord{currentX,currentY+1}, nil
		}
		return Coord{}, fmt.Errorf("invalid approach")
	default:
		return Coord{}, fmt.Errorf("invalid pipe")
	}
}

func FollowPipe (data []string, startingCoord Coord ) PipeMap {
	pipeMap := PipeMap{}
	log.Printf("startingCoord: %v", startingCoord)
	log.Printf("data: %v", data)
	startingCoordOptions := []Coord{
		{0, 1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}
	log.Printf("startingCoordOptions: %v", startingCoordOptions)

	for _, coord := range startingCoordOptions {
		log.Printf("starting coord modifier: %v", coord)
		firstCoord := startingCoord
		secondCoord := Coord{startingCoord.X+coord.X, startingCoord.Y+coord.Y}
		pipeMap[firstCoord] = string(data[firstCoord.Y][firstCoord.X])
		log.Printf(`firstCoord: %v, secondCoord: %v`, firstCoord, secondCoord)
		Subloop:
		for i:=0; i<100000; i++ {
			log.Printf("current pipe is %v", string(data[secondCoord.Y][secondCoord.X]))
			newCoord, err := FollowNextPipe(firstCoord, secondCoord, string(data[secondCoord.Y][secondCoord.X]))
			if string(data[secondCoord.Y][secondCoord.X]) == "S" {
				log.Printf("found S at %v", secondCoord)
				return pipeMap
			}
			if err != nil {
				log.Printf("error: %v", err)
				break Subloop
			}
			firstCoord = secondCoord
			secondCoord = newCoord
			pipeMap[firstCoord] = string(data[firstCoord.Y][firstCoord.X])
			log.Printf("firstCoord: %v, secondCoord: %v", firstCoord, secondCoord)
		}
	}
	return pipeMap

}

func FindS (data []string) Coord {
	for y, row := range data {
		for x, char := range row {
			if string(char) == "S" {
				return Coord{x,y}
			}
		}
	}
	return Coord{}
}

func TransformOthersToBlanks (data []string, realPipes PipeMap) []string {
	cols := len(data[0])
	rows := len(data)

	// create a blank map filled with .	
	blankMap := []string{}
	for i:=0; i<rows; i++ {
		blankMap = append(blankMap, "")
		for j:=0; j<cols; j++ {
			blankMap[i] += "."
		}
	}

	// loop through realPipes and replace the blankMap with the realPipes
	for coord, pipe := range realPipes {
		blankMap[coord.Y] = blankMap[coord.Y][:coord.X] + pipe + blankMap[coord.Y][coord.X+1:]
	}

	return blankMap
}

// walk the maze, saving ghe coords of each pipe including the S
// turn everything else into .
// now walk the graph, from left to right, creating a double sized graph for each char it runs into

func HorizontalExpand(s string) (string) {
	switch s {
	case "S":
		return "S-"
	case "|":
		return "|."
	case "-":
		return "--"
	case "L":
		return "L-"
	case "J":
		return "J."
	case "7":
		return "7."
	case "F":
		return "F-"
	default:
		return ".."
	}
}

func VerticalExpand (s string) (string, string) {
	switch s {
	case "S":
		return "S", "|"
	case "|":
		return "|", "|"
	case "-":
		return "-", "."
	case "L":
		return "L", "."
	case "J":
		return "J", "."
	case "7":
		return "7", "|"
	case "F":
		return "F", "|"
	default:
		return ".", "."
	}
}


var mods = [...]struct {
    X, Y int
}{
    {-1, 0}, {1, 0}, {0, -1}, {0, 1},
}


func FloodFill(graph [][]string, origin Coord) []Coord {
    val := graph[origin.Y][origin.X]

    seen := make([][]bool, len(graph))
    for i, row := range graph {
        seen[i] = make([]bool, len(row))
    }

    // let go sort out the appended size.
    fill := []Coord{}

    // go will shuffle memory too when adding/removing items from q
    q := []Coord{origin}

    for len(q) > 0 {

        // shift the q
        op := q[0]
        q = q[1:]

        if seen[op.Y][op.X] {
            continue
        }

        seen[op.Y][op.X] = true
        fill = append(fill, op)

        for _, mod := range mods {
            newx := op.X + mod.X
            newy := op.Y + mod.Y
            if 0 <= newy && newy < len(graph) && 0 <= newx && newx < len(graph[newy]) {
                if graph[newy][newx] == val {
                    q = append(q, Coord{newx, newy})
                }
            }
        }
    }
    return fill
}