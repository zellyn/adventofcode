package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([][2]geom.Vec2, error) {
	ints := util.ParseRegexInts(inputs, false)
	res := make([][2]geom.Vec2, 0, len(ints))

	for i, ii := range ints {
		if len(ii) != 4 {
			return nil, fmt.Errorf("want 4 ints, got %d on line %q", len(ii), inputs[i])
		}
		res = append(res, [2]geom.Vec2{geom.V2(ii[0], ii[1]), geom.V2(ii[2], ii[3])})
	}
	return res, nil
}

func drawLine(m map[geom.Vec2]int, start, end geom.Vec2) {
	inc := end.Sub(start).Sgn()

	for pos := start; ; pos = pos.Add(inc) {
		m[pos]++
		if pos == end {
			return
		}
	}
}

func drawLines(m map[geom.Vec2]int, endpoints [][2]geom.Vec2) {
	for _, points := range endpoints {
		drawLine(m, points[0], points[1])
	}
}

func part1(inputs []string) (int, error) {
	endpoints, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	endpoints = util.Filter(endpoints, func(points [2]geom.Vec2) bool {
		return points[0].X == points[1].X || points[0].Y == points[1].Y
	})

	m := make(map[geom.Vec2]int)
	drawLines(m, endpoints)
	count := 0
	for _, val := range m {
		if val > 1 {
			count++
		}
	}
	return count, nil
}

func part2(inputs []string) (int, error) {
	endpoints, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	m := make(map[geom.Vec2]int)
	drawLines(m, endpoints)
	count := 0
	for _, val := range m {
		if val > 1 {
			count++
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
