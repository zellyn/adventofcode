package geom

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

var (
	LEFT  = Vec2{X: -1}
	RIGHT = Vec2{X: 1}
	UP    = Vec2{Y: -1}
	DOWN  = Vec2{Y: 1}
	W     = LEFT
	E     = RIGHT
	N     = UP
	S     = DOWN
	NW    = Vec2{X: -1, Y: -1}
	NE    = Vec2{X: 1, Y: -1}
	SW    = Vec2{X: -1, Y: 1}
	SE    = Vec2{X: 1, Y: 1}
)

// Compass4 map the 4 cardinal compass directions to corresponding vectors.
var Compass4 = map[string]Vec2{
	"N": N,
	"S": S,
	"E": E,
	"W": W,
}

// NameToDir maps the compass directions and U,D,L,R to corresponding vectors.
var NameToDir = map[string]Vec2{
	"N":  N,
	"S":  S,
	"E":  E,
	"W":  W,
	"NW": NW,
	"NE": NE,
	"SW": SW,
	"SE": SE,
	"U":  N,
	"D":  S,
	"L":  W,
	"R":  E,
}

// DirToName maps direction vectors to compass directions.
var DirToName = map[Vec2]string{
	N:  "N",
	S:  "S",
	E:  "E",
	W:  "W",
	NW: "NW",
	NE: "NE",
	SW: "SW",
	SE: "SE",
}

// Vec2 is a two-element vector.
type Vec2 struct {
	X int
	Y int
}

// Rect represents a rectangle reaching from (min.X, min.Y) to (max.X, max.Y), inclusive.
type Rect struct {
	Min Vec2
	Max Vec2
}

// Vec3 is a three-element vector.
type Vec3 struct {
	X int
	Y int
	Z int
}

// Vec2f is a float64-based two-element vector.
type Vec2f struct {
	X float64
	Y float64
}

// Vec3f is a float64-based three-element vector.
type Vec3f struct {
	X float64
	Y float64
	Z float64
}

// Vec4 is a three-element vector.
type Vec4 struct {
	W int
	X int
	Y int
	Z int
}

const EPSILON = 1e-6

// Z2 is a shortcut for an empty Vec2.
var Z2 = Vec2{}

// Z3 is a shortcut for an empty Vec3.
var Z3 = Vec3{}

func Close(a, b, within float64) bool {
	diff := fabs(a - b)
	return -within <= diff && diff <= within
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func fabs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// ToF converts a Vec3 to a Vec3f.
func (v Vec3) ToF() Vec3f {
	return Vec3f{X: float64(v.X), Y: float64(v.Y), Z: float64(v.Z)}
}

// ToF converts a Vec2 to a Vec2f.
func (v Vec2) ToF() Vec2f {
	return Vec2f{X: float64(v.X), Y: float64(v.Y)}
}

// XY converts a Vec3f to a Vec2f containing just the X and Y components.
func (v Vec3f) XY() Vec2f {
	return Vec2f{X: v.X, Y: v.Y}
}

// XZ converts a Vec3f to a Vec2f containing just the X and Z components.
func (v Vec3f) XZ() Vec2f {
	return Vec2f{X: v.X, Y: v.Z}
}

// YZ converts a Vec3f to a Vec2f containing just the Y and Z components.
func (v Vec3f) YZ() Vec2f {
	return Vec2f{X: v.Y, Y: v.Z}
}

// WithZ converts a Vec2f to a Vec3f with matching X and Y components,
// and the specified Z component.
func (v Vec2f) WithZ(z float64) Vec3f {
	return Vec3f{X: v.X, Y: v.Y, Z: z}
}

// String does the usual.
func (v Vec2) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}

// String returns a representation of a Rect.
func (r Rect) String() string {
	return fmt.Sprintf("[%d-%d]", r.Min, r.Max)
}

// MakeRect turns two points into a rectangle, ensuring that they are ordered properly.
func MakeRect(pos1, pos2 Vec2) Rect {
	if pos1.X > pos2.X {
		pos1.X, pos2.X = pos2.X, pos1.X
	}
	if pos1.Y > pos2.Y {
		pos1.Y, pos2.Y = pos2.Y, pos1.Y
	}
	return Rect{Min: pos1, Max: pos2}
}

// MakeRectXYs turns two pairs of X,Y coordinates into a rectangle,
// ensuring that they are ordered properly.
func MakeRectXYs(x1, y1, x2, y2 int) Rect {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return Rect{Min: Vec2{X: x1, Y: y1}, Max: Vec2{X: x2, Y: y2}}
}

// Contains returns true if the given position is within the rectangle.
func (r Rect) Contains(pos Vec2) bool {
	return pos.X >= r.Min.X && pos.X <= r.Max.X && pos.Y >= r.Min.Y && pos.Y <= r.Max.Y
}

// Width does the obvious.
func (r Rect) Width() int {
	return r.Max.X - r.Min.X + 1
}

// Height does the obvious.
func (r Rect) Height() int {
	return r.Max.Y - r.Min.Y + 1
}

// BL returns the Bottom Left corner of the Rect
func (r Rect) BL() Vec2 {
	return r.Min.WithY(r.Max.Y)
}

// TR returns the Top Right corner of the Rect
func (r Rect) TR() Vec2 {
	return r.Min.WithX(r.Max.X)
}

// Positions returns a slice holding all the positions included by the
// rect (inclusive of edges).
func (r Rect) Positions() []Vec2 {
	res := make([]Vec2, 0, r.Width()*r.Height())

	for y := r.Min.Y; y <= r.Max.Y; y++ {
		for x := r.Min.X; x <= r.Max.X; x++ {
			res = append(res, Vec2{X: x, Y: y})
		}
	}

	return res
}

// String does the usual.
func (v Vec3) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v.X, v.Y, v.Z)
}

// WithX returns a new Vec2, but replaces the X coordinate with the one given.
func (v Vec2) WithX(x int) Vec2 {
	return Vec2{X: x, Y: v.Y}
}

// WithY returns a new Vec2, but replaces the Y coordinate with the one given.
func (v Vec2) WithY(y int) Vec2 {
	return Vec2{X: v.X, Y: y}
}

// WithZ returns a new Vec3, with the given Z coordinate, and the same
// X and Y coordinates as v.
func (v Vec2) WithZ(z int) Vec3 {
	return Vec3{X: v.X, Y: v.Y, Z: z}
}

// WithX returns a new Vec3, but replaces the X coordinate with the one given.
func (v Vec3) WithX(x int) Vec3 {
	return Vec3{X: x, Y: v.Y, Z: v.Z}
}

// WithY returns a new Vec3, but replaces the Y coordinate with the one given.
func (v Vec3) WithY(y int) Vec3 {
	return Vec3{X: v.X, Y: y, Z: v.Z}
}

// WithZ returns a new Vec3, but replaces the Z coordinate with the one given.
func (v Vec3) WithZ(z int) Vec3 {
	return Vec3{X: v.X, Y: v.Y, Z: z}
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

// Add adds two vectors.
func (v Vec3f) Add(w Vec3f) Vec3f {
	return Vec3f{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub subtracts a vector from this one, returning the result.
func (v Vec3) Sub(w Vec3) Vec3 {
	return Vec3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// Sub subtracts a vector from this one, returning the result.
func (v Vec3f) Sub(w Vec3f) Vec3f {
	return Vec3f{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
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

// IntDiv returns the vector with each component integer-divided by a scalar.
func (v Vec2) IntDiv(factor int) Vec2 {
	return Vec2{X: v.X / factor, Y: v.Y / factor}
}

// IntDiv returns the vector with each component integer-divided by a scalar.
func (v Vec3) IntDiv(factor int) Vec3 {
	return Vec3{X: v.X / factor, Y: v.Y / factor, Z: v.Z / factor}
}

// Mul returns the vector multiplied by a scalar.
func (v Vec3f) Mul(factor float64) Vec3f {
	return Vec3f{X: v.X * factor, Y: v.Y * factor, Z: v.Z * factor}
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

// ParseCSVec3Lines parses Vec3s, one per line.
func ParseCSVec3Lines(inputs []string) ([]Vec3, error) {
	result := make([]Vec3, 0, len(inputs))
	ints, err := util.ParseLinesOfInts(inputs, ",")
	if err != nil {
		return nil, err
	}
	for i, coords := range ints {
		if len(coords) != 3 {
			return nil, fmt.Errorf("weird input on line %d; wanted 3 ints, got %d: %q", i+1, len(coords), inputs[i])
		}
		result = append(result, Vec3{X: coords[0], Y: coords[1], Z: coords[2]})
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

// Taxi returns the taxicab distance between two positions.
func (v Vec2) Taxi(w Vec2) int {
	return abs(v.X-w.X) + abs(v.Y-w.Y)
}

// Transpose returns the same vector, but with X and Y flipped.
func (v Vec2) Transpose() Vec2 {
	return Vec2{X: v.Y, Y: v.X}
}

// Min returns the minimum of X and Y.
func (v Vec2) Min() int {
	if v.X < v.Y {
		return v.X
	}
	return v.Y
}

// Max returns the maximum of X and Y.
func (v Vec2) Max() int {
	if v.X > v.Y {
		return v.X
	}
	return v.Y
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

// Mul returns the vector multiplied by a scalar.
func (v Vec2f) Mul(factor float64) Vec2f {
	return Vec2f{X: v.X * factor, Y: v.Y * factor}
}

// Div returns the vector divided by a scalar, using integer division.
func (v Vec2) Div(factor int) Vec2 {
	return Vec2{X: v.X / factor, Y: v.Y / factor}
}

// EachDiv returns a new vector with each coordinate (integer) divided
// by the corresponding coordinate of the argument.
func (v Vec2) EachDiv(other Vec2) Vec2 {
	return Vec2{X: v.X / other.X, Y: v.Y / other.Y}
}

// EachMod returns a new vector with each coordinate (integer) modded
// by the corresponding coordinate of the argument.
func (v Vec2) EachMod(other Vec2) Vec2 {
	return Vec2{X: v.X % other.X, Y: v.Y % other.Y}
}

// Within returns true if the Vec2 is within the area specified by min and max (inclusive on all sides).
func (v Vec2) Within(min, max Vec2) bool {
	return v.X >= min.X && v.X <= max.X && v.Y >= min.Y && v.Y <= max.Y
}

// Adjacent4 returns true if the two inputs are adjacent in the four cardinal directions.
func (v Vec2) Adjacent4(other Vec2) bool {
	diff := v.Sub(other).Abs()
	return diff.Sum() == 1
}

// Adjacent8 returns true if the two inputs are adjacent in all eight directions.
func (v Vec2) Adjacent8(other Vec2) bool {
	diff := v.Sub(other).Abs()
	if diff.Sum() == 1 {
		return true
	}
	if diff.X == 1 && diff.Y == 1 {
		return true
	}
	return false
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

// Neighbors4 returns the four orthogonally adjacent positions of a Vec2 position.
func (v Vec2) Neighbors4() []Vec2 {
	return Neighbors4(v)
}

// Neighbors8 returns the four orthogonally and diagonally adjacent positions of a Vec2 position.
func (v Vec2) Neighbors8() []Vec2 {
	return Neighbors8(v)
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

// Neighbors6 returns the 6 orthogonal neighbors of this Vec3.
func (v Vec3) Neighbors6() []Vec3 {
	return Neighbors6(v)
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

// N returns the position north of this one.
func (v Vec2) N() Vec2 {
	return v.Add(N)
}

// NE returns the position northeast of this one.
func (v Vec2) NE() Vec2 {
	return v.Add(NE)
}

// NW returns the position northwest of this one.
func (v Vec2) NW() Vec2 {
	return v.Add(NW)
}

// S returns the position south of this one.
func (v Vec2) S() Vec2 {
	return v.Add(S)
}

// SE returns the position southeast of this one.
func (v Vec2) SE() Vec2 {
	return v.Add(SE)
}

// SW returns the position southwest of this one.
func (v Vec2) SW() Vec2 {
	return v.Add(SW)
}

// E returns the position east of this one.
func (v Vec2) E() Vec2 {
	return v.Add(E)
}

// W returns the position west of this one.
func (v Vec2) W() Vec2 {
	return v.Add(W)
}

// V2 succinctly builds a Vec2.
func V2(x, y int) Vec2 {
	return Vec2{X: x, Y: y}
}

// V2f succinctly builds a Vec2f.
func V2f(x, y float64) Vec2f {
	return Vec2f{X: x, Y: y}
}

// V3 succinctly builds a Vec3.
func V3(x, y, z int) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

// V3f succinctly builds a Vec3f.
func V3f(x, y, z float64) Vec3f {
	return Vec3f{X: x, Y: y, Z: z}
}

// XY returns a Vec2 holding just the X and Y coordinates of v.
func (v Vec3) XY() Vec2 {
	return Vec2{X: v.X, Y: v.Y}
}

// XZ returns a Vec2 holding just the X and Z coordinates of v.
func (v Vec3) XZ() Vec2 {
	return Vec2{X: v.X, Y: v.Z}
}

// YZ returns a Vec2 holding just the Y and Z coordinates of v.
func (v Vec3) YZ() Vec2 {
	return Vec2{X: v.Y, Y: v.Z}
}

// Add adds two vectors.
func (v Vec2f) Add(w Vec2f) Vec2f {
	return Vec2f{v.X + w.X, v.Y + w.Y}
}

// Sub subtracts a vector from this one, returning the result.
func (v Vec2f) Sub(w Vec2f) Vec2f {
	return Vec2f{v.X - w.X, v.Y - w.Y}
}

// Mag calculates the magnitude of the vector.
func (v Vec2f) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Mag calculates the magnitude of the vector.
func (v Vec3f) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// MagSq calculates the magnitude of the vector, squared.
func (v Vec2) MagSq() int {
	return v.X*v.X + v.Y*v.Y
}

// MagSq calculates the magnitude of the vector, squared.
func (v Vec3) MagSq() int {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Norm returns a normalized version of the vector, or (0,0) if it's the zero vector already.
func (v Vec2f) Norm() Vec2f {
	if v == (Vec2f{X: 0, Y: 0}) {
		return v
	}
	mag := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vec2f{X: v.X / mag, Y: v.Y / mag}
}

// Cross computes the cross product (in 2D, aka perpendicular dot product) with another vector.
func (v1 Vec2f) Cross(v2 Vec2f) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

// Dot computes the dot product with another vector.
func (v1 Vec2f) Dot(v2 Vec2f) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v Vec2f) String() string {
	return fmt.Sprintf("(%g,%g)", v.X, v.Y)
}

func (v Vec2f) Close(other Vec2f, within float64) bool {
	return Close(v.X, other.X, within) && Close(v.Y, other.Y, within)
}

func (v Vec3f) Close(other Vec3f, within float64) bool {
	return Close(v.X, other.X, within) && Close(v.Y, other.Y, within) && Close(v.Z, other.Z, within)
}

// PosVel2f represents a 2-D position and direction, in float64.
type PosVel2f struct {
	Pos Vec2f
	Vel Vec2f
}

// PosVel3f represents a 3-D position and direction, in float64.
type PosVel3f struct {
	Pos Vec3f
	Vel Vec3f
}

// PosVel2 represents a 2-D position and direction.
type PosVel2 struct {
	Pos Vec2
	Vel Vec2
}

// WithPos returns a new PosVel with Pos replaced.
func (pv PosVel2) WithPos(pos Vec2) PosVel2 {
	return PosVel2{
		Pos: pos,
		Vel: pv.Vel,
	}
}

// WithVel returns a PosVel from the current Vec2, with the given Vec2
// as its velocity.
func (v Vec2) WithVel(vel Vec2) PosVel2 {
	return PosVel2{
		Pos: v,
		Vel: vel,
	}
}

// WithVel returns a new PosVel with Vel replaced.
func (pv PosVel2) WithVel(vel Vec2) PosVel2 {
	return PosVel2{
		Pos: pv.Pos,
		Vel: vel,
	}
}

// Step returns a new PosVel, but with the position incremented by one addition of the velocity.
func (pv PosVel2) Step() PosVel2 {
	return PosVel2{
		Vec2{X: pv.Pos.X + pv.Vel.X, Y: pv.Pos.Y + pv.Vel.Y},
		pv.Vel,
	}
}

// PosVel3 represents a 3-D position and direction.
type PosVel3 struct {
	Pos Vec3
	Vel Vec3
}

// PV2 constructs a PosVel2 from x and y position and vx and vy velocity components.
func PV2(x, y, vx, vy int) PosVel2 {
	return PosVel2{
		Pos: Vec2{X: x, Y: y},
		Vel: Vec2{X: vx, Y: vy},
	}
}

// PV3 constructs a PosVel3 from x, y, and z position and vx, vy, and vz velocity components.
func PV3(x, y, z, vx, vy, vz int) PosVel3 {
	return PosVel3{
		Pos: Vec3{X: x, Y: y, Z: z},
		Vel: Vec3{X: vx, Y: vy, Z: vz},
	}
}

// XY returns a PosVel2 from a PosVel3, using just the X and Y components.
func (pv PosVel3) XY() PosVel2 {
	return PosVel2{
		Pos: Vec2{X: pv.Pos.X, Y: pv.Pos.Y},
		Vel: Vec2{X: pv.Vel.X, Y: pv.Vel.Y},
	}
}

// XZ returns a PosVel2 from a PosVel3, using just the X and Z components.
func (pv PosVel3) XZ() PosVel2 {
	return PosVel2{
		Pos: Vec2{X: pv.Pos.X, Y: pv.Pos.Z},
		Vel: Vec2{X: pv.Vel.X, Y: pv.Vel.Z},
	}
}

// YZ returns a PosVel2 from a PosVel3, using just the Y and Z components.
func (pv PosVel3) YZ() PosVel2 {
	return PosVel2{
		Pos: Vec2{X: pv.Pos.Y, Y: pv.Pos.Z},
		Vel: Vec2{X: pv.Vel.Y, Y: pv.Vel.Z},
	}
}

// ToF converts a PosVel3 to a PosVel3f.
func (pv PosVel3) ToF() PosVel3f {
	return PosVel3f{
		Pos: Vec3f{X: float64(pv.Pos.X), Y: float64(pv.Pos.Y), Z: float64(pv.Pos.Z)},
		Vel: Vec3f{X: float64(pv.Vel.X), Y: float64(pv.Vel.Y), Z: float64(pv.Vel.Z)},
	}
}

// XY returns a PosVel2f from a PosVel3f, using just the X and Y components.
func (pv PosVel3f) XY() PosVel2f {
	return PosVel2f{
		Pos: Vec2f{X: pv.Pos.X, Y: pv.Pos.Y},
		Vel: Vec2f{X: pv.Vel.X, Y: pv.Vel.Y},
	}
}

// XZ returns a PosVel2f from a PosVel3, using just the X and Z components.
func (pv PosVel3f) XZ() PosVel2f {
	return PosVel2f{
		Pos: Vec2f{X: pv.Pos.X, Y: pv.Pos.Z},
		Vel: Vec2f{X: pv.Vel.X, Y: pv.Vel.Z},
	}
}

// YZ returns a PosVel2f from a PosVel3, using just the Y and Z components.
func (pv PosVel3f) YZ() PosVel2f {
	return PosVel2f{
		Pos: Vec2f{X: pv.Pos.Y, Y: pv.Pos.Z},
		Vel: Vec2f{X: pv.Vel.Y, Y: pv.Vel.Z},
	}
}

// Zero returns true if X and Y are both 0.
func (v Vec2) Zero() bool {
	return v.X == 0 && v.Y == 0
}

// Clockwise90 rotates the given direction by 90°
func (v Vec2) Clockwise90() Vec2 {
	switch v {
	case N:
		return E
	case NE:
		return SE
	case E:
		return S
	case SE:
		return SW
	case S:
		return W
	case SW:
		return NW
	case W:
		return N
	case NW:
		return NE
	}
	return v
}

// Clockwise45 rotates the given direction by 45° (and toggles length between 1 and sqrt2)
func (v Vec2) Clockwise45() Vec2 {
	switch v {
	case N:
		return NE
	case NE:
		return E
	case E:
		return SE
	case SE:
		return S
	case S:
		return SW
	case SW:
		return W
	case W:
		return NW
	case NW:
		return N
	}
	return v
}

// CounterClockwise90 rotates the given direction by 90°
func (v Vec2) CounterClockwise90() Vec2 {
	switch v {
	case E:
		return N
	case SE:
		return NE
	case S:
		return E
	case SW:
		return SE
	case W:
		return S
	case NW:
		return SW
	case N:
		return W
	case NE:
		return NW
	}
	return v
}

// CounterClockwise45 rotates the given direction by 45° (and toggles length between 1 and sqrt2)
func (v Vec2) CounterClockwise45() Vec2 {
	switch v {
	case E:
		return NE
	case SE:
		return E
	case S:
		return SE
	case SW:
		return S
	case W:
		return SW
	case NW:
		return W
	case N:
		return NW
	case NE:
		return N
	}
	return v
}

// Zero returns true if X and Y are both 0.
func (v Vec2f) Zero() bool {
	return v.X == 0 && v.Y == 0
}

// NearZero returns true if X and Y are both within `within` of 0.
func (v Vec2f) NearZero(within float64) bool {
	return v.Close(Vec2f{}, within)
}

// Zero returns true if X, Y and Z are all 0.
func (v Vec3) Zero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// Zero returns true if X, Y and Z are all 0.
func (v Vec3f) Zero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// NearZero returns true if X, Y, and Z are all within `within` of 0.
func (v Vec3f) NearZero(within float64) bool {
	return v.Close(Vec3f{}, within)
}
