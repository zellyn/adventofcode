package main

import (
	"testing"
)

func TestBattle(t *testing.T) {
	player := stats{
		name:      "player",
		hitpoints: 8,
		damage:    5,
		armor:     5,
	}
	enemy := stats{
		name:      "enemy",
		hitpoints: 12,
		damage:    7,
		armor:     2,
	}

	got := battle(player, enemy)
	want := 2
	if got != want {
		t.Errorf("want battle(player, enemy)=%d; got %d", want, got)
	}
}

func TestPart1(t *testing.T) {
	weapons, armor, rings := gear()
	hp := 100
	enemy := stats{
		hitpoints: 103,
		damage:    9,
		armor:     2,
	}
	got := lowestWin(weapons, armor, rings, hp, enemy)
	want := 121
	if got != want {
		t.Errorf("want lowestWin(...)=%d; got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	weapons, armor, rings := gear()
	hp := 100
	enemy := stats{
		hitpoints: 103,
		damage:    9,
		armor:     2,
	}
	got := highestLoss(weapons, armor, rings, hp, enemy)
	want := 201
	if got != want {
		t.Errorf("want highestLoss(...)=%d; got %d", want, got)
	}
}
