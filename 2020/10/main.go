package main

import (
	"fmt"
	"os"
	"sort"
)

func prod(inputs []int) int {
	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)

	gaps := make(map[int]int)
	last := 0
	for _, i := range inputs {
		gaps[i-last]++
		last = i
	}
	return gaps[1] * gaps[3]
}

func ways(inputs []int) int {
	inputs = append(inputs, 0)
	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)

	from := 0
	last := 0
	result := 1
	for i, input := range inputs {
		if input-last == 3 {
			result *= combos(inputs[from:i])
			from = i
		}
		last = input
	}

	return result
}

func combos(inputs []int) int {
	switch len(inputs) {
	case 1, 2:
		return 1
	case 3:
		/*
			1 2 3
			1 3
		*/
		return 2
	case 4:
		/*
			1 2 3 4
			1 2 4
			1 3 4
			1 4
		*/
		return 4
	case 5:
		/*
			1 2 3 4 5
			1 2 3 5
			1 2 4 5
			1 3 4 5
			1 2 5
			1 3 5
			1 4 5
		*/
		return 7
	default:
		panic("foo")
	}
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
