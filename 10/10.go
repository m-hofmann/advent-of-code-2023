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

func main() {
	fmt.Println("Starting day 10 ... ")

	f, err := os.OpenFile("./data/part1b.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	grid := make([][]byte, 0)

	startPos := pos{-1, -1}
	adjacency := make(map[pos]map[pos]struct{})
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

	for y := range grid {
		for x := range grid[y] {
			currPos := pos{y, x}
			currC := grid[y][x]
			if currC == '.' {
				continue
			}

			for _, reachable := range getReachable(currPos, currC) {
				if reachable.x < 0 || reachable.x >= len(grid[y]) ||
					reachable.y < 0 || reachable.y >= len(grid) {
					// out of bounds coordinate
					continue
				}

				// only add pipe-tiles to the adjacency matrix
				if grid[reachable.y][reachable.x] == '.' {
					continue
				}

				if _, ok := adjacency[currPos]; !ok {
					adjacency[currPos] = make(map[pos]struct{})
				}
				adjacency[currPos][reachable] = struct{}{}
				if _, ok := adjacency[reachable]; !ok {
					adjacency[reachable] = make(map[pos]struct{})
				}
				adjacency[reachable][currPos] = struct{}{}
			}
		}
	}

	dist := make(map[pos]int)
	dist[startPos] = 0
	prev := make(map[pos]pos)
	toCheck := make(map[pos]struct{})
	toCheck[startPos] = struct{}{}

	for len(toCheck) != 0 {
		minDist := math.MaxInt
		var cand pos
		for curr := range toCheck {
			if dist[curr] < minDist {
				minDist = dist[curr]
				cand = curr
			}
		}
		delete(toCheck, cand)

		for neighbor := range adjacency[cand] {
			newDist := dist[cand] + 1
			neighborDist := math.MaxInt
			if val, ok := dist[neighbor]; ok {
				neighborDist = val
			}
			if newDist < neighborDist {
				prev[neighbor] = cand
				dist[neighbor] = newDist
				toCheck[neighbor] = struct{}{}
			}
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if val, ok := dist[pos{y, x}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	maxDist := math.MinInt

	for _, dist := range dist {
		maxDist = max(dist, maxDist)
	}

	fmt.Println("Part 1 solution:", maxDist)

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
		return []pos{}
	default:
		return []pos{}
	}
}
