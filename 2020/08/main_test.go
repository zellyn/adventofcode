package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name      string
		input     []string
		wantAcc   int
		wantFixed int
	}{
		{
			name:      "example",
			input:     example,
			wantAcc:   5,
			wantFixed: 8,
		},
		{
			name:      "input",
			input:     input,
			wantAcc:   1709,
			wantFixed: 1976,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotAcc, err := accBeforeLoop(tt.input)
			if err != nil {
				t.Error(err)
			}
			if gotAcc != tt.wantAcc {
				t.Errorf("Want accBeforeLoop(tt.input)=%d; got %d", tt.wantAcc, gotAcc)
			}

			gotFixed, err := accFixed(tt.input)
			if err != nil {
				t.Error(err)
			}
			if gotFixed != tt.wantFixed {
				t.Errorf("Want accFixed(tt.input)=%d; got %d", tt.wantFixed, gotFixed)
			}
		})
	}
}
