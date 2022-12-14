package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/graph"
)

type spellInfo struct {
	name     string
	cost     int
	duration int

	// Instantaneous stats
	hit  int
	heal int

	// Each turn stats
	armor  int
	damage int
	mana   int
}

var spells = [5]spellInfo{
	{
		name: "Magic Missile",
		cost: 53,
		hit:  4,
	},
	{
		name: "Drain",
		cost: 73,
		hit:  2,
		heal: 2,
	},
	{
		name:     "Shield",
		cost:     113,
		duration: 6,
		armor:    7,
	},
	{
		name:     "Poison",
		cost:     173,
		duration: 6,
		damage:   3,
	},
	{
		name:     "Recharge",
		cost:     229,
		duration: 5,
		mana:     101,
	},
}

type state struct {
	playerHealth int
	bossHealth   int
	mana         int
	bossDamage   int
	turnsLeft    [5]int
	hard         bool
}

func (s *state) End() bool {
	return s.bossHealth <= 0
}

func (s *state) Key() string {
	return fmt.Sprintf("%v", *s)
}

func (s *state) Neighbors() []graph.CostedNode {
	var result []graph.CostedNode
	for i := range spells {
		if s.turnsLeft[i] > 1 {
			continue
		}
		next := s.twoStep(i)
		if next.Steps > 0 {
			result = append(result, next)
		}
	}
	return result
}

var _ graph.Node = &state{}

func (s *state) applyEffects() {
	for i, spell := range spells {
		if s.turnsLeft[i] > 0 {
			s.turnsLeft[i]--
			s.bossHealth -= spell.damage
			s.mana += spell.mana
		}
	}
}

func (s state) activeArmor() int {
	armor := 0
	for i, spell := range spells {
		if s.turnsLeft[i] > 0 {
			armor += spell.armor
		}
	}
	return armor
}

func (s state) twoStep(spellIndex int) graph.CostedNode {
	// Player turn
	if s.hard {
		s.playerHealth--
		if s.playerHealth <= 0 {
			return graph.CostedNode{}
		}
	}

	(&s).applyEffects()

	spell := spells[spellIndex]
	s.mana -= spell.cost
	if s.mana < 0 {
		return graph.CostedNode{} // Empty result
	}
	s.playerHealth += spell.heal
	s.bossHealth -= spell.hit
	if s.bossHealth <= 0 {
		return graph.CostedNode{
			N:     &s,
			Steps: spell.cost,
		}
	}
	s.turnsLeft[spellIndex] = spell.duration

	// Boss turn
	(&s).applyEffects()
	if s.bossHealth <= 0 {
		return graph.CostedNode{
			N:     &s,
			Steps: spell.cost,
		}
	}
	damage := s.bossDamage - s.activeArmor()
	if damage < 1 {
		damage = 1
	}
	s.playerHealth -= damage
	if s.playerHealth <= 0 {
		return graph.CostedNode{}
	}
	return graph.CostedNode{
		N:     &s,
		Steps: spell.cost,
	}
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
