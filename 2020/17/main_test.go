package main

import (
	"testing"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/util"
)

var example = charmap.Parse(util.TrimmedLines(`
.#.
..#
###
`))

var input = charmap.MustRead("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input charmap.M
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  112,
		},
		{
			name:  "input",
			input: input,
			want:  276,
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
		input charmap.M
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  848,
		},
		{
			name:  "input",
			input: input,
			want:  2136,
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
