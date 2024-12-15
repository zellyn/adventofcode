package main

import (
	"testing"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name    string
		input   []string
		size    geom.Vec2
		seconds int
		want    int
	}{
		{
			name:    "example",
			input:   example,
			size:    geom.V2(11, 7),
			seconds: 100,
			want:    12,
		},
		{
			name:    "input",
			input:   input,
			size:    geom.V2(101, 103),
			seconds: 100,
			want:    224554908,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.size, tt.seconds)
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
		name    string
		input   []string
		size    geom.Vec2
		seconds int
		want    int
	}{
		{
			name:  "input",
			input: input,
			size:  geom.V2(101, 103),
			want:  6644,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.size)
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
