package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type interval struct {
	dest int
	src  int
	size int
}

type conversionMap struct {
	from      string
	to        string
	intervals []interval
}

func main() {
	fmt.Println("Starting day 04 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	preamble := true
	// parsed data
	var seeds []int
	typeToConversions := make(map[string]*conversionMap, 0)

	// parsing data
	var currentMap *conversionMap = nil
	for sc.Scan() {
		line := sc.Text()

		if preamble {
			if remainder, ok := strings.CutPrefix(line, "seeds: "); ok {
				seeds = extractNumList(remainder)
				preamble = false
			}
		} else {
			if strings.TrimSpace(line) != "" {
				if unicode.IsLetter(rune(line[0])) {
					// this is the start of a new map
					if remainder, ok := strings.CutSuffix(line, " map:"); ok {
						types := strings.Split(remainder, "-to-")
						currentMap = &conversionMap{
							from:      types[0],
							to:        types[1],
							intervals: make([]interval, 0),
						}

						typeToConversions[currentMap.from] = currentMap
					}
				} else {
					values := extractNumList(line)
					currentMap.intervals = append(currentMap.intervals, interval{
						dest: values[0],
						src:  values[1],
						size: values[2],
					})
				}
			}
		}
	}

	// part 1 - walk the translation tables
	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		currType := "seed"
		value := seed
		fmt.Println()
		fmt.Print("Seed ", value)
		for currType != "location" {
			if conversionMap, ok := typeToConversions[currType]; ok {
				for _, conversion := range conversionMap.intervals {
					if conversion.src <= value && value < conversion.src+conversion.size {
						value = value - conversion.src + conversion.dest
						break
					}
				}
				currType = conversionMap.to
			}
			fmt.Print(", ", currType, " ", value)
		}

		lowestLocation = min(lowestLocation, value)
	}
	fmt.Println()
	fmt.Println("Part 1 solution:", lowestLocation)

	// part 2
	lowestLocation = math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		for offset := 0; offset < seeds[i+1]; offset++ {
			currType := "seed"
			value := seeds[i] + offset
			for currType != "location" {
				if conversionMap, ok := typeToConversions[currType]; ok {
					for _, conversion := range conversionMap.intervals {
						if conversion.src <= value && value < conversion.src+conversion.size {
							value = value - conversion.src + conversion.dest
							break
						}
					}
					currType = conversionMap.to
				}
			}

			lowestLocation = min(lowestLocation, value)
		}
	}
	fmt.Println()
	fmt.Println("Part 2 solution:", lowestLocation)
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
