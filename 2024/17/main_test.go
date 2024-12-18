package main

import (
	"slices"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`)

var example2 = util.TrimmedLines(`
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  []int
	}{
		{
			name:  "example",
			input: example1,
			want:  []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0},
		},
		{
			name:  "input",
			input: input,
			want:  []int{7, 4, 2, 5, 1, 4, 6, 0, 4},
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
				return
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("Want part1(tt.input)=%q; got %q", intsToStr(tt.want), intsToStr(got))
			}
		})
	}
}

func XTestYields(t *testing.T) {
	s, err := parse(example2)
	if err != nil {
		t.Fatal(err)
	}

	got := s.yields(117440, s.program)
	if got != true {
		t.Errorf("want s.yields(117440, s.program) == true; got false")
	}

	got = s.yields(42, s.program)
	if got != false {
		t.Errorf("want s.yields(42, s.program) == false; got true")
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		bits  int
		want  int
	}{
		{
			name:  "example",
			input: example2,
			bits:  2,
			want:  117440,
		},
		// {
		// 	name:  "input",
		// 	input: input,
		// 	bits:  3,
		// 	want:  164278764924605,
		// },
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.bits)
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
