package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
)

func TestParts(t *testing.T) {
	input, err := ioutil.ReadFileString("input")
	if err != nil {
		t.Fatal(err)
	}
	testdata := []struct {
		name      string
		input     string
		wantEnd   int
		wantTwice int
	}{
		{
			name:      "example1",
			input:     "R2, L3",
			wantEnd:   5,
			wantTwice: -1,
		},
		{
			name:      "example2",
			input:     "R2, R2, R2",
			wantEnd:   2,
			wantTwice: -1,
		},
		{
			name:      "example3",
			input:     "R5, L5, R5, R3",
			wantEnd:   12,
			wantTwice: -1,
		},
		{
			name:      "example4",
			input:     "R8, R4, R4, R8",
			wantEnd:   8,
			wantTwice: 4,
		},
		{
			name:      "input",
			input:     input,
			wantEnd:   287,
			wantTwice: 133,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotEnd, gotTwice, err := distance(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if gotEnd != tt.wantEnd || gotTwice != tt.wantTwice {
				t.Errorf("Want distance(input)=%d,%d; got %d,%d", tt.wantEnd, tt.wantTwice, gotEnd, gotTwice)
			}
		})
	}
}
