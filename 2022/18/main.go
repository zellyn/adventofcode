package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/cubemap"
	"github.com/zellyn/adventofcode/geom"
)

func within(pos, min, max geom.Vec3) bool {
	return pos.X >= min.X && pos.Y >= min.Y && pos.Z >= min.Z && pos.X <= max.X && pos.Y <= max.Y && pos.Z <= max.Z
}

func floodFill(c cubemap.M, r rune) {
	one := geom.Vec3{X: 1, Y: 1, Z: 1}
	min, max := c.MinMax()
	min = min.Sub(one)
	max = max.Add(one)
	todo := []geom.Vec3{min}

	for len(todo) > 0 {
		pos := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		c[pos] = r
		for _, n := range pos.Neighbors6() {
			if c.Has(n) {
				continue
			}
			if !within(n, min, max) {
				continue
			}
			todo = append(todo, n)
		}
	}
}

func part1(inputs []string) (int, error) {
	c, err := cubemap.ParseCSVLines(inputs, '#')
	if err != nil {
		return 0, err
	}

	area := 0
	for v := range c {
		for _, n := range v.Neighbors6() {
			if c[n] != '#' {
				area++
			}
		}
	}
	return area, nil
}

func part2(inputs []string) (int, error) {
	c, err := cubemap.ParseCSVLines(inputs, '#')
	if err != nil {
		return 0, err
	}
	floodFill(c, '.')

	area := 0
	for v, r := range c {
		if r != '#' {
			continue
		}
		for _, n := range v.Neighbors6() {
			if c[n] == '.' {
				area++
			}
		}
	}
	return area, nil
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
