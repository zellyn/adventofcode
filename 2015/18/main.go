package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

type vec2 = geom.Vec2

const light = '#'
const space = '.'

func next(m map[vec2]rune) map[vec2]rune {
	mm := make(map[vec2]rune, len(m))

	for pos, r := range m {
		neighbors := 0
		for _, nn := range geom.Neighbors8(pos) {
			if m[nn] == light {
				neighbors++
			}
		}
		if (r == space && neighbors == 3) || (r == light && (neighbors == 2 || neighbors == 3)) {
			mm[pos] = light
		} else {
			mm[pos] = space
		}
	}
	return mm
}

func countOn(m map[vec2]rune) int {
	sum := 0
	for _, val := range m {
		if val == light {
			sum++
		}
	}
	return sum
}

func lightsAfter(input []string, steps int) int {
	m := charmap.Parse(input)
	// charmap.Draw(m, '?')
	for i := 0; i < steps; i++ {
		m = next(m)
		// charmap.Draw(m, '?')
	}
	return countOn(m)
}

func stuckLightsAfter(input []string, steps int) int {
	m := charmap.Parse(input)
	min, max := charmap.MinMax(m)
	fix := func(m map[vec2]rune) {
		m[vec2{X: min.X, Y: min.Y}] = light
		m[vec2{X: max.X, Y: min.Y}] = light
		m[vec2{X: max.X, Y: max.Y}] = light
		m[vec2{X: min.X, Y: max.Y}] = light
	}
	fix(m)
	// charmap.Draw(m, '?')
	for i := 0; i < steps; i++ {
		m = next(m)
		fix(m)
		// charmap.Draw(m, '?')
	}
	return countOn(m)
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
