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

// start and end inclusive
type interval struct {
	start int
	end   int
}

type state struct {
	workflow string
	ranges   map[byte]interval
}

func (s *state) copy() state {
	newState := state{s.workflow, make(map[byte]interval)}
	for k, v := range s.ranges {
		newState.ranges[k] = v
	}
	return newState
}

func (s *state) assertInvariants() {
	for _, r := range s.ranges {
		if r.end < r.start {
			log.Fatalln("Invalid range", s)
		}
	}
}

func (s *state) getRangeCoverage(name byte) int64 {
	// +1 to compensate for the fact that all start, end are inclusive
	return int64(s.ranges[name].end - s.ranges[name].start + 1)
}

func (s *state) getCombinations() int64 {
	return s.getRangeCoverage('x') * s.getRangeCoverage('m') *
		s.getRangeCoverage('a') * s.getRangeCoverage('s')
}

func (s *state) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("@%s:", s.workflow))
	for k, v := range s.ranges {
		sb.WriteString(fmt.Sprintf("%s -> [%4d, %4d], ", k, v.start, v.end))
	}
	return sb.String()
}

func main() {
	fmt.Println("Starting day 19 ... ")

	f, err := os.OpenFile("./data/part1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to read input file!")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	workflows, parts := parseInput(sc)

	allAcceptedValues := getAcceptedPartValues(parts, workflows)
	fmt.Printf("Part 1 solution %d\n", allAcceptedValues)

	finalAcceptedRanges := make([]state, 0)
	finalRejectedRanges := make([]state, 0)
	unfinished := make([]state, 1)
	unfinished[0] = state{
		"in",
		map[byte]interval{
			'x': {1, 4000},
			'm': {1, 4000},
			'a': {1, 4000},
			's': {1, 4000},
		}}

	for len(unfinished) > 0 {
		curr := unfinished[0]
		unfinished = unfinished[1:]
		curr.assertInvariants()
		if curr.workflow == "A" {
			finalAcceptedRanges = append(finalAcceptedRanges, curr)
			continue
		} else if curr.workflow == "R" {
			finalRejectedRanges = append(finalRejectedRanges, curr)
			continue
		}

		wf, ok := workflows[curr.workflow]
		if !ok {
			log.Fatalln("Could not find workflow", curr)
		}
		for _, rule := range wf.rules {
			currRange := curr.ranges[rule.variable]
			switch rule.comparator {
			case '<':
				if currRange.end < rule.value {
					// case 1: currRange \in rule range
					curr.workflow = rule.target
					unfinished = append(unfinished, curr)
					goto out
				} else if currRange.start < rule.value && rule.value < currRange.end {
					// case 2: overlap - a part matches. Interval must be split
					second := curr.copy()
					// create second interval that should have matched rule and moves to new workflow
					oldRange := second.ranges[rule.variable]
					oldRange.end = rule.value - 1
					second.ranges[rule.variable] = oldRange
					second.workflow = rule.target
					unfinished = append(unfinished, second)

					oldRange = curr.ranges[rule.variable]
					oldRange.start = rule.value
					curr.ranges[rule.variable] = oldRange
				} else if rule.value <= currRange.start && rule.value <= currRange.end {
					// case 3: no match -> next rule
				}
			case '>':
				if rule.value < currRange.start && rule.value < currRange.end {
					// case 1: currRange \in rule range
					curr.workflow = rule.target
					unfinished = append(unfinished, curr)
					goto out
				} else if currRange.start < rule.value && rule.value < currRange.end {
					// case 2: overlap - a part matches. Interval must be split
					second := curr.copy()
					// create second interval that should have matched rule and moves to new workflow
					oldRange := second.ranges[rule.variable]
					oldRange.start = rule.value + 1
					second.ranges[rule.variable] = oldRange
					second.workflow = rule.target
					unfinished = append(unfinished, second)

					oldRange = curr.ranges[rule.variable]
					oldRange.end = rule.value
					curr.ranges[rule.variable] = oldRange
				} else if rule.value > currRange.start && rule.value > currRange.end {
					// no match -> next rule
				}
			default:
				curr.workflow = rule.target
				unfinished = append(unfinished, curr)
				goto out
			}
		}
	out:
	}

	totalAccCombinations := int64(0)

	for _, acceptedState := range finalAcceptedRanges {
		localCombinations := acceptedState.getCombinations()
		totalAccCombinations += localCombinations
	}

	fmt.Println("Part 2 solution", totalAccCombinations)
}

func getAcceptedPartValues(parts []part, workflows map[string]workflow) int {
	allAcceptedValues := 0
	for _, part := range parts {
		state := "in"

		for state != "A" && state != "R" {
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
	return allAcceptedValues
}

func parseInput(sc *bufio.Scanner) (map[string]workflow, []part) {
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
	return workflows, parts
}

func partValue(p part) int {
	result := 0
	for _, v := range p.values {
		result += v
	}
	return result

}
