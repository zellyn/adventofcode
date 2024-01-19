package main

import (
	"fmt"
	"os"
	"slices"
	"sort"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"golang.org/x/exp/maps"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	sum := 0

OUTER:
	for pos, c := range m {
		for _, nPos := range pos.Neighbors4() {
			n, ok := m[nPos]
			if !ok {
				continue
			}
			if n <= c {
				continue OUTER
			}
		}
		sum += int(c-'0') + 1
	}

	return sum, nil
}

func getSink(pos geom.Vec2, m charmap.M, sinks map[geom.Vec2]geom.Vec2) geom.Vec2 {
	known, ok := sinks[pos]
	if ok {
		return known
	}

	c := m[pos]

	for _, nPos := range pos.Neighbors4() {
		n, ok := m[nPos]
		if !ok || n >= c {
			continue
		}
		sink := getSink(nPos, m, sinks)
		sinks[pos] = sink
		return sink
	}

	sinks[pos] = pos
	return pos
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	sinks := make(map[geom.Vec2]geom.Vec2)

	for pos, c := range m {
		if c == '9' {
			continue
		}
		getSink(pos, m, sinks)
	}

	basins := make(map[geom.Vec2]int)
	for _, sink := range sinks {
		basins[sink]++
	}
	basinSizes := maps.Values(basins)
	sort.Ints(basinSizes)
	slices.Reverse(basinSizes)

	return basinSizes[0] * basinSizes[1] * basinSizes[2], nil
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
