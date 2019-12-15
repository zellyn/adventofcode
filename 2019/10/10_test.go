package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGCD(t *testing.T) {
	testdata := []struct {
		a   int
		b   int
		gcd int
	}{
		{21, 7, 7},
		{1, 3, 1},
		{12, 9, 3},
		{12, 8, 4},
	}
	for _, tt := range testdata {
		g := gcd(tt.a, tt.b)
		if g != tt.gcd {
			t.Errorf("want gcd(%d,%d)==%d; got %d", tt.a, tt.b, tt.gcd, g)
		}
	}
}

func TestRatio(t *testing.T) {
	testdata := []struct {
		a     int
		b     int
		ratio [2]int
	}{
		{
			a:     -21,
			b:     14,
			ratio: [2]int{-3, 2},
		},
		{
			a:     15,
			b:     -5,
			ratio: [2]int{3, -1},
		},
		{
			a:     -2,
			b:     -2,
			ratio: [2]int{-1, -1},
		},
		{
			a:     -10,
			b:     0,
			ratio: [2]int{-1, 0},
		},
		{
			a:     0,
			b:     3,
			ratio: [2]int{0, 1},
		},
		{
			a:     -1,
			b:     -2,
			ratio: [2]int{-1, -2},
		},
	}

	for _, tt := range testdata {
		r := ratio(tt.a, tt.b)
		if r != tt.ratio {
			t.Errorf("want ratio(%d,%d)==%d; got %d", tt.a, tt.b, tt.ratio, r)
		}
	}
}

var map1 = `.#..#
.....
#####
....#
...##`

var map2 = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

var map3 = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

var map4 = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

var map5 = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

var mapTestData = []struct {
	name    string
	data    string
	best    [2]int
	visible int
}{
	{
		name:    "map1",
		data:    map1,
		best:    [2]int{3, 4},
		visible: 8,
	},
	{
		name:    "map2",
		data:    map2,
		best:    [2]int{5, 8},
		visible: 33,
	},
	{
		name:    "map3",
		data:    map3,
		best:    [2]int{1, 2},
		visible: 35,
	},
	{
		name:    "map4",
		data:    map4,
		best:    [2]int{6, 3},
		visible: 41,
	},
	{
		name:    "map5",
		data:    map5,
		best:    [2]int{11, 13},
		visible: 210,
	},
}

func TestParseMap(t *testing.T) {
	testdata := []struct {
		mapData string
		asts    [][2]int
		space   [][2]int
	}{
		{
			mapData: map1,
			asts:    [][2]int{{3, 4}, {1, 0}, {4, 0}},
			space:   [][2]int{{0, 0}, {1, 1}, {0, 1}},
		},
		{
			mapData: map5,
			asts:    [][2]int{{11, 13}, {1, 0}, {0, 1}},
			space:   [][2]int{{0, 0}, {12, 14}, {9, 12}},
		},
	}

	for i, tt := range testdata {
		t.Run(fmt.Sprintf("map%d", i+1), func(t *testing.T) {
			p := parseMap(tt.mapData)
			for _, a := range tt.asts {
				if !p[a] {
					t.Errorf("want asteroid at %v; found none", a)
				}
			}
			for _, s := range tt.space {
				if p[s] {
					t.Errorf("want space at %v; found asteroid", s)
				}
			}
		})
	}
}

func TestVisible(t *testing.T) {
	for _, tt := range mapTestData {
		t.Run(tt.name, func(t *testing.T) {
			m := parseMap(tt.data)
			v := visible(m, tt.best)
			if v != tt.visible {
				t.Errorf("want visible(%v)==%d; got %d", tt.best, tt.visible, v)
			}
		})
	}
}

func TestBest(t *testing.T) {
	for _, tt := range mapTestData {
		t.Run(tt.name, func(t *testing.T) {
			m := parseMap(tt.data)
			from, v := best(m)
			if from != tt.best {
				t.Errorf("want best(m).from==%v; got %v", tt.best, from)
			}
			if v != tt.visible {
				t.Errorf("want best(m).visible==%d; got %d", tt.visible, v)
			}
		})
	}
}

func TestStacks(t *testing.T) {
	got := stacks(parseMap(map1), [2]int{3, 4})
	want := map[[2]int][]info{
		[2]int{0, -1}:  []info{{[2]int{3, 2}, 4}},
		[2]int{1, -4}:  []info{{[2]int{4, 0}, 17}},
		[2]int{1, -2}:  []info{{[2]int{4, 2}, 5}},
		[2]int{1, -1}:  []info{{[2]int{4, 3}, 2}},
		[2]int{1, 0}:   []info{{[2]int{4, 4}, 1}},
		[2]int{-3, -2}: []info{{[2]int{0, 2}, 13}},
		[2]int{-1, -1}: []info{{[2]int{1, 2}, 8}},
		[2]int{-1, -2}: []info{{[2]int{2, 2}, 5}, {[2]int{1, 0}, 20}},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want stacks(map1, %v)==\n%v;got\n%v", [2]int{3, 4}, want, got)
	}
}

func TestByAngle(t *testing.T) {
	ratios := [][2]int{
		{1, -4},
		{1, 0},
		{0, -1},
		{1, -2},
		{1, -1},
		{-1, -1},
		{-3, -2},
		{-1, -2},
	}
	want := [][2]int{
		{0, -1},
		{1, -4},
		{1, -2},
		{1, -1},
		{1, 0},
		{-3, -2},
		{-1, -1},
		{-1, -2},
	}
	got := make([][2]int, len(ratios))
	copy(got, ratios)
	sort.Sort(byAngle(got))
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want sorted ratios == %v; got %v", want, got)
	}
}

func TestSequence(t *testing.T) {
	got := sequence(parseMap(map1), [2]int{3, 4})
	want := [][2]int{{3, 2}, {4, 0}, {4, 2}, {4, 3}, {4, 4}, {0, 2}, {1, 2}, {2, 2}, {1, 0}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want sequence(map1, %v)==\n%v;got\n%v", [2]int{3, 4}, want, got)
	}
}
