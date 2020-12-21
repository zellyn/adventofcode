package main

import (
	"reflect"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
`)

var example2 = util.TrimmedLines(`
shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name     string
		input    []string
		wantWays int
		wantSize int
	}{
		{
			name:     "example1",
			input:    example1,
			wantWays: 4,
			wantSize: 32,
		},
		{
			name:     "example2",
			input:    example2,
			wantWays: 0,
			wantSize: 126,
		},
		{
			name:     "input",
			input:    input,
			wantWays: 177,
			wantSize: 34988,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotWays, err := waysToGold(tt.input)
			if err != nil {
				t.Error(err)
			}
			if gotWays != tt.wantWays {
				t.Errorf("Want waysToGold(tt.input)=%d; got %d", tt.wantWays, gotWays)
			}

			gotSize, err := sizeOfGold(tt.input)
			if err != nil {
				t.Error(err)
			}
			if gotSize != tt.wantSize {
				t.Errorf("Want sizeOfGold(tt.input)=%d; got %d", tt.wantSize, gotSize)
			}
		})
	}
}

func TestParseBag(t *testing.T) {
	testdata := []struct {
		input string
		want  baginfo
	}{
		{
			input: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			want: baginfo{
				name:     "light red",
				contains: []bagcount{{name: "bright white", count: 1}, {name: "muted yellow", count: 2}},
			},
		},
		{
			input: "bright white bags contain 1 shiny gold bag.",
			want: baginfo{
				name:     "bright white",
				contains: []bagcount{{name: "shiny gold", count: 1}},
			},
		},
		{
			input: "faded blue bags contain no other bags.",
			want: baginfo{
				name: "faded blue",
			},
		},
	}

	for _, tt := range testdata {
		got, err := parseBag(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("want parseBag(%q)=%v; got %v", tt.input, tt.want, got)
		}
	}
}
