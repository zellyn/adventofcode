package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/geom"
)

var dirs = map[rune]geom.Vec2{
	'U': geom.Vec2{X: 0, Y: -1},
	'D': geom.Vec2{X: 0, Y: 1},
	'L': geom.Vec2{X: -1, Y: 0},
	'R': geom.Vec2{X: 1, Y: 0},
}

var digits2 = map[geom.Vec2]string{
	geom.Vec2{X: 0, Y: -2}:  "1",
	geom.Vec2{X: -1, Y: -1}: "2",
	geom.Vec2{X: 0, Y: -1}:  "3",
	geom.Vec2{X: 1, Y: 0}:   "4",
	geom.Vec2{X: -2, Y: 0}:  "5",
	geom.Vec2{X: -1, Y: 0}:  "6",
	geom.Vec2{X: 0, Y: 0}:   "7",
	geom.Vec2{X: 1, Y: 0}:   "8",
	geom.Vec2{X: 2, Y: 0}:   "9",
	geom.Vec2{X: -1, Y: 1}:  "A",
	geom.Vec2{X: 0, Y: 1}:   "B",
	geom.Vec2{X: 1, Y: 1}:   "C",
	geom.Vec2{X: 0, Y: 2}:   "D",
}

func code(inputs []string) (string, string, error) {
	result1 := ""
	pos1 := geom.Vec2{X: 0, Y: 0}
	result2 := ""
	pos2 := geom.Vec2{X: -2, Y: 0}
	for _, input := range inputs {
		for _, ch := range input {
			delta, ok := dirs[ch]
			if !ok {
				return "", "", fmt.Errorf("Weird direction: %c", ch)
			}
			new1 := pos1.Add(delta)
			new2 := pos2.Add(delta)
			if new1.Abs().X <= 1 && new1.Abs().Y <= 1 {
				pos1 = new1
			}
			if new2.Abs().Sum() <= 2 {
				pos2 = new2
			}
		}
		digit := pos1.Y*3 + pos1.X + 5
		result1 = result1 + strconv.Itoa(digit)
		digit2, ok := digits2[pos2]
		if !ok {
			digit2 = "X"
		}
		result2 += digit2
	}
	_ = pos2
	return result1, result2, nil
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
