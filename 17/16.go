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
	pos pos
	dir int
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

	dist, previous := findShortestPathLength(pos{0, 0}, &grid)
	target := pos{len(grid) - 1, len(grid[0]) - 1}
	curr := target
	for !(curr.x == 0 && curr.y == 0) {
		fmt.Println(curr)
		dirChar := 'X'
		prev := previous[curr.y][curr.x]
		switch prev.dir {
		case N:
			dirChar = '^'
		case W:
			dirChar = '<'
		case S:
			dirChar = 'v'
		case E:
			dirChar = '>'
		}
		grid[curr.y][curr.x] = int(dirChar - '0')
		curr = prev.pos
	}
	fmt.Println("Distances ---")
	for y := range dist {
		for _, c := range dist[y] {
			fmt.Printf("%3d ", c)
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
	fmt.Println("Part 1 solution:", dist[target.y][target.x])

}

// returns tuple of (distance matrix, prev/dir matrix)
func findShortestPathLength(from pos, grid *[][]int) ([][]int, [][]*posVector) {
	// heat shed
	dist := make([][]int, len(*grid))
	// source node + incoming direction
	prev := make([][]*posVector, len(*grid))
	for y := range *grid {
		dist[y] = make([]int, len((*grid)[0]))
		prev[y] = make([]*posVector, len((*grid)[0]))
		for x := range (*grid)[y] {
			dist[y][x] = math.MaxInt
			prev[y][x] = nil
		}
	}
	dist[from.y][from.x] = 0
	prev[from.y][from.x] = &posVector{pos: pos{-1, -1}, dir: E}
	unvisited := make(map[pos]struct{})
	unvisited[from] = struct{}{}

	for len(unvisited) > 0 {
		node, _ := selectMin(&unvisited, &dist)
		delete(unvisited, node)

		for _, next := range possibleNeighbors(node, prev[node.y][node.x].dir, grid) {
			// prune > 3x same direction
			pred := node
			dir := 0
			i := 0
			for ; (pred.y != -1 && pred.x != -1) && i < 3; i++ {
				vec := prev[pred.y][pred.x]
				if vec != nil {
					dir |= vec.dir
					pred = vec.pos
				} else {
					break
				}
			}
			if i == 3 && (dir&(dir-1)) == 0 && dir == next.dir {
				continue
			}

			alt := dist[node.y][node.x] + (*grid)[next.pos.y][next.pos.x]
			if alt < dist[next.pos.y][next.pos.x] {
				dist[next.pos.y][next.pos.x] = alt
				prev[next.pos.y][next.pos.x] = &posVector{node, next.dir}
				unvisited[next.pos] = struct{}{}
			}
		}
	}

	return dist, prev
}

func selectMin(unvisited *map[pos]struct{}, dist *[][]int) (pos, bool) {
	minVal := math.MaxInt
	minNode := pos{-1, -1}
	for k := range *unvisited {
		if (*dist)[k.y][k.x] < minVal {
			minVal = (*dist)[k.y][k.x]
			minNode = k
		}
	}

	if minVal != math.MaxInt {
		return minNode, true
	} else {
		return minNode, false
	}
}

func possibleNeighborsUnconstrained(from pos, dir int, grid *[][]int) []posVector {
	uncheckedPos := make([]posVector, 0)
	uncheckedPos = append(uncheckedPos, posVector{pos{from.y, from.x - 1}, W},
		posVector{pos{from.y - 1, from.x}, N},
		posVector{pos{from.y, from.x + 1}, E},
		posVector{pos{from.y + 1, from.x}, S})

	checkedPos := make([]posVector, 0)
	for _, elem := range uncheckedPos {
		if elem.pos.x >= 0 && elem.pos.x < len((*grid)[0]) && elem.pos.y >= 0 && elem.pos.y < len(*grid) {
			checkedPos = append(checkedPos, elem)
		}
	}
	return checkedPos
}

// coming from direction dir onto from, get valid neighbor cells (coordinate safe)
func possibleNeighbors(from pos, dir int, grid *[][]int) []posVector {
	uncheckedPos := make([]posVector, 0)
	switch dir {
	case N:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.y, from.x - 1}, W}, posVector{pos{from.y - 1, from.x}, N}, posVector{pos{from.y, from.x + 1}, E})
	case W:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.y, from.x - 1}, W}, posVector{pos{from.y - 1, from.x}, N}, posVector{pos{from.y + 1, from.x}, W})
	case S:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.y, from.x - 1}, W}, posVector{pos{from.y + 1, from.x}, S}, posVector{pos{from.y, from.x + 1}, E})
	case E:
		uncheckedPos = append(uncheckedPos, posVector{pos{from.y, from.x + 1}, E}, posVector{pos{from.y + 1, from.x}, S}, posVector{pos{from.y - 1, from.x}, N})
	default:
	}

	checkedPos := make([]posVector, 0)
	for _, elem := range uncheckedPos {
		if elem.pos.x >= 0 && elem.pos.x < len((*grid)[0]) && elem.pos.y >= 0 && elem.pos.y < len(*grid) {
			checkedPos = append(checkedPos, elem)
		}
	}
	return checkedPos
}
