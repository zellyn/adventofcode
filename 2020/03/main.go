package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

func treeCount(m map[geom.Vec2]rune, step geom.Vec2) int {
	var pos geom.Vec2
	count := 0
	_, max := charmap.MinMax(m)
	for ; pos.Y <= max.Y; pos = pos.Add(step) {
		pos.X = pos.X % (max.X + 1)
		if m[pos] == '#' {
			count++
		}
	}
	return count
}

func multiTreeCount(m map[geom.Vec2]rune, steps []geom.Vec2) int {
	prod := 1
	for _, step := range steps {
		prod *= treeCount(m, step)
	}
	return prod
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
