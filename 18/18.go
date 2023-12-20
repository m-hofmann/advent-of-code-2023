package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	dir   int
	dist  int
	color string
}

type coord struct {
	x int
	y int
}

type edge struct {
	from  coord
	to    coord
	color string
}

const (
	U = 1
	L = 1 << 1
	D = 1 << 2
	R = 1 << 3
)

func main() {
	fmt.Println("Starting day 18 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	instructions := make([]instruction, 0)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, " ")

		dist, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			log.Fatalln("Could not parse distance for digging instruction", err)
		}
		instructions = append(instructions,
			instruction{
				dir:   parseDir(parts[0][0]),
				dist:  int(dist),
				color: parts[2][1:8],
			})
	}

	totalArea := getTotalAreaPart1(instructions)

	fmt.Println("Part 1 solution:", totalArea)

	newInstructions := make([]instruction, len(instructions))
	for i, instr := range instructions {
		dist, err := strconv.ParseInt(instr.color[1:6], 16, 0)
		if err != nil {
			log.Fatalln("Could not parse hex distance from color", err)
		}

		dir := 0
		lastChar := instr.color[len(instr.color)-1]
		switch lastChar {
		case '0':
			dir = R
		case '1':
			dir = D
		case '2':
			dir = L
		case '3':
			dir = U
		}
		newInstructions[i] = instruction{dir: dir, color: instr.color, dist: int(dist)}
		fmt.Println("New true instruction", newInstructions[i])
	}
	totalArea = getTotalAreaPart1(newInstructions)

	fmt.Printf("Part 2 solution %f\n", totalArea)
}

func getTotalAreaPart1(instructions []instruction) float64 {
	edges := make([]edge, 0)
	currPos := coord{0, 0}
	bounds := coord{}
	for i := 0; i < len(instructions); i++ {
		instr := instructions[i]
		target := coord{}
		switch instr.dir {
		case U:
			target = coord{currPos.x, currPos.y - instr.dist}
		case L:
			target = coord{currPos.x - instr.dist, currPos.y}
		case D:
			target = coord{currPos.x, currPos.y + instr.dist}
		case R:
			target = coord{currPos.x + instr.dist, currPos.y}
		}

		edge := edge{
			from:  currPos,
			to:    target,
			color: instr.color,
		}
		edges = append(edges, edge)
		currPos = target

		bounds.y = max(currPos.y, bounds.y)
		bounds.x = max(currPos.x, bounds.x)
	}

	currPos = coord{}
	coveredByTrench := outlineLength(instructions)
	points := make([]coord, 0)
	for i := 0; i < len(edges); i++ {
		points = append(points, edges[i].from)
	}
	areaCovered := shoelace(points)
	// Pick's theorem
	// A from shoelace theorem
	// A = i + b/2 - 1
	// have A, b, want i
	// A - b/2 +1 = i
	i := areaCovered - float64(coveredByTrench)/2 + 1
	totalArea := i + float64(coveredByTrench)
	return totalArea
}

func outlineLength(instructions []instruction) int {
	coveredByTrench := 0
	for _, instr := range instructions {
		coveredByTrench += instr.dist
	}
	return coveredByTrench
}

func shoelace(points []coord) float64 {
	areaCovered := 0
	fmt.Println(len(points))
	for i := 0; i < len(points); i++ {
		a := points[i%len(points)]
		b := points[(i+1)%(len(points))]
		areaCovered += determinant(a, b)
	}

	return float64(areaCovered) / 2.0
}

func determinant(a, b coord) int {
	return a.x*b.y - a.y*b.x
}

func parseDir(c byte) int {
	switch c {
	case 'U':
		return U
	case 'L':
		return L
	case 'D':
		return D
	case 'R':
		return R
	default:
		return -1
	}
}
