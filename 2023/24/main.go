package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/linalg"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf
var debugf = func(debug bool, format string, args ...any) {
	if !debug {
		return
	}
	printf(format, args...)
}

func parse(inputs []string) ([]geom.PosVel3, error) {
	var res []geom.PosVel3

	for _, input := range inputs {
		ints, err := util.ParseInts(strings.ReplaceAll(strings.ReplaceAll(input, " @ ", ","), " ", ""), ",")
		if err != nil {
			return nil, fmt.Errorf("weird input %q: %w", input, err)
		}

		res = append(res, geom.PosVel3{
			Pos: geom.Vec3{X: ints[0], Y: ints[1], Z: ints[2]},
			Vel: geom.Vec3{X: ints[3], Y: ints[4], Z: ints[5]},
		})
	}

	return res, nil
}

func part1(inputs []string, minCoord, maxCoord float64) (int, error) {
	posVels3, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	posVels := util.Map(posVels3, geom.PosVel3.XY)

	count := 0
	for i, a := range posVels[:len(posVels)-1] {
		for j := i + 1; j < len(inputs); j++ {
			b := posVels[j]

			// printf("A=%s\nB=%s\n", a, b)

			pos, t1, t2, ok := a.Intersect(b)
			if !ok {
				continue
			}
			// printf(" intersection at %s t1=%g, t2=%g\n", pos, t1, t2)

			if t1 < 0 {
				continue
			}
			if t2 < 0 {
				continue
			}
			if pos.X < minCoord || pos.X > maxCoord || pos.Y < minCoord || pos.Y > maxCoord {
				continue
			}
			count++
		}
	}
	return count, nil
}

// Names of coefficients we're going to be building up.
var names = strings.Split("x y z u v w xv yu xw zu yw zv", " ")

func lineToRows(pv geom.PosVel3) ([]*linalg.Row, geom.Vec3) {
	a, b, c, d, e, f := pv.Pos.X, pv.Pos.Y, pv.Pos.Z, pv.Vel.X, pv.Vel.Y, pv.Vel.Z

	rows := []*linalg.Row{
		linalg.NewRow([]int{e, -d, 0, -b, a, 0, -1, 1, 0, 0, 0, 0}, a*e-b*d),
		linalg.NewRow([]int{f, 0, -d, -c, 0, a, 0, 0, -1, 1, 0, 0}, a*f-c*d),
		linalg.NewRow([]int{0, f, -e, 0, -c, b, 0, 0, 0, 0, -1, 1}, b*f-c*e),
	}

	// u ≠ d, v ≠ e, w ≠ f
	return rows, geom.V3(d, e, f)
}

func getGroup(posVels []geom.PosVel3) *linalg.Matrix {
	m := &linalg.Matrix{}
	for _, pv := range posVels[:4] {
		rows, _ := lineToRows(pv)
		m.AddRows(rows)
	}
	return m
}

func matrixToResult(m *linalg.Matrix) int {
	if m.Impossible() {
		return -1
	}
	result := &big.Rat{}

	cs, known := m.KnownCoefficients()

	for i := 0; i < 3; i++ {
		if !known[i] {
			return -1
		}

		result.Add(result, cs[i])
	}

	if !result.IsInt() || !result.Num().IsInt64() {
		return -1
	}

	return int(result.Num().Int64())
}

func part2(inputs []string) (int, error) {
	posVels3, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	m := getGroup(posVels3)
	return matrixToResult(m), nil
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
