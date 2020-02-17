package main

import "testing"

func TestParts(t *testing.T) {
	testdata := []struct {
		doorID string
		want1  string
		want2  string
	}{
		{
			doorID: "abc",
			want1:  "18f47a30",
			want2:  "05ace8e3",
		},
		{
			doorID: "reyedfim",
			want1:  "f97c354d",
			want2:  "863dde27",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.doorID, func(t *testing.T) {
			got1 := password(tt.doorID)
			if got1 != tt.want1 {
				t.Errorf("Want password(%q)=%q; got %q", tt.doorID, tt.want1, got1)
			}

			got2 := password2(tt.doorID)
			if got2 != tt.want2 {
				t.Errorf("Want password2(%q)=%q; got %q", tt.doorID, tt.want2, got2)
			}
		})
	}
}

func TestHash(t *testing.T) {
	got := hash("abc", 3231929)[:6]
	want := "000001"
	if got != want {
		t.Errorf(`want hash("abc", 3231929)[:6]==%q; got %q`, want, got)
	}
}
