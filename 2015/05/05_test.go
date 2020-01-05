package main

import "testing"

func TestNice2(t *testing.T) {
	testdata := []struct {
		s    string
		nice bool
	}{
		{
			s:    "qjhvhtzxzqqjkmpb",
			nice: true,
		},
		{
			s:    "xxyxx",
			nice: true,
		},
		{
			s:    "uurcxstgmygtbstg",
			nice: false,
		},
		{
			s:    "ieodomkazucvgmuy",
			nice: false,
		},
	}

	for _, tt := range testdata {
		got := isNice2(tt.s)
		if got != tt.nice {
			t.Errorf("Want isNice2(%q)=%v; got %v", tt.s, tt.nice, got)
		}
	}
}
