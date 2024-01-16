package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/2016/assembunny"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func part1(inputs []string) (int, error) {
	s, err := assembunny.Parse(inputs)
	if err != nil {
		return 0, err
	}
	done := false
	for !done {
		done = s.Step()
	}
	return s.GetRegister("a")
}

func part2(inputs []string) (int, error) {
	s, err := assembunny.Parse(inputs)
	if err != nil {
		return 0, err
	}
	err = s.SetRegister("c", 1)
	if err != nil {
		return 0, err
	}
	done := false
	for !done {
		done = s.Step()
	}
	return s.GetRegister("a")
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
