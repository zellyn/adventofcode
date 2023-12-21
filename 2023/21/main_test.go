package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		steps int
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			steps: 6,
			want:  16,
		},
		{
			name:  "input",
			input: input,
			steps: 64,
			want:  3740,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.steps)
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
		steps int
		want  int
	}{
		{
			name:  "input",
			input: input,
			steps: 458,
			want:  186292,
		},
		{
			name:  "input",
			input: input,
			steps: 589,
			want:  307795,
		},
		{
			name:  "input",
			input: input,
			steps: 26501365,
			want:  620962518745459,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.steps)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
