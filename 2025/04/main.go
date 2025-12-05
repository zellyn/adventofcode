package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	count := 0

	for pos, r := range m {
		if r != '@' {
			continue
		}

		neighbors := 0
		for _, nn := range pos.Neighbors8() {
			if m[nn] == '@' {
				neighbors += 1
			}
		}
		if neighbors < 4 {
			count++
		}
	}

	return count, nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	count := 0

	done := false
	for !done {
		done = true
		var poses []geom.Vec2
		for pos, r := range m {
			if r != '@' {
				continue
			}

			neighbors := 0
			for _, nn := range pos.Neighbors8() {
				if m[nn] == '@' {
					neighbors += 1
				}
			}
			if neighbors < 4 {
				poses = append(poses, pos)
				done = false
				count += 1
			}
		}

		for _, pos := range poses {
			delete(m, pos)
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
