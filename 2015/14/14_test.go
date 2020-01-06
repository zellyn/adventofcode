package main

import "testing"

import "github.com/zellyn/adventofcode/util"

import "github.com/zellyn/adventofcode/ioutil"

func TestStuff(t *testing.T) {
	example := util.TrimmedLines(`
		Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
    	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`)
	input, err := ioutil.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	_ = input
	testdata := []struct {
		name  string
		input []string
		time  int
		want1 int
		want2 int
	}{
		{
			name:  "example",
			input: example,
			time:  1000,
			want1: 1120,
			want2: 689,
		},
		{
			name:  "real input",
			input: input,
			time:  2503,
			want1: 2696,
			want2: 1084,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotDist, err := dist(tt.input, tt.time)
			if err != nil {
				t.Fatal(err)
			}
			if gotDist != tt.want1 {
				t.Errorf("want dist(input, seconds)=%d; got %d", tt.want1, gotDist)
			}

			gotPoints, err := points(tt.input, tt.time)
			if err != nil {
				t.Fatal(err)
			}
			if gotPoints != tt.want2 {
				t.Errorf("want points(input, seconds)=%d; got %d", tt.want2, gotPoints)
			}
		})
	}
}
