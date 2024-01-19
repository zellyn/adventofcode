package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`)

var example2 = util.TrimmedLines(`
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`)

var example3 = util.TrimmedLines(`
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  10,
		},
		{
			name:  "example2",
			input: example2,
			want:  19,
		},
		{
			name:  "example3",
			input: example3,
			want:  226,
		},
		{
			name:  "input",
			input: input,
			want:  3713,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
				return
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
			name:  "example1",
			input: example1,
			want:  36,
		},
		{
			name:  "example2",
			input: example2,
			want:  103,
		},
		{
			name:  "example3",
			input: example3,
			want:  3509,
		},
		{
			name:  "input",
			input: input,
			want:  91292,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
