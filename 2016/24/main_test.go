package main

import "testing"

func TestParts(t *testing.T) {
	testdata := []struct {
		s    string
		want int
	}{
		{
			s:    "42",
			want: 42,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.s, func(t *testing.T) {
			got, err := foo(tt.s)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want foo(%q)=%d; got %d", tt.s, tt.want, got)
			}
		})
	}
}
