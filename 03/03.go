package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coord struct {
	y int
	x int
}

func main() {
	fmt.Println("Starting day 02 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	board := make([][]byte, 0)
	for sc.Scan() {
		line := sc.Text()
		board = append(board, []byte(line))
	}

	accu := 0
	for y := 0; y < len(board); y++ {
		numStartX := -1
		inNum := false
		for x := 0; x < len(board[y]); x++ {
			current := board[y][x]
			if isDigit(current) && !inNum {
				inNum = true
				numStartX = x
			}

			if inNum {
				endX := -1
				if !isDigit(current) {
					endX = x
				} else if x == len(board[y])-1 {
					endX = len(board[y])
				}

				if endX != -1 {
					numberStr := string(board[y][numStartX:endX])
					num, err := strconv.ParseInt(numberStr, 10, 0)
					if err != nil {
						log.Fatalln("Could not parse number", err)
					}
					inNum = false
					if isAdjacentToSymbol(board, y, numStartX, endX-1) {
						accu += int(num)
					}
				}
			}
		}
	}

	fmt.Println("Part 1 solution:", accu)

	gearAdjacentNumbers := make(map[coord][]int)
	for y := 0; y < len(board); y++ {
		numStartX := -1
		inNum := false
		for x := 0; x < len(board[y]); x++ {
			current := board[y][x]
			if isDigit(current) && !inNum {
				inNum = true
				numStartX = x
			}

			if inNum {
				endX := -1
				if !isDigit(current) {
					endX = x
				} else if x == len(board[y])-1 {
					endX = len(board[y])
				}

				if endX != -1 {
					numberStr := string(board[y][numStartX:endX])
					num, err := strconv.ParseInt(numberStr, 10, 0)
					if err != nil {
						log.Fatalln("Could not parse number", err)
					}
					inNum = false
					for _, coord := range getAdjacentGearCoords(board, y, numStartX, endX-1) {
						gearAdjacentNumbers[coord] = append(gearAdjacentNumbers[coord], int(num))
					}
				}
			}
		}
	}

	accu2 := 0
	for _, nums := range gearAdjacentNumbers {
		if len(nums) == 2 {
			accu2 += nums[0] * nums[1]
		}
	}

	fmt.Println("Part 2 solution:", accu2)
}

// check if there is an adjacent symbol
// startX, endX _inclusive_ indices of number position
func isAdjacentToSymbol(board [][]byte, numberY, startX, endX int) bool {
	minY := max(numberY-1, 0)
	minX := max(startX-1, 0)
	maxY := min(numberY+1, len(board)-1)
	maxX := min(endX+1, len(board[numberY])-1)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if !isDigit(board[y][x]) && board[y][x] != '.' {
				return true
			}
		}
	}
	return false
}

// check if there is an adjacent symbol
// startX, endX _inclusive_ indices of number position
func getAdjacentGearCoords(board [][]byte, numberY, startX, endX int) []coord {
	minY := max(numberY-1, 0)
	minX := max(startX-1, 0)
	maxY := min(numberY+1, len(board)-1)
	maxX := min(endX+1, len(board[numberY])-1)

	coords := make([]coord, 0)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if board[y][x] == '*' {
				coords = append(coords, coord{y: y, x: x})
			}
		}
	}
	return coords
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
