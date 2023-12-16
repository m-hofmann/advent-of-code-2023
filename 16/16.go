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

const (
	N = 1
	W = 1 << 1
	S = 1 << 2
	E = 1 << 3
)

func main() {
	fmt.Println("Starting day 16 ... ")

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

	energized := make([][]int, len(grid))
	for y := range grid {
		energized[y] = make([]int, len(grid[0]))
		for x := range energized[y] {
			energized[y][x] = 0
		}
	}

	for _, line := range grid {
		for _, c := range line {
			fmt.Print(string(c))
		}
		fmt.Println()
	}

	traceBeam(pos{0, 0}, E, grid, &energized)

	energizedCount := 0
	for y, line := range energized {
		for x := range line {
			if energized[y][x] > 0 {
				fmt.Print(string('#'))
				energizedCount++
			} else {
				fmt.Print(string('.'))
			}
		}
		fmt.Println()
	}
	fmt.Println("Part 1 solution:", energizedCount)

}

func traceBeam(curr pos, dir int, grid [][]byte, energized *[][]int) {
	if curr.y < 0 || curr.y >= len(grid) || curr.x < 0 || curr.x >= len(grid[0]) {
		return
	}

	if (*energized)[curr.y][curr.x]&dir != 0 {
		return
	}

	(*energized)[curr.y][curr.x] |= dir

	switch grid[curr.y][curr.x] {
	case '.':
		passThrough(curr, dir, grid, energized)
	case '-':
		switch dir {
		case W, E:
			passThrough(curr, dir, grid, energized)
		case N, S:
			traceBeam(pos{curr.y, curr.x + 1}, E, grid, energized)
			traceBeam(pos{curr.y, curr.x - 1}, W, grid, energized)
		}
	case '|':
		switch dir {
		case W, E:
			traceBeam(pos{curr.y + 1, curr.x}, S, grid, energized)
			traceBeam(pos{curr.y - 1, curr.x}, N, grid, energized)
		case N, S:
			passThrough(curr, dir, grid, energized)
		}
	case '\\':
		switch dir {
		case N:
			traceBeam(pos{curr.y, curr.x - 1}, W, grid, energized)
		case W:
			traceBeam(pos{curr.y - 1, curr.x}, N, grid, energized)
		case S:
			traceBeam(pos{curr.y, curr.x + 1}, E, grid, energized)
		case E:
			traceBeam(pos{curr.y + 1, curr.x}, S, grid, energized)
		}
	case '/':
		switch dir {
		case N:
			traceBeam(pos{curr.y, curr.x + 1}, E, grid, energized)
		case W:
			traceBeam(pos{curr.y + 1, curr.x}, S, grid, energized)
		case S:
			traceBeam(pos{curr.y, curr.x - 1}, W, grid, energized)
		case E:
			traceBeam(pos{curr.y - 1, curr.x}, N, grid, energized)
		}
	}
}

func passThrough(curr pos, dir int, grid [][]byte, energized *[][]int) {
	switch dir {
	case N:
		traceBeam(pos{curr.y - 1, curr.x}, N, grid, energized)
	case W:
		traceBeam(pos{curr.y, curr.x - 1}, W, grid, energized)
	case S:
		traceBeam(pos{curr.y + 1, curr.x}, S, grid, energized)
	case E:
		traceBeam(pos{curr.y, curr.x + 1}, E, grid, energized)
	}
}
