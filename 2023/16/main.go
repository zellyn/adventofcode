package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type pair [2]geom.Vec2

func trace(m charmap.M, pos, dir geom.Vec2) int {
	queue := []pair{
		{pos, dir},
	}

	seen := make(map[geom.Vec2]bool)
	seenSplit := make(map[geom.Vec2]bool)

	for len(queue) > 0 {
		beam := queue[len(queue)-1]
		pos, dir := beam[0], beam[1]
		queue = queue[:len(queue)-1]

		pos = pos.Add(dir)
		contents, ok := m[pos]
		if !ok {
			// out of bounds
			continue
		}

		seen[pos] = true

		switch contents {
		case '.':
			queue = append(queue, pair{pos, dir})
		case '/':
			switch dir {
			case geom.N:
				queue = append(queue, pair{pos, geom.E})
			case geom.S:
				queue = append(queue, pair{pos, geom.W})
			case geom.E:
				queue = append(queue, pair{pos, geom.N})
			case geom.W:
				queue = append(queue, pair{pos, geom.S})
			}
		case '\\':
			switch dir {
			case geom.N:
				queue = append(queue, pair{pos, geom.W})
			case geom.S:
				queue = append(queue, pair{pos, geom.E})
			case geom.E:
				queue = append(queue, pair{pos, geom.S})
			case geom.W:
				queue = append(queue, pair{pos, geom.N})
			}

		case '-':
			switch dir {
			case geom.E, geom.W:
				queue = append(queue, pair{pos, dir})
			default:
				if !seenSplit[pos] {
					seenSplit[pos] = true
					queue = append(queue, pair{pos, geom.E}, pair{pos, geom.W})
				}
			}
		case '|':
			switch dir {
			case geom.N, geom.S:
				queue = append(queue, pair{pos, dir})
			default:
				if !seenSplit[pos] {
					seenSplit[pos] = true
					queue = append(queue, pair{pos, geom.N}, pair{pos, geom.S})
				}
			}
		}
	}

	return len(seen)
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	return trace(m, geom.Vec2{X: -1, Y: 0}, geom.E), nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	min, max := m.MinMax()

	var starts []pair

	for y := min.Y; y <= max.Y; y++ {
		starts = append(starts,
			pair{geom.Vec2{X: min.X - 1, Y: y}, geom.E},
			pair{geom.Vec2{X: max.X + 1, Y: y}, geom.W},
		)
	}

	for x := min.X; x <= max.X; x++ {
		starts = append(starts,
			pair{geom.Vec2{X: x, Y: min.Y - 1}, geom.S},
			pair{geom.Vec2{X: x, Y: max.Y + 1}, geom.N},
		)
	}

	best := 0
	for _, start := range starts {
		score := trace(m, start[0], start[1])
		if score > best {
			best = score
		}
	}

	return best, nil
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
