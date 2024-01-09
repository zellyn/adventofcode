package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestDistances(t *testing.T) {
	example := util.TrimmedLines(`
		London to Dublin = 464
		London to Belfast = 518
		Dublin to Belfast = 141`)
	input, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	_ = input
	testdata := []struct {
		name         string
		input        []string
		wantShortest int
		wantLongest  int
	}{
		{
			name:         "example",
			input:        example,
			wantShortest: 605,
			wantLongest:  982,
		},
		{
			name:         "real input",
			input:        input,
			wantShortest: 251,
			wantLongest:  898,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotShortest, err := shortestDistance(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if gotShortest != tt.wantShortest {
				t.Errorf("want shortestDistance(input)=%d; got %d", tt.wantShortest, gotShortest)
			}

			gotLongest, err := longestDistance(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if gotLongest != tt.wantLongest {
				t.Errorf("want LongestDistance(input)=%d; got %d", tt.wantLongest, gotLongest)
			}

		})
	}
}
