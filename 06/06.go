package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	dest int
	src  int
	size int
}

func main() {
	fmt.Println("Starting day 06 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	// parsed data
	times := make([]int, 0)
	distances := make([]int, 0)

	// parsing data

	sc.Scan()
	firstLine := sc.Text()
	times = extractNumList(strings.TrimPrefix(firstLine, "Time:"))
	sc.Scan()
	secondLine := sc.Text()
	distances = extractNumList(strings.TrimPrefix(secondLine, "Distance:"))

	if len(times) != len(distances) {
		log.Fatalln("times must have same length as distances")
	}

	accu := 1
	for race := 0; race < len(times); race++ {
		record := distances[race]
		winning := 0
		for t := 0; t < times[race]; t++ {
			dist := (times[race] - t) * t
			if dist > record {
				winning++
			}

			if winning > 0 && dist <= record {
				break
			}
		}
		accu *= winning
	}

	fmt.Println("Part 1 solution:", accu)

	totalTime, err := strconv.ParseInt(strings.ReplaceAll(strings.TrimPrefix(firstLine, "Time:"), " ", ""), 10, 0)
	if err != nil {
		log.Fatalln("Could not read totalTime", err)
	}
	record, err := strconv.ParseInt(strings.ReplaceAll(strings.TrimPrefix(secondLine, "Distance:"), " ", ""), 10, 0)
	if err != nil {
		log.Fatalln("Could not read totalDistance", err)
	}

	x1 := -float64(totalTime) + math.Sqrt(float64(totalTime*totalTime-4*(-1)*-record))/2*(-1)
	x2 := -float64(totalTime) - math.Sqrt(float64(totalTime*totalTime-4*(-1)*-record))/2*(-1)
	fmt.Println("Part 2 solution", math.Round(math.Abs(x1-x2)))
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
