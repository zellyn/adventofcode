package main

import (
	"testing"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

var input = util.MustReadSingleInt("input")
var example = 10

func TestPart1(t *testing.T) {
	testdata := []struct {
		name   string
		input  int
		target geom.Vec2
		want   int
	}{
		{
			name:   "example",
			input:  example,
			target: geom.V2(7, 4),
			want:   11,
		},
		{
			name:   "input",
			input:  input,
			target: geom.V2(31, 39),
			want:   82,
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
		name  string
		input int
		steps int
		want  int
	}{
		{
			name:  "example",
			input: example,
			steps: 50,
			want:  151,
		},
		{
			name:  "input",
			input: input,
			steps: 50,
			want:  138,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.steps)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, tt.steps)=%d; got %d", tt.want, got)
			}
		})
	}
}
