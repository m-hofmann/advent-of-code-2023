package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type draw struct {
	red   int
	green int
	blue  int
}

type game struct {
	number int
	draws  []draw
}

func main() {
	fmt.Println("Starting day 02 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	gameNoRegex := regexp.MustCompile("Game (?P<game>\\d+):")
	colorRegex := regexp.MustCompile(" (?P<count>\\d+) (?P<color>red|green|blue)")
	games := make([]game, 0)
	games = parseInput(sc, gameNoRegex, colorRegex, games)

	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	accu1 := 0
	for _, game := range games {
		possible := true
		for _, draw := range game.draws {
			if draw.red > maxRed || draw.green > maxGreen || draw.blue > maxBlue {
				possible = false
			}
		}
		if possible {
			accu1 += game.number
		}
	}

	fmt.Println("Part 1 solution: ", accu1)

	accu2 := 0
	for _, game := range games {
		minRed := 0
		minGreen := 0
		minBlue := 0
		for _, draw := range game.draws {
			minRed = max(minRed, draw.red)
			minGreen = max(minGreen, draw.green)
			minBlue = max(minBlue, draw.blue)
		}
		power := minRed * minGreen * minBlue
		accu2 += power
	}

	fmt.Println("Part 2 solution:", accu2)
}

func parseInput(sc *bufio.Scanner, gameNoRegex *regexp.Regexp, colorRegex *regexp.Regexp, games []game) []game {
	for sc.Scan() {
		line := sc.Text()
		gameNoMatch := gameNoRegex.FindStringSubmatch(line)
		gameNo, err := strconv.ParseInt(gameNoMatch[gameNoRegex.SubexpIndex("game")], 10, 0)
		if err != nil {
			log.Fatalln("Failed to parse game number ", err)
		}

		draws := make([]draw, 0)
		for _, input := range strings.Split(line, ";") {
			newDraw := draw{}
			for _, color := range strings.Split(input, ",") {
				colorMatch := colorRegex.FindStringSubmatch(color)
				count, err := strconv.ParseInt(colorMatch[colorRegex.SubexpIndex("count")], 10, 0)
				if err != nil {
					log.Fatalln("Failed to parse color count from draw", err)
				}
				switch colorMatch[colorRegex.SubexpIndex("color")] {
				case "blue":
					newDraw.blue = int(count)
				case "red":
					newDraw.red = int(count)
				case "green":
					newDraw.green = int(count)
				default:
					panic("Unknown color code")
				}
			}

			draws = append(draws, newDraw)
		}

		games = append(games, game{draws: draws, number: int(gameNo)})
	}
	return games
}
