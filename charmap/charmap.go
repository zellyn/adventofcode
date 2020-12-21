package charmap

import (
	"fmt"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
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
func Parse(lines []string) M {
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

// SliceAsString follows a horizontal or vertical slice and returns the
// characters as a string.
func (m M) SliceAsString(start, end geom.Vec2, unknown rune) string {
	if start.X != end.X && start.Y != end.Y {
		return ""
	}
	inc := end.Sub(start).Sgn()
	end = end.Add(inc)
	result := ""
	for pos := start; pos != end; pos = pos.Add(inc) {
		c, ok := m[pos]
		if !ok {
			c = unknown
		}
		result += string(c)
	}

	return result
}

// Rotate the map clockwise. It's assumed the top left is (0,0).
func (m M) Clockwise() M {
	_, max := m.MinMax()
	mm := make(map[geom.Vec2]rune, len(m))
	for pos, c := range m {
		mm[geom.Vec2{X: max.Y - pos.Y, Y: pos.X}] = c
	}
	return mm
}

// FlipLR mirrors the map left-to-right. It's assumed the top left is (0,0).
func (m M) FlipLR() M {
	_, max := m.MinMax()
	mm := make(map[geom.Vec2]rune, len(m))
	for pos, c := range m {
		mm[geom.Vec2{X: max.X - pos.X, Y: pos.Y}] = c
	}
	return mm
}

// Paste writes the given map into this map, at the given offset.
func (m M) Paste(mm M, offset geom.Vec2) {
	for pos, c := range mm {
		m[pos.Add(offset)] = c
	}
}

// Subset takes the section of m from `tl` to `br` (inclusive) and creates a new
// map containing just that rectangle, shifting it to (0,0).
func (m M) Subset(tl, br geom.Vec2) M {
	mm := make(map[geom.Vec2]rune, len(m))
	for pos, c := range m {
		if pos.X >= tl.X && pos.X <= br.X && pos.Y >= tl.Y && pos.Y <= br.Y {
			mm[geom.Vec2{X: pos.X - tl.X, Y: pos.Y - tl.Y}] = c
		}
	}
	return mm
}

// Without creates a copy of a charmap, but with every instance of the given
// rune removed.
func (m M) Without(r rune) M {
	mm := make(map[geom.Vec2]rune, len(m))

	for k, v := range m {
		if v != r {
			mm[k] = v
		}
	}

	return mm
}

// AllInstances finds all offsets at which `m` contains `mm`. Both are assumed
// to start at (0,0).
func (m M) AllInstances(mm M) []geom.Vec2 {
	_, max := m.MinMax()
	_, mmax := mm.MinMax()
	var result []geom.Vec2

	for x := 0; x <= max.X-mmax.X; x++ {
	NEXT:
		for y := 0; y <= max.Y-mmax.Y; y++ {
			offset := geom.Vec2{X: x, Y: y}
			for pos, c := range mm {
				if m[pos.Add(offset)] != c {
					continue NEXT
				}
			}
			result = append(result, offset)
		}
	}
	return result
}

// Replacing creates a copy of a charmap, but with every instance of `from`
// replaced with `to`.
func (m M) Replacing(from, to rune) map[geom.Vec2]rune {
	mm := make(map[geom.Vec2]rune, len(m))

	for k, v := range m {
		if v != from {
			mm[k] = v
		} else {
			mm[k] = to
		}
	}

	return mm
}

// Read reads a two-dimensional map of characters from a file.
func Read(filename string) (map[geom.Vec2]rune, error) {
	lines, err := util.ReadLines(filename)
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
