package main

import (
	"reflect"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
	Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`)

func TestStuff(t *testing.T) {
	input, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	_ = input
	testdata := []struct {
		name  string
		input []string
		want1 int
		want2 int
	}{
		{
			name:  "example",
			input: example,
			want1: 62842880,
			want2: 57600000,
		},
		{
			name:  "real input",
			input: input,
			want1: 222870,
			want2: 117936,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := best(tt.input, 0)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want1 {
				t.Errorf("want best(input, 0)=%d; got %d", tt.want1, got)
			}

			got, err = best(tt.input, 500)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want2 {
				t.Errorf("want best(input, 500)=%d; got %d", tt.want2, got)
			}
		})
	}
}

func TestWaysToAddTo(t *testing.T) {
	want := [][]int{
		{0, 0, 3},
		{0, 1, 2},
		{0, 2, 1},
		{0, 3, 0},
		{1, 0, 2},
		{1, 1, 1},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
		{3, 0, 0},
	}
	got := waysToAddTo(3, 3)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want waysToAddTo(3,3)=%v; got %v", want, got)
	}
}

func TestScore(t *testing.T) {
	parsed, err := parseInput(example)
	if err != nil {
		t.Fatal(err)
	}
	got := score([]int{44, 56}, parsed, 0)
	want := 62842880
	if got != want {
		t.Errorf("want score([]int{44,56}, parsed)=%d; got %d", want, got)
	}
}
