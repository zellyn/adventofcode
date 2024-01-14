package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = "abc"

var input = util.MustReadFileString("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name   string
		input  string
		target int
		want   int
	}{
		{
			name:   "example",
			input:  example,
			target: 64,
			want:   22728,
		},
		{
			name:   "input",
			input:  input,
			target: 64,
			want:   15168,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.target)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.target)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name   string
		input  string
		target int
		want   int
	}{
		{
			name:   "example",
			input:  example,
			target: 64,
			want:   22551,
		},
		{
			name:   "input",
			input:  input,
			target: 64,
			want:   20864,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.target)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
