package main

import (
	"fmt"
	"os"
)

func part1(inputs []string) (int, error) {
	return 42, nil
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
