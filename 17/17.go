package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type pos struct {
	y int
	x int
}

type posVector struct {
	pos       pos
	dir       int
	runLength int
}

const (
	N = 1
	W = 1 << 1
	S = 1 << 2
	E = 1 << 3
)

func main() {
	fmt.Println("Starting day 17 ... ")

	f, err := os.OpenFile("./data/demo.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	y := 0
	grid := make([][]int, 0)
	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, make([]int, len(line)))
		for x, c := range line {
			grid[y][x] = int(c - '0')
		}
		y++
	}

	dist, previous, shortest := findShortestPathLength(posVector{pos{0, 0}, E, 0}, pos{len(grid) - 1, len(grid[0]) - 1}, &grid)
	curr := shortest[len(grid)-1][len(grid[0])-1]
	for curr != previous[curr] {
		fmt.Println(curr)
		dirChar := 'X'
		prev := previous[curr]
		switch curr.dir {
		case N:
			dirChar = '^'
		case W:
			dirChar = '<'
		case S:
			dirChar = 'v'
		case E:
			dirChar = '>'
		}
		grid[curr.pos.y][curr.pos.x] = int(dirChar - '0')
		curr = prev
	}
	fmt.Println("Distances ---")
	for y := range shortest {
		for x := range shortest[y] {
			fmt.Printf("%3d ", dist[shortest[y][x].pos])
		}
		fmt.Println()
	}

	fmt.Println("Path ---")
	for y := range grid {
		for _, c := range grid[y] {
			fmt.Print(string(c + '0'))
		}
		fmt.Println()
	}
	fmt.Println("Part 1 solution:", dist[shortest[len(grid)-1][len(grid[0])-1].pos])
}

// returns tuple of (distance matrix, prev/dir matrix)
func findShortestPathLength(from posVector, to pos, grid *[][]int) (map[pos]int, map[posVector]posVector, [][]posVector) {
	// heat shed
	dist := make(map[pos]int)
	// source node + incoming direction
	prev := make(map[posVector]posVector)
	shortest := make([][]posVector, len(*grid))
	for y := range *grid {
		shortest[y] = make([]posVector, len((*grid)[0]))
	}
	dist[from.pos] = 0
	prev[from] = from
	unvisited := make([]posVector, 0)
	unvisited = append(unvisited, from)

	for len(unvisited) > 0 {
		source := unvisited[0]
		unvisited = unvisited[1:]
		if source.pos.y == 11 && source.pos.x == 11 {
			fmt.Println("break", source)
			fmt.Println("Neighbors:", possibleNeighbors(source, grid))
		}

		for _, next := range possibleNeighbors(source, grid) {
			//fmt.Println("For", source.pos, "next is", next.pos)

			alt := dist[source.pos] + (*grid)[next.pos.y][next.pos.x]
			currentDist := math.MaxInt
			if val, ok := dist[next.pos]; ok {
				currentDist = val
			}
			if alt < currentDist {
				dist[next.pos] = alt
				shortest[next.pos.y][next.pos.x] = next
				prev[next] = source
				unvisited = append(unvisited, next)
			}
		}
	}

	return dist, prev, shortest
}

func selectMin(unvisited *[]posVector, dist *map[posVector]int) (posVector, bool) {
	minVal := math.MaxInt
	var minNode *posVector
	for _, k := range *unvisited {
		if (*dist)[k] < minVal {
			minVal = (*dist)[k]
			minNode = &k
		}
	}

	if minVal != math.MaxInt {
		return *minNode, true
	} else {
		return *minNode, false
	}
}

func selectAny(unvisited *map[posVector]struct{}, dist *map[posVector]int) (posVector, bool) {
	for k := range *unvisited {
		return k, true
	}
	return posVector{}, false
}

// coming from direction dir onto from, get valid neighbor cells (coordinate safe)
func possibleNeighbors(from posVector, grid *[][]int) []posVector {
	uncheckedPos := make([]posVector, 0)
	switch from.dir {
	case N:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.pos.y, from.pos.x - 1}, W, 1},
			posVector{pos{from.pos.y - 1, from.pos.x}, N, from.runLength + 1},
			posVector{pos{from.pos.y, from.pos.x + 1}, E, 1})
	case W:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.pos.y, from.pos.x - 1}, W, from.runLength + 1},
			posVector{pos{from.pos.y - 1, from.pos.x}, N, 1},
			posVector{pos{from.pos.y + 1, from.pos.x}, S, 1})
	case S:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.pos.y, from.pos.x - 1}, W, 1},
			posVector{pos{from.pos.y + 1, from.pos.x}, S, from.runLength + 1},
			posVector{pos{from.pos.y, from.pos.x + 1}, E, 1})
	case E:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.pos.y, from.pos.x + 1}, E, from.runLength + 1},
			posVector{pos{from.pos.y + 1, from.pos.x}, S, 1},
			posVector{pos{from.pos.y - 1, from.pos.x}, N, 1})
	default:
	}

	checkedPos := make([]posVector, 0)
	for _, elem := range uncheckedPos {
		if elem.pos.x >= 0 && elem.pos.x < len((*grid)[0]) && elem.pos.y >= 0 && elem.pos.y < len(*grid) && elem.runLength <= 3 {
			checkedPos = append(checkedPos, elem)
		}
	}
	return checkedPos
}
