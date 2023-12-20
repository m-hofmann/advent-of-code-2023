package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	variable   byte
	comparator byte
	value      int
	target     string
}
type workflow struct {
	rules []rule
}

type part struct {
	values map[byte]int
}

func main() {
	fmt.Println("Starting day 19 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	workflows := make(map[string]workflow, 0)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			// switch to part parsing
			break
		}
		tokens := strings.Split(line, "{")
		rules := make([]rule, 0)
		for _, c := range strings.Split(tokens[1][:len(tokens[1])-1], ",") {
			ruleParts := strings.Split(c, ":")
			if len(ruleParts) > 1 {
				// rule with evaluation
				val, err := strconv.ParseInt(ruleParts[0][2:], 10, 0)
				if err != nil {
					log.Fatalln("Could not parse int", err)
				}
				rules = append(rules, rule{variable: ruleParts[0][0], comparator: ruleParts[0][1], value: int(val), target: ruleParts[1]})
			} else {
				// end rule, just route
				rules = append(rules, rule{target: c})
			}
		}
		workflows[tokens[0]] = workflow{rules}

	}

	parts := make([]part, 0)
	for sc.Scan() {
		line := sc.Text()
		line = strings.Trim(line, "{}")
		part := part{values: make(map[byte]int)}
		for _, categoryValue := range strings.Split(line, ",") {
			category := categoryValue[0]
			val, err := strconv.ParseInt(categoryValue[2:len(categoryValue)], 10, 0)
			if err != nil {
				log.Fatalln("Could not parse category value", err)
			}
			part.values[category] = int(val)
		}
		parts = append(parts, part)
	}

	allAcceptedValues := 0
	for _, part := range parts {
		state := "in"
		fmt.Println("=== Part", part, "===")

		for state != "A" && state != "R" {
			fmt.Println("Now in state", state)
			wf, ok := workflows[state]
			if !ok {
				log.Fatalln("Could not find workflow for state", state)
			}

			for _, rule := range wf.rules {
				switch rule.comparator {
				case '<':
					partVal := part.values[rule.variable]
					if partVal < rule.value {
						state = rule.target
						goto out
					}
				case '>':
					partVal := part.values[rule.variable]
					if partVal > rule.value {
						state = rule.target
						goto out
					}
				default:
					state = rule.target
					goto out
				}
			}
		out:
		}

		if state == "A" {
			allAcceptedValues += partValue(part)
		}
	}

	fmt.Printf("Part 1 solution %d\n", allAcceptedValues)
}

func partValue(p part) int {
	result := 0
	for _, v := range p.values {
		result += v
	}
	return result

}
