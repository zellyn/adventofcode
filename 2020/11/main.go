package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

func part1(input []string) int {
	m := charmap.M(charmap.Parse(input))
	_ = m

	for {
		n := m.Copy()

	OUTER:
		for k, v := range m {
			switch v {
			case '.':
				continue
			case 'L':
				for _, pos := range geom.Neighbors8(k) {
					if m[pos] == '#' {
						continue OUTER
					}
				}
				n[k] = '#'
			case '#':
				count := 0
				for _, pos := range geom.Neighbors8(k) {
					if m[pos] == '#' {
						count++
					}
				}
				if count >= 4 {
					n[k] = 'L'
				}
			default:
				panic(k)
			}
		}
		if m.Equal(n) {
			break
		}
		m = n
	}

	count := 0
	for _, v := range m {
		if v == '#' {
			count++
		}
	}

	return count
}

func neighborcount(m charmap.M, pos geom.Vec2) int {
	count := 0
OUTER:
	for _, dir := range geom.Dirs8 {
		for i := 1; ; i++ {
			switch m[pos.Add(dir.Mul(i))] {
			case '.':
				continue
			case 0:
				continue OUTER
			case 'L':
				continue OUTER
			case '#':
				count++
				continue OUTER
			}
		}
	}
	return count
}

func part2(input []string) int {
	m := charmap.M(charmap.Parse(input))
	_ = m

	for {
		n := m.Copy()

		for k, v := range m {
			switch v {
			case '.':
				continue
			case 'L':
				if neighborcount(m, k) == 0 {
					n[k] = '#'
				}
			case '#':
				if neighborcount(m, k) >= 5 {
					n[k] = 'L'
				}
			default:
				panic(k)
			}
		}
		if m.Equal(n) {
			break
		}
		m = n
	}

	count := 0
	for _, v := range m {
		if v == '#' {
			count++
		}
	}

	return count
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
