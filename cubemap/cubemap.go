package cubemap

import (
	"fmt"
	"strings"

	"github.com/zellyn/adventofcode/geom"
)

// M is a map of geom.Vec3 to rune.
type M map[geom.Vec3]rune

// ParseCSVLines parses a slice of comma-separated x,y,z lines into a cubemap.M.
func ParseCSVLines(lines []string, r rune) (M, error) {
	vec3s, err := geom.ParseCSVec3Lines(lines)
	if err != nil {
		return nil, err
	}

	m := make(map[geom.Vec3]rune, len(lines))
	for _, v := range vec3s {
		m[v] = r
	}

	return m, nil
}

// MinMax returns the minimum and maximum X, Y, and Z coordinates.
func (m M) MinMax() (geom.Vec3, geom.Vec3) {
	var min, max geom.Vec3
	first := true
	for pos := range m {
		if first {
			first = false
			min, max = pos, pos
		} else {
			if pos.X < min.X {
				min.X = pos.X
			}
			if pos.Y < min.Y {
				min.Y = pos.Y
			}
			if pos.Z < min.Z {
				min.Z = pos.Z
			}
			if pos.X > max.X {
				max.X = pos.X
			}
			if pos.Y > max.Y {
				max.Y = pos.Y
			}
			if pos.Z > max.Z {
				max.Z = pos.Z
			}
		}
	}
	return min, max
}

// Has returns true if the cubemap has an entry at pos.
func (m M) Has(pos geom.Vec3) bool {
	_, ok := m[pos]
	return ok
}

func (m M) AsString(empty rune) string {
	var result []string
	min, max := m.MinMax()
	var pos geom.Vec3
	for z := min.Z; z <= max.Z; z++ {
		pos.Z = z
		if z != min.Z {
			result = append(result, "")
		}
		result = append(result, fmt.Sprintf("Z=%d:", z))
		for y := min.Y; y <= max.Y; y++ {
			pos.Y = y
			row := ""
			for x := min.X; x <= max.X; x++ {
				pos.X = x
				r, ok := m[pos]
				if !ok {
					r = empty
				}
				row = row + string(r)
			}
			result = append(result, row)
		}
	}
	return strings.Join(result, "\n")
}
