package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestPart1(t *testing.T) {
	input, err := util.ReadFileString("input")
	if err != nil {
		t.Fatal(err)
	}
	expanded, err := expand(input)
	if err != nil {
		t.Fatal(err)
	}
	want1 := 99145
	got1 := len(expanded)
	if got1 != want1 {
		t.Errorf("want len(expanded)=%d; got %d", want1, got1)
	}
	want2 := 10943094568
	got2, err := expandedLength(input)
	if err != nil {
		t.Fatal(err)
	}
	if got2 != want2 {
		t.Errorf("want expandedLength(input)=%d; got %d", want2, got2)
	}
}

func TestExpand(t *testing.T) {
	testdata := []struct {
		s    string
		want string
	}{
		{
			s:    "ADVENT",
			want: "ADVENT",
		},
		{
			s:    "A(1x5)BC",
			want: "ABBBBBC",
		},
		{
			s:    "(3x3)XYZ",
			want: "XYZXYZXYZ",
		},
		{
			s:    "A(2x2)BCD(2x2)EFG",
			want: "ABCBCDEFEFG",
		},
		{
			s:    "(6x1)(1x3)A",
			want: "(1x3)A",
		},
		{
			s:    "X(8x2)(3x3)ABCY",
			want: "X(3x3)ABC(3x3)ABCY",
		},
	}

	for _, tt := range testdata {
		got, err := expand(tt.s)
		if err != nil {
			t.Errorf("expand(%q) gave error: %v", tt.s, err)
		}
		if got != tt.want {
			t.Errorf("Want expand(%q)=%q; got %q", tt.s, tt.want, got)
		}
	}
}
