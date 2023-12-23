package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var dirMap = map[rune]geom.Vec2{
	'^': geom.N,
	'v': geom.S,
	'>': geom.E,
	'<': geom.W,
}

func getNeighbors(m charmap.M, pos geom.Vec2) []geom.Vec2 {
	at := m[pos]
	if dir, ok := dirMap[at]; ok {
		return []geom.Vec2{dir}
	}
	res := make([]geom.Vec2, 0, 4)
	for _, dir := range geom.Compass4 {
		nn := pos.Add(dir)
		if m[nn] != '#' {
			res = append(res, nn)
		}
	}
	return res
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	// m.Surround('#')
	_, max := m.MinMax()
	start := geom.Vec2{X: 1, Y: 0}
	end := geom.Vec2{X: max.X - 1, Y: max.Y}
	m[start] = 'S'
	m[end] = 'E'
	charmap.Draw(m, ' ')

	// paths := allPaths(m, start, end, []geom.Vec2{start.N()})
	// return util.Max(util.Map(paths, func(l []geom.Vec2) int { return len(l) })), nil
	return 42, nil
}

func part2(inputs []string) (int, error) {
	return 42, nil
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
