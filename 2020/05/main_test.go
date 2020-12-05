package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL
`)

var input = ioutil.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  820,
		},
		{
			name:  "input",
			input: input,
			want:  871,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got := maxParse(tt.input)
			if got != tt.want {
				t.Errorf("Want maxParse(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
func TestPart2(t *testing.T) {
	got := missing(input)
	want := 640
	if got != want {
		t.Errorf("want missing(input)=%d; got %d", want, got)
	}
}

func TestParse(t *testing.T) {
	testdata := []struct {
		pass string
		want int
	}{
		{
			pass: "FBFBBFFRLR",
			want: 357,
		},
		{
			pass: "BFFFBBFRRR",
			want: 567,
		},
		{
			pass: "FFFBBBFRRR",
			want: 119,
		},
		{
			pass: "BBFFBBFRLL",
			want: 820,
		},
	}

	for _, tt := range testdata {
		got := parse(tt.pass)
		if got != tt.want {
			t.Errorf("want parse(%q)=%d; got %d", tt.pass, tt.want, got)
		}
	}
}
