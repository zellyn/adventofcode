package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name      string
		input     []string
		wantCount int
		wantList  string
	}{
		{
			name:      "example",
			input:     example,
			wantCount: 5,
			wantList:  "mxmxvkd,sqjhc,fvjkl",
		},
		{
			name:      "input",
			input:     input,
			wantCount: 1913,
			wantList:  "foo",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotList, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if gotCount != tt.wantCount {
				t.Errorf("Want part1(tt.input)=%d,?; got %d,?", tt.wantCount, gotCount)
			}

			if gotList != tt.wantList {
				t.Errorf("Want part1(tt.input)=?,%q; got ?,%q", tt.wantList, gotList)
			}
		})
	}
}
