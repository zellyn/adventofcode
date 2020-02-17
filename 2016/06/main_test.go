package main

import "testing"

func TestParts(t *testing.T) {
	testdata := []struct {
		filename  string
		wantMost  string
		wantLeast string
	}{
		{
			filename:  "example1",
			wantMost:  "easter",
			wantLeast: "advent",
		},
		{
			filename:  "input",
			wantMost:  "ikerpcty",
			wantLeast: "uwpfaqrq",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.filename, func(t *testing.T) {
			gotMost, gotLeast, err := decode(tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			if gotMost != tt.wantMost || gotLeast != tt.wantLeast {
				t.Errorf("Want foo(%q)=%q,%q; got %q,%q", tt.filename, tt.wantMost, tt.wantLeast, gotMost, gotLeast)
			}
		})
	}
}
