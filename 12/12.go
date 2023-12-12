package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
		sum += getPermCount(record.report, record.checksums)
	}

	fmt.Println("Part 1 solution:", sum)

}

func getPermCount(report []byte, checksums []int) int {
	perms := make([][]byte, 0)
	perms = append(perms, report)
	for i := 0; i < len(report); i++ {
		newPerms := make([][]byte, 0)
		for _, perm := range perms {
			if perm[i] == '?' {
				for _, c := range []byte{'.', '#'} {
					tmp := make([]byte, len(perm))
					copy(tmp, perm)
					tmp[i] = c
					newPerms = append(newPerms, tmp)
				}
			} else {
				newPerms = append(newPerms, perm)
			}
		}
		perms = newPerms

	}

	valid := 0
	for _, perm := range perms {
		if reflect.DeepEqual(getChecksum(perm), checksums) {
			valid++
		}
	}
	return valid
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
