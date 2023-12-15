package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type record struct {
	report    []byte
	checksums []int
}

func main() {
	fmt.Println("Starting day 14 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	grid := make([][]byte, 0)
	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, []byte(line))
	}
	fmt.Println("--- input ---")
	for _, line := range grid {
		fmt.Println(string(line))
	}
	tilted := tiltGridNorthNaive(grid)
	fmt.Println("--- Tilted ---")
	for _, line := range tilted {
		fmt.Println(string(line))
	}

	totalLoad := 0

	totalLoad = calcLoad(tilted)

	fmt.Println("Part 1 solution:", totalLoad)
}

func tiltGridNorthNaive(grid [][]byte) [][]byte {
	movement := 1
	for movement > 0 {
		movement = 0
		newGrid := make([][]byte, len(grid))
		for y, line := range grid {
			newGrid[y] = make([]byte, len(grid[y]))
			for x, c := range line {
				if c == 'O' {
					newX := x
					newY := y - 1
					if newY >= 0 && newGrid[newY][newX] == '.' {
						movement++
						newGrid[newY][newX] = 'O'
						newGrid[y][x] = '.'
					} else {
						newGrid[y][x] = 'O'
					}
				} else {
					newGrid[y][x] = grid[y][x]
				}
			}
		}
		grid = newGrid
	}
	return grid
}

func calcLoad(grid [][]byte) int {
	height := len(grid)
	sum := 0
	for y, line := range grid {
		for _, c := range line {
			if c == 'O' {
				sum += height - y
			}
		}
	}
	return sum
}
