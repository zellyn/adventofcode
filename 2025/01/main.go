package main

import (
	"fmt"
	"os"
	"strconv"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]int, error) {
	var res []int

	for line, input := range inputs {
		dir := input[0]
		i, err := strconv.Atoi(input[1:])
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
		switch dir {
		case 'R':
			i = -i
		case 'L':
			i = i
		default:
			return nil, fmt.Errorf("weird direction on line %d: %c", line, dir)
		}

		res = append(res, i)
	}

	return res, nil
}

func part1(inputs []string) (int, error) {
	nums, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	pointer := 50
	count := 0
	for _, num := range nums {
		pointer += num
		pointer = ((pointer % 100) + 100) % 100
		if pointer == 0 {
			count++
		}
	}

	return count, nil
}

func part2(inputs []string) (int, error) {
	nums, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	pointer := 50
	count := 0
	for _, num := range nums {
		for num != 0 {
			if num > 0 {
				pointer = (pointer + 1) % 100
				num--
			} else {
				pointer = (pointer + 99) % 100
				num++
			}
			if pointer == 0 {
				count++
			}
		}
	}

	return count, nil
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
