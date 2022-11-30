package main

import (
	"fmt"
	"os"
)

func runs(ii []int) []int {
	var result []int
	for i := 0; i < len(ii)-2; i++ {
		result = append(result, ii[i]+ii[i+1]+ii[i+2])
	}
	return result
}

func part1(inputs []int) (int, error) {
	increased := 0
	last := inputs[0]
	for _, i := range inputs[1:] {
		if i > last {
			increased++
		}
		last = i
	}
	return increased, nil
}

func part2(inputs []int) (int, error) {
	return part1(runs(inputs))
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
