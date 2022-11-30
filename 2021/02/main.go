package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

func part1(inputs []string) (int, error) {
	si, err := util.ParseStringsAndInts(inputs, 2, []int{0}, []int{1})
	if err != nil {
		return 0, err
	}
	pos := geom.Vec2{}
	for _, elem := range si {
		s := elem.Strings[0]
		i := elem.Ints[0]
		switch s {
		case "forward":
			pos.X += i
		case "down":
			pos.Y += i
		case "up":
			pos.Y -= i
		default:
			return 0, fmt.Errorf("unknown direction: %s", s)
		}
	}
	return pos.X * pos.Y, nil
}

func part2(inputs []string) (int, error) {
	si, err := util.ParseStringsAndInts(inputs, 2, []int{0}, []int{1})
	if err != nil {
		return 0, err
	}
	pos := geom.Vec2{}
	aim := 0
	for _, elem := range si {
		s := elem.Strings[0]
		i := elem.Ints[0]
		switch s {
		case "forward":
			pos.X += i
			pos.Y += i * aim
		case "down":
			aim += i
		case "up":
			aim -= i
		default:
			return 0, fmt.Errorf("unknown direction: %s", s)
		}
	}
	return pos.X * pos.Y, nil
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
