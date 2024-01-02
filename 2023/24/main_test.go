package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name     string
		input    []string
		min, max float64
		want     int
	}{
		{
			name:  "example",
			input: example,
			min:   7,
			max:   27,
			want:  2,
		},
		{
			name:  "input",
			input: input,
			min:   200000000000000,
			max:   400000000000000,
			want:  13754,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.min, tt.max)
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
			name:  "example",
			input: example,
			want:  47,
		},
		{
			name:  "input",
			input: input,
			want:  711031616315001,
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
