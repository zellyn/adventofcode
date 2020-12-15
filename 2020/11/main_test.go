package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`)

var input = ioutil.MustReadLines("input")

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
			want1: 37,
			want2: 26,
		},
		{
			name:  "input",
			input: input,
			want1: 2316,
			want2: 2128,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got1 := part1(tt.input)
			if got1 != tt.want1 {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want1, got1)
			}

			got2 := part2(tt.input)
			if got2 != tt.want2 {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want2, got2)
			}
		})
	}
}
