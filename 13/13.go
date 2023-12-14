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

		vSymmetry := getSymm(vsig)
		hSymmetry := getSymm(hsig)
		if vSymmetry > -1 && hSymmetry == -1 {
			sum += vSymmetry
		} else if vSymmetry == -1 && hSymmetry > -1 {
			sum += 100 * hSymmetry
		} else {
			fmt.Errorf("Pattern %d: vsig %v, hsig %v\n", i, vsig, hsig)
			fmt.Errorf("Vsig symmetry %v\n", vSymmetry)
			fmt.Errorf("Hsig symmetry %v\n", hSymmetry)
			return
		}
	}
	fmt.Println("Part 1 solution:", sum)

	sum = 0
	for i, grid := range grids {
		vsig := getVSig(grid)
		hsig := getHSig(grid)

		vSymmetry := getSymmWithSmudge(vsig)
		hSymmetry := getSymmWithSmudge(hsig)
		if vSymmetry > -1 && hSymmetry == -1 {
			sum += vSymmetry
		} else if vSymmetry == -1 && hSymmetry > -1 {
			sum += 100 * hSymmetry
		} else {
			fmt.Printf("Pattern %d: vsig %v, hsig %v\n", i, vsig, hsig)
			fmt.Printf("Vsig symmetry %v\n", vSymmetry)
			fmt.Printf("Hsig symmetry %v\n", hSymmetry)
		}
	}
	fmt.Println("Part 2 solution:", sum)
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
func getSymm(arr []int) int {
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

// returns -1 if no symmetry, x axis else
func getSymmWithSmudge(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		j := i
		k := i + 1
		smudges := 0

		for arr[j] == arr[k] || smudgeEqual(arr[j], arr[k]) {
			if smudgeEqual(arr[j], arr[k]) {
				smudges++
			}
			if j == 0 || k == len(arr)-1 {
				if smudges == 1 {
					return i + 1
				} else {
					goto out
				}
			}
			j--
			k++
		}
	out:
	}
	return -1
}

func smudgeEqual(a, b int) bool {
	n := a ^ b
	// exactly one bit (the smudge) different
	if a != b && (n&(n-1)) == 0 {
		return true
	}
	return false
}
