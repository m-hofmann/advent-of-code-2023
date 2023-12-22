package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

const (
	OFF = iota
	ON
	FLIPFLOP
	CONJUNCTION
	BROADCASTER
	OUTPUT
)

const (
	LOW  = 1 << 3
	HIGH = 1 << 4
)

type module interface {
	evaluate(string, int) int
}

type flipFlop struct {
	state int
}

func NewFlipFlop() flipFlop {
	return flipFlop{OFF}
}

func (f *flipFlop) evaluate(input string, in int) int {
	if in == HIGH {
		return 0
	} else if in == LOW {
		if f.state == OFF {
			f.state = ON
			return HIGH
		} else if f.state == ON {
			f.state = OFF
			return LOW
		}
	}
	return 0
}

type conjunction struct {
	values map[string]int
}

type event struct {
	source string
	pulse  int
	target string
}

func (e *event) String() string {

	pulse := "???"
	if e.pulse == LOW {
		pulse = "low"
	} else if e.pulse == HIGH {
		pulse = "high"
	}
	return fmt.Sprintf("%s -%s-> %s", e.source, pulse, e.target)
}

func main() {
	fmt.Println("Starting day 20 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	// OOP modeling? screw that, we have maps
	connected := make(map[string][]string)
	moduleType := map[string]int{"output": OUTPUT}
	flipFlopState := make(map[string]int)
	conjunctionState := make(map[string]map[string]int)
	for sc.Scan() {
		line := sc.Text()
		tokens := strings.Split(line, " -> ")
		modType := 0
		modName := ""
		connectedTo := strings.Split(tokens[1], ", ")
		if tokens[0] == "broadcaster" {
			modType = BROADCASTER
			modName = tokens[0]
		} else if tokens[0][0] == '%' {
			modType = FLIPFLOP
			modName = tokens[0][1:]
			flipFlopState[modName] = OFF
		} else if tokens[0][0] == '&' {
			modType = CONJUNCTION
			modName = tokens[0][1:]

		}
		moduleType[modName] = modType
		connected[modName] = connectedTo
	}

	// fix inputs of conjunction modules
	for k, v := range moduleType {
		if v == CONJUNCTION {
			inputs := make([]string, 0)
			remembered := make(map[string]int, len(inputs))
			for from, to := range connected {
				if slices.Contains(to, k) {
					remembered[from] = LOW
				}
			}
			conjunctionState[k] = remembered
		}
	}

	events := make([]event, 0)
	cLow := 0
	cHigh := 0
	for i := 0; i < 1000; i++ {
		fmt.Println()
		fmt.Printf("=== Button press %d ===\n", i+1)
		events = append(events, event{"button", LOW, "broadcaster"})

		// classic event loop

		for len(events) != 0 {
			curr := events[0]
			events = events[1:]
			fmt.Println(curr.String())
			if curr.pulse == LOW {
				cLow++
			} else if curr.pulse == HIGH {
				cHigh++
			}

			switch moduleType[curr.target] {
			case BROADCASTER:
				for _, conn := range connected[curr.target] {
					events = append(events, event{curr.target, curr.pulse, conn})
				}
			case FLIPFLOP:
				if curr.pulse == LOW {
					if flipFlopState[curr.target] == OFF {
						flipFlopState[curr.target] = ON
						for _, conn := range connected[curr.target] {
							events = append(events, event{curr.target, HIGH, conn})
						}
					} else if flipFlopState[curr.target] == ON {
						flipFlopState[curr.target] = OFF
						for _, conn := range connected[curr.target] {
							events = append(events, event{curr.target, LOW, conn})
						}
					}
				}
			case CONJUNCTION:
				// update remembered input state
				conjunctionState[curr.target][curr.source] = curr.pulse
				// decide if we need to send out pulse
				state := 0
				for _, last := range conjunctionState[curr.target] {
					state |= last
				}
				pulse := HIGH
				if state == HIGH {
					pulse = LOW
				}
				for _, conn := range connected[curr.target] {
					events = append(events, event{curr.target, pulse, conn})
				}
			case OUTPUT:

			default:
			}
		}
	}

	fmt.Println("Part 1 solution", cHigh*cLow)
}
