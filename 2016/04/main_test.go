package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

func TestPart1(t *testing.T) {
	example1 := util.TrimmedLines(`
		aaaaa-bbb-z-y-x-123[abxyz]
		a-b-c-d-e-f-g-h-987[abcde]
		not-a-real-room-404[oarel]
		totally-real-room-200[decoy]`)
	input, err := ioutil.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	testdata := []struct {
		name  string
		rooms []string
		want  int
	}{
		{
			name:  "example",
			rooms: example1,
			want:  1514,
		},
		{
			name:  "input",
			rooms: input,
			want:  185371,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validSectorSum(tt.rooms)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want validSectorSum(tt.rooms)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	room, err := parseRoom("qzmt-zixmtkozy-ivhz-343[zimth]")
	if err != nil {
		t.Fatal(err)
	}
	want := "very encrypted name"
	got := room.decrypt()
	if got != want {
		t.Errorf("want room.decrypt()=%q; got %q", want, got)
	}
}

func TestPart2(t *testing.T) {
	input, err := ioutil.ReadLines("input")
	if err != nil {
		t.Fatal(err)
	}
	rooms, err := parseRooms(input)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	wantSector := 984
	for _, r := range rooms {
		if r.decrypt() == "northpole object storage" {
			found = true
			if r.sector != wantSector {
				t.Errorf("want sector for northpole object storage = %d; got %d", wantSector, r.sector)
			}
		}
	}
	if !found {
		t.Errorf("not found")
	}
}
