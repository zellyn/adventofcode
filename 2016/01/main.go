package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/geom"
)

func distance(input string) (int, int, error) {
	parts := strings.Split(input, ", ")
	direction := 0
	pos := geom.Vec2{}
	twice := -1
	seen := map[geom.Vec2]bool{}
	for _, part := range parts {
		distance, err := strconv.Atoi(part[1:])
		if err != nil {
			return 0, 0, err
		}
		switch part[0] {
		case 'L':
			direction = (direction + 3) % 4
		case 'R':
			direction = (direction + 1) % 4
		default:
			panic(fmt.Sprintf("weird input: %q", part))
		}
		for i := 0; i < distance; i++ {
			pos = pos.Add(geom.Dirs4[direction])
			if seen[pos] && twice == -1 {
				twice = pos.AbsSum()
			} else {
				seen[pos] = true
			}
		}
	}
	return pos.AbsSum(), twice, nil
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
