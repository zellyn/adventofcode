package main

import (
	"fmt"
	"os"
)

func parse(input string) int {
	if input == "" {
		return 0
	}
	rest := 2 * parse(input[:len(input)-1])
	switch input[len(input)-1] {
	case 'F', 'L':
		return rest
	case 'B', 'R':
		return rest + 1
	default:
		panic(fmt.Sprintf("weird input: %q", input))
	}
}

func maxParse(inputs []string) int {
	max := 0
	for _, input := range inputs {
		p := parse(input)
		if p > max {
			max = p
		}
	}
	return max
}

func missing(inputs []string) int {
	min := 1024
	max := 0
	seen := make(map[int]bool)
	for _, input := range inputs {
		p := parse(input)
		if p > max {
			max = p
		}
		if p < min {
			min = p
		}
		seen[p] = true
	}

	for i := min; i <= max; i++ {
		if !seen[i] {
			return i
		}
	}
	return -1
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
