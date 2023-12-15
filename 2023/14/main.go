package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func roll(m charmap.M, dirName string) {
	dir := geom.Compass4[dirName]
	min, max := m.MinMax()

	if dirName == "N" || dirName == "W" {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				pos := geom.Vec2{X: x, Y: y}
				if m[pos] == 'O' {
					next := pos.Add(dir)
					for next.Y >= min.Y && next.X >= min.X && m[next] == '.' {
						m[pos] = '.'
						m[next] = 'O'
						pos = next
						next = next.Add(dir)
					}
				}
			}
		}
	} else {
		for y := max.Y; y >= min.Y; y-- {
			for x := max.X; x >= min.X; x-- {
				pos := geom.Vec2{X: x, Y: y}
				if m[pos] == 'O' {
					next := pos.Add(dir)
					for next.Y <= max.Y && next.X <= max.X && m[next] == '.' {
						m[pos] = '.'
						m[next] = 'O'
						pos = next
						next = next.Add(dir)
					}
				}
			}
		}
	}
}

func cycle(m charmap.M) {
	roll(m, "N")
	roll(m, "W")
	roll(m, "S")
	roll(m, "E")
}

func rollNorth(m charmap.M) {
	roll(m, "N")
}

func score(m charmap.M) int {
	_, max := m.MinMax()
	// charmap.Draw(m, '/')
	sum := 0

	for pos, char := range m {
		if char == 'O' {
			sum += max.Y - pos.Y + 1
		}
	}
	return sum
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	rollNorth(m)
	// charmap.Draw(m, '/')
	return score(m), nil
}

func part2(inputs []string) (int, error) {
	target := 1_000_000_000
	m := charmap.Parse(inputs)
	count := 0
	for i := 0; i < 100; i++ {
		cycle(m)
		count++
	}
	mark1 := count
	s := m.AsString('/')

	for {
		cycle(m)
		count++
		if m.AsString('/') == s {
			break
		}
	}
	mark2 := count

	leftOver := (target - mark2) % (mark2 - mark1)

	for i := 0; i < leftOver; i++ {
		cycle(m)
	}
	return score(m), nil
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
