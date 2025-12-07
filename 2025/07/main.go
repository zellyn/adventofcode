package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func count(pos geom.Vec2, m charmap.M, cache map[geom.Vec2]int) int {
	if res, ok := cache[pos]; ok {
		return res
	}
	c := m[pos]
	switch c {
	case '.', 'S':
		res := count(pos.S(), m, cache)
		cache[pos] = res
		return res
	case 0:
		cache[pos] = 1
		return 1
	case '^':
		res := count(pos.W(), m, cache) + count(pos.E(), m, cache)
		cache[pos] = res
		return res
	default:
		panic(fmt.Sprintf("weird character: '%c'", c))
	}
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	start, found := m.Find('S')
	if !found {
		return 0, fmt.Errorf("Cannot find 'S'")
	}

	splits := 0
	row := start.Y
	cols := map[int]bool{
		start.X: true,
	}
	row++
	for m.Has(geom.V2(0, row)) {
		newCols := make(map[int]bool)
		for col := range cols {
			c := m[geom.V2(col, row)]
			switch c {
			case '.':
				newCols[col] = true
			case '^':
				splits++
				newCols[col+1] = true
				newCols[col-1] = true
			default:
				return 0, fmt.Errorf("weird char: '%c'", c)
			}
		}

		cols = newCols
		row++
	}

	return splits, nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	cache := make(map[geom.Vec2]int)
	start, found := m.Find('S')
	if !found {
		return 0, fmt.Errorf("Cannot find 'S'")
	}
	return count(start, m, cache), nil
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
