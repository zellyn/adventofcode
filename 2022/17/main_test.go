package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

var input = util.MustReadFileString("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  3068,
		},
		{
			name:  "input",
			input: input,
			want:  3065,
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
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  1514285714288,
		},
		{
			name:  "input",
			input: input,
			want:  1562536022966,
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
