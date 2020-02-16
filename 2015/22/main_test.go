package main

import (
	"testing"

	"github.com/zellyn/adventofcode/graph"
)

func TestParts(t *testing.T) {
	testdata := []struct {
		name   string
		hp     int
		mana   int
		bossHp int
		damage int
		hard   bool
		want   int
	}{
		{
			name:   "example 1",
			hp:     10,
			mana:   250,
			bossHp: 13,
			damage: 8,
			want:   226,
		},
		{
			name:   "example 2",
			hp:     10,
			mana:   250,
			bossHp: 14,
			damage: 8,
			want:   641,
		},
		{
			name:   "part 1",
			hp:     50,
			mana:   500,
			bossHp: 51,
			damage: 9,
			want:   900,
		},
		{
			name:   "part 2",
			hp:     50,
			mana:   500,
			bossHp: 51,
			damage: 9,
			hard:   true,
			want:   1216,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			s := &state{
				playerHealth: tt.hp,
				bossHealth:   tt.bossHp,
				mana:         tt.mana,
				bossDamage:   tt.damage,
				hard:         tt.hard,
			}
			got, err := graph.Dijkstra(s)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want graph.Dijkstra(s)=%d; got %d", tt.want, got)
			}
		})
	}
}
