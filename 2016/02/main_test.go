package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestParts(t *testing.T) {
	example1 := util.TrimmedLines(`
		ULL
		RRDDD
		LURDL
		UUUUD`)
	input, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	testdata := []struct {
		name  string
		input []string
		want1 string
		want2 string
	}{
		{
			name:  "example1",
			input: example1,
			want1: "1985",
			want2: "5DB3",
		},
		{
			name:  "input",
			input: input,
			want1: "12578",
			want2: "516DD",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, err := code(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("Want code(tt.input)=%q,%q; got %q,%q", tt.want1, tt.want2, got1, got2)
			}
		})
	}
}
