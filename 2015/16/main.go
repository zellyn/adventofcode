package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

var readout = util.TrimmedLines(`
	children: 3
	cats: 7
	samoyeds: 2
	pomeranians: 3
	akitas: 0
	vizslas: 0
	goldfish: 5
	trees: 3
	cars: 2
	perfumes: 1`)

func parseReadout(readout []string) (map[string]int, error) {
	m := map[string]int{}
	for _, part := range readout {
		lr := strings.Split(part, ": ")
		if len(lr) != 2 {
			return nil, fmt.Errorf("weird input: %q", part)
		}
		ii, err := strconv.Atoi(lr[1])
		if err != nil {
			return nil, fmt.Errorf("weird input: %q: %v", part, err)
		}
		m[lr[0]] = ii
	}
	return m, nil
}

func parseInput(input []string) (map[int]map[string]int, error) {
	result := map[int]map[string]int{}
	for i, line := range input {
		m := map[string]int{}
		halves := strings.SplitN(line, ": ", 2)
		parts := strings.Split(halves[1], ", ")
		for _, part := range parts {
			lr := strings.Split(part, ": ")
			if len(lr) != 2 {
				return nil, fmt.Errorf("weird input: %q", line)
			}
			ii, err := strconv.Atoi(lr[1])
			if err != nil {
				return nil, fmt.Errorf("weird input: %q: %v", line, err)
			}
			m[lr[0]] = ii
		}
		result[i+1] = m
	}

	return result, nil
}

func which() (int, error) {
	input, err := util.ReadLines("input")
	if err != nil {
		return 0, err
	}
	aunts, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	need, err := parseReadout(readout)
	if err != nil {
		return 0, err
	}

OUTER:
	for key, val := range aunts {
		for thing, count := range val {
			if need[thing] != count {
				continue OUTER
			}
		}
		return key, nil
	}
	return 0, nil
}

func which2() (int, error) {
	less := map[string]bool{
		"pomeranians": true,
		"goldfish":    true,
	}
	more := map[string]bool{
		"cats":  true,
		"trees": true,
	}
	input, err := util.ReadLines("input")
	if err != nil {
		return 0, err
	}
	aunts, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	need, err := parseReadout(readout)
	if err != nil {
		return 0, err
	}

OUTER:
	for key, val := range aunts {
		for thing, count := range val {
			if less[thing] {
				if count >= need[thing] {
					continue OUTER
				}
			} else if more[thing] {
				if count <= need[thing] {
					continue OUTER
				}
			} else if need[thing] != count {
				continue OUTER
			}
		}
		return key, nil
	}
	return 0, nil
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
