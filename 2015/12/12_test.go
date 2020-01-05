package main

import "testing"

import "github.com/zellyn/adventofcode/ioutil"

func TestSum(t *testing.T) {
	testdata := []struct {
		input string
		want  int
	}{
		{
			input: `{"a":2,"b":4}`,
			want:  6,
		},
		{
			input: `[1,2,3]`,
			want:  6,
		},
		{
			input: `[[[3]]]`,
			want:  3,
		},
		{
			input: `{"a":{"b":4},"c":-1}`,
			want:  3,
		},
		{
			input: `{"a":[-1,1]}`,
			want:  0,
		},
		{
			input: `[-1,{"a":1}]`,
			want:  0,
		},
		{
			input: `[]`,
			want:  0,
		},
		{
			input: `{}`,
			want:  0,
		},
		{
			input: ioutil.MustReadFileString("input"),
			want:  119433,
		},
	}

	for _, tt := range testdata {
		got, err := sum(tt.input)
		if err != nil {
			t.Error(err)
		} else {
			if got != tt.want {
				t.Errorf("Want sum(%q)=%d; got %d", tt.input, tt.want, got)
			}
		}
	}
}

func TestSumRed(t *testing.T) {
	testdata := []struct {
		input string
		want  int
	}{
		{
			input: `[1,2,3]`,
			want:  6,
		},
		{
			input: `[1,{"c":"red","b":2},3]`,
			want:  4,
		},
		{
			input: `{"d":"red","e":[1,2,3,4],"f":5}`,
			want:  0,
		},
		{
			input: `[1,"red",5]`,
			want:  6,
		},
		{
			input: ioutil.MustReadFileString("input"),
			want:  68466,
		},
	}

	for _, tt := range testdata {
		got, err := sumRed(tt.input)
		if err != nil {
			t.Error(err)
		} else {
			if got != tt.want {
				s := tt.input
				if len(s) > 50 {
					s = s[:50] + "..."
				}
				t.Errorf("Want sum(`%s`)=%d; got %d", s, tt.want, got)
			}
		}
	}
}
