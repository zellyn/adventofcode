package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]geom.Vec2, error) {
	ints, err := util.ParseLinesOfInts(inputs, ",")
	if err != nil {
		return nil, err
	}
	return util.MapE(ints, func(ints []int) (geom.Vec2, error) {
		if len(ints) != 2 {
			return geom.Vec2{}, fmt.Errorf("weird input: not two ints: %v", ints)
		}
		return geom.V2(ints[0], ints[1]), nil
	})

}

type posStep struct {
	pos  geom.Vec2
	step int
}

func findPathLength(size geom.Vec2, coords []geom.Vec2) (int, error) {
	endPos := size.NW()
	m := charmap.New(size.X, size.Y, '.')

	for _, pos := range coords {
		m[pos] = '#'
	}

	todo := make([]posStep, 1)
	seen := make(map[geom.Vec2]bool)
	for len(todo) > 0 {
		ps := todo[0]
		todo = todo[1:]
		if ps.pos == endPos {
			return ps.step, nil
		}
		if seen[ps.pos] {
			continue
		}
		seen[ps.pos] = true

		for _, nn := range ps.pos.Neighbors4() {
			if m[nn] == '.' && !seen[nn] {
				todo = append(todo, posStep{pos: nn, step: ps.step + 1})
			}
		}
	}

	return 0, fmt.Errorf("no path")

}

func part1(inputs []string, steps int, size geom.Vec2) (int, error) {
	coords, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	return findPathLength(size, coords[:steps])
}

func part2(inputs []string, size geom.Vec2) (geom.Vec2, error) {
	coords, err := parse(inputs)
	if err != nil {
		return geom.Vec2{}, nil
	}

	pos, _ := sort.Find(len(coords), func(i int) int {
		_, err := findPathLength(size, coords[:i+1])
		if err == nil {
			return 1
		}
		return -1
	})

	return coords[pos], nil
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
