package main

import (
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/util"
)

func TestPart1(t *testing.T) {
	m := charmap.New(50, 6, '.')
	lines, err := util.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	err = runCommands(m, lines)
	if err != nil {
		t.Fatal(err)
	}
	wantCount := 123
	gotCount := m.Count('#')
	if wantCount != gotCount {
		t.Errorf("want count=%d; got %d", wantCount, gotCount)
	}

	// ZFHFSFOGPO
	wantStr := strings.Join(util.TrimmedLines(`
	.##..####.###..#..#.###..####.###....##.###...###.
	#..#.#....#..#.#..#.#..#....#.#..#....#.#..#.#....
	#..#.###..###..#..#.#..#...#..###.....#.#..#.#....
	####.#....#..#.#..#.###...#...#..#....#.###...##..
	#..#.#....#..#.#..#.#....#....#..#.#..#.#.......#.
	#..#.#....###...##..#....####.###...##..#....###..`), "\n")
	if got := charmap.String(m, '?'); got != wantStr {
		t.Errorf("final value: want \n%sgot \n%s", wantStr, got)
	}
}

func TestCommands(t *testing.T) {
	want := strings.Join(util.TrimmedLines(`
		.#..#.#
		#.#....
		.#.....`), "\n")

	commands := util.TrimmedLines(`
		rect 3x2
		rotate column x=1 by 1
		rotate row y=0 by 4
		rotate column x=1 by 1`)
	m := charmap.New(7, 3, '.')
	err := runCommands(m, commands)
	if err != nil {
		t.Fatal(err)
	}
	if got := charmap.String(m, '?'); got != want {
		t.Errorf("want: want \n%sgot \n%s", want, got)
	}
}

func TestOps(t *testing.T) {
	m := charmap.New(7, 3, '.')
	zero := strings.Join(util.TrimmedLines(`
		.......
		.......
		.......`), "\n")
	if got := charmap.String(m, '?'); got != zero {
		t.Errorf("zero: want \n%sgot \n%s", zero, got)
	}
	one := strings.Join(util.TrimmedLines(`
		###....
		###....
		.......`), "\n")
	rect(m, 3, 2)
	if got := charmap.String(m, '?'); got != one {
		t.Errorf("one: want \n%sgot \n%s", one, got)
	}

	two := strings.Join(util.TrimmedLines(`
		#.#....
		###....
		.#.....`), "\n")
	rotateCol(m, 1, 1)
	if got := charmap.String(m, '?'); got != two {
		t.Errorf("two: want \n%sgot \n%s", two, got)
	}

	three := strings.Join(util.TrimmedLines(`
		....#.#
		###....
		.#.....`), "\n")
	rotateRow(m, 0, 4)
	if got := charmap.String(m, '?'); got != three {
		t.Errorf("three: want \n%sgot \n%s", three, got)
	}

	four := strings.Join(util.TrimmedLines(`
		.#..#.#
		#.#....
		.#.....`), "\n")
	rotateCol(m, 1, 1)
	if got := charmap.String(m, '?'); got != four {
		t.Errorf("four: want \n%sgot \n%s", four, got)
	}
}
