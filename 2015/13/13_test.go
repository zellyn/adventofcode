package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestDistances(t *testing.T) {
	example := util.TrimmedLines(`
		Alice would gain 54 happiness units by sitting next to Bob.
		Alice would lose 79 happiness units by sitting next to Carol.
		Alice would lose 2 happiness units by sitting next to David.
		Bob would gain 83 happiness units by sitting next to Alice.
		Bob would lose 7 happiness units by sitting next to Carol.
		Bob would lose 63 happiness units by sitting next to David.
		Carol would lose 62 happiness units by sitting next to Alice.
		Carol would gain 60 happiness units by sitting next to Bob.
		Carol would gain 55 happiness units by sitting next to David.
		David would gain 46 happiness units by sitting next to Alice.
		David would lose 7 happiness units by sitting next to Bob.
		David would gain 41 happiness units by sitting next to Carol.`)
	input, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	_ = input
	testdata := []struct {
		name        string
		input       []string
		wantBest    int
		wantWithYou int
	}{
		{
			name:        "example",
			input:       example,
			wantBest:    330,
			wantWithYou: 286,
		},
		{
			name:        "real input",
			input:       input,
			wantBest:    709,
			wantWithYou: 668,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotBest, err := best(tt.input, false)
			if err != nil {
				t.Fatal(err)
			}
			if gotBest != tt.wantBest {
				t.Errorf("want best(input, false)=%d; got %d", tt.wantBest, gotBest)
			}

			gotWithYou, err := best(tt.input, true)
			if err != nil {
				t.Fatal(err)
			}
			if gotWithYou != tt.wantWithYou {
				t.Errorf("want best(input, true)=%d; got %d", tt.wantWithYou, gotWithYou)
			}

		})
	}
}
