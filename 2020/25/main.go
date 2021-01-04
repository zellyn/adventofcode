package main

import (
	"fmt"
	"os"
)

const mod = 20201227

func loops(target int) int {
	value := 1
	for result := 0; ; result++ {
		if value == target {
			return result
		}
		value = (value * 7) % mod
	}
}

func pow(sn int, exponent int) int {
	value := 1
	for i := 0; i < exponent; i++ {
		value = (value * sn) % mod
	}
	return value
}

func part1(card int, door int) int {
	return pow(card, loops(door))
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
