package charmap

import (
	"fmt"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/ioutil"
)

// M is a map of geom.Vec2 to rune.
type M map[geom.Vec2]rune

// MinMax returns a geom.Vec2 for minimum coordinates, and one for maximum.
func (m M) MinMax() (geom.Vec2, geom.Vec2) {
	return MinMax(m)
}

// AsString stringifies a charmap.
// It uses the `unknown` param for positions within the min/max range
// of X and Y, but not in the map. Rows are terminated by newlines.
func (m M) AsString(unknown rune) string {
	return String(m, unknown)
}

// Count returns a count of the number of cells in m that hold the matching rune.
func (m M) Count(which rune) int {
	count := 0
	for _, ch := range m {
		if ch == which {
			count++
		}
	}
	return count
}

// MinMax returns a geom.Vec2 for minimum coordinates, and one for maximum.
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
	fmt.Print(String(drawing, unknown))
}

// String takes a map of `geom.Vec2` to `rune`, and stringifies it out.
// It uses the `unknown` param for positions within the min/max range
// of X and Y, but not in the map. Rows are terminated by newlines.
func String(drawing map[geom.Vec2]rune, unknown rune) string {
	result := ""
	min, max := MinMax(drawing)
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			c, ok := drawing[geom.Vec2{X: x, Y: y}]
			if !ok {
				c = unknown
			}
			result += string(c)
		}
		result += "\n"
	}
	return result
}

// Parse parses a map from a list of strings.
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

// New creates a new charmap, filled with the given fill rune.
func New(width, height int, fill rune) M {
	m := map[geom.Vec2]rune{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			m[geom.Vec2{X: x, Y: y}] = fill
		}
	}
	return m
}

// Copy creates a copy of a charmap.
func (m M) Copy() map[geom.Vec2]rune {
	c := make(map[geom.Vec2]rune, len(m))

	for k, v := range m {
		c[k] = v
	}

	return c
}

// Equal tests two charmaps for equality.
func (m M) Equal(n M) bool {
	if len(m) != len(n) {
		return false
	}

	for k, v := range m {
		v2, ok := n[k]
		if !ok {
			return false
		}
		if v != v2 {
			return false
		}
	}

	return true
}

// Read reads a two-dimensional map of characters from a file.
func Read(filename string) (map[geom.Vec2]rune, error) {
	lines, err := ioutil.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	return Parse(lines), nil
}

// MustRead reads a two-dimensional map of characters from a file, or dies.
func MustRead(filename string) map[geom.Vec2]rune {
	m, err := Read(filename)
	if err != nil {
		panic(err)
	}
	return m
}
