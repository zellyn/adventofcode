package main

import (
	"fmt"
	"os"
)

func printf(format string, args ...any) {
	// fmt.Printf(format, args...)
}

var digitValues = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'-': -1,
	'=': -2,
}

var valueDigits = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	4: "-",
	3: "=",
}

func fromSnafu(s string) int {
	result := 0
	for _, r := range s {
		result *= 5
		result += digitValues[r]
	}
	return result
}

func toSnafu(i int) string {
	printf("toSnafu(%d):\n", i)
	if i == 0 {
		return "0"
	}
	result := ""

	for i > 0 {
		printf("i=%d\n", i)
		d := i % 5
		result = valueDigits[d] + result
		if d > 2 {
			i += 5
		}
		i /= 5
	}
	return result
}

func part1(inputs []string) (string, error) {
	sum := 0
	for _, input := range inputs {
		sum += fromSnafu(input)
	}
	printf("sum=%d\n", sum)
	return toSnafu(sum), nil
}

func part2(inputs []string) (int, error) {
	return 42, nil
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
