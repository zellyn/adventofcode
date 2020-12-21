package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
F10
N3
F7
R90
F11
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want1 int
		want2 int
	}{
		{
			name:  "example",
			input: example,
			want1: 25,
			want2: 286,
		},
		{
			name:  "input",
			input: input,
			want1: 1133,
			want2: 61053,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got1, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got1 != tt.want1 {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want1, got1)
			}

			got2, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got2 != tt.want2 {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want2, got2)
			}
		})
	}
}
