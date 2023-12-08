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
	fmt.Println("Starting day 08 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	choices := make(map[string][]string)
	sc.Scan()
	movement := []byte(sc.Text())
	sc.Scan()
	for sc.Scan() {
		line := sc.Text()
		tokens := strings.Split(line, " = ")
		reachable := strings.Split(strings.Trim(tokens[1], "()"), ", ")
		choices[tokens[0]] = reachable
	}

	current := "AAA"
	accu := 0
	i := 0
	for current != "ZZZ" {
		direction := movement[i]
		i = (i + 1) % len(movement)
		idx := -1
		if direction == 'L' {
			idx = 0
		} else if direction == 'R' {
			idx = 1
		} else {
			log.Fatalln("Unknown direction", direction)
		}
		if possible, ok := choices[current]; ok {
			current = possible[idx]
		} else {
			fmt.Println("No successor for node", current)
			break
		}
		accu++
	}
	fmt.Println("Part 1 solution:", accu)

	startNodes := make([]string, 0)
	for k := range choices {
		if k[2] == 'A' {
			startNodes = append(startNodes, k)
		}
	}

	results := make(chan int)
	for _, startNode := range startNodes {
		go func(start string) {
			current := start
			accu := 0
			i := 0
			for current[2] != 'Z' {
				direction := movement[i]
				i = (i + 1) % len(movement)
				idx := -1
				if direction == 'L' {
					idx = 0
				} else if direction == 'R' {
					idx = 1
				} else {
					log.Fatalln("Unknown direction", direction)
				}
				next := choices[current][idx]
				//fmt.Println("Going from", current, "to", next)
				current = next
				accu++
			}
			results <- accu
		}(startNode)
	}

	minStepsPerNode := make([]int, len(startNodes))
	for i := 0; i < len(startNodes); i++ {
		result := <-results
		fmt.Println("Worker for", startNodes[i], "took", result, "steps")
		minStepsPerNode[i] = result
	}
	fmt.Println("Part 2 solution:", leastCommonMultiple(minStepsPerNode[0], minStepsPerNode[1], minStepsPerNode[2:]...))

}

func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)
	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}
	return result
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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
