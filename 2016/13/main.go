package main

import (
	"fmt"
	"math/bits"
	"os"

	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func isWall(pos geom.Vec2, favorite int) bool {
	if pos.X < 0 || pos.Y < 0 {
		return true
	}
	num := pos.X*pos.X + 3*pos.X + 2*pos.X*pos.Y + pos.Y + pos.Y*pos.Y + favorite
	return (bits.OnesCount(uint(num)) % 2) == 1
}

type pair struct {
	pos   geom.Vec2
	steps int
}

func part1(favorite int, target geom.Vec2) (int, error) {
	seen := make(map[geom.Vec2]bool)
	todo := []pair{{pos: geom.V2(1, 1), steps: 0}}

	for len(todo) > 0 {
		here := todo[0]
		todo = todo[1:]
		pos, steps := here.pos, here.steps
		if pos == target {
			return steps, nil
		}
		if seen[pos] {
			continue
		}
		seen[pos] = true

		for _, nn := range pos.Neighbors4() {
			if seen[nn] || isWall(nn, favorite) {
				continue
			}
			todo = append(todo, pair{pos: nn, steps: steps + 1})
		}

	}
	return 42, nil
}

func part2(favorite int, steps int) (int, error) {
	seen := make(map[geom.Vec2]bool)
	todo := []pair{{pos: geom.V2(1, 1), steps: 0}}

	for len(todo) > 0 {
		here := todo[0]
		todo = todo[1:]
		pos, steps := here.pos, here.steps
		if seen[pos] {
			continue
		}
		seen[pos] = true
		if steps >= 50 {
			continue
		}

		for _, nn := range pos.Neighbors4() {
			if seen[nn] || isWall(nn, favorite) {
				continue
			}
			todo = append(todo, pair{pos: nn, steps: steps + 1})
		}

	}
	return len(seen), nil
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
