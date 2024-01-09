package main

import "testing"

func TestParts(t *testing.T) {
	got1, got2, err := countValid("input")
	if err != nil {
		t.Fatal(err)
	}
	want1 := 115
	want2 := 231
	if got1 != want1 || got2 != want2 {
		t.Errorf("want countValid(%q)=%d,%d; got %d,%d", "input", want1, want2, got1, got2)
	}
}

func TestTLS(t *testing.T) {
	testdata := []struct {
		s    string
		want bool
	}{
		{
			s:    "abba[mnop]qrst",
			want: true,
		},
		{
			s:    "abcd[bddb]xyyx",
			want: false,
		},
		{
			s:    "aaaa[qwer]tyui",
			want: false,
		},
		{
			s:    "ioxxoj[asdfgh]zxcvbn",
			want: true,
		},
		{
			s:    "ioxxoj[abba]zxcvbn",
			want: false,
		},
	}

	for _, tt := range testdata {
		got := tls(tt.s)
		if got != tt.want {
			t.Errorf("want tls(%q)=%v; got %v", tt.s, tt.want, got)
		}
	}
}

func TestSSL(t *testing.T) {
	testdata := []struct {
		s    string
		want bool
	}{
		{
			s:    "aba[bab]xyz",
			want: true,
		},
		{
			s:    "xyx[xyx]xyx",
			want: false,
		},
		{
			s:    "aaa[kek]eke",
			want: true,
		},
		{
			s:    "zazbz[bzb]cdb",
			want: true,
		},
	}

	for _, tt := range testdata {
		got := ssl(tt.s)
		if got != tt.want {
			t.Errorf("want ssl(%q)=%v; got %v", tt.s, tt.want, got)
		}
	}
}
