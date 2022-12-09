package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/geom"
)

var dirs = map[byte]geom.Vec2{
	'U': {X: 0, Y: -1},
	'D': {X: 0, Y: 1},
	'L': {X: -1, Y: 0},
	'R': {X: 1, Y: 0},
}

func follow(head, tail geom.Vec2) geom.Vec2 {
	if head == tail || head.Adjacent8(tail) {
		return tail
	}
	return tail.Add(head.Sub(tail).Sgn())
}

func move(head, tail, dir geom.Vec2) (geom.Vec2, geom.Vec2) {
	head = head.Add(dir)
	tail = follow(head, tail)

	return head, tail
}

func part1(inputs []string) (int, error) {
	var head geom.Vec2
	var tail geom.Vec2
	seen := map[geom.Vec2]bool{
		{X: 0, Y: 0}: true,
	}
	for _, input := range inputs {
		dir := dirs[input[0]]
		steps, err := strconv.Atoi(input[2:])
		if err != nil {
			return 0, err
		}
		for i := 0; i < steps; i++ {
			head, tail = move(head, tail, dir)
			seen[tail] = true
		}
	}
	return len(seen), nil
}

func part2(inputs []string) (int, error) {
	var knots [10]geom.Vec2
	seen := map[geom.Vec2]bool{
		{X: 0, Y: 0}: true,
	}
	for _, input := range inputs {
		dir := dirs[input[0]]
		steps, err := strconv.Atoi(input[2:])
		if err != nil {
			return 0, err
		}
		for i := 0; i < steps; i++ {
			knots[0] = knots[0].Add(dir)
			for j := 1; j < 10; j++ {
				knots[j] = follow(knots[j-1], knots[j])
			}
			seen[knots[9]] = true
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
