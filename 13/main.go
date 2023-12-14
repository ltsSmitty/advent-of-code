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

type Puzzle []string

func Part1(data []string) int {
	
	puzzles := []Puzzle{}
	p:= Puzzle{}
	for _, line := range data {
		if line == "" {
			puzzles = append(puzzles, p)
			p = Puzzle{}
		} else {
			p = append(p, line)
		}
	}
	puzzles = append(puzzles, p) // do a final append for the last puzzle

	counter:=0
	for i, puzzle := range puzzles {
		matchingRow, matchingColumn := -1, -1
		matchingRow = FindMatchingPuzzleRow(puzzle)
		matchingColumn = FindMatchingPuzzleColumn(puzzle)
		// log.Printf("matchingRow: %v, matchingColumn: %v", matchingRow, matchingColumn)

		if (matchingRow == -1 && matchingColumn == -1) {
			log.Printf("no match found for puzzle %d", i)
			continue
		}

		if matchingRow != -1 {
			counter+= 100*matchingRow
		} 
		if matchingColumn != -1 {
			counter += matchingColumn
		}
	}


	return counter
}

func FindMatchingPuzzleRow (puzzle Puzzle) int {
	MainLoop:
	for i := 0; i < len(puzzle)-1; i++ {
		if puzzle[i] == puzzle[i+1] {
			// continue stepping outward to make sure the altering rows match too
			// see if i is closer to the top or bottom
			numRowsToCheck := 0
			if i < len(puzzle)/2 {
				numRowsToCheck = i
			} else {
				numRowsToCheck = len(puzzle)-i-2
			}
			for j := 1; j <= numRowsToCheck; j++ {
				if puzzle[i-j] != puzzle[i+1+j] {
					continue MainLoop
				}
			}
			return i+1
		}
	}
	return -1
}

func FindMatchingPuzzleRow2 (puzzle Puzzle) int {
	MainLoop:
	for i := 0; i < len(puzzle)-1; i++ {
		foundSmudge := false
		numDifferent := NumDifferentBytes([]byte(puzzle[i]), []byte(puzzle[i+1]))

		// two options: look for starting rows that are identical and look for a single smudge walking away
		// or look for starting rows that are one different and look for exactness to the end
		// need to do both. how to do both?

		if numDifferent == 0 {
			// need to check that the rows are identical or 1 different. if they're different and havent found smudge, mark smudge and continue
			// else that's a second smudge so continue back to MainLoop
			numRowsToCheck := 0
			if i < len(puzzle)/2 {
				numRowsToCheck = i
			} else {
				numRowsToCheck = len(puzzle)-i-2
			}
			for j := 1; j <= numRowsToCheck; j++ {
				n := NumDifferentBytes([]byte(puzzle[i-j]), []byte(puzzle[i+1+j]))
				if n == 0 {
					// continue
					// log.Printf(`Rows %d and %d are identical: %v vs %v`, i-j, i+1+j, puzzle[i-j], puzzle[i+1+j])
				} else if n == 1 {
					if !foundSmudge {
						log.Printf("Smudge found at row %d and %d: %v vs %v", i-j, i+1+j, puzzle[i-j], puzzle[i+1+j])
						foundSmudge = true
					} else {
						continue MainLoop
					}
				} else {
					continue MainLoop
				}
			}
			if foundSmudge {				
			return i+1 
		} else {
			continue MainLoop
		}

		} else if numDifferent == 1 {
			log.Printf("Smudge found at row %d and %d: %v vs %v", i, i+1, puzzle[i], puzzle[i+1])
			// already got the smudge out of the way, so walk the rest like from part 1
			// continue stepping outward to make sure the altering rows match too
			// see if i is closer to the top or bottom
			numRowsToCheck := 0
			if i < len(puzzle)/2 {
				numRowsToCheck = i
			} else {
				numRowsToCheck = len(puzzle)-i-2
			}
			for j := 1; j <= numRowsToCheck; j++ {
				if puzzle[i-j] != puzzle[i+1+j] {
					continue MainLoop
				}
			}
			return i+1

		} else { 
			continue MainLoop
		}
	}
	return -1
}

func FindMatchingPuzzleColumn (puzzle Puzzle) int {
	MainLoop:
	for i := 0; i < len(puzzle[0])-1; i++ {
		col1 := GetColumnFrom2dSlice(puzzle, i)
		col2 := GetColumnFrom2dSlice(puzzle, i+1)
		if ByteSlicesMatch(col1, col2) {
			log.Printf("col %d: %v, col %d: %v", i,string(col1),i+1, string(col2))
			// continue stepping outward to make sure the alternating cols match too
			// see if i is closer to the left or right
			numColsToCheck := 0
			if i < len(puzzle[0])/2 {
				numColsToCheck = i
			} else {
				numColsToCheck = len(puzzle[0])-i-2
			}

			for j := 1; j <= numColsToCheck; j++ {
				col1 := GetColumnFrom2dSlice(puzzle, i-j)
				col2 := GetColumnFrom2dSlice(puzzle, i+1+j)
				if !ByteSlicesMatch(col1, col2) {
					log.Printf("columns stop matching at %d and %d:\n%v\n%v", i-j, i+1+j,col1,col2)
					continue MainLoop
				}
			}
			return i+1
		}
	}
	return -1
}
func FindMatchingPuzzleColumn2 (puzzle Puzzle) int {
	MainLoop:
	for i := 0; i < len(puzzle[0])-1; i++ {
		foundSmudge := false
		col1 := GetColumnFrom2dSlice(puzzle, i)
		col2 := GetColumnFrom2dSlice(puzzle, i+1)
		numDifferent := NumDifferentBytes(col1, col2) 
		
		if numDifferent == 0 {
			// need to check that the rows are identical or 1 different. if they're different and havent found smudge, mark smudge and continue
			// else that's a second smudge so continue back to MainLoop
			numColsToCheck := 0
			if i < len(puzzle[0])/2 {
				numColsToCheck = i
			} else {
				numColsToCheck = len(puzzle[0])-i-2
			}

			for j := 1; j <= numColsToCheck; j++ {
				col1 := GetColumnFrom2dSlice(puzzle, i-j)
				col2 := GetColumnFrom2dSlice(puzzle, i+1+j)
				n := NumDifferentBytes(col1, col2)

				if n == 0 {
					// continue
				} else if n == 1 {
					if !foundSmudge {
						log.Printf("Smudge found at col %d and %d: %v vs %v", i-j, i+1+j, col1, col2)
						foundSmudge = true
					} else {
						continue MainLoop
					}
				} else {
					continue MainLoop
				}
			}
			if foundSmudge {
				// make sure we've found the smudge
				return i+1
			} else {
				continue MainLoop
			}
		} else if numDifferent == 1 {
			// already got the smudge out of the way, so walk the rest like from part 1
			// continue stepping outward to make sure the altering rows match too
			// see if i is closer to the top or bottom
			numColsToCheck := 0
			if i < len(puzzle[0])/2 {
				numColsToCheck = i
			} else {
				numColsToCheck = len(puzzle[0])-i-2
			}

			for j := 1; j <= numColsToCheck; j++ {
				col1 := GetColumnFrom2dSlice(puzzle, i-j)
				col2 := GetColumnFrom2dSlice(puzzle, i+1+j)
				if !ByteSlicesMatch(col1, col2) {
					log.Printf("columns stop matching at %d and %d:\n%v\n%v", i-j, i+1+j,col1,col2)
					continue MainLoop
				}
			}
			return i+1
		} else {
			continue MainLoop
		}
	}
	return -1
}

func GetColumnFrom2dSlice (slice []string, column int) []byte {
	col := []byte{}
	for i := 0; i < len(slice); i++ {
		v := (slice[i][column])
		col = append(col, v)
	}
	return col
}

func ByteSlicesMatch (slice1 []byte, slice2 []byte) bool {
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func NumDifferentBytes (slice1 []byte, slice2 []byte) int {
	diff := 0
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			diff++
		}
	}
	return diff
}

func Part2(data []string) int {
	
	puzzles := []Puzzle{}
	p:= Puzzle{}
	for _, line := range data {
		if line == "" {
			puzzles = append(puzzles, p)
			p = Puzzle{}
		} else {
			p = append(p, line)
		}
	}
	puzzles = append(puzzles, p) // do a final append for the last puzzle

	counter:=0
	for i, puzzle := range puzzles {
		matchingRow, matchingColumn := -1, -1
		matchingRow = FindMatchingPuzzleRow2(puzzle)
		matchingColumn = FindMatchingPuzzleColumn2(puzzle)
		// log.Printf("matchingRow: %v, matchingColumn: %v", matchingRow, matchingColumn)

		if (matchingRow == -1 && matchingColumn == -1) {
			log.Printf("no match found for puzzle %d", i)
			continue
		}

		if matchingRow != -1 {
			counter+= 100*matchingRow
		} 
		if matchingColumn != -1 {
			counter += matchingColumn
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