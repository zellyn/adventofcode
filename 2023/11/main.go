package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var p1 = geom.Vec2{X: 1, Y: 5}
var p2 = geom.Vec2{X: 4, Y: 9}

// var p1 = geom.Vec2{X: 0, Y: 9}
//var p2 = geom.Vec2{X: 4, Y: 9}

func distance(pos1, pos2 geom.Vec2, seenX, seenY map[int]bool, expand int) int {
	if pos1.X > pos2.X {
		pos1.X, pos2.X = pos2.X, pos1.X
	}

	if pos1.Y > pos2.Y {
		pos1.Y, pos2.Y = pos2.Y, pos1.Y
	}

	d := pos2.X - pos1.X + pos2.Y - pos1.Y
	for x := pos1.X; x < pos2.X; x++ {
		if !seenX[x] {
			d += expand - 1
		}
	}
	for y := pos1.Y; y < pos2.Y; y++ {
		if !seenY[y] {
			d += expand - 1
		}
	}

	return d
}

func parse(inputs []string) ([]geom.Vec2, map[int]bool, map[int]bool) {
	m := charmap.Parse(inputs)
	seenX := make(map[int]bool)
	seenY := make(map[int]bool)
	var galaxies []geom.Vec2

	for pos, r := range m {
		if r == '#' {
			seenX[pos.X] = true
			seenY[pos.Y] = true
			galaxies = append(galaxies, pos)
		}
	}
	return galaxies, seenX, seenY
}

func sumDistances(galaxies []geom.Vec2, seenX, seenY map[int]bool, expand int) int {
	d := 0
	for i, pos1 := range galaxies {
		for _, pos2 := range galaxies[i+1:] {
			d += distance(pos1, pos2, seenX, seenY, expand)
		}
	}
	return d
}

func part1(inputs []string) (int, error) {
	galaxies, seenX, seenY := parse(inputs)

	return sumDistances(galaxies, seenX, seenY, 2), nil
}

func part2(inputs []string, expand int) (int, error) {
	galaxies, seenX, seenY := parse(inputs)

	return sumDistances(galaxies, seenX, seenY, expand), nil
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
