package main

import "testing"

func TestStuff(t *testing.T) {
	got, err := which()
	if err != nil {
		t.Fatal(err)
	}
	want := 40
	if got != want {
		t.Errorf("want which()=%d; got %d", want, got)
	}

	got, err = which2()
	if err != nil {
		t.Fatal(err)
	}
	want = 241
	if got != want {
		t.Errorf("want which2()=%d; got %d", want, got)
	}
}
