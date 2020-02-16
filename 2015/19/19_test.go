package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPart1(t *testing.T) {
	testdata := []struct {
		name string
		want int
	}{
		{
			name: "example1",
			want: 4,
		},
		{
			name: "example2",
			want: 7,
		},
		{
			name: "input",
			want: 576,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := distinctMolecules(tt.name)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("want distinctMolecules(input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	repls, target, err := parseInput("input")
	if err != nil {
		t.Fatal(err)
	}
	got, err := stepsToReduce(repls, target)
	if err != nil {
		t.Fatal(err)
	}
	want := 207
	if got != want {
		t.Errorf("want stepsToReduce(repls, target)=%d; got %d", want, got)
	}
}

func XTestPart2Floundering(t *testing.T) {
	testdata := []struct {
		name   string
		target string
		start  string
		want   int
		debug  bool
	}{
		{
			name:  "example3",
			start: "e",
			want:  3,
		},
		{
			name:  "example4",
			start: "e",
			want:  6,
		},
		/*
			{
				name: "input",
				start: "e",
				want: 42,
			},
		*/
		{
			name:   "input",
			target: "TiRnFArCaPRnFArSiRnCaCaCaSiThCaRnCaFArYCaSiRnFArBCaCaCaSiThFArPBF",
			start:  "Mg",
			want:   25,
		},
		{
			name:   "input",
			target: "SiRnMgArFYCaSiRnMgArCaCaSiThPRnFArPBCaSiRnMgArCaCaSiThCaSiRnTiMgArF",
			start:  "FYF",
			want:   24,
		},
		/*
			{
				name:   "input",
				target: "ORnPBPMgArCaCaCaSiThCaCaSiThCaCaPBSiRnFArRnFArCaCaSiThCaCaSiThCaCaCaCaCaCaSiRnFYFArSiRnMgArCaSiRnPTiTiBFYPBFArSiRnCaSiRnTiRnFArSiAlArPTiBPTiRnCaSiAlArCaPTiTiBPMgYFArPTiRnFArSiRnCaCaFArRnCaFArCaSiRnFYFArSiThSiThCaCaSiRnMgArCaCaSiRnFArTiBPTiRnCaSiAlArCaPTiRnFArPBPBCaCaSiThCaPBSiThPRnFArSiThCaSiThCaSiThCaPTiBSiRnFYFArCaCaPRnFArPBCaCaPBSiRnMgArCaSiRnFArRnCaCaCaFArSiRnFArTiRnPMgArF",
				start:  "e",
				want:   42,
				debug:  true,
			},
		*/
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			repls, target, err := parseInput(tt.name)
			if err != nil {
				t.Fatal(err)
			}
			repls = inverseReplacements(repls, target)
			if tt.target != "" {
				target = tt.target
			}
			/*
				got, err := fewestSteps(repls, target, tt.start, tt.debug)
				if err != nil {
					t.Fatal(err)
				}
				if got != tt.want {
					t.Errorf("want fewestSteps(input)=%d; got %d", tt.want, got)
				}
			*/

			got2, err := fewestInStages(repls, target, tt.start, tt.debug)
			if err != nil {
				t.Fatalf("fewestInStages error: %v", err)
			}
			if got2 != tt.want {
				t.Errorf("want fewestInStages(input)=%d; got %d", tt.want, got2)
			}
		})
	}
}

func TestPositions(t *testing.T) {
	testdata := []struct {
		s      string
		substr string
		want   []int
	}{
		{
			s:      "aabaaca",
			substr: "a",
			want:   []int{0, 1, 3, 4, 6},
		},
		{
			s:      "AbAbAcdAb",
			substr: "Ab",
			want:   []int{0, 2, 7},
		},
	}

	for _, tt := range testdata {
		got := positions(tt.s, tt.substr)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("want positions(%q, %q)=%v; got %v", tt.s, tt.substr, tt.want, got)
		}
	}
}

func TestReplaceAt(t *testing.T) {
	testdata := []struct {
		s      string
		substr string
		repl   string
		pos    int
		want   string
	}{
		{
			s:      "aabaaca",
			substr: "aa",
			repl:   "XX",
			pos:    0,
			want:   "XXbaaca",
		},
		{
			s:      "aabaaca",
			substr: "ca",
			repl:   "XX",
			pos:    5,
			want:   "aabaaXX",
		},
	}

	for _, tt := range testdata {
		got := replaceAt(tt.s, tt.substr, tt.repl, tt.pos)
		if got != tt.want {
			t.Errorf("want replaceAt(%q, %q, %q, %d)=%q; got %q", tt.s, tt.substr, tt.repl, tt.pos, tt.want, got)
		}
	}
}

func TestOuterRnArs(t *testing.T) {
	testdata := []struct {
		s    string
		want [][2]int
	}{
		{
			s:    "ArBRnCArDRn",
			want: [][2]int{{0, 3}, {6, 9}},
		},
		{
			s:    "XArYArZRnURnV",
			want: [][2]int{{1, 10}},
		},
	}

	for _, tt := range testdata {
		got := outerRnArs(tt.s)
		if !cmp.Equal(tt.want, got) {
			t.Errorf("want outerRnArs(%q)=%v; diff: %s", tt.s, tt.want, cmp.Diff(tt.want, got))
		}
	}
}

func TestHelpers(t *testing.T) {
	repls, target, err := parseInput("input")
	if err != nil {
		t.Fatal(err)
	}

	gotRightOnly := rightOnly(repls)
	wantRightOnly := map[string]bool{
		"C":  true,
		"Ar": true,
		"Rn": true,
		"Y":  true,
	}
	if !cmp.Equal(wantRightOnly, gotRightOnly) {
		t.Errorf("rightOnly(input) diff: %v", cmp.Diff(wantRightOnly, gotRightOnly))
	}

	unused := setDiff(gotRightOnly, atomSet(target))
	wantUnused := map[string]bool{
		"C": true,
	}
	if !cmp.Equal(wantUnused, unused) {
		t.Errorf("unused diff: %v", cmp.Diff(wantUnused, unused))
	}

	inverse := inverseReplacements(repls, target)
	gotInside := insideRnAr(inverse)
	wantInside := map[string]bool{
		"F":   true,
		"FYF": true,
		"Mg":  true,
	}
	if !cmp.Equal(wantInside, gotInside) {
		t.Errorf("inside diff: %v", cmp.Diff(wantInside, gotInside))
	}
}
