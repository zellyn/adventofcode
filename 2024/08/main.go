package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) (charmap.M, map[rune][]geom.Vec2, geom.Rect) {
	m := charmap.Parse(inputs)
	p := make(map[rune][]geom.Vec2)
	for pos, r := range m {
		if r != '.' {
			p[r] = append(p[r], pos)
		}
	}
	pmin, pmax := m.MinMax()
	return m, p, geom.Rect{Min: pmin, Max: pmax}
}

func part1(inputs []string) (int, error) {
	m, p, _ := parse(inputs)
	antinodes := make(map[geom.Vec2]bool)
	for _, poses := range p {
		for i, p1 := range poses {
			for _, p2 := range poses[i+1:] {
				diff := p1.Sub(p2)
				antinodes[p2.Sub(diff)] = true
				antinodes[p1.Add(diff)] = true
			}
		}
	}
	// fmt.Printf("%s\\nn", m.AsString('_'))
	sum := 0
	for pos := range m {
		if antinodes[pos] {
			sum++
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	_, p, bounds := parse(inputs)
	antinodes := make(map[geom.Vec2]bool)
	for _, poses := range p {
		for i, p1 := range poses {
			for _, p2 := range poses[i+1:] {
				diff := p1.Sub(p2)
				for pos := p1; bounds.Contains(pos); pos = pos.Add(diff) {
					antinodes[pos] = true
				}
				for pos := p2; bounds.Contains(pos); pos = pos.Sub(diff) {
					antinodes[pos] = true
				}
			}
		}
	}
	return len(antinodes), nil
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
