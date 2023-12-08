package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	sortedHands := SeparatedHands{}
	for _, hand := range data {
		parsedHand := ParseHand(hand)
		if sortedHands[parsedHand.handType] == nil {
			sortedHands[parsedHand.handType] = []Hand{parsedHand}
			log.Printf("Inserting %v at %v into handType %d", parsedHand, 0, parsedHand.handType)
		} else {
			// find where to insert it then insert it
			// sort using CompareHands
			i := sort.Search(len(sortedHands[parsedHand.handType]), func(i int) bool { 
				compare, err := CompareHands(parsedHand,sortedHands[parsedHand.handType][i])
				if err != nil {
					log.Panic(err)
				}
				return compare !=1
				})
			sortedHands[parsedHand.handType] = insert(sortedHands[parsedHand.handType],parsedHand,i)
		}
	}
	// calculate the score 
	// join the hands together starting from 7 > 1
	joinedHands := []Hand{}
	for handType := 7; handType > 0; handType-- {
		if sortedHands[HandType(handType)] != nil {
			joinedHands = append(joinedHands, sortedHands[HandType(handType)]...)
		}
	}

	counter :=0
	for i, hand := range joinedHands {
		// log.Printf("%v: %v.", i, hand)
		counter += hand.bid * (len(joinedHands)-i)
	}


	log.Printf("%v: %v", counter, joinedHands)
	return counter
}

func Part2(data []string) int {
	sortedHands := SeparatedHands{}
	for _, hand := range data {
		parsedHand := ParseHand2(hand)
		if sortedHands[parsedHand.handType] == nil {
			sortedHands[parsedHand.handType] = []Hand{parsedHand}
			log.Printf("Inserting %v at %v into handType %d", parsedHand, 0, parsedHand.handType)
		} else {
			// find where to insert it then insert it
			// sort using CompareHands
			i := sort.Search(len(sortedHands[parsedHand.handType]), func(i int) bool { 
				compare, err := CompareHands2(parsedHand,sortedHands[parsedHand.handType][i])
				if err != nil {
					log.Panic(err)
				}
				return compare !=1
				})
			sortedHands[parsedHand.handType] = insert(sortedHands[parsedHand.handType],parsedHand,i)
		}
	}
	// calculate the score 
	// join the hands together starting from 7 > 1
	joinedHands := []Hand{}
	for handType := 7; handType > 0; handType-- {
		if sortedHands[HandType(handType)] != nil {
			joinedHands = append(joinedHands, sortedHands[HandType(handType)]...)
		}
	}

	counter :=0
	for i, hand := range joinedHands {
		counter += hand.bid * (len(joinedHands)-i)
	}

	// log.Printf("%v: %v", counter, joinedHands)
	for _, hand := range joinedHands {
		log.Printf("%v: %v", hand, hand.handType)
	}

	return counter
}

type Hand struct {
	cards string
	bid int
	handType HandType
}


func main() {
	startTime := time.Now()
	data := LoadInput("input.txt")
	// fmt.Printf("Part 1: %v\n", Part1(data))
	// fmt.Println(time.Since(startTime))
	fmt.Printf("Part 2: %v\n", Part2(data))
	fmt.Println(time.Since(startTime))

}

type HandType int

const  (
	highCard  HandType = iota+1
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type Ranks int

const (
	T  Ranks = iota+10
	J
	Q
	K
	A
)

var rankMap = map[string]int {
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,	
}

var rankMap2 = map[string]int {
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 12,
	"K": 13,
	"A": 14,	
}


func GetHandType (hand Hand) HandType {
	cards := map[string]int {}
	for _, card := range hand.cards {
		cards[string(card)] += 1		
	}
	if len(cards) == 5 {
		return highCard
	}
	if len(cards) == 4 {
		// no need to tweak for J
		return onePair
	}
	if len(cards) == 3 {
		for _, count := range cards {
			if count == 3 {
				return threeOfAKind
			}
		}
		return twoPair
	}
	if len(cards) == 2 {
		for _, count := range cards {
			if count == 4 {
				return fourOfAKind
			}
		}
		return fullHouse
	}
	return fiveOfAKind
}
func GetHandType2 (hand Hand) HandType {
	cards := map[string]int {}
	numJs := 0
	for _, card := range hand.cards {
		if card == 'J' {
				numJs += 1
			} else {
				cards[string(card)] += 1
		}
	}
	log.Printf("Num Js: %v for %v", numJs, hand.cards)
	if len(cards) == 5 {
		return highCard
	}
	if len(cards) == 4 {
		// no need to tweak for J
		return onePair
	}
	if len(cards) == 3 {
		switch numJs {
			case 2: return threeOfAKind
			case 1: return threeOfAKind
			// if there are no js, then it'll be a two pair or 3 of a kind
			case 0: {
				for _, count := range cards {
					if count == 3 {
						return threeOfAKind
					}
				}
				return twoPair
			}
		}
	}
	if len(cards) == 2 {
		switch numJs {
			case 3: return fourOfAKind
			case 2: return fourOfAKind
			case 1: {
				for _, count := range cards {
					if count == 3 {
						return fourOfAKind
					}
				}
				return fullHouse
			}
			case 0: {
				for _, count := range cards {
					if count == 4 {
						return fourOfAKind
					}
				}
				return fullHouse
			}
		}
	}
	return fiveOfAKind
}

func CompareHands (hand1 Hand, hand2 Hand) (int, error) {
	if hand1.handType > hand2.handType {
		return 1, nil
	}
	if hand1.handType < hand2.handType {
		return -1, nil
	}
	// walk through cards til one is higher
	for i := 0; i < 5; i++ {
		if (rankMap[string(hand1.cards[i])] == 0) {
			log.Printf("This turned out as 0 %v", int(hand1.cards[i]))
		}
		v1 := int(rankMap[string(hand1.cards[i])])
		v2:=  int(rankMap[string(hand2.cards[i])])
		log.Printf("Rank %v vs %v; %v: %v %v", v1, v2, i, hand1.cards, hand2.cards)
		if v1 > v2 {
			log.Printf(`%v (%v) (rank %v) > %v (%v) (rank %v)`, hand1.cards[i],  string(hand1.cards[i]), rankMap[string(hand1.cards[i])], hand2.cards[i], string(hand2.cards[i]), rankMap[string(hand2.cards[i])])
			return -1, nil
		}
		if v1 < v2 {
			log.Printf(`%v < %v`, string(hand1.cards[i]), string(hand2.cards[i]))
			return 1, nil
		}
	}
	return 0, fmt.Errorf("hands are fully equal %v %v", hand1, hand2)
}

func CompareHands2 (hand1 Hand, hand2 Hand) (int, error) {
	if hand1.handType > hand2.handType {
		return 1, nil
	}
	if hand1.handType < hand2.handType {
		return -1, nil
	}
	// walk through cards til one is higher
	for i := 0; i < 5; i++ {
		if (rankMap2[string(hand1.cards[i])] == 0) {
			log.Printf("This turned out as 0 %v", int(hand1.cards[i]))
		}
		v1 := int(rankMap2[string(hand1.cards[i])])
		v2:=  int(rankMap2[string(hand2.cards[i])])
		log.Printf("Rank %v vs %v; %v: %v %v", v1, v2, i, hand1.cards, hand2.cards)
		if v1 > v2 {
			log.Printf(`%v (%v) (rank %v) > %v (%v) (rank %v)`, hand1.cards[i],  string(hand1.cards[i]), rankMap2[string(hand1.cards[i])], hand2.cards[i], string(hand2.cards[i]), rankMap2[string(hand2.cards[i])])
			return -1, nil
		}
		if v1 < v2 {
			log.Printf(`%v < %v`, string(hand1.cards[i]), string(hand2.cards[i]))
			return 1, nil
		}
	}
	return 0, fmt.Errorf("hands are fully equal %v %v", hand1, hand2)
}

func ParseHand (hand string) Hand {
	cards := hand[0:5]
	bid, _ := strconv.Atoi(hand[6:])
	handType := GetHandType(Hand{cards,bid,-1})
	return Hand {cards,bid,handType}
}
func ParseHand2 (hand string) Hand {
	cards := hand[0:5]
	bid, _ := strconv.Atoi(hand[6:])
	handType := GetHandType2(Hand{cards,bid,-1})
	return Hand {cards,bid,handType}
}

type SeparatedHands map[HandType][]Hand

func insert(array []Hand, element Hand, i int) []Hand {
	// log.Printf(">> Inserting %v at %v into %v", element, i, array)
	addedArray := append(array[:i], append([]Hand{element}, array[i:]...)...)
	// log.Printf("<< Inserted %v at %v. Current: %v", element, i, addedArray)
	return addedArray
}