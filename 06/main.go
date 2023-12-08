package main

import (
	"bufio"
	"fmt"
	"log"
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

type Race struct {
	Time int
	Distance int
}

func Part1(data []string) int {
	times := strings.Fields(strings.Split(data[0], ":")[1])
	speed := strings.Fields(strings.Split(data[1], ":")[1])
	raceData :=	MakeRaceData(times, speed)[0]
	
	losses := CalculateLosingRaces(raceData)
	log.Printf("Losses: %d", losses)
	wins := raceData.Time - losses
	log.Printf("Wins: %d", wins)
	// for _, race := range raceData {
	// 	counter:=0
	// 	ways := CalculateWaysToRace(race)
	// 	for _, way := range ways {
	// 		if way>race.Distance {
	// 			counter++
	// 			log.Printf("Increasing counter to %d", counter)
	// 		}
	// 	}
	// 	totalCounter*=counter
	// 	log.Printf("Total counter: %d", totalCounter)
	// }

	return wins
}

func MakeRaceData(times []string, speed []string) []Race {
	raceData := make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		timeInt, _ := strconv.Atoi(times[i])
		speedInt, _ := strconv.Atoi(speed[i])
		raceData[i] = Race{Time: timeInt, Distance: speedInt}
	}
	return raceData
}

func Part2(data []string) int {
	times := strings.Fields(strings.Split(data[0], ":")[1])
	speed := strings.Fields(strings.Split(data[1], ":")[1])
	log.Printf("Times: %v, Speed: %v", times, speed)
	raceData :=	MakeRaceData(times, speed)[0]
	
	losses := CalculateLosingRaces(raceData)
	log.Printf("Losses: %d", losses)
	wins := raceData.Time - losses
	log.Printf("Wins: %d", wins)

	return wins
}

func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	fmt.Printf("Part 1: %v\n", Part1(data))
	fmt.Println(time.Since(startTime))
	// fmt.Printf("Part 2: %v\n", Part2(data))
	// fmt.Println(time.Since(startTime))

}

func CalcTotalDistance(timeCharging int, totalTime int) int {
	speed:=timeCharging
	remainingTime:=totalTime-timeCharging
	log.Printf("speed: %v, remainingTime: %v, total distance: %d", speed, remainingTime, speed*remainingTime)
	return speed*remainingTime
}

func CalculateLosingRaces (race Race) int {
	 losingRace:= 0

	for chargingTime:=1; chargingTime<= race.Time; chargingTime++ {
		racingTime:=(race.Time-chargingTime)
		log.Printf(`Charging time: %d, racing time: %d, Total distance: %d`, chargingTime, racingTime, chargingTime*(race.Time-chargingTime))
		if chargingTime*racingTime < race.Distance {
			losingRace++
		} else {
			break
		}
	}
	return losingRace*2 + 1
}