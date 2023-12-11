package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
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
			input: example,
			want:  374,
		},
		{
			name:  "input",
			input: input,
			want:  9693756,
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
		name   string
		input  []string
		expand int
		want   int
	}{
		{
			name:   "example-10",
			input:  example,
			expand: 10,
			want:   1030,
		},
		{
			name:   "example-100",
			input:  example,
			expand: 100,
			want:   8410,
		},
		{
			name:   "input",
			input:  input,
			expand: 1e6,
			want:   717878258016,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.expand)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, %d)=%d; got %d", tt.expand, tt.want, got)
			}
		})
	}
}
