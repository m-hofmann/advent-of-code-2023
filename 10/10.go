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

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
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

	fmt.Println("Part 1 solution:", (1+getMaxReachableDist(startPos, grid, 0))/2)
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
