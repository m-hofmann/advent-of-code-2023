package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards []byte
	bid   int
}

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

var cardValuesPart1 = map[byte]int{
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

var cardValuesPart2 = map[byte]int{
	'J': 0,
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func main() {
	fmt.Println("Starting day 07 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	// parsed data
	hands := make([]hand, 0)

	// parsing data
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, " ")
		bidValue, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			log.Fatalln("Could not parse bid value")
		}
		hand := hand{
			cards: []byte(parts[0]),
			bid:   int(bidValue),
		}
		hands = append(hands, hand)
	}

	// part 1
	sort.Slice(hands, func(i, j int) bool {
		aHand := hands[i]
		bHand := hands[j]
		return compareHandsPart1(aHand, bHand)
	})
	accu := 0
	for i, v := range hands {
		fmt.Println("Hand of rank", i+1, "is", string(v.cards))
		accu += (i + 1) * v.bid
	}
	fmt.Println("Part 1 solution: ", accu)

	// part 1
	sort.Slice(hands, func(i, j int) bool {
		aHand := hands[i]
		bHand := hands[j]
		return compareHandsPart2(aHand, bHand)
	})
	accu = 0
	for i, v := range hands {
		fmt.Println("Hand of rank", i+1, "is", string(v.cards))
		accu += (i + 1) * v.bid
	}
	fmt.Println("Part 2 solution: ", accu)
}

func compareHandsPart1(aHand hand, bHand hand) bool {
	aType, err := handTypePart1(aHand)
	if err != nil {
		log.Fatalln("Could not get hand type for ", aHand, "because", err)
	}
	bType, err := handTypePart1(bHand)
	if err != nil {
		log.Fatalln("Could not get hand type for ", bHand, "because", err)
	}
	if aType != bType {
		return aType < bType
	} else {
		for k := 0; k < len(aHand.cards); k++ {
			if aHand.cards[k] != bHand.cards[k] {
				return cardValuesPart1[aHand.cards[k]] < cardValuesPart1[bHand.cards[k]]
			}
		}
	}
	return false
}

func handTypePart2(h hand) (int, error) {
	chrCounts := make(map[byte]int)

	for _, c := range h.cards {
		chrCounts[c]++
	}

	jCount := chrCounts['J']
	if jCount > 0 && len(chrCounts) > 1 {
		// if there exists a J, and it's not a full hand of Js
		// assume all Js are now the most counted card
		// this should suffice to always upgrade our hand according to the rules
		delete(chrCounts, 'J')
		var maxChar byte
		maxCount := math.MinInt
		for k, v := range chrCounts {
			if v > maxCount {
				maxCount = v
				maxChar = k
			}
		}

		chrCounts[maxChar] += jCount
	}

	values := make([]int, len(chrCounts))
	i := 0
	for _, v := range chrCounts {
		values[i] = v
		i++
	}

	sort.Ints(values)

	if len(values) == 1 && values[0] == 5 {
		return FIVE_OF_A_KIND, nil
	}

	if len(values) == 2 && values[0] == 1 && values[1] == 4 {
		return FOUR_OF_A_KIND, nil
	}

	if len(values) == 2 && values[0] == 2 && values[1] == 3 {
		return FULL_HOUSE, nil
	}

	if len(values) == 3 && values[0] == 1 && values[1] == 1 && values[2] == 3 {
		return THREE_OF_A_KIND, nil
	}

	if len(values) == 3 && values[0] == 1 && values[1] == 2 && values[2] == 2 {
		return TWO_PAIR, nil
	}

	if len(values) == 4 && values[0] == values[1] && values[1] == values[2] && values[3] == 2 {
		return ONE_PAIR, nil
	}

	if len(values) == 5 {
		return HIGH_CARD, nil
	}

	return -1, fmt.Errorf("unknown hand type %s", string(hand{}.cards))
}

func compareHandsPart2(aHand hand, bHand hand) bool {
	aType, err := handTypePart2(aHand)
	if err != nil {
		log.Fatalln("Could not get hand type for ", aHand, "because", err)
	}
	bType, err := handTypePart2(bHand)
	if err != nil {
		log.Fatalln("Could not get hand type for ", bHand, "because", err)
	}
	if aType != bType {
		return aType < bType
	} else {
		for k := 0; k < len(aHand.cards); k++ {
			if aHand.cards[k] != bHand.cards[k] {
				return cardValuesPart2[aHand.cards[k]] < cardValuesPart2[bHand.cards[k]]
			}
		}
	}
	return false
}
func handTypePart1(h hand) (int, error) {
	chrCounts := make(map[byte]int)

	for _, c := range h.cards {
		chrCounts[c]++
	}

	values := make([]int, len(chrCounts))
	i := 0
	for _, v := range chrCounts {
		values[i] = v
		i++
	}

	sort.Ints(values)

	if len(values) == 1 && values[0] == 5 {
		return FIVE_OF_A_KIND, nil
	}

	if len(values) == 2 && values[0] == 1 && values[1] == 4 {
		return FOUR_OF_A_KIND, nil
	}

	if len(values) == 2 && values[0] == 2 && values[1] == 3 {
		return FULL_HOUSE, nil
	}

	if len(values) == 3 && values[0] == 1 && values[1] == 1 && values[2] == 3 {
		return THREE_OF_A_KIND, nil
	}

	if len(values) == 3 && values[0] == 1 && values[1] == 2 && values[2] == 2 {
		return TWO_PAIR, nil
	}

	if len(values) == 4 && values[0] == values[1] && values[1] == values[2] && values[3] == 2 {
		return ONE_PAIR, nil
	}

	if len(values) == 5 {
		return HIGH_CARD, nil
	}

	return -1, fmt.Errorf("unknown hand type %s", string(hand{}.cards))
}

func extractNumList(nums string) []int {
	numList := make([]int, 0)
	for _, str := range strings.Split(nums, " ") {
		if str == "" {
			continue
		}
		number, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			log.Fatalln("Could not parse number from numlist", err)
		}
		numList = append(numList, int(number))
	}
	return numList
}
