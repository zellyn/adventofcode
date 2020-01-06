package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/graph"
)

func parseInput(input []string) (map[string]map[string]int, error) {
	result := map[string]map[string]int{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 11 {
			return nil, fmt.Errorf("stange input line: %q", line)
		}
		from := parts[0]
		to := parts[10]
		to = strings.Trim(to, ".")
		points, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing input line: %q: %v", line, err)
		}
		if parts[2] == "lose" {
			points = -points
		}
		if result[from] == nil {
			result[from] = map[string]int{}
		}
		result[from][to] = points
	}
	return result, nil
}

func best(input []string, includeYou bool) (int, error) {
	g, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	var people []string
	for c := range g {
		people = append(people, c)
	}
	if includeYou {
		people = append(people, "you")
	}

	perms := graph.PermutationsString(people)
	highest := 0

	for _, perm := range perms {
		score := 0

		for i := 0; i < len(perm); i++ {
			from := perm[i]
			to := perm[(i+1)%len(perm)]
			score += g[from][to]
			score += g[to][from]
		}

		if score > highest {
			highest = score
		}
	}
	return highest, nil
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
