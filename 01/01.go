package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Starting day 01 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	accu1 := 0
	accu2 := 0
	for sc.Scan() {
		line := sc.Text()
		a, b := findFirstLastDigit(line)
		if a != -1 && b != -1 {
			accu1 += a*10 + b
		}

		a, b = findFirstLastNumber(line)
		if a != -1 && b != -1 {
			accu2 += a*10 + b
		}
	}

	fmt.Println("Part 1:", accu1)
	fmt.Println("Part 2:", accu2)
}

func findFirstLastDigit(s string) (int, int) {
	lDigit := -1
	rDigit := -1
	for i := 0; i < len(s); i++ {
		num, err := strconv.ParseInt(s[i:i+1], 10, 0)
		if err == nil {
			lDigit = int(num)
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		num, err := strconv.ParseInt(s[i:i+1], 10, 0)
		if err == nil {
			rDigit = int(num)
			break
		}
	}

	return lDigit, rDigit
}

// finds number regardless whether it is a digit or spelled out
func findFirstLastNumber(s string) (int, int) {
	regex, err := regexp.Compile("(one|1|two|2|three|3|four|4|five|5|six|6|seven|7|eight|8|nine|9)")
	if err != nil {
		log.Fatalln("Could not compile regex")
	}

	// go's FindAll... regex functions don't do overlapping matches. Therefore, we have to proceed manually.
	currIdx := 0
	matches := make([]string, 0)
	for {
		match := regex.FindStringIndex(s[currIdx:])
		if match == nil {
			break
		}
		matches = append(matches, s[currIdx+match[0]:currIdx+match[1]])
		currIdx = currIdx + match[0] + 1
	}
	a := matchToDigit(matches[0])
	b := matchToDigit(matches[len(matches)-1])

	return a, b
}

func matchToDigit(s string) int {
	switch s {
	case "0", "zero":
		return 0
	case "1", "one":
		return 1
	case "2", "two":
		return 2
	case "3", "three":
		return 3
	case "4", "four":
		return 4
	case "5", "five":
		return 5
	case "6", "six":
		return 6
	case "7", "seven":
		return 7
	case "8", "eight":
		return 8
	case "9", "nine":
		return 9
	}
	return -1
}
