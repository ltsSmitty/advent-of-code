package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
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
	seedRow := data[0]
	seedNumbers := GetSeedNumbers(seedRow)
	plantingLocations := []int{}
	sections := LoadSections(data[1:])

	for _, seedNumber := range seedNumbers {
		// translate seed to soil to fertilizer to water to light to temperature to humidity to location
		soil := TranslateInstructions(seedNumber, sections["seed-to-soil"])
		fertilizer := TranslateInstructions(soil, sections["soil-to-fertilizer"])
		water := TranslateInstructions(fertilizer, sections["fertilizer-to-water"])
		light := TranslateInstructions(water, sections["water-to-light"])
		temperature := TranslateInstructions(light, sections["light-to-temperature"])
		humidity := TranslateInstructions(temperature, sections["temperature-to-humidity"])
		location := TranslateInstructions(humidity, sections["humidity-to-location"])
		plantingLocations = append(plantingLocations, location)
	}
	return slices.Min(plantingLocations)
}

func ProcessIntructions(initialRanges []Range, instructions []Instruction) []Range{

	seedRanges := []Range{}
	i := initialRanges

	for k:=0; k< len(i); k++ {
		r := i[k]
		res := TranslateRangeInstructions(r, instructions)
		seedRanges = append(seedRanges, res.TranslatedRange)
		if res.RemainingRanges!=nil {
			log.Printf("adding remainingRanges: %v", res.RemainingRanges)
		}
		i = append(i, res.RemainingRanges...)
	}
	log.Printf("seedRanges: %v", seedRanges)
	return seedRanges
}

func Part2(data []string) int {
	seedRow := data[0]
	log.Printf("\nGetting Seed Ranges")
	seedRanges := GetSeedRanges(seedRow)
	log.Printf("\nLoading Sections")
	sections := LoadSections(data[1:])

	
	log.Printf("\nTranslating Seed Ranges for Seed-to-Soil\n")
	soilRanges := ProcessIntructions(seedRanges, sections["seed-to-soil"])
	log.Printf("\ntranslatedRanges after seed-to-soil: %v", soilRanges)

	log.Printf("\nTranslating Seed Ranges for Soil-to-Fertilizer\n")
	fertilizerRanges:= ProcessIntructions(soilRanges, sections["soil-to-fertilizer"])
	log.Printf("\ntranslatedRanges after soil-to-fertilizer: %v", fertilizerRanges)

	log.Printf("\nTranslating Seed Ranges for Fertilizer-to-Water\n")
	waterRanges:= ProcessIntructions(fertilizerRanges, sections["fertilizer-to-water"])
	log.Printf("\ntranslatedRanges after fertilizer-to-water: %v", waterRanges)

	log.Printf("\nTranslating Seed Ranges for Water-to-Light\n")
	lightRanges:= ProcessIntructions(waterRanges, sections["water-to-light"])
	log.Printf("\ntranslatedRanges after water-to-light: %v", lightRanges)

	log.Printf("\nTranslating Seed Ranges for Light-to-Temperature\n")
	temperatureRanges:= ProcessIntructions(lightRanges, sections["light-to-temperature"])
	log.Printf("\ntranslatedRanges after light-to-temperature: %v", temperatureRanges)

	log.Printf("\nTranslating Seed Ranges for Temperature-to-Humidity\n")
	humidityRanges:= ProcessIntructions(temperatureRanges, sections["temperature-to-humidity"])
	log.Printf("\ntranslatedRanges after temperature-to-humidity: %v", humidityRanges)

	log.Printf("\nTranslating Seed Ranges for Humidity-to-Location\n")
	locationRanges:= ProcessIntructions(humidityRanges, sections["humidity-to-location"])
	log.Printf("\ntranslatedRanges after humidity-to-location: %v", locationRanges)


	return 999

	// plantingLocations := []int{}


	// for _, seedNumber := range seedNumbers {
	// 	// translate seed to soil to fertilizer to water to light to temperature to humidity to location
	// 	soil := TranslateInstructions(seedNumber, sections["seed-to-soil"])
	// 	fertilizer := TranslateInstructions(soil, sections["soil-to-fertilizer"])
	// 	water := TranslateInstructions(fertilizer, sections["fertilizer-to-water"])
	// 	light := TranslateInstructions(water, sections["water-to-light"])
	// 	temperature := TranslateInstructions(light, sections["light-to-temperature"])
	// 	humidity := TranslateInstructions(temperature, sections["temperature-to-humidity"])
	// 	location := TranslateInstructions(humidity, sections["humidity-to-location"])
	// 	plantingLocations = append(plantingLocations, location)
	// }
	// return slices.Min(plantingLocations)
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}

func TranslateInstructions(initialNumber int, instructions []Instruction) int {
	for _, instruction := range instructions {
		if initialNumber>=instruction.Source && initialNumber<instruction.Source+instruction.Length {
			// apply the translation
			delta := initialNumber-instruction.Source
			log.Printf("initialNumber: %v, instruction: %v, delta: %v", initialNumber, instruction, delta)
			return instruction.Destination+delta
		}
	}
	return initialNumber
}

func GetSeedRanges(s string) []Range {
	numStrs := strings.Split(strings.Split(s, ":")[1], " ")

	seedRanges := []Range{}
	for i, r := range numStrs {
		if v, err := strconv.Atoi(r); err == nil {
			if (i%2==1) {
				initial := v
				delta, _ := strconv.Atoi(numStrs[i+1])
				seedRanges = append(seedRanges, Range{Min: initial, Max: initial+delta-1})
			}
		}
	}
	log.Printf("seedRanges: %v", seedRanges)
	return seedRanges
}

func SortRanges(rs []Range) []Range {
	// sort by min
	r:=rs
	sort.Slice(r, func(i, j int) bool {
		return r[i].Min < r[j].Min
	})
	return r

}

func GetSeedNumbers(s string) []int {
	numStrs := strings.Split(strings.Split(s, ":")[1], " ")

	seedNumbers := []int{}
	for _, r := range numStrs {
		if v, err := strconv.Atoi(r); err == nil {
		seedNumbers = append(seedNumbers, int(v))
	}}
	return seedNumbers
}

func LoadSections (data []string) map[string][]Instruction {
	sections := map[string][]Instruction{}
	section := ""

	for _, line := range data {
		if strings.Contains(line, "map:"){
			section = strings.Split(line, " ")[0]
			sections[section] = []Instruction{}
		} else if strings.Contains(line, " ") {
			numStrs := strings.Split(line, " ")
			var instruction Instruction
			for i, r := range numStrs {
				if v, err := strconv.Atoi(r); err == nil {
					switch i {
					case 0:
						instruction.Destination = int(v)
					case 1:
						instruction.Source = int(v)
					case 2:
						instruction.Length = int(v)
					}
				}
			}
			sections[section] = append(sections[section], instruction)
	}}
	log.Printf("sections: %v", sections)
	return sections
}

type Instruction struct {
	Destination int
	Source 	int
	Length int
}

type Range struct {
	Min int
	Max int
}

func TranslateNumAndLengthToRange (num int, length int) Range {
	return Range{Min: num, Max: num+length-1}
}

/*
	take the original range and see if the min is in an instruction 
		if it is, see if the max is in the same instruction
			if it is, return the translated range
			if it isn't, translate the part that is in the range but also return the leftover part to be translated
		if it isn't, find the instruction with the closest min
		& split the beginning part our and translate as much of the remaining range as possible
	if neither the min or max is in an instruction, return that range as is.
*/


/*
	seed 1
	inputting {79, 92}
	seed-to-soil map:
		50 98 2
		52 50 48
	both are within the range of the second instruction
	79 => 79-50 = 29 + 52 = 81
	92 => 92-50 = 42 + 52 = 94
	return {81, 94}
*/

/*
	seed 2
	inputting {55, 67}
	seed-to-soil map:
		50 98 2
		52 50 48
	both are within the range of the second instruction
	55 => 55-50 = 5 + 52 = 57
	67 => 67-50 = 17 + 52 = 69
	return {57, 69}
*/

type RangeTranslationResult struct {
	TranslatedRange Range
	RemainingRanges []Range
}

func TranslateRangeInstructions (initialRange Range, instructions []Instruction) RangeTranslationResult {
	ranges := []Range{}
	closestBiggerInstruction := -1
	isInRange := false
	for i, instruction := range instructions {
		instructionMax := instruction.Source+instruction.Length
		delta := initialRange.Min-instruction.Source

		if  initialRange.Max>=instruction.Source && initialRange.Max<instructionMax {
			log.Printf("updating closestBiggerInstruction. initialRange: %v, instruction: %v, delta: %v", initialRange, instruction, delta)
			isInRange = true
			if closestBiggerInstruction==-1  {
				closestBiggerInstruction = i
				log.Printf("closestBiggerInstruction: %v", closestBiggerInstruction)
				} else if  instruction.Source<instructions[closestBiggerInstruction].Source {
					closestBiggerInstruction = i
					log.Printf("closestBiggerInstruction: %v", closestBiggerInstruction)
			}
		}

		/*
			take the original range and see if the min is in an instruction 
				if it is, see if the max is in the same instruction
					if it is, return the translated range
					if it isn't, translate the part that is in the range but also return the leftover part to be translated
				if it isn't, find the instruction with the closest min
				& split the beginning part our and translate as much of the remaining range as possible
			if neither the min or max is in an instruction, return that range as is.
		*/

		// if the initial range is fully within the instruction range
		if initialRange.Min>=instruction.Source {
			log.Printf("initialRange.Min greater than instruction range min: %d %d", initialRange.Min, instruction.Source)
			if initialRange.Max<=instructionMax {
				// apply the translation
				// log.Printf("initialRange: %v, instruction: %v, delta: %v", initialRange, instruction, delta)
				log.Printf("initialRange.Max less than instruction range max, so fully contained. Success! %d %d", initialRange.Max, instructionMax)
				r := Range{Min: instruction.Destination+delta, Max: instruction.Destination+delta+initialRange.Max-initialRange.Min}
				log.Printf("returning %v", r)
				return RangeTranslationResult{r,nil}
			} else if initialRange.Min>=instructionMax {
				// log.Printf("and initialRange.Min also greater than instruction range, making it totally out of range: %d %d", initialRange.Min, instructionMax)
			} else {
				log.Printf("top is out of bounds: %d > %d", initialRange.Max, instructionMax)
				// apply the translation				
				log.Printf("initialRange: %v, instruction: %v, delta: %v", initialRange, instruction, delta)
				t:= Range{Min: instruction.Destination+delta, Max: instruction.Destination+delta+delta}

				leftovers := Range{Min: instructionMax, Max: initialRange.Max}
				log.Printf("returning %v and %v", t, leftovers)
				return RangeTranslationResult{t,[]Range{leftovers}}
			}
		} 
	}
	if !isInRange {
		log.Printf(`not in range, so returning original range`)
		return RangeTranslationResult{initialRange,nil}
	}
	log.Printf(`needing to split bigger`)
	if closestBiggerInstruction!=-1 {
		instruction := instructions[closestBiggerInstruction]
		delta := initialRange.Min-instruction.Source
		// apply the translation
		// log.Printf("initialRange: %v, instruction: %v, delta: %v", initialRange, instruction, delta)
		ranges = append(ranges, Range{Min: instruction.Destination+delta, Max: instruction.Destination+delta+instruction.Length})
		// return the remaining part
		remainingRange := Range{Min: instruction.Source+instruction.Length, Max: initialRange.Max}
		log.Printf("Splitting below,  %v and %v", ranges[0], remainingRange)
		return RangeTranslationResult{ranges[0],[]Range{remainingRange}}
	} else {
		return RangeTranslationResult{initialRange,nil}
	}
}

