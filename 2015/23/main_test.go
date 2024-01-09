package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestParts(t *testing.T) {
	example1 := util.TrimmedLines(`
		inc a
		jio a, +2
		tpl a
		inc a`)
	input, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	testdata := []struct {
		name    string
		program []string
		startA  int
		startB  int
		wantA   int
		wantB   int
	}{
		{
			name:    "example1",
			program: example1,
			wantA:   2,
			wantB:   0,
		},
		{
			name:    "part 1",
			program: input,
			wantA:   1,
			wantB:   170,
		},
		{
			name:    "part 2",
			program: input,
			startA:  1,
			wantA:   1,
			wantB:   247,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB, err := runProgram(tt.program, tt.startA, tt.startB)
			if err != nil {
				t.Fatal(err)
			}
			if gotA != tt.wantA {
				t.Errorf("Want gotA=%d; got %d", tt.wantA, gotA)
			}
			if gotB != tt.wantB {
				t.Errorf("Want gotB=%d; got %d", tt.wantB, gotB)
			}
		})
	}
}
