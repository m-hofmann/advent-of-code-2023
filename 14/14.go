package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

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

	tilted := tiltGridNorthNaive(grid)

	totalLoad := 0
	totalLoad = calcLoad(tilted)

	fmt.Println("Part 1 solution:", totalLoad)

	tilted = grid
	start := time.Now()
	seenAtCycle := make(map[string]int)
	cycleAndLoad := make(map[int]int)
	iterations := 1000000000
	cycleLength := -1
	for i := 0; i < iterations; i++ {
		tilted = tiltGridNorthNaive(tilted)
		tilted = tiltGridWestNaive(tilted)
		tilted = tiltGridSouthNaive(tilted)
		tilted = tiltGridEastNaive(tilted)
		if i > 0 && i%100000 == 0 {
			deltaT := time.Since(start)
			estimatedLeft := int64(iterations/i)*deltaT.Abs().Nanoseconds() - deltaT.Abs().Nanoseconds()
			fmt.Printf("Reached %6d after %s (estimated left: %s)\n", i, time.Since(start), time.Duration(estimatedLeft))
		}
		cycleAndLoad[i] = calcLoad(tilted)
		asString := stringify(tilted)
		if val, ok := seenAtCycle[asString]; ok {
			cycleLength = i - val
			fmt.Printf("In Cycle %d: Already saw this grid at cycle %d!\n\n", i, val)
			fmt.Println("Cycle length is", cycleLength)
			futureLoad := cycleAndLoad[val+(iterations-val-1)%cycleLength]
			fmt.Println("Part 2 solution:", futureLoad)
			os.Exit(0)
		} else {
			seenAtCycle[asString] = i
		}
	}
	fmt.Println("Part 2 solution: UNKNOWN")
}

func stringify(grid [][]byte) string {
	bldr := strings.Builder{}
	for _, line := range grid {
		bldr.Write(line)
	}
	return bldr.String()
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

func tiltGridWestNaive(grid [][]byte) [][]byte {
	movement := 1
	for movement > 0 {
		movement = 0
		newGrid := make([][]byte, len(grid))
		for y, line := range grid {
			newGrid[y] = make([]byte, len(grid[y]))
			for x, c := range line {
				if c == 'O' {
					newX := x - 1
					newY := y
					if newX >= 0 && newGrid[newY][newX] == '.' {
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

func tiltGridSouthNaive(grid [][]byte) [][]byte {
	movement := 1
	for movement > 0 {
		movement = 0
		newGrid := make([][]byte, len(grid))
		for y := len(grid) - 1; y >= 0; y-- {
			line := grid[y]
			newGrid[y] = make([]byte, len(grid[y]))
			for x, c := range line {
				if c == 'O' {
					newX := x
					newY := y + 1
					if newY < len(grid) && newGrid[newY][newX] == '.' {
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

func tiltGridEastNaive(grid [][]byte) [][]byte {
	movement := 1
	for movement > 0 {
		movement = 0
		newGrid := make([][]byte, len(grid))
		for y, line := range grid {
			newGrid[y] = make([]byte, len(grid[y]))
			for x := len(line) - 1; x >= 0; x-- {
				if grid[y][x] == 'O' {
					newX := x + 1
					newY := y
					if newX < len(line) && newGrid[newY][newX] == '.' {
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
