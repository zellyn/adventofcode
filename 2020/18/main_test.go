package main

import (
	"strconv"
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
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
			want:  26457,
		},
		{
			name:  "input",
			input: input,
			want:  464478013511,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  694173,
		},
		{
			name:  "input",
			input: input,
			want:  85660197232452,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestExprs(t *testing.T) {
	testdata := []struct {
		expr string
		prec bool
		want int
	}{
		{
			expr: "2 * 3 + (4 * 5)",
			want: 26,
		},
		{
			expr: "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			want: 437,
		},
		{
			expr: "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			want: 12240,
		},
		{
			expr: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			want: 13632,
		},
		{
			expr: "1 + 2 * 3 + 4 * 5 + 6",
			prec: true,
			want: 231,
		},
		{
			expr: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			prec: true,
			want: 23340,
		},
	}

	for i, tt := range testdata {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := eval(tt.expr, tt.prec)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want eval(%q)=%d; got %d", tt.expr, tt.want, got)
			}
		})
	}

}
