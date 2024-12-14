package main

import (
	"testing"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`)

var input = util.MustReadLines("input")

func XTestPresses(t *testing.T) {
	machines, err := parse(example)
	if err != nil {
		t.Fatal(err)
	}

	wants := []int{280, 0, 200, 0}

	for i := range machines {
		m := machines[i]
		want := wants[i]
		if got := presses2(m); got != want {
			t.Errorf("want presses(%v) == %d; got %d", m, want, got)
		}
	}
}

func TestPresses2Math(t *testing.T) {
	// m := machine{
	// 	a:     geom.V2(20, 84),
	// 	b:     geom.V2(35, 18),
	// 	prize: geom.V2(10000000002260, 10000000003300),
	// }
	// {{66 17} {17 68} {19741 12161}}
	m := machine{
		a:     geom.V2(66, 17),
		b:     geom.V2(17, 68),
		prize: geom.V2(19741, 12161),
	}
	presses := presses2(m)

	_ = presses
}

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  480,
		},
		{
			name:  "input",
			input: input,
			want:  31623,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  875318608908,
		},
		{
			name:  "input",
			input: input,
			want:  93209116744825,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
