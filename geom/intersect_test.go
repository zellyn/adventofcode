package geom

import (
	"testing"
)

func TestIntersect2D(t *testing.T) {
	testdata := []struct {
		name           string
		pv1, pv2       PosVel2
		wantPos        Vec2f
		wantT1, wantT2 float64
		wantOk         bool
	}{
		{
			name:    "perpendicular",
			pv1:     PV2(1, -1, 1, 2),
			pv2:     PV2(1, 4, 2, -1),
			wantPos: V2f(3, 3),
			wantT1:  2,
			wantT2:  1,
			wantOk:  true,
		},
		// Different cases I made up.
		{
			name:    "equal positions, zero velocities",
			pv1:     PV2(10, 9, 0, 0),
			pv2:     PV2(10, 9, 0, 0),
			wantPos: V2f(10, 9),
			wantT1:  0,
			wantT2:  0,
			wantOk:  true,
		},
		{
			name:    "equal positions, different velocities",
			pv1:     PV2(10, 9, 1, 2),
			pv2:     PV2(10, 9, -1, -2),
			wantPos: V2f(10, 9),
			wantT1:  0,
			wantT2:  0,
			wantOk:  true,
		},
		{
			name:    "both zero velocity",
			pv1:     PV2(10, 9, 0, 0),
			pv2:     PV2(8, 7, 0, 0),
			wantPos: V2f(0, 0),
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "b is on a's line",
			pv1:     PV2(10, 11, 0, 4),
			pv2:     PV2(10, 19, 7, 8),
			wantPos: V2f(10, 19),
			wantT1:  2,
			wantT2:  0,
			wantOk:  true,
		},
		{
			name:    "b is on a's line, in the past",
			pv1:     PV2(1, 2, -3, 4),
			pv2:     PV2(10, -10, 7, 8),
			wantPos: V2f(10, -10),
			wantT1:  -3,
			wantT2:  0,
			wantOk:  true,
		},
		{
			name:    "a is on b's line",
			pv1:     PV2(-10, 6, 3, 4),
			pv2:     PV2(5, 6, -3, 0),
			wantPos: V2f(-10, 6),
			wantT1:  0,
			wantT2:  5,
			wantOk:  true,
		},
		{
			name:    "a is on b's line, in the past",
			pv1:     PV2(-9, 22, 3, 4),
			pv2:     PV2(5, 6, 7, -8),
			wantPos: V2f(-9, 22),
			wantT1:  0,
			wantT2:  -2,
			wantOk:  true,
		},
		{
			name:    "a has zero velocity, not colinear",
			pv1:     PV2(1, 2, 0, 0),
			pv2:     PV2(5, 6, 7, 8),
			wantPos: V2f(0, 0),
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "b has zero velocity, not colinear",
			pv1:     PV2(1, 2, 3, 4),
			pv2:     PV2(5, 6, 0, 0),
			wantPos: V2f(0, 0),
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "parallel, not colinear",
			pv1:     PV2(1, 2, 3, 4),
			pv2:     PV2(5, 6, 3, 4),
			wantPos: V2f(0, 0),
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "parallel, colinear",
			pv1:     PV2(2, 3, 1, 2),
			pv2:     PV2(3, 5, 2, 4),
			wantPos: V2f(3, 5),
			wantT1:  1,
			wantT2:  0,
			wantOk:  true,
		},
		{
			name:    "vertical parallel, not colinear",
			pv1:     PV2(1, 2, 0, 4),
			pv2:     PV2(5, 6, 0, 4),
			wantPos: V2f(0, 0),
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "horizontal parallel, not colinear",
			pv1:     PV2(1, 2, 3, 0),
			pv2:     PV2(5, 6, 3, 0),
			wantPos: V2f(0, 0),
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "a is vertical",
			pv1:     PV2(1, 2, 0, 4),
			pv2:     PV2(9, 6, 2, 8),
			wantPos: V2f(1, -26),
			wantT1:  -7,
			wantT2:  -4,
			wantOk:  true,
		},
		{
			name:    "b is vertical",
			pv1:     PV2(-10, 2, 2, 1),
			pv2:     PV2(0, 0, 0, -1),
			wantPos: V2f(0, 7),
			wantT1:  5,
			wantT2:  -7,
			wantOk:  true,
		},
		{
			name:    "basic intersection",
			pv1:     PV2(1, 2, 2, 1),
			pv2:     PV2(1, 6, 2, -1),
			wantPos: V2f(5, 4),
			wantT1:  2,
			wantT2:  2,
			wantOk:  true,
		},
		{
			name:    "a is horizontal",
			pv1:     PV2(1, 2, 3, 0),
			pv2:     PV2(8, 6, 2, 2),
			wantPos: V2f(4, 2),
			wantT1:  1,
			wantT2:  -2,
			wantOk:  true,
		},
		{
			name:    "b is horizontal",
			pv1:     PV2(1, 2, 3, 4),
			pv2:     PV2(2, 8, 7, 0),
			wantPos: V2f(5.5, 8),
			wantT1:  1.5,
			wantT2:  0.5,
			wantOk:  true,
		},
		{
			name:    "all zeros",
			pv1:     PosVel2{},
			pv2:     PosVel2{},
			wantPos: Vec2f{},
			wantT1:  0,
			wantT2:  0,
			wantOk:  true,
		},

		// AdventOfCode 2023, day 24 example values ------------------------------
		{
			name:    "aoc-example1",
			pv1:     PV2(19, 13, -2, 1),
			pv2:     PV2(18, 19, -1, -1),
			wantPos: Vec2f{X: 14.3333, Y: 15.3333},
			wantT1:  2.3333,
			wantT2:  3.6666,
			wantOk:  true,
		},
		{
			name:    "aoc-example2",
			pv1:     PV2(19, 13, -2, 1),
			pv2:     PV2(20, 25, -2, -2),
			wantPos: Vec2f{X: 11.6667, Y: 16.6667},
			wantT1:  3.6666,
			wantT2:  4.1666,
			wantOk:  true,
		},
		{
			name:    "aoc-example3",
			pv1:     PV2(19, 13, -2, 1),
			pv2:     PV2(12, 31, -1, -2),
			wantPos: Vec2f{X: 6.2, Y: 19.4},
			wantT1:  6.3999,
			wantT2:  5.8,
			wantOk:  true,
		},
		{
			name:    "aoc-example4",
			pv1:     PV2(19, 13, -2, 1),
			pv2:     PV2(20, 19, 1, -5),
			wantPos: Vec2f{X: 21.4444, Y: 11.7777},
			wantT1:  -1.2222,
			wantT2:  1.4444,
			wantOk:  true,
		},
		{
			name:    "aoc-example5",
			pv1:     PV2(18, 19, -1, -1),
			pv2:     PV2(20, 25, -2, -2),
			wantPos: Vec2f{X: 0, Y: 0},
			wantT1:  0,
			wantT2:  0,
			wantOk:  false,
		},
		{
			name:    "aoc-example6",
			pv1:     PV2(18, 19, -1, -1),
			pv2:     PV2(12, 31, -1, -2),
			wantPos: Vec2f{X: -6, Y: -5},
			wantT1:  24,
			wantT2:  18,
			wantOk:  true,
		},
		{
			name:    "aoc-example7",
			pv1:     PV2(18, 19, -1, -1),
			pv2:     PV2(20, 19, 1, -5),
			wantPos: Vec2f{X: 19.6666, Y: 20.6666},
			wantT1:  -1.6666,
			wantT2:  -0.3333,
			wantOk:  true,
		},
		{
			name:    "aoc-example8",
			pv1:     PV2(20, 25, -2, -2),
			pv2:     PV2(12, 31, -1, -2),
			wantPos: Vec2f{X: -2, Y: 3},
			wantT1:  11,
			wantT2:  14,
			wantOk:  true,
		},
		{
			name:    "aoc-example9",
			pv1:     PV2(20, 25, -2, -2),
			pv2:     PV2(20, 19, 1, -5),
			wantPos: Vec2f{X: 19, Y: 24},
			wantT1:  0.5,
			wantT2:  -1,
			wantOk:  true,
		},
		{
			name:    "aoc-example10",
			pv1:     PV2(12, 31, -1, -2),
			pv2:     PV2(20, 19, 1, -5),
			wantPos: Vec2f{X: 16, Y: 39},
			wantT1:  -4,
			wantT2:  -4,
			wantOk:  true,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {

			logErr := func(field string, wantVal, gotVal any) {
				t.Errorf("want %v.Intersect(%v) field %s == %v; got %v", tt.pv1, tt.pv2, field, wantVal, gotVal)
			}

			gotPos, gotT1, gotT2, gotOk := tt.pv1.Intersect(tt.pv2)

			if gotOk != tt.wantOk {
				logErr("ok", tt.wantOk, gotOk)
				t.Fail()
			}

			if !tt.wantPos.Close(gotPos, 1e-3) {
				logErr("pos", tt.wantPos, gotPos)
			}
			if !Close(tt.wantT1, gotT1, 1e-3) {
				logErr("t1", tt.wantT1, gotT1)
			}
			if !Close(tt.wantT2, gotT2, 1e-3) {
				logErr("t2", tt.wantT2, gotT2)
			}
		})
	}
}

func TestIntersect3D(t *testing.T) {
	testdata := []struct {
		name           string
		pv1, pv2       PosVel3
		wantPos        Vec3f
		wantT1, wantT2 float64
		wantOk         bool
	}{
		{
			name:    "2d-perpendicular,z=17",
			pv1:     PV3(1, -1, 17, 1, 2, 0),
			pv2:     PV3(1, 4, 17, 2, -1, 0),
			wantPos: V3f(3, 3, 17),
			wantT1:  2,
			wantT2:  1,
			wantOk:  true,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {

			logErr := func(field string, wantVal, gotVal any) {
				t.Errorf("want %v.Intersect(%v) field %s == %v; got %v", tt.pv1, tt.pv2, field, wantVal, gotVal)
			}

			gotPos, gotT1, gotT2, gotOk := tt.pv1.Intersect(tt.pv2)

			if gotOk != tt.wantOk {
				logErr("ok", tt.wantOk, gotOk)
				t.Fail()
			}

			if !tt.wantPos.Close(gotPos, 1e-3) {
				logErr("pos", tt.wantPos, gotPos)
			}
			if !Close(tt.wantT1, gotT1, 1e-3) {
				logErr("t1", tt.wantT1, gotT1)
			}
			if !Close(tt.wantT2, gotT2, 1e-3) {
				logErr("t2", tt.wantT2, gotT2)
			}
		})
	}
}
