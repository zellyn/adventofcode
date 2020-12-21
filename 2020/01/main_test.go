package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.MustStringsToInts(util.TrimmedLines(`
	1721
	979
	366
	299
	675
	1456`))

var input = util.MustReadFileInts("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name string
		i    []int
		want int
	}{
		{
			name: "example",
			i:    example,
			want: 514579,
		},
		{
			name: "input",
			i:    input,
			want: 842016,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := twoAddTo(2020, tt.i)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want addTo(2020, %v...)=%d; got %d", tt.i[:5], tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name string
		i    []int
		want int
	}{
		{
			name: "example",
			i:    example,
			want: 241861950,
		},
		{
			name: "input",
			i:    input,
			want: 9199664,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := threeAddTo(2020, tt.i)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want threeTo(2020, %v...)=%d; got %d", tt.i[:5], tt.want, got)
			}
		})
	}
}
