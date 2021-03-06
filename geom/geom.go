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

// Vec4 is a three-element vector.
type Vec4 struct {
	W int
	X int
	Y int
	Z int
}

// String does the usual.
func (v Vec2) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}

// String does the usual.
func (v Vec3) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v.X, v.Y, v.Z)
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

// Mul returns the vector multiplied by a scalar.
func (v Vec3) Mul(factor int) Vec3 {
	return Vec3{X: v.X * factor, Y: v.Y * factor, Z: v.Z * factor}
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

// Sum returns the x+y.
func (v Vec2) Sum() int {
	return v.X + v.Y
}

// AbsSum returns |x| + |y|.
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

// Mul returns the vector multiplied by a scalar.
func (v Vec2) Mul(factor int) Vec2 {
	return Vec2{X: v.X * factor, Y: v.Y * factor}
}

// Min2 returns the minimum of two vectors in both X and Y.
func Min2(a, b Vec2) Vec2 {
	if a.X > b.X {
		a.X = b.X
	}
	if a.Y > b.Y {
		a.Y = b.Y
	}
	return a
}

// Max2 returns the minimum of two vectors in both X and Y.
func Max2(a, b Vec2) Vec2 {
	if a.X < b.X {
		a.X = b.X
	}
	if a.Y < b.Y {
		a.Y = b.Y
	}
	return a
}

// Add adds two vectors.
func (v Vec4) Add(w Vec4) Vec4 {
	return Vec4{v.W + w.W, v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Dirs4 are the four cardinal direction length-1 Vec2s.
var Dirs4 = []Vec2{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

// Neighbors4 returns the four orthogonally adjacent positions of a Vec2 position.
func Neighbors4(pos Vec2) []Vec2 {
	return []Vec2{
		{pos.X - 1, pos.Y},
		{pos.X, pos.Y + 1},
		{pos.X + 1, pos.Y},
		{pos.X, pos.Y - 1},
	}
}

// Dirs8 are the eight neighboring Vec2s to the given Vec2.
var Dirs8 = []Vec2{
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

// Neighbors8 returns the four orthogonally and diagonally adjacent positions of a Vec2 position.
func Neighbors8(pos Vec2) []Vec2 {
	return []Vec2{
		{pos.X, pos.Y - 1},
		{pos.X + 1, pos.Y - 1},
		{pos.X + 1, pos.Y},
		{pos.X + 1, pos.Y + 1},
		{pos.X, pos.Y + 1},
		{pos.X - 1, pos.Y + 1},
		{pos.X - 1, pos.Y},
		{pos.X - 1, pos.Y - 1},
	}
}

// Dirs6 are the four cardinal direction length-1 Vec3s.
var Dirs6 = []Vec3{
	{0, -1, 0},
	{1, 0, 0},
	{0, 1, 0},
	{-1, 0, 0},
	{0, 0, -1},
	{0, 0, 1},
}

// Neighbors6 returns the six orthogonally adjacent positions of a Vec3 position.
func Neighbors6(pos Vec3) []Vec3 {
	return []Vec3{
		{pos.X - 1, pos.Y, pos.Z},
		{pos.X, pos.Y + 1, pos.Z},
		{pos.X + 1, pos.Y, pos.Z},
		{pos.X, pos.Y - 1, pos.Z},
		{pos.X, pos.Y, pos.Z - 1},
		{pos.X, pos.Y, pos.Z + 1},
	}
}

// Dirs26 are the 26 neighboring Vec3s to the given Vec3.
var Dirs26 = []Vec3{
	{0, -1, 0},
	{1, -1, 0},
	{1, 0, 0},
	{1, 1, 0},
	{0, 1, 0},
	{-1, 1, 0},
	{-1, 0, 0},
	{-1, -1, 0},

	{0, 0, -1},
	{0, -1, -1},
	{1, -1, -1},
	{1, 0, -1},
	{1, 1, -1},
	{0, 1, -1},
	{-1, 1, -1},
	{-1, 0, -1},
	{-1, -1, -1},

	{0, 0, 1},
	{0, -1, 1},
	{1, -1, 1},
	{1, 0, 1},
	{1, 1, 1},
	{0, 1, 1},
	{-1, 1, 1},
	{-1, 0, 1},
	{-1, -1, 1},
}

// Neighbors26 returns the 26 orthogonally and diagonally adjacent positions of a Vec3 position.
func Neighbors26(pos Vec3) []Vec3 {
	return []Vec3{
		{pos.X, pos.Y - 1, pos.Z},
		{pos.X + 1, pos.Y - 1, pos.Z},
		{pos.X + 1, pos.Y, pos.Z},
		{pos.X + 1, pos.Y + 1, pos.Z},
		{pos.X, pos.Y + 1, pos.Z},
		{pos.X - 1, pos.Y + 1, pos.Z},
		{pos.X - 1, pos.Y, pos.Z},
		{pos.X - 1, pos.Y - 1, pos.Z},

		{pos.X, pos.Y, pos.Z - 1},
		{pos.X, pos.Y - 1, pos.Z - 1},
		{pos.X + 1, pos.Y - 1, pos.Z - 1},
		{pos.X + 1, pos.Y, pos.Z - 1},
		{pos.X + 1, pos.Y + 1, pos.Z - 1},
		{pos.X, pos.Y + 1, pos.Z - 1},
		{pos.X - 1, pos.Y + 1, pos.Z - 1},
		{pos.X - 1, pos.Y, pos.Z - 1},
		{pos.X - 1, pos.Y - 1, pos.Z - 1},

		{pos.X, pos.Y, pos.Z + 1},
		{pos.X, pos.Y - 1, pos.Z + 1},
		{pos.X + 1, pos.Y - 1, pos.Z + 1},
		{pos.X + 1, pos.Y, pos.Z + 1},
		{pos.X + 1, pos.Y + 1, pos.Z + 1},
		{pos.X, pos.Y + 1, pos.Z + 1},
		{pos.X - 1, pos.Y + 1, pos.Z + 1},
		{pos.X - 1, pos.Y, pos.Z + 1},
		{pos.X - 1, pos.Y - 1, pos.Z + 1},
	}
}

// Dirs80 are the 80 neighboring Vec4s to the given Vec4.
var Dirs80 []Vec4

func init() {
	for w := -1; w <= 1; w++ {
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				for z := -1; z <= 1; z++ {
					if w == 0 && x == 0 && y == 0 && z == 0 {
						continue
					}
					Dirs80 = append(Dirs80, Vec4{W: w, X: x, Y: y, Z: z})
				}

			}

		}

	}
}

// Neighbors80 returns the 80 orthogonally and diagonally adjacent positions of a Vec4 position.
func Neighbors80(pos Vec4) []Vec4 {
	result := make([]Vec4, 80)
	for i, v := range Dirs80 {
		result[i] = v.Add(pos)
	}
	return result
}
