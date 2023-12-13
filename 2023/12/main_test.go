package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  21,
		},
		{
			name:  "input",
			input: input,
			want:  7407, // 7991 is too high
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
			want:  525152,
		},
		{
			name:  "input",
			input: input,
			want:  30568243604962,
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

func TestCompare(t *testing.T) {
	for _, line := range input {
		row, err := parseRow(line)
		if err != nil {
			t.Error(err)
			continue
		}
		efficient := row.combinations()
		ugly := row.uglyCombinations()

		if efficient != ugly {
			t.Errorf("Want combinations(%q) == %d; got %d", line, ugly, efficient)
		}
	}
}

func TestIndividual(t *testing.T) {
	line := ".#???.?#..#???# 2,1,1,3"
	row, err := parseRow(line)
	if err != nil {
		t.Error(err)
		return
	}
	efficient := row.combinations()
	ugly := row.uglyCombinations()

	if efficient != ugly {
		t.Errorf("Want combinations(%q) == %d; got %d", line, ugly, efficient)
	}
}
