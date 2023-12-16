package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func main() {
	fmt.Println("Starting day 15 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	steps := make([]string, 0)
	sc.Scan()
	line := sc.Text()
	steps = strings.Split(line, ",")

	sum := int64(0)
	for _, step := range steps {
		sum += HASH(step)
	}
	fmt.Println("Part 1 solution:", sum)

	boxes := make(map[int64][]lens)
	for _, step := range steps {
		if step[len(step)-1] == '-' {
			label, _ := strings.CutSuffix(step, "-")
			boxId := HASH(label)
			newBox := make([]lens, 0)
			for _, lens := range boxes[boxId] {
				if lens.label != label {
					newBox = append(newBox, lens)
				}
			}
			boxes[boxId] = newBox
		} else {
			parts := strings.Split(step, "=")
			boxId := HASH(parts[0])
			focalLength, err := strconv.ParseInt(parts[1], 10, 0)
			if err != nil {
				log.Fatalln(err, "Could not parse number")
			}
			newLens := lens{label: parts[0], focalLength: int(focalLength)}
			if box, ok := boxes[boxId]; ok {
				for i, curr := range box {
					if curr.label == parts[0] {
						box[i] = newLens
						goto out
					}
				}
				box = append(box, newLens)
				boxes[boxId] = box
			out:
			} else {
				boxes[boxId] = make([]lens, 0)
				boxes[boxId] = append(boxes[boxId], newLens)
			}
		}
	}

	focusingPower := int64(0)
	for box, lenses := range boxes {
		for slot, lens := range lenses {
			power := (box + 1) * int64(slot+1) * int64(lens.focalLength)
			fmt.Println(lens.label, ":", power)
			focusingPower += power
		}
	}
	fmt.Println("Part 2 solution:", focusingPower)
}

func HASH(s string) int64 {
	val := int64(0)
	for _, c := range s {
		val += int64(c)
		val *= 17
		val %= 256
	}
	return val
}

func printBoxes(b map[int64][]lens) {
	for k, v := range b {
		fmt.Printf("Box %3d: %v\n", k, v)
	}
}
