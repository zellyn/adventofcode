package charvol

import (
	"fmt"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// V is a map of geom.Vec3 to rune.
type V map[geom.Vec3]rune

// V4 is a map of geom.Vec4 to rune.
type V4 map[geom.Vec4]rune

// MinMax returns a geom.Vec3 for minimum coordinates, and one for maximum.
func (v V) MinMax() (geom.Vec3, geom.Vec3) {
	return MinMax(v)
}

// MinMax4 returns a geom.Vec4 for minimum coordinates, and one for maximum.
func (v V4) MinMax() (geom.Vec4, geom.Vec4) {
	return MinMax4(v)
}

// AsString stringifies a charvol.
// It uses the `unknown` param for positions within the min/max range
// of X/Y/Z, but not in the map. Rows are terminated by newlines.
func (v V) AsString(unknown rune) string {
	return String(v, unknown)
}

// Count returns a count of the number of cells in v that hold the matching rune.
func (v V4) Count(which rune) int {
	count := 0
	for _, ch := range v {
		if ch == which {
			count++
		}
	}
	return count
}

// Count returns a count of the number of cells in v that hold the matching rune.
func (v V) Count(which rune) int {
	count := 0
	for _, ch := range v {
		if ch == which {
			count++
		}
	}
	return count
}

// MinMax returns a geom.Vec3 for minimum coordinates, and one for maximum.
func MinMax(vol map[geom.Vec3]rune) (geom.Vec3, geom.Vec3) {
	minx, miny, minz, maxx, maxy, maxz := 0, 0, 0, 0, 0, 0
	for k := range vol {
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
		if k.Z < minz {
			minz = k.Z
		}
		if k.Z > maxz {
			maxz = k.Z
		}
	}
	return geom.Vec3{X: minx, Y: miny, Z: minz}, geom.Vec3{X: maxx, Y: maxy, Z: maxz}
}

// MinMax4 returns a geom.Vec4 for minimum coordinates, and one for maximum.
func MinMax4(vol map[geom.Vec4]rune) (geom.Vec4, geom.Vec4) {
	var minw, minx, miny, minz, maxw, maxx, maxy, maxz int
	for k := range vol {
		if k.W < minw {
			minw = k.W
		}
		if k.W > maxw {
			maxw = k.W
		}
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
		if k.Z < minz {
			minz = k.Z
		}
		if k.Z > maxz {
			maxz = k.Z
		}
	}
	return geom.Vec4{W: minw, X: minx, Y: miny, Z: minz}, geom.Vec4{W: maxw, X: maxx, Y: maxy, Z: maxz}
}

// Draw takes a map of `geom.Vec3` to `rune`, and prints it out.
// It uses the `unknown` param for positions within the min/max range
// of X/Y/Z, but not in the map.
func Draw(vol map[geom.Vec3]rune, unknown rune) {
	fmt.Print(String(vol, unknown))
}

// String takes a map of `geom.Vec3` to `rune`, and stringifies it out.
// It uses the `unknown` param for positions within the min/max range
// of X/Y/Z, but not in the map. Rows are terminated by newlines.
func String(vol map[geom.Vec3]rune, unknown rune) string {
	result := ""
	min, max := MinMax(vol)
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				c, ok := vol[geom.Vec3{X: x, Y: y, Z: z}]
				if !ok {
					c = unknown
				}
				result += string(c)
			}
			result += "\n"
		}
		result += "\n"
	}
	return result
}

// New creates a new charvol, filled with the given fill rune.
func New(width, height int, depth int, fill rune) V {
	v := map[geom.Vec3]rune{}
	for z := 0; z < depth; z++ {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				v[geom.Vec3{X: x, Y: y, Z: z}] = fill
			}
		}
	}
	return v
}

// Copy creates a copy of a charvol.
func (v V) Copy() map[geom.Vec3]rune {
	cp := make(map[geom.Vec3]rune, len(v))

	for k, c := range v {
		cp[k] = c
	}

	return cp
}

// Equal tests two charvols for equality.
func (v V) Equal(w V) bool {
	if len(v) != len(w) {
		return false
	}

	for k, c := range v {
		v2, ok := w[k]
		if !ok {
			return false
		}
		if c != v2 {
			return false
		}
	}

	return true
}

// FromCharmap promotes a charmap.M to a charvol.V, at the given z-index.
func FromCharmap(m charmap.M, z int) V {
	v := make(map[geom.Vec3]rune, len(m))

	for k, c := range m {
		v[geom.Vec3{X: k.X, Y: k.Y, Z: z}] = c
	}

	return v
}

// FromCharmap4 promotes a charmap.M to a charvol.V4, at the given w- and z-index.
func FromCharmap4(m charmap.M, w int, z int) V4 {
	v := make(map[geom.Vec4]rune, len(m))

	for k, c := range m {
		v[geom.Vec4{W: w, X: k.X, Y: k.Y, Z: z}] = c
	}

	return v
}
