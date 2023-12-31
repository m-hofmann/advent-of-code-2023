package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	y int
	x int
}

func main() {
	fmt.Println("Starting day 10 ... ")

	f, err := os.OpenFile("./data/demo4.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	grid := make([][]byte, 0)

	startPos := pos{-1, -1}
	y := 0
	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, []byte(line))
		for x, c := range []byte(line) {
			curr := pos{y, x}
			if c == 'S' {
				startPos = curr
			}
		}
		y++
	}
	gridCopy := make([][]byte, len(grid))
	for i := range grid {
		gridCopy[i] = make([]byte, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	fmt.Println("Part 1 solution:", (1+getMaxReachableDist(startPos, gridCopy, 0))/2)
	for i := range grid {
		gridCopy[i] = make([]byte, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	// gridCopy: Contains traced-out 'X' path
	// grid: Contains original tiles for looking up "walls"
	for i := range grid {
		gridCopy[i] = make([]byte, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	traceLoop(startPos, &gridCopy, 0)

	for _, line := range gridCopy {
		for _, c := range line {
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	fmt.Println()
	totalEnclosed := 0
	for y, line := range grid {
		within := false
		last := '.'
		for x, c := range line {
			if gridCopy[y][x] == 'X' {
				if c == '|' {
					within = !within
				} else if c == 'F' {
					last = 'F'
				} else if c == 'L' {
					last = 'L'
				} else if c == 'J' && last == 'F' {
					within = !within
				} else if c == '7' && last == 'L' {
					within = !within
				}
			} else if within {
				totalEnclosed++
				grid[y][x] = 'I'
			} else if !within {
			}
			fmt.Print(string(grid[y][x]))
		}
		fmt.Println()
	}
	fmt.Println("Part 2 solution:", totalEnclosed)
}

func getMaxReachableDist(from pos, grid [][]byte, level int) int {
	maxDist := level
	//fmt.Println("Nesting level", level, "at", from)
	for _, reachable := range getReachable(from, grid[from.y][from.x]) {
		grid[from.y][from.x] = '.'
		if grid[reachable.y][reachable.x] != 'S' && grid[reachable.y][reachable.x] != '.' &&
			reachable.y >= 0 && reachable.y < len(grid) && reachable.x >= 0 && reachable.x < len(grid[reachable.y]) {
			return getMaxReachableDist(reachable, grid, level+1)
		}
	}
	return maxDist
}

func traceLoop(from pos, grid *[][]byte, level int) int {
	maxDist := level
	//fmt.Println("Nesting level", level, "at", from)
	for _, reachable := range getReachable(from, (*grid)[from.y][from.x]) {
		(*grid)[from.y][from.x] = 'X'
		if (*grid)[reachable.y][reachable.x] != 'S' && (*grid)[reachable.y][reachable.x] != 'X' && (*grid)[reachable.y][reachable.x] != '.' &&
			reachable.y >= 0 && reachable.y < len(*grid) && reachable.x >= 0 && reachable.x < len((*grid)[reachable.y]) {
			return traceLoop(reachable, grid, level+1)
		}
	}
	return maxDist
}
func traceLoopOld(from pos, grid *[][]byte) bool {
	//fmt.Println("Nesting level", level, "at", from)
	for _, reachable := range getReachable(from, (*grid)[from.y][from.x]) {
		if (*grid)[reachable.y][reachable.x] == 'S' {
			return true
		}
		(*grid)[reachable.y][reachable.x] = 'X'
		if (*grid)[reachable.y][reachable.x] != '.' && (*grid)[reachable.y][reachable.x] != 'X' &&
			reachable.y >= 0 && reachable.y < len(*grid) && reachable.x >= 0 && reachable.x < len((*grid)[reachable.y]) {
			if traceLoopOld(reachable, grid) {
				return true
			}
		}
	}
	return false
}

func getReachable(from pos, tile byte) []pos {
	switch tile {
	case '|':
		return []pos{{from.y - 1, from.x}, {from.y + 1, from.x}}
	case '-':
		return []pos{{from.y, from.x - 1}, {from.y, from.x + 1}}
	case 'L':
		return []pos{{from.y - 1, from.x}, {from.y, from.x + 1}}
	case 'J':
		return []pos{{from.y - 1, from.x}, {from.y, from.x - 1}}
	case '7':
		return []pos{{from.y + 1, from.x}, {from.y, from.x - 1}}
	case 'F':
		return []pos{{from.y + 1, from.x}, {from.y, from.x + 1}}
	case '.':
		return []pos{}
	case 'S':
		// ignore for now, we treat that separately
		// by using the fact the if S reachable from neighbor, then neighbor reachable from S
		return []pos{{from.y - 1, from.x}, {from.y + 1, from.x}, {from.y, from.x - 1}, {from.y, from.x + 1}}
	default:
		return []pos{}
	}
}
