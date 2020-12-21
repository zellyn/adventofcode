package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
939
7,13,x,x,59,x,31,19
`)

var example2 = util.TrimmedLines(`
0
17,x,13,19
`)

var example3 = util.TrimmedLines(`
0
67,7,59,61
`)

var example4 = util.TrimmedLines(`
0
67,x,7,59,61
`)

var example5 = util.TrimmedLines(`
0
67,7,x,59,61
`)

var example6 = util.TrimmedLines(`
0
1789,37,47,1889
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example1,
			want:  295,
		},
		{
			name:  "input",
			input: input,
			want:  3246,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
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
			name:  "example1",
			input: example1,
			want:  1068781,
		},
		{
			name:  "example2",
			input: example2,
			want:  3417,
		},
		{
			name:  "example3",
			input: example3,
			want:  754018,
		},
		{
			name:  "example4",
			input: example4,
			want:  779210,
		},
		{
			name:  "example5",
			input: example5,
			want:  1261476,
		},
		{
			name:  "example6",
			input: example6,
			want:  1202161486,
		},
		{
			name:  "input",
			input: input,
			want:  1010182346291467,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
