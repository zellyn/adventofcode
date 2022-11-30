package main

import (
	"fmt"
	"os"
)

func getCounts(inputs []string) ([]int, error) {
	counts := make([]int, len(inputs[0]))
	for _, input := range inputs {
		for i, c := range input {
			switch c {
			case '0':
				counts[i]--
			case '1':
				counts[i]++
			default:
				return nil, fmt.Errorf("weird input '%c' in %q", c, input)
			}
		}
	}
	return counts, nil
}

func findOne(inputs []string, sign int) (int, error) {
	pos := 0
	for len(inputs) > 1 {
		counts, err := getCounts(inputs)
		if err != nil {
			return 0, err
		}
		next := make([]string, 0, len(inputs))
		for _, input := range inputs {
			c := counts[pos] * sign
			if c == 0 {
				c += sign
			}
			if c > 0 && input[pos] == '1' || c < 0 && input[pos] == '0' {
				next = append(next, input)
			}
		}
		inputs = next
		// fmt.Println(inputs)
		pos++
	}
	result := 0
	for _, c := range inputs[0] {
		result = result*2 + int(c-'0')
	}
	return result, nil
}

func part1(inputs []string) (int, error) {
	counts, err := getCounts(inputs)
	if err != nil {
		return 0, err
	}
	var gamma, epsilon int
	for _, count := range counts {
		gamma *= 2
		epsilon *= 2
		if count == 0 {
			return 0, fmt.Errorf("zero count")
		} else if count > 0 {
			gamma++
		} else {
			epsilon++
		}
	}
	return gamma * epsilon, nil
}

func part2(inputs []string) (int, error) {
	o2, err := findOne(inputs, 1)
	if err != nil {
		return 0, err
	}
	co2, err := findOne(inputs, -1)
	if err != nil {
		return 0, err
	}
	return o2 * co2, nil
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
