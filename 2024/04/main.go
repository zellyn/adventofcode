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
	count := 0
	m := charmap.Parse(inputs)
	for pos, ch := range m {
		if ch != 'X' {
			continue
		}
		for _, dir := range geom.Dirs8 {
			if m[pos.Add(dir)] == 'M' &&
				m[pos.Add(dir.Mul(2))] == 'A' &&
				m[pos.Add(dir.Mul(3))] == 'S' {

				count++
			}
		}
	}
	return count, nil
}

func part2(inputs []string) (int, error) {
	count := 0
	m := charmap.Parse(inputs)
	for pos, ch := range m {
		if ch != 'A' {
			continue
		}
		c1, c2 := m[pos.NE()], m[pos.SW()]
		if (c1 != 'M' && c1 != 'S') || (c2 != 'M' && c2 != 'S') || (c1 == c2) {
			continue
		}
		c1, c2 = m[pos.SE()], m[pos.NW()]
		if (c1 != 'M' && c1 != 'S') || (c2 != 'M' && c2 != 'S') || (c1 == c2) {
			continue
		}
		count++
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
