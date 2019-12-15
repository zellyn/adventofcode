package geom

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Vec2 is a two-element vector.
type Vec2 struct {
	X int
	Y int
}

// Vec3 is a three-element vector.
type Vec3 struct {
	X int
	Y int
	Z int
}

// Abs returns the same vector, but with negative coordinates replaced by their positive values.
func (v Vec3) Abs() Vec3 {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
	if v.Z < 0 {
		v.Z = -v.Z
	}
	return v
}

// Sum returns the x+y+z.
func (v Vec3) Sum() int {
	return v.X + v.Y + v.Z
}

// AbsSum returns |x| + |y| + |z|.
func (v Vec3) AbsSum() int {
	return v.Abs().Sum()
}

// Add adds two vectors.
func (v Vec3) Add(w Vec3) Vec3 {
	return Vec3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub subtracts a vector from this one, returning the result.
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Neg negates a vector.][]
func (v Vec3) Neg() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Sgn replaces each element of a vector with -1, 0, 1, depending on its sign.
func (v Vec3) Sgn() Vec3 {
	var result Vec3

	if v.X < 0 {
		result.X = -1
	} else if v.X > 0 {
		result.X = 1
	}
	if v.Y < 0 {
		result.Y = -1
	} else if v.Y > 0 {
		result.Y = 1
	}
	if v.Z < 0 {
		result.Z = -1
	} else if v.Z > 0 {
		result.Z = 1
	}

	return result
}

var vec3regex = regexp.MustCompile(`<x=(-?[0-9]+), *y=(-?[0-9]+), *z=(-?[0-9]+)>`)

// ParseVec3 parses a string vec3 in format "<x=17,y=42,z=-1>".
func ParseVec3(s string) (Vec3, error) {
	s = strings.TrimSpace(s)
	parts := vec3regex.FindStringSubmatch(s)
	if parts == nil {
		return Vec3{}, fmt.Errorf("ParseVec3: weird input: %q", s)
	}
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		return Vec3{}, fmt.Errorf("cannot parse x coordinate %q (in vector %q)", parts[1], s)
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return Vec3{}, fmt.Errorf("cannot parse y coordinate %q (in vector %q)", parts[2], s)
	}
	z, err := strconv.Atoi(parts[3])
	if err != nil {
		return Vec3{}, fmt.Errorf("cannot parse z coordinate %q (in vector %q)", parts[3], s)
	}
	return Vec3{x, y, z}, nil
}

// ParseVec3Lines parses Vec3s, one per line.
func ParseVec3Lines(s string) ([]Vec3, error) {
	var result []Vec3
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		v, err := ParseVec3(line)
		if err != nil {
			return nil, err
		}
		result = append(result, v)
	}

	return result, nil
}

// Hash3 takes a shitty hash of a slice of Vec3s.
func Hash3(vecs []Vec3) uint {
	res := uint(1)
	for _, v := range vecs {
		res = res*31 + uint(v.X)
		res = res*31 + uint(v.Y)
		res = res*31 + uint(v.Z)
	}
	return res
}

// Abs returns the same vector, but with negative coordinates replaced by their positive values.
func (v Vec2) Abs() Vec2 {
	if v.X < 0 {
		v.X = -v.X
	}
	if v.Y < 0 {
		v.Y = -v.Y
	}
	return v
}

// Sum returns the x+y+z.
func (v Vec2) Sum() int {
	return v.X + v.Y
}

// AbsSum returns |x| + |y| + |z|.
func (v Vec2) AbsSum() int {
	return v.Abs().Sum()
}

// Add adds two vectors.
func (v Vec2) Add(w Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}

// Sub subtracts a vector from this one, returning the result.
func (v Vec2) Sub(w Vec2) Vec2 {
	return Vec2{v.X - w.X, v.Y - w.Y}
}

// Neg negates a vector.][]
func (v Vec2) Neg() Vec2 {
	return Vec2{-v.X, -v.Y}
}

// Sgn replaces each element of a vector with -1, 0, 1, depending on its sign.
func (v Vec2) Sgn() Vec2 {
	var result Vec2

	if v.X < 0 {
		result.X = -1
	} else if v.X > 0 {
		result.X = 1
	}
	if v.Y < 0 {
		result.Y = -1
	} else if v.Y > 0 {
		result.Y = 1
	}
	return result
}
