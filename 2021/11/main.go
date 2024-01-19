package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func iterate(m charmap.M) int {
	flashes := 0
	flashed := make(map[geom.Vec2]bool, len(m))

	for pos := range m {
		m[pos]++
	}

	for done := false; !done; {
		done = true
		for pos := range m {
			if m[pos] > '9' && !flashed[pos] {
				flashed[pos] = true
				flashes++
				done = false
				for _, nPos := range pos.Neighbors8() {
					_, ok := m[nPos]
					if ok {
						m[nPos]++
					}
				}
			}
		}
	}

	for pos := range flashed {
		m[pos] = '0'
	}

	return flashes
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += iterate(m)
	}
	return flashes, nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	for step := 1; ; step++ {
		if flashes := iterate(m); flashes == 100 {
			return step, nil
		}
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
