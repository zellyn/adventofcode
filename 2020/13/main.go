package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(inputs []string) (int, []int, error) {
	i, err := strconv.Atoi(inputs[0])
	if err != nil {
		return 0, nil, err
	}

	parts := strings.Split(inputs[1], ",")
	pieces := make([]int, 0, len(parts))
	for _, part := range parts {
		if part == "x" {
			pieces = append(pieces, -1)
		} else {
			j, err := strconv.Atoi(part)
			if err != nil {
				return 0, nil, err
			}
			pieces = append(pieces, j)
		}
	}

	return i, pieces, nil
}

func part1(inputs []string) (int, error) {
	i, buses, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	minWait := i
	minPeriod := 0

	for _, j := range buses {
		if j == -1 {
			continue
		}
		until := (j - i%j) % j
		if until < minWait {
			minWait = until
			minPeriod = j
		}
	}
	return minWait * minPeriod, nil
}

func part2(inputs []string) (int, error) {
	_, buses, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	add := 0
	mul := 1

	for i, bus := range buses {
		j := i + 1
		if bus == -1 {
			continue
		}
		wantMod := bus - j%bus
		for ; add%bus != wantMod; add += mul {
		}
		mul *= bus
	}

	return add + 1, nil
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
