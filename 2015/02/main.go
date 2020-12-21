package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/math"
	"github.com/zellyn/adventofcode/util"
)

func run() error {
	lines, err := util.ReadLines("input")
	if err != nil {
		return err
	}

	paper := 0
	ribbon := 0

	for _, line := range lines {
		ints, err := util.ParseInts(line, "x")
		if err != nil {
			return err
		}
		a, b, c := math.Sort3(ints[0], ints[1], ints[2])
		paper += 3*a*b + 2*b*c + 2*a*c
		ribbon += 2*(a+b) + a*b*c
	}

	fmt.Println("Paper:", paper)
	fmt.Println("Ribbon:", ribbon)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
