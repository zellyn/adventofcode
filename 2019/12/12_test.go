package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

var mainInput = util.MustReadFileString("input")

var testdata = []struct {
	input  string
	steps  int
	pos    []vec3
	vel    []vec3
	energy int
	repeat int
	xl     int
	yl     int
	zl     int
}{
	{
		input:  "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>",
		steps:  10,
		pos:    []vec3{{2, 1, -3}, {1, -8, 0}, {3, -6, 1}, {2, 0, 4}},
		vel:    []vec3{{-3, -2, 1}, {-1, 1, 3}, {3, 2, -3}, {1, -1, -1}},
		energy: 179,
		repeat: 2772,
		xl:     18, // 2 * 3 * 3
		yl:     28, // 2 * 2 * 7
		zl:     44, // 2 * 2 * 11
	},
	{
		input:  "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>",
		steps:  100,
		pos:    []vec3{{8, -12, -9}, {13, 16, -3}, {-29, -11, -1}, {16, -13, 23}},
		vel:    []vec3{{-7, 3, 0}, {3, -11, -5}, {-3, 7, 4}, {7, 1, 1}},
		energy: 1940,
		repeat: 4686774924,
		xl:     2028,
		yl:     5898,
		zl:     4702,
	},
	{
		input:  mainInput,
		steps:  1000,
		energy: 7098,
		repeat: 400128139852752,
		xl:     135024,
		yl:     231614,
		zl:     102356,
	},
}

func TestEnergy(t *testing.T) {
	for i, tt := range testdata {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			pos, err := geom.ParseVec3Lines(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			vel := make([]vec3, len(pos))
			for j := 0; j < tt.steps; j++ {
				pos, vel = step(pos, vel)
			}
			if len(tt.pos) > 0 {
				if !reflect.DeepEqual(pos, tt.pos) {
					t.Errorf("want pos==%v; got %v", tt.pos, pos)
				}
			}
			if len(tt.vel) > 0 {
				if !reflect.DeepEqual(vel, tt.vel) {
					t.Errorf("want vel==%v; got %v", tt.vel, vel)
				}
			}
			e := energy(pos, vel)
			if e != tt.energy {
				t.Errorf("want energy==%d; got %d", e, tt.energy)
			}
		})
	}
}

func TestPeriod(t *testing.T) {
	for i, tt := range testdata {
		t.Run(fmt.Sprintf("test%d", i), func(t *testing.T) {
			pos, err := geom.ParseVec3Lines(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			xs := make([]int, 4)
			ys := make([]int, 4)
			zs := make([]int, 4)
			for i, vec := range pos {
				xs[i] = vec.X
				ys[i] = vec.Y
				zs[i] = vec.Z
			}
			xlen := period(xs, 1000000)
			if xlen != tt.xl {
				t.Errorf("Want xlen=%d; got xlen=%d", tt.xl, xlen)
			}
			ylen := period(ys, 1000000)
			if ylen != tt.yl {
				t.Errorf("Want ylen=%d; got ylen=%d", tt.yl, ylen)
			}
			zlen := period(zs, 1000000)
			if zlen != tt.zl {
				t.Errorf("Want zlen=%d; got zlen=%d", tt.zl, zlen)
			}

			period := multmaps(factors(xlen), factors(ylen), factors(zlen))
			if period != tt.repeat {
				t.Errorf("Want repeat=%d; got %d", tt.repeat, period)
			}
		})
	}

}

func TestFactors(t *testing.T) {
	n := 167624
	got := factors(n)
	want := map[int]int{
		2:   3,
		23:  1,
		911: 1,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want factors(%d)==%v; got %v", n, want, got)
	}
}
