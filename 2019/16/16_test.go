package main

import (
	"fmt"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var fileInput = util.MustReadFileString("input")

var testdata = []struct {
	input  string
	reps   int
	result string
}{
	{
		input:  "12345678",
		reps:   1,
		result: "48226158",
	},
	{
		input:  "12345678",
		reps:   4,
		result: "01029498",
	},
	{
		input:  "80871224585914546619083218645595",
		reps:   100,
		result: "24176176",
	},
	{
		input:  "80871224585914546619083218645595",
		reps:   1,
		result: "24706861",
	},
	{
		input:  "19617804207202209144916044189917",
		reps:   100,
		result: "73745418",
	},
	{
		input:  "69317163492948606335995924319873",
		reps:   100,
		result: "52432133",
	},
	{
		input:  fileInput,
		reps:   100,
		result: "40580215",
	},
}

func TestReps(t *testing.T) {
	for i, tt := range testdata {
		t.Run(fmt.Sprintf("%d:%s", i, tt.input[:8]), func(t *testing.T) {
			ints := split(tt.input)
			for i := 0; i < tt.reps; i++ {
				ints = phase(ints)
			}
			res := join(ints)
			got := res[:8]
			if got != tt.result {
				t.Errorf("Want %d phases of %q to be %s; got %s", tt.reps, tt.input[:8], tt.result, got)
			}
		})
	}
}

func TestSubRep(t *testing.T) {
	var testdata = []struct {
		input  string
		offset int
		result string
	}{
		{
			input:  "12345678",
			offset: 0,
			result: "48226158",
		},
		{
			input:  "12345678",
			offset: 1,
			result: "8226158",
		},
		{
			input:  "12345678",
			offset: 4,
			result: "6158",
		},
		{
			input:  "80871224585914546619083218645595",
			offset: 24,
			result: "32484945",
		},
		{
			input:  fileInput + fileInput + fileInput + fileInput + fileInput + fileInput + fileInput + fileInput + fileInput + fileInput,
			offset: 2000,
			result: "81832980",
		},
	}

	// Initial sanity check.
	ints := splitRep("80871224585914546619083218645595", 8, 24)
	want := "18645595"
	if join(ints) != want {
		t.Errorf("want %s; got %s", join(ints), want)
	}

	for i, tt := range testdata {
		t.Run(fmt.Sprintf("%d:%s", i, tt.input[:8]), func(t *testing.T) {
			ints := splitRep(tt.input, len(tt.input)-tt.offset, tt.offset)
			got := subRep(ints, tt.offset)
			if len(got) > 8 {
				got = got[:8]
			}
			if join(got) != tt.result {
				t.Errorf("want subRep(%v, %d) == %s; got %s", ints[:8], tt.offset, tt.result, join(got))
			}
		})
	}
}
