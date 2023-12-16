package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting day 15 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	steps := make([]string, 0)
	sc.Scan()
	line := sc.Text()
	steps = strings.Split(line, ",")

	sum := int64(0)
	for _, step := range steps {
		sum += HASH(step)
	}
	fmt.Println("Part 1 solution:", sum)
}

func HASH(s string) int64 {
	val := int64(0)
	for _, c := range s {
		val += int64(c)
		val *= 17
		val %= 256
	}
	return val
}
