package main

import (
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = strings.Trim(`
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`, "\n")

var input = util.MustReadFileString("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  6032,
		},
		{
			name:  "input",
			input: input,
			want:  64256,
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
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  5031,
		},
		{
			name:  "input",
			input: input,
			want:  109224,
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

/*
func TestFindFace(t *testing.T) {
	map1 := strings.Trim(`
  A
BCD
  EF
`, "\n")

	map2 := strings.Trim(`
  AB
  C
 DE
 F
`, "\n")

	map3 := strings.Trim(`
 AB
 C
 D
EF
`, "\n")

	map4 := strings.Trim(`
 A
BCD
 E
 F
`, "\n")

	_, _, _ = map1, map2, map3

	testdata := []struct {
		input string
		want  []string
	}{
		{
			input: map1,
			want: []string{
				"AFDC",
				"BCEFA",
				"CDEBA",
				"DFECA",
				"EFBCD",
				"FABED",
			},
		},
		{
			input: map2,
			want: []string{
				"ABCDD",
				"BECAF",
				"CBEDA",
				"DEFAC",
				"EBFDC",
				"FEBAD",
			},
		},
		{
			input: map3,
			want: []string{
				"ABCEF",
				"BDCAF",
				"CBDEA",
				"DBFEC",
				"EFACD",
				"FBAED",
			},
		},
		{
			input: map4,
			want: []string{
				"ADCBF",
				"BCEFA",
				"CDEBA",
				"DFECA",
				"EDFBC",
				"FDABE",
			},
		},
	}

	for i, tt := range testdata {
		if i != 0 {
			continue
		}
		t.Run(fmt.Sprintf("map%d", i+1), func(t *testing.T) {
			m := charmap.ParseWithBackground(strings.Split(tt.input, "\n"), ' ')
			for _, s := range tt.want {
				startSide := rune(s[0])
				wantSides := s[1:]
				pos, ok := m.Find(startSide)
				if !ok {
					t.Errorf("couldn't find side '%c' in map", startSide)
				}
				for j, want := range wantSides {
					fmt.Printf("%d: going %c from %c\n", i+1, compassChars[j], startSide)
					newPos, _ := findFace(m, pos, j, 1)
					if got := m[newPos]; got != want {
						fmt.Printf("  we want %c, but we got '%c'", want, got)
						t.Fatalf("going %c from %c, we want %c, but we got '%c'", compassChars[j], startSide, want, got)
					}
				}
			}
		})
	}
}
*/
