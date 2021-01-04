package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

func neighbors(pos geom.Vec2) []geom.Vec2 {
	return []geom.Vec2{
		{X: pos.X - 1, Y: pos.Y},
		{X: pos.X + 1, Y: pos.Y},
		{X: pos.X, Y: pos.Y - 1},
		{X: pos.X + 1, Y: pos.Y - 1},
		{X: pos.X - 1, Y: pos.Y + 1},
		{X: pos.X, Y: pos.Y + 1},
	}
}

func count(m charmap.M, pos geom.Vec2) int {
	c := 0
	for _, n := range neighbors(pos) {
		if m[n] == 'b' {
			c++
		}
	}
	return c
}

func iter(m charmap.M) charmap.M {
	mm := make(charmap.M)

	todo := make(map[geom.Vec2]bool)

	for p := range m {
		todo[p] = true
		for _, n := range neighbors(p) {
			todo[n] = true
		}
	}

	for p := range todo {
		c := count(m, p)
		if m[p] == 'b' {
			if c == 1 || c == 2 {
				mm[p] = 'b'
			}
		} else {
			if c == 2 {
				mm[p] = 'b'
			}
		}
	}

	return mm
}

func split(input string) []string {
	var result []string
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'e', 'w':
			result = append(result, input[i:i+1])
		default:
			result = append(result, input[i:i+2])
			i++
		}

	}
	return result
}

func flip(m charmap.M, input string) {
	pos := geom.Vec2{}
	for _, dir := range split(input) {
		switch dir {
		case "w":
			pos.X--
		case "e":
			pos.X++
		case "nw":
			pos.Y--
		case "ne":
			pos.Y--
			pos.X++
		case "sw":
			pos.Y++
			pos.X--
		case "se":
			pos.Y++
		default:
			panic(fmt.Sprintf("weird direction: %q", dir))
		}
	}
	if m[pos] == 'b' {
		delete(m, pos)
	} else {
		m[pos] = 'b'
	}
}

func part1(inputs []string) (int, error) {
	m := make(charmap.M)
	for _, input := range inputs {
		flip(m, input)
	}

	return m.Count('b'), nil
}

func part2(inputs []string) (int, error) {
	m := make(charmap.M)
	for _, input := range inputs {
		flip(m, input)
	}

	for i := 0; i < 100; i++ {
		m = iter(m)
	}

	return m.Count('b'), nil
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
