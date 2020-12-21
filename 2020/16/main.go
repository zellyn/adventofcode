package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type rule []clause

type clause struct {
	min int
	max int
}

func (c clause) valid(i int) bool {
	return i >= c.min && i <= c.max
}

func (r rule) valid(i int) bool {
	for _, c := range r {
		if c.valid(i) {
			return true
		}
	}
	return false
}

func (r rule) allValid(is []int) bool {
	for _, i := range is {
		if !r.valid(i) {
			return false
		}
	}
	return true
}

func parseRules(input string) (string, rule, error) {
	parts := strings.Split(input, ": ")
	if len(parts) != 2 {
		return "", nil, fmt.Errorf("weird rule: %q", input)
	}
	var rs rule
	for _, s := range strings.Split(parts[1], " or ") {
		ii, err := util.ParseInts(s, "-")
		if err != nil {
			return "", nil, err
		}
		if len(ii) != 2 {
			return "", nil, fmt.Errorf("weird rule clause %q in rule %q", s, input)
		}
		rs = append(rs, clause{min: ii[0], max: ii[1]})
	}
	return parts[0], rs, nil
}

func parseAllRules(inputs []string) (map[string]rule, error) {
	result := make(map[string]rule, len(inputs))
	for _, input := range inputs {
		name, r, err := parseRules(input)
		if err != nil {
			return nil, err
		}
		result[name] = r
	}
	return result, nil
}

func parse(inputs []string) (map[string]rule, []int, [][]int, error) {
	parts := util.LinesByParagraph(inputs)
	if len(parts) != 3 {
		return nil, nil, nil, fmt.Errorf("want 3 paragraphs; got %d", len(parts))
	}
	rules, err := parseAllRules(parts[0])
	if err != nil {
		return nil, nil, nil, err
	}
	if parts[1][0] != "your ticket:" {
		return nil, nil, nil, fmt.Errorf("want para2 line1 = %q; got %q", "your ticket:", parts[1][0])
	}
	if parts[2][0] != "nearby tickets:" {
		return nil, nil, nil, fmt.Errorf("want para3 line1 = %q; got %q", "your ticket:", parts[2][0])
	}
	yours, err := util.ParseInts(parts[1][1], ",")
	if err != nil {
		return nil, nil, nil, err
	}
	nearby, err := util.ParseLinesOfInts(parts[2][1:], ",")
	if err != nil {
		return nil, nil, nil, err
	}
	return rules, yours, nearby, nil
}

func completelyInvalid(ticket []int, rules map[string]rule) []int {
	var result []int
	for _, i := range ticket {
		valid := false
		for _, r := range rules {
			if r.valid(i) {
				valid = true
				break
			}
		}
		if !valid {
			result = append(result, i)
		}
	}
	return result
}

func part1(inputs []string) (int, error) {
	rules, _, nearby, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	sum := 0

	for _, ticket := range nearby {
		for _, i := range completelyInvalid(ticket, rules) {
			sum += i
		}
	}
	return sum, nil
}

func validFor(col []int, rules map[string]rule) map[string]bool {
	result := make(map[string]bool)
	for name, r := range rules {
		if r.allValid(col) {
			result[name] = true
		}
	}
	return result
}

func firstVal(m map[string]bool) string {
	for k := range m {
		return k
	}
	return ""
}

func sift(vv []map[string]bool, unknown map[string]bool) (done bool, progress bool) {
	done = true
	progress = false

	for thisCol, names := range vv {
		if len(names) == 1 {
			name := firstVal(names)
			if unknown[name] {
				progress = true
				delete(unknown, name)

				for eachCol, eachName := range vv {
					if eachCol == thisCol {
						continue
					}
					delete(eachName, name)
				}
			}
		}
	}

	for _, v := range vv {
		if len(v) > 1 {
			done = false
		}
	}

	return done, progress
}

func part2(inputs []string, prefix string) (int, error) {
	rules, mine, nearbyCandidates, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	nearby := make([][]int, 0, len(nearbyCandidates))

	for _, ticket := range nearbyCandidates {
		if len(completelyInvalid(ticket, rules)) == 0 {
			nearby = append(nearby, ticket)
		}
	}

	cols := util.Transpose(nearby)

	var vv []map[string]bool
	unknown := make(map[string]bool)
	for name := range rules {
		unknown[name] = true
	}

	for _, c := range cols {
		vv = append(vv, validFor(c, rules))
	}

	for {
		done, progress := sift(vv, unknown)
		if done {
			break
		}
		if !progress {
			return 0, fmt.Errorf("Cannot progress simply")
		}
	}

	prod := 1

	for i, names := range vv {
		if strings.HasPrefix(firstVal(names), prefix) {
			prod *= mine[i]
		}
	}

	return prod, nil
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
