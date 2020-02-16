package main

import (
	"fmt"
	"os"
	"regexp"
)

type stats struct {
	name      string
	cost      int
	damage    int
	armor     int
	hitpoints int
}

var inputRE = regexp.MustCompile(`^(\w.*\S)  +([0-9]+)  +([0-9]+)  +([0-9]+)$`)

func gear() (weapons, armor, rings []stats) {
	weapons = []stats{
		{name: "Dagger", cost: 8, damage: 4, armor: 0},
		{name: "Shortsword", cost: 10, damage: 5, armor: 0},
		{name: "Warhammer", cost: 25, damage: 6, armor: 0},
		{name: "Longsword", cost: 40, damage: 7, armor: 0},
		{name: "Greataxe", cost: 74, damage: 8, armor: 0},
	}

	armor = []stats{
		{name: "Leather", cost: 13, damage: 0, armor: 1},
		{name: "Chainmail", cost: 31, damage: 0, armor: 2},
		{name: "Splintmail", cost: 53, damage: 0, armor: 3},
		{name: "Bandedmail", cost: 75, damage: 0, armor: 4},
		{name: "Platemail", cost: 102, damage: 0, armor: 5},
	}

	rings = []stats{
		{name: "Damage +1", cost: 25, damage: 1, armor: 0},
		{name: "Damage +2", cost: 50, damage: 2, armor: 0},
		{name: "Damage +3", cost: 100, damage: 3, armor: 0},
		{name: "Defense +1", cost: 20, damage: 0, armor: 1},
		{name: "Defense +2", cost: 40, damage: 0, armor: 2},
		{name: "Defense +3", cost: 80, damage: 0, armor: 3},
	}

	return weapons, armor, rings
}

func ringOptions(max int, rings []stats) []stats {
	if max == 0 || len(rings) == 0 {
		return []stats{{}}
	}
	optionsWithoutFirst := ringOptions(max, rings[1:])
	optionsWithFirst := ringOptions(max-1, rings[1:])
	result := make([]stats, len(optionsWithoutFirst), len(optionsWithFirst)+len(optionsWithoutFirst))

	copy(result, optionsWithoutFirst)
	for _, s := range optionsWithFirst {
		result = append(result, stats{
			cost:   s.cost + rings[0].cost,
			damage: s.damage + rings[0].damage,
			armor:  s.armor + rings[0].armor,
		})
	}
	return result
}

func allOptions(weapons, armor, rings []stats) []stats {
	var result []stats

	ro := ringOptions(2, rings)

	for _, weapon := range weapons {
		for _, s := range ro {
			s.cost += weapon.cost
			s.damage += weapon.damage
			result = append(result, s)
		}
		for _, armor := range armor {
			for _, s := range ro {
				s.cost += weapon.cost + armor.cost
				s.armor += armor.armor
				s.damage += weapon.damage
				result = append(result, s)
			}
		}
	}

	return result
}

func battle(player, enemy stats) int {
	playerDamage := player.damage - enemy.armor
	if playerDamage < 1 {
		playerDamage = 1
	}
	enemyDamage := enemy.damage - player.armor
	if enemyDamage < 1 {
		enemyDamage = 1
	}
	for {
		enemy.hitpoints -= playerDamage
		if enemy.hitpoints <= 0 {
			return player.hitpoints
		}
		player.hitpoints -= enemyDamage
		if player.hitpoints <= 0 {
			return -enemy.hitpoints
		}
	}
}

func lowestWin(weapons, armor, rings []stats, hp int, enemy stats) int {
	min := 999999
	for _, s := range allOptions(weapons, armor, rings) {
		s.hitpoints = hp
		if battle(s, enemy) < 0 {
			continue
		}
		if s.cost < min {
			min = s.cost
		}
	}
	return min
}

func highestLoss(weapons, armor, rings []stats, hp int, enemy stats) int {
	max := 0
	for _, s := range allOptions(weapons, armor, rings) {
		s.hitpoints = hp
		if battle(s, enemy) > 0 {
			continue
		}
		if s.cost > max {
			max = s.cost
		}
	}
	return max
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
