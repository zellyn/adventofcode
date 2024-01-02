package geom

import (
	"math/big"
	"strconv"

	"github.com/zellyn/adventofcode/linalg"
	mymath "github.com/zellyn/adventofcode/math"
)

// ReduceDir computes the rise and run of the slope of the vector, in
// lowest terms, with a positive X. If X and Y are both zero, it
// returns (0,0). If the X component of the vector is 0, it returns
// (0,1), and if the Y component is zero, it returns (1,0).
func (v Vec2) ReduceDir() Vec2 {
	if v.X == 0 {
		if v.Y == 0 {
			return v
		}
		return Vec2{X: 0, Y: 1}
	}
	if v.Y == 0 {
		return Vec2{X: 1, Y: 0}
	}

	if v.X < 0 {
		v.X, v.Y = -v.X, -v.Y
	}

	gcd := mymath.GCD(v.X, abs(v.Y))
	return v.IntDiv(gcd)
}

// Parallel returns true if the two velocity vectors are parallel.
func (pv PosVel2) Parallel(other PosVel2) bool {
	if pv.Vel.Zero() || other.Vel.Zero() {
		return false
	}
	return pv.Vel.ReduceDir() == other.Vel.ReduceDir()
}

// Parallel returns true if the two velocity vectors are parallel.
func (pv PosVel3) Parallel(other PosVel3) bool {
	if pv.Vel.Zero() || other.Vel.Zero() {
		return false
	}
	return pv.Vel.ReduceDir() == other.Vel.ReduceDir()
}

// ReduceDir returns X, Y, Z in reduced form: first non-zero element
// is positive, no common factors.
func (v Vec3) ReduceDir() Vec3 {
	if v.Zero() {
		return v
	}
	if v.X == 0 && v.Y == 0 {
		return Vec3{Z: 1}
	}
	if v.Y == 0 && v.Z == 0 {
		return Vec3{X: 1}
	}
	if v.Z == 0 && v.X == 0 {
		return Vec3{Y: 1}
	}

	if v.X < 0 || (v.X == 0 && v.Y < 0) {
		v = v.Mul(-1)
	}

	gcd := mymath.MultiGCD(v.X, v.Y, v.Z)
	return v.IntDiv(gcd)
}

// GetCoord returns a coordinate by index: X=0, Y=1, Z=2.
func (v Vec3) GetCoord(i int) int {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic("Invalid coordinate for Vec3: " + strconv.Itoa(i))
	}
}

// GetCoord returns a coordinate by index: X=0, Y=1, Z=2.
func (v Vec3f) GetCoord(i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic("Invalid coordinate for Vec3f: " + strconv.Itoa(i))
	}
}

// MinCoord returns a minimum coordinate value and index. X=0, Y=1, Z=2.
func (v Vec3) MinCoord() (value int, index int) {
	value = v.X
	index = 0
	if v.Y < value {
		value = v.Y
		index = 1
	}
	if v.Z < value {
		value = v.Z
		index = 2
	}
	return value, index
}

// MaxCoord returns a maximum coordinate value and index. X=0, Y=1, Z=2.
func (v Vec3) MaxCoord() (value int, index int) {
	value = v.X
	index = 0
	if v.Y > value {
		value = v.Y
		index = 1
	}
	if v.Z > value {
		value = v.Z
		index = 2
	}
	return value, index
}

// SameLine returns true if the two positions and directions represent
// the same line.
func (pv PosVel2) SameLine(other PosVel2) bool {
	if pv.Vel.Zero() || other.Vel.Zero() {
		return false
	}
	rr := pv.Vel.ReduceDir()
	return rr == other.Vel.ReduceDir() && rr == pv.Pos.Sub(other.Pos).ReduceDir()
}

// SameLine returns true if the two positions and directions represent
// the same line.
func (pv PosVel3) SameLine(other PosVel3) bool {
	if pv.Vel.Zero() || other.Vel.Zero() {
		return false
	}
	rr := pv.Vel.ReduceDir()
	return rr == other.Vel.ReduceDir() && rr == pv.Pos.Sub(other.Pos).ReduceDir()
}

// Colinear returns true if the given point lies on the line
// represented by the PosVel.
func (pv PosVel2) Colinear(pt Vec2) bool {
	return pv.Vel.ReduceDir() == pt.Sub(pv.Pos).ReduceDir()
}

// Colinear returns true if the given point lies on the line
// represented by the PosVel.
func (pv PosVel3) Colinear(pt Vec3) bool {
	return pv.Vel.ReduceDir() == pt.Sub(pv.Pos).ReduceDir()
}

// isRatZero checks if the given rational is zero.
func isRatZero(x *big.Rat) bool {
	return x.Num().BitLen() == 0
}

// divideOrZero is a helper to divide a numerator by denominator, but
// return 0 if the denominator is zero.
func divideOrZero(num, denom *big.Rat) (f float64, exact bool) {
	if isRatZero(denom) {
		return 0, true
	}

	res := &big.Rat{}
	return res.Inv(denom).Mul(res, num).Float64()
}

// Intersect calculates the intersection point of the lines defined by
// the two PosVels. If they are parallel and not co-linear, it returns
// `(0,0), 0, 0, false`. Otherwise it returns the intersection
// point, the time at which each of the vectors would reach (or would
// have reached) that point, and true. If they are co-linear, it
// returns one of the vector's positions, that vector's time=0, and
// the other vector's time what it would take to get there.
func (pv1 PosVel2) Intersect(pv2 PosVel2) (Vec2f, float64, float64, bool) {
	rows := [][]int{
		{pv1.Vel.X, -pv2.Vel.X, pv2.Pos.X - pv1.Pos.X},
		{pv1.Vel.Y, -pv2.Vel.Y, pv2.Pos.Y - pv1.Pos.Y},
	}

	m := linalg.NewMatrix(rows)
	rank := m.Rank()

	// No intersection.
	if rank == -1 {
		return Vec2f{}, 0, 0, false
	}

	// identical vectors with all-zero velocities.
	if rank == 0 {
		return pv1.Pos.ToF(), 0, 0, true
	}

	coeffs, known := m.KnownCoefficientFloats()
	knownCount := 0
	for _, k := range known {
		if k {
			knownCount++
		}
	}

	if knownCount == 2 {
		t1 := coeffs[0]
		t2 := coeffs[1]
		return pv1.Pos.ToF().Add(pv1.Vel.ToF().Mul(t1)), t1, t2, true
	}

	// sanity check.
	if rank != 1 {
		panic("rank should be 1 by here!")
	}

	coefficients := m.Rows()[0].Coefficients()
	k := m.Rows()[0].Constant()

	s, exactS := divideOrZero(coefficients[0], k)
	t, exactT := divideOrZero(coefficients[1], k)

	if !exactS || !exactT {
		return Vec2f{}, 0, 0, false
	}

	useT := s == 0 || (s > 0 && t > 0 && t < s) || (s < 0 && t > s)
	if useT {
		return pv1.Pos.ToF(), 0, t, true
	} else {
		return pv2.Pos.ToF(), s, 0, true
	}
}

// Intersect calculates the intersection point of the lines defined by
// the two PosVels. If they are parallel and not co-linear, it returns
// `(0,0,0), 0, 0, false`. Otherwise it returns the intersection
// point, the time at which each of the vectors would reach (or would
// have reached) that point, and true. If they are co-linear, it
// returns one of the vector's positions, that vector's time=0, and
// the other vector's time what it would take to get there.
func (pv1 PosVel3) Intersect(pv2 PosVel3) (Vec3f, float64, float64, bool) {
	rows := [][]int{
		{pv1.Vel.X, -pv2.Vel.X, pv2.Pos.X - pv1.Pos.X},
		{pv1.Vel.Y, -pv2.Vel.Y, pv2.Pos.Y - pv1.Pos.Y},
		{pv1.Vel.Z, -pv2.Vel.Z, pv2.Pos.Z - pv1.Pos.Z},
	}

	m := linalg.NewMatrix(rows)
	rank := m.Rank()

	// No intersection.
	if rank == -1 {
		return Vec3f{}, 0, 0, false
	}

	// identical vectors with all-zero velocities.
	if rank == 0 {
		return pv1.Pos.ToF(), 0, 0, true
	}

	coeffs, known := m.KnownCoefficientFloats()
	knownCount := 0
	for _, k := range known {
		if k {
			knownCount++
		}
	}

	if knownCount == 2 {
		t1 := coeffs[0]
		t2 := coeffs[1]
		return pv1.Pos.ToF().Add(pv1.Vel.ToF().Mul(t1)), t1, t2, true
	}

	// sanity check.
	if rank != 1 {
		panic("rank should be 1 by here!")
	}

	coefficients := m.Rows()[0].Coefficients()
	k := m.Rows()[0].Constant()

	s, exactS := divideOrZero(coefficients[0], k)
	t, exactT := divideOrZero(coefficients[1], k)

	if !exactS || !exactT {
		return Vec3f{}, 0, 0, false
	}

	useT := s == 0 || (s > 0 && t > 0 && t < s) || (s < 0 && t > s)
	if useT {
		return pv1.Pos.ToF(), 0, t, true
	} else {
		return pv2.Pos.ToF(), s, 0, true
	}
}

// Intersect calculates the intersection point of the lines defined by
// the two PosVels. If they are parallel and not co-linear, it returns
// `(0,0,0), 0, 0, false`. Otherwise it returns the intersection point,
// the time at which each of the vectors would reach (or would have
// reached) that point, and true. If they are co-linear, it returns
// `first_vector_position, 0, 0, true`.
func (pv1 PosVel3) IntersectOld(pv2 PosVel3, within float64) (Vec3f, float64, float64, bool) {
	// Degenerate case of same points.
	if pv1.Pos == pv2.Pos {
		return pv1.Pos.ToF(), 0, 0, true
	}

	// Zero velocity?
	if pv1.Vel.Zero() && pv2.Vel.Zero() {
		return Vec3f{}, 0, 0, false
	}

	okVal := func(i int, ok bool) float64 {
		if !ok {
			return 0
		}
		return float64(i)
	}

	// Matching X velocity components are zero.
	if pv1.Vel.X == 0 && pv2.Vel.X == 0 {
		// Different X position components? Nope.
		if pv1.Pos.X != pv2.Pos.X {
			return Vec3f{}, 0, 0, false
		}

		// Same X. Ok, let's just do the 2d intersect for YZ.
		pos2, t1, t2, ok := pv1.YZ().Intersect(pv2.YZ())
		return Vec3f{X: okVal(pv1.Pos.X, ok), Y: pos2.X, Z: pos2.Y}, t1, t2, ok
	}

	// Matching Y velocity components are zero.
	if pv1.Vel.Y == 0 && pv2.Vel.Y == 0 {
		// Different Y position components? Nope.
		if pv1.Pos.Y != pv2.Pos.Y {
			return Vec3f{}, 0, 0, false
		}

		// Same Y. Ok, let's just do the 2d intersect for XZ.
		pos2, t1, t2, ok := pv1.XZ().Intersect(pv2.XZ())
		return Vec3f{X: pos2.X, Y: okVal(pv1.Pos.Y, ok), Z: pos2.Y}, t1, t2, ok
	}

	// Matching Z velocity components are zero.
	if pv1.Vel.Z == 0 && pv2.Vel.Z == 0 {
		// Different Z position components? Nope.
		if pv1.Pos.Z != pv2.Pos.Z {
			return Vec3f{}, 0, 0, false
		}

		// Same Z. Ok, let's just do the 2d intersect for XY.
		pos2, t1, t2, ok := pv1.XY().Intersect(pv2.XY())
		return pos2.WithZ(okVal(pv1.Pos.Z, ok)), t1, t2, ok
	}

	// Degenerate case: pv2 lies on pv1's line.
	if pv1.Colinear(pv2.Pos) {
		_, index := pv1.Vel.Abs().MaxCoord()
		return pv2.Pos.ToF(), pv2.Pos.ToF().Sub(pv1.Pos.ToF()).GetCoord(index) / pv1.Vel.ToF().GetCoord(index), 0, true
	}

	// Degenerate case: pv1 lies on pv2's line.
	if pv2.Colinear(pv1.Pos) {
		_, index := pv2.Vel.Abs().MaxCoord()
		return pv1.Pos.ToF(), 0, pv1.Pos.ToF().Sub(pv2.Pos.ToF()).GetCoord(index) / pv2.Vel.ToF().GetCoord(index), true
	}

	// They aren't colinear, so if one has zero velocity, there's no way.
	if pv1.Vel.Zero() || pv2.Vel.Zero() {
		return Vec3f{}, 0, 0, false
	}

	// They aren't colinear, so if they're parallel, there's no way.
	if pv1.Parallel(pv2) {
		return Vec3f{}, 0, 0, false
	}

	return intersectHelper3(pv1, pv2, within)
}

func intersectHelper3(pv1, pv2 PosVel3, within float64) (Vec3f, float64, float64, bool) {
	px1, py1, pz1, vx1, vy1, vz1 := pv1.Pos.X, pv1.Pos.Y, pv1.Pos.Z, pv1.Vel.X, pv1.Vel.Y, pv1.Vel.Z
	px2, py2, pz2, vx2, vy2, vz2 := pv2.Pos.X, pv2.Pos.Y, pv2.Pos.Z, pv2.Vel.X, pv2.Vel.Y, pv2.Vel.Z
	/*
	   x₁ + s·dx₁ = x₂ + t·dx₂   s·dx₁ - t·dx₂ = x₂ - x₁
	   y₁ + s·dy₁ = y₂ + t·dy₂   s·dy₁ - t·dy₂ = y₂ - y₁
	   z₁ + s·dz₁ = z₂ + t·dz₂   s·dz₁ - t·dz₂ = z₂ - z₁
	*/

	rows := make([]Vec3, 0, 3)

	if row := V3(vx1, -vx2, px2-px1); !row.XY().Zero() {
		rows = append(rows, row)
	}
	if row := V3(vy1, -vy2, py2-py1); !row.XY().Zero() {
		rows = append(rows, row)
	}
	if row := V3(vz1, -vz2, pz2-pz1); !row.XY().Zero() {
		rows = append(rows, row)
	}

	panic("not imlemented")
}
