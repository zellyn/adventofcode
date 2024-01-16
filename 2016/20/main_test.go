package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
5-8
0-2
4-7
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
			want:  3,
		},
		{
			name:  "input",
			input: input,
			want:  4793564,
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
		last  int
		want  int
	}{
		{
			name:  "example",
			input: example,
			last:  9,
			want:  2,
		},
		{
			name:  "input",
			input: input,
			last:  (1 << 32) - 1,
			want:  146,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.last)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, tt.last)=%d; got %d", tt.want, got)
			}
		})
	}
}
