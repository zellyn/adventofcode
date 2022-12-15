package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/lists"
	"github.com/zellyn/adventofcode/util"
)

func drawLine(m charmap.M, from, to geom.Vec2) {
	inc := to.Sub(from).Sgn()
	for at := from; at != to; at = at.Add(inc) {
		m[at] = '#'
	}
	m[to] = '#'
}

func draw(inputs []string) charmap.M {
	m := make(charmap.M)
	for _, input := range inputs {
		coords := lists.Map(strings.Split(input, " -> "), func(s string) geom.Vec2 {
			ints := util.MustParseInts(s, ",")
			return geom.Vec2{X: ints[0], Y: ints[1]}
		})
		for i, to := range coords[1:] {
			from := coords[i]
			drawLine(m, from, to)
		}
	}
	return m
}

var DOWN = geom.Vec2{X: 0, Y: 1}
var DOWNLEFT = geom.Vec2{X: -1, Y: 1}
var DOWNRIGHT = geom.Vec2{X: 1, Y: 1}
var DIRS = [3]geom.Vec2{DOWN, DOWNLEFT, DOWNRIGHT}

func restingPos(m charmap.M, pos geom.Vec2, floorY int, useFloor bool) (geom.Vec2, bool) {
OUTER:
	for pos.Y < floorY-1 {
		for _, dir := range DIRS {
			if _, occupied := m[pos.Add(dir)]; !occupied {
				pos = pos.Add(dir)
				continue OUTER
			}
		}
		return pos, true

	}
	return pos, useFloor
}

func part1(inputs []string) (int, error) {
	m := draw(inputs)
	start := geom.Vec2{X: 500, Y: 0}
	m[start] = '+'
	_, max := m.MinMax()
	maxY := max.Y + 2

	for count := 0; ; count++ {
		// fmt.Printf("%s\n\n", m.AsString('.'))
		newPos, cameToRest := restingPos(m, start, maxY, false)
		if !cameToRest {
			return count, nil
		}
		m[newPos] = 'o'
	}
	// return 0, nil
}

func part2(inputs []string) (int, error) {
	m := draw(inputs)
	start := geom.Vec2{X: 500, Y: 0}
	m[start] = '+'
	_, max := m.MinMax()
	maxY := max.Y + 2

	for count := 0; ; count++ {
		// fmt.Printf("%s\n\n", m.AsString('.'))
		newPos, _ := restingPos(m, start, maxY, true)
		if newPos == start {
			return count + 1, nil
		}
		m[newPos] = 'o'
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
