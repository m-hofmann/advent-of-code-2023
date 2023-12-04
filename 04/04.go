package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type scratchcard struct {
	number  int
	winning map[int]bool
	have    map[int]bool
}

func main() {
	fmt.Println("Starting day 04 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	cards := make([]scratchcard, 0)
	for sc.Scan() {
		line := sc.Text()
		parseCard := strings.Split(line, ":")
		gameNo, err := strconv.ParseInt(strings.TrimFunc(parseCard[0], func(j rune) bool { return !unicode.IsDigit(j) }), 10, 0)
		if err != nil {
			log.Fatalln("Could not parse game number", err)
		}
		parseLists := strings.Split(parseCard[1], "|")
		cards = append(cards, scratchcard{
			number:  int(gameNo),
			winning: extractNumSet(parseLists[0]),
			have:    extractNumSet(parseLists[1]),
		})
	}

	accu := 0
	cardNumberIntersectionCount := make(map[int]int)
	for _, card := range cards {
		value := 0
		intersec := setIntersectionCount(card.winning, card.have)

		if intersec > 0 {
			value = 1 << (intersec - 1)
		}

		fmt.Println("Card", card.number, "is worth", value)
		cardNumberIntersectionCount[card.number] = intersec
		accu += value
	}

	fmt.Println("Part 1 solution:", accu)

	// count how many of each card we have: 1 original + n won copies
	cardCounts := make(map[int]int)
	for _, card := range cards {
		// original card
		cardCounts[card.number] = cardCounts[card.number] + 1
		for i := 1; i <= cardNumberIntersectionCount[card.number]; i++ {
			cardCounts[card.number+i] += cardCounts[card.number]
		}
	}

	accu2 := 0

	for card, count := range cardCounts {
		fmt.Println("Card number ", card, ":", count, "instances")
		accu2 += count
	}

	fmt.Println("Part 2 solution:", accu2)
}

func extractNumSet(nums string) map[int]bool {
	numberSet := make(map[int]bool)
	for _, str := range strings.Split(nums, " ") {
		if str == "" {
			continue
		}
		number, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			log.Fatalln("Could not parse number from card", err)
		}
		numberSet[int(number)] = true
	}
	return numberSet
}

func setIntersectionCount(a, b map[int]bool) int {
	count := 0
	for fromA := range a {
		if _, ok := b[fromA]; ok {
			count++
		}

	}
	return count
}
