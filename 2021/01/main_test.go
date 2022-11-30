package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.MustStringsToInts(util.TrimmedLines(`
199
200
208
210
200
207
240
269
260
263
`))

var input = util.MustReadFileInts("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  7,
		},
		{
			name:  "input",
			input: input,
			want:  1387,
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
		input []int
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  5,
		},
		{
			name:  "input",
			input: input,
			want:  1362,
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
