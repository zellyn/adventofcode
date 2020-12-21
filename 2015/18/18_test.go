package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestStuff(t *testing.T) {
	example := util.TrimmedLines(`
		.#.#.#
		...##.
		#....#
		..#...
		#.#..#
		####..`)

	input, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}

	testdata := []struct {
		name   string
		input  []string
		steps1 int
		steps2 int
		want1  int
		want2  int
	}{
		{
			name:   "example",
			input:  example,
			steps1: 4,
			steps2: 5,
			want1:  4,
			want2:  17,
		},
		{
			name:   "real input",
			input:  input,
			steps1: 100,
			steps2: 100,
			want1:  768,
			want2:  781,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got := lightsAfter(tt.input, tt.steps1)
			if got != tt.want1 {
				t.Errorf("want lightsAfter(input, %d)=%d; got %d", tt.steps1, tt.want1, got)
			}

			got = stuckLightsAfter(tt.input, tt.steps2)
			if got != tt.want2 {
				t.Errorf("want stuckLightsAfter(input, %d)=%d; got %d", tt.steps1, tt.want2, got)
			}
		})
	}

}
