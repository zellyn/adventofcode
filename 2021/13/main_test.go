package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
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
			want:  17,
		},
		{
			name:  "input",
			input: input,
			want:  743,
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
