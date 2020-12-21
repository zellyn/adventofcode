package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

func rect(m charmap.M, x, y int) {
	for xx := 0; xx < x; xx++ {
		for yy := 0; yy < y; yy++ {
			m[geom.Vec2{X: xx, Y: yy}] = '#'
		}
	}
}

func rotateRow(m charmap.M, y, a int) {
	c := m.Copy()
	_, max := charmap.MinMax(m)
	for x := 0; x <= max.X; x++ {
		newX := (x + a) % (max.X + 1)
		m[geom.Vec2{X: newX, Y: y}] = c[geom.Vec2{X: x, Y: y}]
	}
}

func rotateCol(m charmap.M, x, a int) {
	c := m.Copy()
	_, max := charmap.MinMax(m)
	for y := 0; y <= max.Y; y++ {
		newY := (y + a) % (max.Y + 1)
		m[geom.Vec2{X: x, Y: newY}] = c[geom.Vec2{X: x, Y: y}]
	}
}

func matchParse(s, prefix, separator string, num1, num2 *int) bool {
	if !strings.HasPrefix(s, prefix) {
		return false
	}
	ints, err := util.ParseInts(s[len(prefix):], separator)
	if err != nil {
		return false
	}
	*num1 = ints[0]
	*num2 = ints[1]
	return true
}

func runCommands(m charmap.M, commands []string) error {
	var g, h int
	for i, line := range commands {
		switch {
		// This is not beautiful 😂
		case matchParse(line, "rect ", "x", &g, &h):
			rect(m, g, h)
		case matchParse(line, "rotate column x=", " by ", &g, &h):
			rotateCol(m, g, h)
		case matchParse(line, "rotate row y=", " by ", &g, &h):
			rotateRow(m, g, h)
		default:
			return fmt.Errorf("weird line %d: %q", i+1, line)
		}
	}
	return nil
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
