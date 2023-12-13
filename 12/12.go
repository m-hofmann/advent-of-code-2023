package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type record struct {
	report    []byte
	checksums []int
}

func main() {
	fmt.Println("Starting day 12 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	records := make([]record, 0)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, " ")
		nums := extractNumList(parts[1])
		records = append(records, record{[]byte(parts[0]), nums})
	}

	sum := 0
	for _, record := range records {
		sum += memoizedGetPermCount(record)
	}

	fmt.Println("Part 1 solution:", sum)

	sum = 0
	for _, curr := range records {
		report := fmt.Sprintf("%s?%s?%s?%s?%s", string(curr.report), string(curr.report), string(curr.report), string(curr.report), string(curr.report))
		checksums := make([]int, len(curr.checksums)*5)
		for i := 0; i < len(checksums); i++ {
			checksums[i] = curr.checksums[i%len(curr.checksums)]
		}
		sum += memoizedGetPermCount(record{[]byte(report), checksums})
	}

	fmt.Println("Part 2 solution:", sum)

}

var cache = make(map[string]int)

func memoizedGetPermCount(r record) int {
	// Golang memoization with arbitrary structs (possible containing slices) is a pain
	// this must be sooo inefficient
	key := string(r.report) + "/" + fmt.Sprintf("%v", r.checksums)
	if _, ok := cache[key]; !ok {
		cache[key] = getPermCount(r)
		return cache[key]
	} else {
		return cache[key]
	}
}

func getPermCount(r record) int {
	// took the liberty of brushing up on my dynamic programming using this tutorial
	// https://old.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/
	if len(r.checksums) == 0 {
		if !slices.Contains(r.report, '#') {
			return 1
		} else {
			return 0
		}
	}

	if len(r.report) == 0 {
		return 0
	}

	nxt := r.report[0]
	out := 0
	switch nxt {
	case '.':
		out = handleDot(r.report, r.checksums)
	case '#':
		out = handleHash(r.report, r.checksums)
	case '?':
		out = handleDot(r.report, r.checksums) + handleHash(r.report, r.checksums)
	}

	// fmt.Println(string(report), checksums, "->", out)

	return out
}

func handleDot(report []byte, checksums []int) int {
	return memoizedGetPermCount(record{report[1:], checksums})
}

func handleHash(report []byte, checksums []int) int {
	nextChecksum := checksums[0]
	currSprings := make([]byte, nextChecksum)
	hashCount := 0
	if len(report) < nextChecksum {
		return 0
	}
	for i, c := range report[:nextChecksum] {
		if c == '?' {
			currSprings[i] = '#'
			hashCount++
		} else if c == '#' {
			currSprings[i] = c
			hashCount++
		} else {
			currSprings[i] = c
		}
	}
	if hashCount != nextChecksum {
		return 0
	}

	if len(report) == nextChecksum {
		// last n bytes fits last checksum, we're done
		if len(checksums) == 1 {
			return 1
		} else {
			return 0
		}
	}
	if report[nextChecksum] == '?' || report[nextChecksum] == '.' {
		// can be separator, skip and go to next group
		return memoizedGetPermCount(record{report[nextChecksum+1:], checksums[1:]})
	}
	return 0
}

func getChecksum(report []byte) []int {
	checksums := make([]int, 0)
	cons := 0
	for i := 0; i < len(report); i++ {
		if report[i] == '#' {
			cons++
		} else if report[i] == '.' && cons != 0 {
			checksums = append(checksums, cons)
			cons = 0
		}
	}
	if cons != 0 {
		checksums = append(checksums, cons)
	}
	return checksums
}

func extractNumList(nums string) []int {
	numList := make([]int, 0)
	for _, str := range strings.Split(nums, ",") {
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
