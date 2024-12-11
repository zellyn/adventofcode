package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parseAndMapConnections(inputs []string) (charmap.M, map[geom.Vec2]map[geom.Vec2]bool) {
	m := charmap.Parse(inputs)
	connected := make(map[geom.Vec2]map[geom.Vec2]bool)
	for _, pos := range m.FindAll('9') {
		connected[pos] = map[geom.Vec2]bool{pos: true}
	}

	for i := 0; i < 10; i++ {
		for pos, r := range m {
			for _, nn := range pos.Neighbors4() {
				if m[nn] != r+1 {
					continue
				}
				for hilltop := range connected[nn] {
					if connected[pos] == nil {
						connected[pos] = make(map[geom.Vec2]bool)
					}
					connected[pos][hilltop] = true
				}
			}
		}
	}
	return m, connected
}

func countPaths(pos geom.Vec2, m charmap.M, connected map[geom.Vec2]map[geom.Vec2]bool) int {
	r := m[pos]
	if r == '9' {
		return 1
	}
	count := 0
	for _, nn := range pos.Neighbors4() {
		if m[nn] != r+1 {
			continue
		}
		if len(connected[nn]) == 0 {
			continue
		}

		count += countPaths(nn, m, connected)
	}
	return count
}

func part1(inputs []string) (int, error) {
	m, connected := parseAndMapConnections(inputs)

	sum := 0

	for _, pos := range m.FindAll('0') {
		sum += len(connected[pos])
	}

	return sum, nil
}

func part2(inputs []string) (int, error) {
	m, connected := parseAndMapConnections(inputs)
	sum := 0

	for _, pos := range m.FindAll('0') {
		sum += countPaths(pos, m, connected)
	}

	return sum, nil
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
