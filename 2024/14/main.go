package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parseOne(input string) (geom.PosVel2, error) {
	// p=0,4 v=3,-3
	var px, py, vx, vy int
	n, err := fmt.Sscanf(input, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
	if n != 4 || err != nil {
		return geom.PosVel2{}, err
	}
	return geom.PV2(px, py, vx, vy), err
}

func part1(inputs []string, size geom.Vec2, iters int) (int, error) {
	defs, err := util.MapE(inputs, parseOne)
	if err != nil {
		return 0, err
	}

	var counts [4]int

	for _, def := range defs {
		inc := size.Add(def.Vel).EachMod(size)
		finalPos := def.Pos.Add(inc.Mul(iters)).EachMod(size)

		if finalPos.X == size.X/2 || finalPos.Y == size.Y/2 {
			continue
		}

		index := 0
		if finalPos.X > size.X/2 {
			index++
		}
		if finalPos.Y > size.Y/2 {
			index += 2
		}
		counts[index]++
	}

	return counts[0] * counts[1] * counts[2] * counts[3], nil
}

func part2(inputs []string, size geom.Vec2) (int, error) {
	defs, err := util.MapE(inputs, parseOne)
	if err != nil {
		return 0, err
	}

	maxCount := 0
	maxIter := 0

	for i := range 6644 {
		m := charmap.Empty()
		for i, def := range defs {
			def.Pos = def.Pos.Add(def.Vel).Add(size).EachMod(size)
			defs[i] = def
			m[def.Pos] = '#'
		}

		// fmt.Println(i)
		// fmt.Printf("%s\n\n\n\n\n", m.AsString(' '))

		count := 0
		for pos := range m {
			if m[pos.N()] == '#' {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			maxIter = i + 1

			fmt.Println(i + 1)
			fmt.Printf("%s\n\n\n\n\n\n", m.AsString(' '))
		}
	}

	fmt.Printf("maxIter: %d\n", maxIter)

	return maxIter, nil
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
