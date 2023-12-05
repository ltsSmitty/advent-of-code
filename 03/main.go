package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	hasSymbol := false
	counter := 0
	activeNum := []string{}

	for row, rowData := range data {
		for column, columnData := range rowData {
			// check if it's a number
			log.Printf("Row: %s, element: %s \nActive number: %s, hasSymbol: %t", rowData, string(columnData), activeNum, hasSymbol)
			
			if numValue, err := strconv.Atoi(string(columnData)); err == nil {
				activeNum = append(activeNum, string(columnData))

				if !hasSymbol {
					// check for hasSymbol by looking at the 8 around for a char that's not a . or a number
					// if it's not a number or a ".", then it's a symbol
					// if it's a symbol, then we can stop checking
					indices := GetIndiciesToCheckForSymbol(Val{Val: numValue, X: column, Y: row})
					Subloop:
					for _, index := range indices {
						if index.Y < len(data) && index.X < len(data[0]) {
							if _, err := strconv.Atoi(string(data[index.Y][index.X])); err != nil && string(data[index.Y][index.X]) != "." {
								hasSymbol = true
								break Subloop
							}
						}					
					}
			}
		} else {
			// it's not a number.
			// so if hasSymbol, then we can add it to the counter
			// then reset hasSymbol and activeNum
			if hasSymbol {
				log.Printf("Adding %d to counter", turnArrayOfNumberStringsIntoNDigitNumber(activeNum))
				counter += turnArrayOfNumberStringsIntoNDigitNumber(activeNum)
				log.Printf("Counter is now %d", counter)
			}
			hasSymbol = false
			activeNum = []string{}
		}
	
	}

	}
	return counter
}

func Part2(data []string) int {
	counter := 0
	
	for row, rowData := range data {
		for column, columnData := range rowData {
			// if it's a "*", get the 8 around it
			validNumberVals := []Val{}
			if string(columnData) == "*" {
				numberCoords := []StringCoord{}
				var leftmost, rightMost, topMost, bottomMost int
				if column > 0 {
					leftmost = column - 1
				} else {
					leftmost = 0
				}
				if column < len(data[0]) {
					rightMost = column + 1
				} else {
					rightMost = len(data[0])
				}
				if row > 0 {
					topMost = row - 1
				} else {
					topMost = 0
				}
				if row < len(data) {
					bottomMost = row + 1
				} else {
					bottomMost = len(data)
				}
				// log.Printf("Leftmost: %d, rightMost: %d, topMost: %d, bottomMost: %d", leftmost, rightMost, topMost, bottomMost)

				for i := leftmost; i <= rightMost; i++ {
					for j := topMost; j <= bottomMost; j++ {
						numberCoords = append(numberCoords, StringCoord{string(data[j][i]),i, j})
						
					}
				}
				log.Printf("Number coords: %v", numberCoords)
				// check if there are two unique numbers from the coords
				// if there are, then add the product of the two numbers to the counter
				// if there aren't, then do nothing
	
				for _, numberCoord := range numberCoords {
					if _, err := strconv.Atoi(numberCoord.Val); err == nil {
						validNumberVals=append(validNumberVals, GetValFromDigit(numberCoord, &data))
					}
				}
			}
			// log.Printf("Valid number vals: %v", validNumberVals)

			// Create a map to store the unique values
			uniqueNumbers := make(map[int]Val)
			// Iterate over the slice of Person structs
			for _, v := range validNumberVals {
				// If the name is not already in the map, add it
				if _, ok := uniqueNumbers[v.Val]; !ok {
					uniqueNumbers[v.Val] = Val{Val: v.Val, X: v.X, Y: v.Y}
				}
			}
			if len(uniqueNumbers) == 2 {
				multiple :=1
				for _, v := range uniqueNumbers {
					multiple *= v.Val
					log.Printf("Unique numbers: %v", uniqueNumbers)
				}
				counter+=multiple
				log.Printf("Counter is now %d", counter)
			}
		}

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

type StringCoord struct {
	Val string
	X int
	Y int
}

type Val struct {
	Val int
	X int
	Y int
}

func SplitRemovingEmpty(s string, sep string) []string {
	result := []string{}
	substrs := strings.Split(s, sep)
	for _, c := range substrs {
		if c != "" {
			// split every character and then append together each number
			num:=[]string{}
			for _, v := range c {
				if _, err := strconv.Atoi(string(v)); err == nil {
					result = append(num, string(v))
				}
			}

			result = append(result, strings.Join(num, ""))
			log.Printf("Result : %s",result)
		}
	}
	return result
}

func RemoveNonNumeric(s []string) []int {
	result := []int{}
	for _, c := range s {
		if i, err := strconv.Atoi(c); err == nil {
			result = append(result, i)
		}
	}
	return result
}

func ExtractVals(s []string) []Val {
	result := []Val{}
	for rowIndex, rowValue := range s {
		substrs := SplitRemovingEmpty(rowValue, ".")
		nums:= RemoveNonNumeric(substrs)

		for i, c := range nums {
			log.Printf("Chunk %d: %d", i, c)
			index := strings.Index(rowValue, strconv.Itoa((c)))
			result = append(result, Val{Val: c, X: index,Y: rowIndex})
		}
	}
	return result
}

type Coord struct {
	X int
	Y int
}

func GetIndiciesToCheckForSymbol(val Val) []Coord {
	result := []Coord{}
	maxLeft := math.Max(float64(val.X - 1),0)
	// unable to validate out of range right
	maxRight := float64( val.X + len(strconv.Itoa(val.Val)))
	minY := math.Max(float64(val.Y - 1),0)
	maxY := float64(val.Y + 1)

	for i := maxLeft; i <= maxRight; i++ {
		for j := minY; j <= maxY; j++ {
			// exclude the values we are checking
			if int(i) >= int(maxLeft)+1 && int(i) <= int(maxRight-1) && int(j) == val.Y {
				// log.Println("skipping")
			} else {
				result = append(result, Coord{X: int(i), Y: int(j)})
			}
		}
	}
	return result
}

func turnArrayOfNumberStringsIntoNDigitNumber(arrayOfNumbers []string) int {
	// Check if the array is empty.
	if len(arrayOfNumbers) == 0 {
		return 0
	}

	// Create a string builder.
	builder := strings.Builder{}

	// Iterate over the array and append each number to the string builder.
	for _, number := range arrayOfNumbers {
		builder.WriteString( number)
	}

	// Return the string builder.
	num, err := strconv.Atoi(builder.String())
	check(err)
	return num
}


func GetValFromDigit(val StringCoord, data *[]string) Val {
	// walk both sides of the val appendinding the numbers
	// then return the leftmost coordinate and the val
	leftmost :=-1
	numberString := []string{}
	currentX := val.X
	var currentXValue byte
	// move left
	// log.Printf("CurrentX: %d", currentX)
	Subloop:
	for currentX>=0 {
	currentXValue = (*data)[val.Y][currentX]
	// log.Printf("CurrentXValue: %s", string(currentXValue))
	if _, err := strconv.Atoi(string(currentXValue)); err == nil {
		leftmost = currentX
		numberString = append([]string{string(currentXValue)}, numberString...)
	} else {
		break Subloop
	}
	currentX--
	}
	// move right
	currentX = val.X+1
	Subloop2:
	for currentX<len((*data)[0]) {
	currentXValue = (*data)[val.Y][currentX]
	if _, err := strconv.Atoi(string(currentXValue)); err == nil {
		numberString = append(numberString, string(currentXValue))
	} else {
		break Subloop2
	}
	currentX++
}
// log.Printf("Number string: %v", numberString)
return Val{Val: turnArrayOfNumberStringsIntoNDigitNumber(numberString), X: leftmost, Y: val.Y}
}