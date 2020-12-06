package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
abc

a
b
c

ab
ac

a
a
a
a

b
`)

var input = ioutil.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name    string
		input   []string
		wantAny int
		wantAll int
	}{
		{
			name:    "example",
			input:   example,
			wantAny: 11,
			wantAll: 6,
		},
		{
			name:    "input",
			input:   input,
			wantAny: 6549,
			wantAll: 42,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotAny := sumAny(tt.input)
			if gotAny != tt.wantAny {
				t.Errorf("Want sumAny(tt.input)=%d; got %d", tt.wantAny, gotAny)
			}

			gotAll := sumAll(tt.input)
			if gotAll != tt.wantAll {
				t.Errorf("Want sumAll(tt.input)=%d; got %d", tt.wantAll, gotAll)
			}
		})
	}
}
