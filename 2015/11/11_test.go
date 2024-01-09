package main

import "testing"

func TestIncrement(t *testing.T) {
	testdata := []struct {
		input string
		want  string
	}{
		{
			input: "abc",
			want:  "abd",
		},
		{
			input: "abz",
			want:  "aca",
		},
		{
			input: "ahz",
			want:  "aja",
		},
	}

	for _, tt := range testdata {
		got := increment(tt.input)
		if got != tt.want {
			t.Errorf("want increment(%q)=%q; got %q", tt.input, tt.want, got)
		}
	}
}

func TestNextValid(t *testing.T) {
	testdata := []struct {
		input string
		want  string
	}{
		{
			input: "abcdefgh",
			want:  "abcdffaa",
		},
		{
			input: "ghhzzzzz",
			want:  "ghjaabcc",
		},
		{
			input: "hepxcrrq",
			want:  "hepxxyzz",
		},
		{
			input: "hepxxyzz",
			want:  "heqaabcc",
		},
		{
			input: "hxbxwxba",
			want:  "hxbxxyzz",
		},
		{
			input: "hxbxxyzz",
			want:  "hxcaabcc",
		},
	}

	for _, tt := range testdata {
		got := nextValid(tt.input)
		if got != tt.want {
			t.Errorf("want nextValid(%q)=%q; got %q", tt.input, tt.want, got)
		}
	}
}
