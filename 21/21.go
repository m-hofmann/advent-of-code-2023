package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const (
	UNREACHABLE = math.MaxInt
	UNSEEN      = math.MaxInt - 1
)

type coord struct {
	x int
	y int
}

type state struct {
	dist int
	node coord
}

func main() {
	fmt.Println("Starting day 21 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	grid := make([][]byte, 0)

	startPos := coord{-1, -1}
	y := 0
	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, []byte(line))
		for x, c := range []byte(line) {
			curr := coord{y, x}
			if c == 'S' {
				startPos = curr
			}
		}
		y++
	}
	gridCopy := make([][]byte, len(grid))
	dist := make([][]int, len(grid))
	for y := range grid {
		gridCopy[y] = make([]byte, len(grid[y]))
		dist[y] = make([]int, len(grid[y]))
		copy(gridCopy[y], grid[y])
		for x, c := range grid[y] {
			if c == '#' {
				dist[y][x] = UNREACHABLE
			} else {
				dist[y][x] = UNSEEN
			}
		}
	}
	dist[startPos.y][startPos.x] = 0

	reachableExactlySteps := make(map[int]map[coord]struct{})
	unvisited := make([]state, 0)
	unvisited = append(unvisited, state{0, startPos})
	safeAssign(&reachableExactlySteps, 0, startPos)

	for len(unvisited) > 0 {
		curr := unvisited[0]
		unvisited = unvisited[1:]

		for _, neighbor := range reachableNeighbors(curr.node, &grid) {
			if grid[neighbor.y][neighbor.x] != '#' {
				newDist := dist[curr.node.y][curr.node.x] + 1
				oldDist := dist[neighbor.y][neighbor.x]
				if newDist < oldDist {
					if oldDist < UNSEEN {
						delete(reachableExactlySteps[oldDist], neighbor)
					}
					safeAssign(&reachableExactlySteps, newDist, neighbor)
					dist[neighbor.y][neighbor.x] = newDist

					unvisited = append(unvisited, state{newDist, neighbor})
				}
			}
		}
	}

	stepsExactly := 64
	reachableCount := 0
	for i := 0; i <= stepsExactly; i++ {
		diff := stepsExactly - i
		// if difference between number of steps and current dist to node (= i) is even
		// then we can go diff/2 steps to another node and return in diff/2 steps -> current node is reachable
		if diff%2 == 0 {
			reachableExactly := len(reachableExactlySteps[i])
			fmt.Println("At", i, "steps we add", reachableExactly, "fields")
			reachableCount += reachableExactly
		}
	}

	fmt.Println("Part 1 solution", reachableCount)
}

func printReachableInStep(i int, grid *[][]byte, m *map[int]map[coord]struct{}) {
	for y := range *grid {
		for x := range (*grid)[y] {
			if _, ok := (*m)[i][coord{x, y}]; ok {
				fmt.Print("O")
			} else {
				fmt.Print(string((*grid)[y][x]))
			}
		}
		fmt.Println()
	}

}

func reachableNeighbors(from coord, grid *[][]byte) []coord {
	coords := []coord{
		{from.x + 1, from.y},
		{from.x - 1, from.y},
		{from.x, from.y + 1},
		{from.x, from.y - 1},
	}

	legalCoords := make([]coord, 0)

	for _, cell := range coords {
		if cell.x >= 0 && cell.x < len((*grid)[0]) && cell.y >= 0 && cell.y < len(*grid) {
			legalCoords = append(legalCoords, cell)
		}
	}

	return legalCoords
}

func debugPrint(grid *[][]byte) {
	for _, line := range *grid {
		fmt.Println(string(line))
	}
}

func safeAssign(m *map[int]map[coord]struct{}, steps int, pos coord) {
	if val, ok := (*m)[steps]; ok {
		val[pos] = struct{}{}
	} else {
		(*m)[steps] = make(map[coord]struct{})
		(*m)[steps][pos] = struct{}{}
	}
}
