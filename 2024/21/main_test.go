package main

import (
	"testing"

	"github.com/zellyn/adventofcode/dgraph"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
029A
980A
179A
456A
379A
`)

var input = util.MustReadLines("input")

func TestSimple(t *testing.T) {
	s := state{
		stack: stack{
			arm{numeric: true},
		},
		want: []rune("029A"),
	}
	steps, err := dgraph.Dijkstra(s)
	if err != nil {
		t.Fatal(err)
	}

	if steps != 12 {
		t.Errorf("want steps=8; got %d", steps)
	}
}

func TestNumericMovesToMulti(t *testing.T) {
	numerics := "0123456789A"

	for _, from := range numerics {
		fromPos := posNumeric[from]
	TO_LOOP:
		for _, to := range numerics {
			moveses, toPos := numericMovesToMulti(to, fromPos)
			for _, moves := range moveses {
				a := arm{
					pos:     fromPos,
					numeric: true,
				}
				for _, m := range moves {
					var r rune
					a, r = a.simulate(m)
					if r == -1 {
						t.Errorf("moves %q from '%c' to '%c' go out of bounds", moves, from, to)
						continue TO_LOOP
					}
				}
				if a.pos != toPos {
					t.Errorf("moves %q from '%c' to '%c' end up at %c instead", moves, from, to, a.over())
				}
			}
		}
	}
}

func TestDirMovesToMulti(t *testing.T) {
	dirs := "<>^vA"

	for _, from := range dirs {
		fromPos := posDirectional[from]
	TO_LOOP:
		for _, to := range dirs {
			moveses, toPos := dirMovesToMulti(to, fromPos)
			for _, moves := range moveses {
				a := arm{
					pos:     fromPos,
					numeric: false,
				}
				for _, m := range moves {
					var r rune
					a, r = a.simulate(m)
					if r == -1 {
						t.Errorf("moves %q from '%c' to '%c' go out of bounds", moves, from, to)
						continue TO_LOOP
					}
				}
				if a.pos != toPos {
					t.Errorf("moves %q from '%c' to '%c' end up at %c instead", moves, from, to, a.over())
				}
			}
		}
	}
}

func TestDirMovesTo(t *testing.T) {
	testdata := []struct {
		from rune
		to   rune
		want string
	}{
		{
			from: 'A',
			to:   '<',
			want: "v<<",
		},
		{
			from: '<',
			to:   'A',
			want: ">>^",
		},
	}

	for _, tt := range testdata {
		startPos := posDirectional[tt.from]
		endPos := posDirectional[tt.to]
		gotMoves, gotPos := dirMovesTo(tt.to, startPos)
		if gotMoves != tt.want || gotPos != endPos {
			t.Errorf("want dirMovesTo('%c', posDirectional['%c'])=%q, %s; got %q, %s", tt.to, tt.from, tt.want, endPos, gotMoves, gotPos)
		}
	}
}

func TestNumericMovesFor(t *testing.T) {
	target := "029A"
	got := numericMovesFor(target)
	want := "<A^A>^^AvvvA"
	if got != want {
		t.Errorf("want numericMovesFor(%q)=%q; got %q", target, want, got)
	}
}

func TestMovesForCode(t *testing.T) {
	testdata := []struct {
		code  string
		depth int
		want  int
	}{
		{
			code:  "029A",
			depth: 0,
			want:  4,
		},
		{
			code:  "029A",
			depth: 1,
			want:  12,
		},
		{
			code:  "029A",
			depth: 2,
			want:  28,
		},
		{
			code:  "029A",
			depth: 3,
			want:  68,
		},
		{
			code:  "980A",
			depth: 3,
			want:  60,
		},
		{
			code:  "179A",
			depth: 3,
			want:  68,
		},
		{
			code:  "456A",
			depth: 3,
			want:  64,
		},

		{
			code:  "379A",
			depth: 0,
			want:  4,
		},
		{
			code:  "379A",
			depth: 1,
			want:  14, // ^A <<^^A >>A vvvA
		},
		{
			code:  "379A",
			depth: 2,
			want:  28,
		},
		{
			code:  "379A",
			depth: 2,
			want:  28,
			// 12            23                          11            18
			// 3             7                           9             A
			// ^A            <<^^A                       >>A           vvvA
			// <A>A          v<<AA>^AA>A                 vAA^A         v<AAA>^A
			// <A       >A   v<<A       A >^A     A >A   vA     A ^A   v<A       A A >^A
			// v<<A>>^A vA^A v<A<AA>>^A A vA<^A>A A vA^A v<A>^A A <A>A v<A<A>>^A A A vA^<A>A

			// 12            27                                   11              18
			// 3             7                                    9               A
			// ^A            ^^<<A                                >>A             vvvA
			// <A       >A   <A       A v<A         A >>^A        vA      A ^A    v<A       A A >^A
			// v<<A>>^A vA^A ........ A v<A <A >>^A . vA A ^<A >A v<A >^A A <A >A ......... . . vA ^<A >A

			//

			// 23
			// 7
			// <<^^A
			// v<<AA>^AA>A
			// v<A<AA>>^AAvA<^A>AAvA^A

			// 27
			// 7
			// ^^<<A
			// <AAv<AA>>^A
			// ........Av<A<A>>^A.vAA^<A>A

			// <v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
			// <v<A>>^A vA^A <vA<AA>>^A A vA<^A>A A vA^A <vA>^A A <A>A <v<A>A>^A A A vA<^A>A
			// <A       >A   v<<A       A >^A     A >A
		},
		{
			code:  "379A",
			depth: 3,
			want:  64,
		},
	}

	cache := make(map[cacheKey]int)

	for _, tt := range testdata {
		got := movesForCode(tt.code, tt.depth, cache)

		if got != tt.want {
			t.Errorf("want movesForCode(%q, %d)=%d; got %d", tt.code, tt.depth, tt.want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	testdata := []struct {
		name      string
		input     []string
		robotPads int
		want      int
	}{
		{
			name:  "example",
			input: example,
			want:  126384,
		},
		{
			name:  "input",
			input: input,
			want:  155252,
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
			name:  "input-long",
			input: input,
			want:  195664513288128,
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
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
