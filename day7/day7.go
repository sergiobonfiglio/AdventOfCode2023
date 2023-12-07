package day7

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() []string {

	readFile, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var data []string

	for fileScanner.Scan() {

		for fileScanner.Text() != "" {
			data = append(data, fileScanner.Text())
			fileScanner.Scan()
		}

	}

	err = readFile.Close()
	if err != nil {
		panic(err)
	}

	return data
}

func SolveDay() {

	var input = readInput()

	sol1, sol2 := solveParts(input)

	fmt.Println("Part 1:", sol1)

	fmt.Println("Part 2:", sol2)
}

type Hand struct {
	hand string
	bid  int
	rank int
}

var mappedCards = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}
var mappedCards2 = map[rune]int{
	'J': -1,
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func cmpHandsPart1(a, b Hand) int {

	aType := a.getType()
	bType := b.getType()

	if aType == bType {
		// check highest card
		for i := 0; i < len(a.hand); i++ {
			aValue := mappedCards[rune(a.hand[i])]
			bValue := mappedCards[rune(b.hand[i])]
			if aValue != bValue {
				return aValue - bValue
			}
		}
		return 0
	} else {
		return int(aType - bType)
	}

}

func cmpHandsPart2(a, b Hand) int {

	aType := a.getType2()
	bType := b.getType2()

	if aType == bType {
		// check highest card
		for i := 0; i < len(a.hand); i++ {
			aValue := mappedCards2[rune(a.hand[i])]
			bValue := mappedCards2[rune(b.hand[i])]
			if aValue != bValue {
				return aValue - bValue
			}
		}
		return 0
	} else {
		return int(aType - bType)
	}

}

type HandType int

const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OAK
	FULL_HOUSE
	FOUR_OAK
	FIVE_OAK
)

func (h Hand) getType2() HandType {
	cardMap := map[rune]int{}
	for _, c := range h.hand {
		cardMap[c]++
	}

	jokers := cardMap['J']

	if jokers == 5 {
		return FIVE_OAK
	}

	threeNum, pairNum := 0, 0
	rawType := HIGH_CARD
	for card, count := range cardMap {
		if count == 5 {
			return FIVE_OAK
		}
		if card != 'J' {
			if count == 4 {
				rawType = FOUR_OAK
				break
			}
			if count == 3 {
				threeNum++
			}
			if count == 2 {
				pairNum++
			}
		}
	}

	if threeNum == 1 && pairNum == 1 {
		rawType = FULL_HOUSE
	} else if threeNum == 1 {
		rawType = THREE_OAK
	} else if pairNum == 2 {
		rawType = TWO_PAIR
	} else if pairNum == 1 {
		rawType = ONE_PAIR
	}

	if jokers > 0 {
		if rawType == FOUR_OAK {
			return FIVE_OAK
		}
		if rawType == THREE_OAK {
			if jokers == 1 {
				return FOUR_OAK
			}
			return FIVE_OAK
		}
		if rawType == TWO_PAIR {
			return FULL_HOUSE
		}
		if rawType == ONE_PAIR {
			if jokers == 1 {
				return THREE_OAK
			} else if jokers == 2 {
				return FOUR_OAK
			}
			return FIVE_OAK
		}
		if rawType == HIGH_CARD {
			if jokers == 1 {
				return ONE_PAIR
			} else if jokers == 2 {
				return THREE_OAK
			} else if jokers == 3 {
				return FOUR_OAK
			}
			return FIVE_OAK
		}
	}
	return rawType
}

func (h Hand) getType() HandType {

	cardMap := map[rune]int{}
	for _, c := range h.hand {
		cardMap[c]++
	}

	threeNum, pairNum := 0, 0
	for _, count := range cardMap {
		if count == 5 {
			return FIVE_OAK
		}
		if count == 4 {
			return FOUR_OAK
		}
		if count == 3 {
			threeNum++
		}
		if count == 2 {
			pairNum++
		}
	}

	if threeNum == 1 && pairNum == 1 {
		return FULL_HOUSE
	}
	if threeNum == 1 {
		return THREE_OAK
	}
	if pairNum == 2 {
		return TWO_PAIR
	}
	if pairNum == 1 {
		return ONE_PAIR
	}
	return HIGH_CARD
}

func solveParts(input []string) (int, int64) {

	var hands []Hand
	for _, row := range input {
		fields := strings.Fields(row)
		hand := fields[0]
		bid, err := strconv.Atoi(fields[1])
		if err != nil {
			panic("err")
		}
		hands = append(hands, Hand{
			hand: hand,
			bid:  bid,
		})
	}

	slices.SortFunc(hands, cmpHandsPart1)

	winnings1 := 0
	for rank, hand := range hands {
		winnings1 += (rank + 1) * hand.bid
	}

	slices.SortFunc(hands, cmpHandsPart2)

	winnings2 := int64(0)
	for rank, hand := range hands {
		winnings2 += int64(rank+1) * int64(hand.bid)
	}

	return winnings1, winnings2
}
