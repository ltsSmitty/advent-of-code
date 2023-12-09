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

type MapData struct {
	Left string
	Right string
}

// func Part1(data []string) int {
// 	directions := strings.Split(data[0], "")
// 	mapLines := data[2:]
// 	mapData := make(map[string]MapData, len(mapLines))
// 	for _, line := range mapLines {
// 		k := line[:3]
// 		l := line[7:10]
// 		r := line[12:15]
// 		mapData[k] = MapData{Left: l, Right: r}
// 	}
// 	log.Printf("mapData: %v", mapData)

// 	locationsEndingInA:= []string{}
// 	for l, _ := range mapData {
// 		if string(l[2]) == "A" {
// 			log.Printf("found XXA at %v", l)
// 			locationsEndingInA = append(locationsEndingInA, l)
// 		}
// 	}

// 	log.Printf("locationsEndingInA: %v", locationsEndingInA)

// 	// location := startingLocation
// 	steps :=0
// 	for i:= 0; location != "ZZZ"; i++ {
// 	// for i:= 0; i<500; i++ {
// 		if (steps % 10000 == 0) {
// 			log.Printf("steps: %v", steps)
// 			// log.Panic("too many steps")
// 		}
// 		nextTurn := directions[i%len(directions)]
// 		log.Printf("starting location: %v, turning %v", location, nextTurn)
// 		if nextTurn == "L" {
// 			location = mapData[location].Left
// 			log.Printf("new location: %v", location)
// 			} else {
// 				location = mapData[location].Right
// 				log.Printf("new location: %v", location)
// 		}
// 		steps++
// 		if (location == "ZZZ") {
// 			log.Printf("found ZZZ at %v", steps)
// 			break
// 		}
// 	}

	

// 	return steps
// }

	func Part2(data []string) int {
	directions := strings.Split(data[0], "")
	mapLines := data[2:]
	mapData := make(map[string]MapData, len(mapLines))
	for _, line := range mapLines {
		k := line[:3]
		l := line[7:10]
		r := line[12:15]
		mapData[k] = MapData{Left: l, Right: r}
	}
	log.Printf("mapData: %v", mapData)

	locationsEndingInA:= []string{}
	for l := range mapData {
		if string(l[2]) == "A" {
			log.Printf("found XXA at %v", l)
			locationsEndingInA = append(locationsEndingInA, l)
		}
	}
	log.Printf("locationsEndingInA: %v", locationsEndingInA)

	locations := locationsEndingInA
	// for i:= 0;
	counter := []int{}
	
	for _, l := range locations {
		// log.Printf(`Iteration %v`,j)
		steps :=0
		Subloop:
		for i:= 0; i<50000; i++ {

				
				nextTurn := directions[i%len(directions)]
				// log.Printf("starting location: %v, turning %v", l, nextTurn)
				if nextTurn == "L" {
					l = mapData[l].Left
					// locations[j] = mapData[l].Left
					} else {
						l= mapData[l].Right
						// locations[j] = mapData[l].Right
					}
				steps++
				if AtLeastOneLocationEndsInZZZ([]string{l}) {
					log.Printf(`found! %v steps`, steps)
					counter = append(counter, steps)
					break Subloop
					} 
				}
			}
			
	log.Printf("counter: %v", counter)
	lcm := lcmOfSlice(counter)
	return lcm
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}

func AtLeastOneLocationEndsInZZZ (locations []string) bool {
	for _, l := range locations {
		if string(l[2]) == "Z" {
			log.Printf("Found Zs %v", locations)
			return true
		}
	}
	return false
}

func LocationsAllEqualZZZ (locations []string) bool {
	for _, l := range locations {
		if string(l[2]) != "Z" {
			// log.Printf("Not Zs, %v", locations)
			return false
		}
	}
	log.Printf("All Zs, %v", locations)
	return true
}

// Function to calculate GCD (Greatest Common Divisor)
func gcd(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}

// Function to calculate LCM (Least Common Multiple)
func lcm(a, b int) int {
    return a * b / gcd(a, b)
}

// Function to find LCM of a slice of integers
func lcmOfSlice(numbers []int) int {
    result := numbers[0]
    for i := 1; i < len(numbers); i++ {
        result = lcm(result, numbers[i])
    }
    return result
}