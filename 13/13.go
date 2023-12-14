package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type record struct {
	report    []byte
	checksums []int
}

func main() {
	fmt.Println("Starting day 13 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	grids := make([][][]byte, 0)
	grid := make([][]byte, 0)
	for sc.Scan() {
		line := sc.Text()
		if line != "" {
			grid = append(grid, []byte(line))
		} else {
			grids = append(grids, grid)
			grid = make([][]byte, 0)
		}
	}
	grids = append(grids, grid)

	sum := 0
	for i, grid := range grids {
		vsig := getVSig(grid)
		hsig := getHSig(grid)
		fmt.Printf("Pattern %d: vsig %v, hsig %v\n", i, vsig, hsig)
		fmt.Printf("Vsig symmetry %v\n", getVSymmAxis(vsig))
		fmt.Printf("Hsig symmetry %v\n", getHSymmAxis(hsig))

		vSymmetry := getVSymmAxis(vsig)
		hSymmetry := getHSymmAxis(hsig)
		if vSymmetry > -1 && hSymmetry == -1 {
			sum += vSymmetry
		} else if vSymmetry == -1 && hSymmetry > -1 {
			sum += 100 * hSymmetry
		} else {
			fmt.Errorf("Pattern %d: vsig %v, hsig %v\n", i, vsig, hsig)
			fmt.Errorf("Vsig symmetry", getVSymmAxis(vsig))
			fmt.Errorf("Hsig symmetry", getHSymmAxis(hsig))
			return
		}
	}
	fmt.Println("Part 1 solution:", sum)
}

func getHSig(grid [][]byte) []int {
	hsig := make([]int, len(grid))
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			hsig[y] <<= 1
			if grid[y][x] == '#' {
				hsig[y]++
			}
		}
	}
	return hsig
}

func getVSig(grid [][]byte) []int {
	vsig := make([]int, len(grid[0]))
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			vsig[x] <<= 1
			if grid[y][x] == '#' {
				vsig[x]++
			}
		}
	}
	return vsig
}

// returns -1 if no symmetry, x axis else
func getVSymmAxis(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		j := i
		k := i + 1
		for arr[j] == arr[k] {
			if j == 0 || k == len(arr)-1 {
				return i + 1
			}
			j--
			k++
		}
	}
	return -1
}

// returns -1 if no symmetry, y axis else
func getHSymmAxis(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		j := i
		k := i + 1
		for arr[j] == arr[k] {
			if j == 0 || k == len(arr)-1 {
				return i + 1
			}
			j--
			k++
		}
	}
	return -1
}
