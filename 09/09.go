package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting day 09 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	readings := make([][]int, 0)
	for sc.Scan() {
		line := sc.Text()
		readings = append(readings, extractNumList(line))
	}

	accu := 0
	for i, reading := range readings {
		fmt.Println("Dealing with", reading)
		nextVal := extrapolatePart1(reading)

		fmt.Println("Next value for", i+1, "series is", nextVal)
		accu += nextVal
	}
	fmt.Println("Part 1 solution:", accu)

	accu = 0
	for i, reading := range readings {
		fmt.Println("Dealing with", reading)
		nextVal := extrapolatePart2(reading)

		fmt.Println("Next value for", i+1, "series is", nextVal)
		accu += nextVal
	}
	fmt.Println("Part 2 solution:", accu)

}

func extrapolatePart1(values []int) int {
	deltas := make([]int, len(values)-1)

	onlyZeroes := true
	for i := 1; i < len(values); i++ {
		deltas[i-1] = values[i] - values[i-1]
		if deltas[i-1] != 0 {
			onlyZeroes = false
		}
	}

	if onlyZeroes {
		return values[len(values)-1]
	} else {
		return values[len(values)-1] + extrapolatePart1(deltas)
	}
}

func extrapolatePart2(values []int) int {
	deltas := make([]int, len(values)-1)

	onlyZeroes := true
	for i := 1; i < len(values); i++ {
		deltas[i-1] = values[i] - values[i-1]
		if deltas[i-1] != 0 {
			onlyZeroes = false
		}
	}

	if onlyZeroes {
		return values[0]
	} else {
		return values[0] - extrapolatePart2(deltas)
	}
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
