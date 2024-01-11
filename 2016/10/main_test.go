package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2
`)

var input = util.MustReadLines("input")

func TestParts(t *testing.T) {
	testdata := []struct {
		name           string
		input          []string
		value1, value2 int
		wantIndex      int
		wantProduct    int
	}{
		{
			name:        "example",
			input:       example,
			value1:      2,
			value2:      5,
			wantIndex:   2,
			wantProduct: 30,
		},
		{
			name:        "input",
			input:       input,
			value1:      61,
			value2:      17,
			wantIndex:   98,
			wantProduct: 4042,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotProduct, err := parts(tt.input, tt.value1, tt.value2)
			if err != nil {
				t.Error(err)
				return
			}

			if gotIndex != tt.wantIndex || gotProduct != tt.wantProduct {
				t.Errorf("Want parts(tt.input)=%d, %d; got %d, %d", tt.wantIndex, tt.wantProduct, gotIndex, gotProduct)
			}
		})
	}
}
