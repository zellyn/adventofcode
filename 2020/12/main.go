package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/geom"
)

func part1(inputs []string) (int, error) {
	pos := geom.Vec2{}
	dir := 90

	for _, input := range inputs {
		cmd := input[0]
		i, err := strconv.Atoi(input[1:])
		if err != nil {
			return 0, err
		}
		var inc geom.Vec2
		switch cmd {
		case 'L':
			dir -= i
			dir = (dir + 360) % 360
		case 'R':
			dir += i
			dir = dir % 360
		case 'F':
			index := dir / 90
			inc = geom.Dirs4[index]
		case 'N':
			inc = geom.Vec2{X: 0, Y: -1}
		case 'S':
			inc = geom.Vec2{X: 0, Y: 1}
		case 'E':
			inc = geom.Vec2{X: 1, Y: 0}
		case 'W':
			inc = geom.Vec2{X: -1, Y: 0}
		default:
			return 0, fmt.Errorf("unknown command: '%c'", cmd)
		}
		pos = pos.Add(inc.Mul(i))
	}
	return pos.AbsSum(), nil
}

func part2(inputs []string) (int, error) {
	pos := geom.Vec2{}
	wp := geom.Vec2{X: 10, Y: -1}

	for _, input := range inputs {
		cmd := input[0]
		i, err := strconv.Atoi(input[1:])
		if err != nil {
			return 0, err
		}
		var inc geom.Vec2
		var wpinc geom.Vec2
		rots := 0
		switch cmd {
		case 'L':
			rots = 4 - i/90
		case 'R':
			rots = i / 90
		case 'F':
			inc = wp
		case 'N':
			wpinc = geom.Vec2{X: 0, Y: -1}
		case 'S':
			wpinc = geom.Vec2{X: 0, Y: 1}
		case 'E':
			wpinc = geom.Vec2{X: 1, Y: 0}
		case 'W':
			wpinc = geom.Vec2{X: -1, Y: 0}
		default:
			return 0, fmt.Errorf("unknown command: '%c'", cmd)
		}
		for j := 0; j < rots; j++ {
			wp = geom.Vec2{X: -wp.Y, Y: wp.X}
		}
		pos = pos.Add(inc.Mul(i))
		wp = wp.Add(wpinc.Mul(i))
	}
	return pos.AbsSum(), nil
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
