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
	fmt.Println("Starting day 11 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	grid := make([][]byte, 0)

	y := 0
	// raw arrays would be more efficient, but I'd need to now sizes beforehand
	// so - what gives?
	hasXGalaxy := make(map[int]bool)
	hasYGalaxy := make(map[int]bool)
	originalGalaxyPos := make(map[int]pos)
	galaxyCount := 0
	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, []byte(line))
		hasYGalaxy[y] = false
		for x, c := range line {
			if c == '#' {
				hasXGalaxy[x] = true
				hasYGalaxy[y] = true
				originalGalaxyPos[galaxyCount] = pos{y, x}
				galaxyCount++
			}
		}
		y++
	}

	yGalaxyCount := 0
	for _, v := range hasYGalaxy {
		if v {
			yGalaxyCount++
		}
	}
	xGalaxyCount := 0
	for i := 0; i < len(grid[0]); i++ {
		if hasXGalaxy[i] {
			xGalaxyCount++
		} else {
			hasXGalaxy[i] = false
		}
	}

	// expand universe
	ySize := len(grid) + (len(grid) - yGalaxyCount)
	xSize := len(grid[0]) + (len(grid[0]) - xGalaxyCount)
	expanded := make([][]byte, ySize)
	actualY := 0
	for y := 0; y < len(grid); y++ {
		expanded[actualY] = make([]byte, xSize)
		if !hasYGalaxy[y] {
			for x := 0; x < xSize; x++ {
				expanded[actualY][x] = '.'
			}
			actualY++
			expanded[actualY] = make([]byte, xSize)
			for x := 0; x < xSize; x++ {
				expanded[actualY][x] = '.'
			}
			actualY++
			continue
		}
		actualX := 0
		for x := 0; x < len(grid[0]); x++ {
			if hasXGalaxy[x] {
				expanded[actualY][actualX] = grid[y][x]
			} else {
				expanded[actualY][actualX] = grid[y][x]
				actualX++
				expanded[actualY][actualX] = '.'
			}
			actualX++
		}
		actualY++
	}

	galaxyIdx := 0
	galaxies := make(map[int]pos)
	for y := range expanded {
		for x, c := range expanded[y] {
			if c == '#' {
				galaxies[galaxyIdx] = pos{y, x}
				galaxyIdx++
			}
		}
	}

	fmt.Println("have:", galaxyIdx, "galaxies")
	sumLength := 0
	for i := 0; i < galaxyCount; i++ {
		for k := 0; k < galaxyCount; k++ {
			if k == i {
				continue
			}
			sumLength += manhattanDist(galaxies[i], galaxies[k])
		}
	}
	fmt.Println("Part 1 solution:", sumLength/2)

	for i := 0; i < galaxyCount; i++ {
		curr := originalGalaxyPos[i]
		xOffset := curr.x
		yOffset := curr.y
		for y := 0; y < curr.y; y++ {
			if !hasYGalaxy[y] {
				yOffset += 1000000 - 1
			}
		}
		for x := 0; x < curr.x; x++ {
			if !hasXGalaxy[x] {
				xOffset += 1000000 - 1
			}
		}
		galaxies[i] = pos{yOffset, xOffset}
	}
	sumLength = 0
	for i := 0; i < galaxyCount; i++ {
		for k := 0; k < galaxyCount; k++ {
			if k == i {
				continue
			}
			sumLength += manhattanDist(galaxies[i], galaxies[k])
		}
	}
	fmt.Println("Part 2 solution:", sumLength/2)

}

func manhattanDist(a, b pos) int {
	return abs(a.y-b.y) + abs(a.x-b.x)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
