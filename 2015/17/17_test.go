package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
	20
	15
	10
	5
	5`)

func TestStuff(t *testing.T) {
	input, err := ioutil.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}

	testdata := []struct {
		name  string
		input []string
		goal  int
		want1 int
		want2 int
	}{
		{
			name:  "example",
			input: example,
			goal:  25,
			want1: 4,
			want2: 3,
		},
		{
			name:  "real input",
			input: input,
			goal:  150,
			want1: 1304,
			want2: 18,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := combinations(tt.input, tt.goal)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want1 {
				t.Errorf("want combinations(input, %d)=%d; got %d", tt.goal, tt.want1, got)
			}

			got, err = smallestCombinations(tt.input, tt.goal)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want2 {
				t.Errorf("want smallestCombinations(input, %d)=%d; got %d", tt.goal, tt.want2, got)
			}
		})
	}
}
