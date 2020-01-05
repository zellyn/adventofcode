package main

import "testing"

func TestPermute(t *testing.T) {
	testdata := []struct {
		input  string
		output string
	}{
		{
			input:  "1",
			output: "11",
		},
		{
			input:  "11",
			output: "21",
		},
		{
			input:  "21",
			output: "1211",
		},
		{
			input:  "1211",
			output: "111221",
		},
		{
			input:  "111221",
			output: "312211",
		},
	}

	for _, tt := range testdata {
		got := iterate(tt.input)
		if tt.output != got {
			t.Errorf("want iterate(%q)=%q; got %q", tt.input, tt.output, got)
		}
	}
}

func TestIterated(t *testing.T) {
	testdata := []struct {
		input   string
		count   int
		want    string
		wantLen int
	}{
		{
			input:   "1",
			count:   5,
			want:    "312211",
			wantLen: 6,
		},
		{
			input:   "1321131112",
			count:   40,
			wantLen: 492982,
		},
		{
			input:   "1321131112",
			count:   50,
			wantLen: 6989950,
		},
	}

	for _, tt := range testdata {
		s := tt.input
		for i := 0; i < tt.count; i++ {
			s = iterate2(s)
		}
		if tt.want != "" && tt.want != s {
			t.Errorf("want %dxiterate(%q)=%q; got %q", tt.count, tt.input, tt.want, s)
		}
		if tt.wantLen != len(s) {
			t.Errorf("want len(%dxiterate(%q))=%d; got %d", tt.count, tt.input, tt.wantLen, len(s))
		}
	}
}
