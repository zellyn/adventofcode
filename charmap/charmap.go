package charmap

import (
	"fmt"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/ioutil"
)

func MinMax(drawing map[geom.Vec2]rune) (geom.Vec2, geom.Vec2) {
	minx, miny, maxx, maxy := 0, 0, 0, 0
	for k := range drawing {
		if k.X < minx {
			minx = k.X
		}
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y < miny {
			miny = k.Y
		}
		if k.Y > maxy {
			maxy = k.Y
		}
	}
	return geom.Vec2{X: minx, Y: miny}, geom.Vec2{X: maxx, Y: maxy}
}

// Draw takes a map of `geom.Vec2` to `rune`, and prints it out.
// It uses the `unknown` param for positions within the min/max range
// of X and Y, but not in the map.
func Draw(drawing map[geom.Vec2]rune, unknown rune) {
	min, max := MinMax(drawing)
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			c, ok := drawing[geom.Vec2{X: x, Y: y}]
			if !ok {
				c = unknown
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func Parse(lines []string) map[geom.Vec2]rune {
	m := map[geom.Vec2]rune{}
	for y, line := range lines {
		for x, ch := range line {
			pos := geom.Vec2{X: x, Y: y}
			m[pos] = ch
		}
	}
	return m
}

// Read reads a two-dimensional map of characters from a file.
func Read(filename string) (map[geom.Vec2]rune, error) {
	lines, err := ioutil.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	return Parse(lines), nil
}
