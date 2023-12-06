package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type card struct {
	num     int
	winners []int
	have    []int
}

func (c card) matches() int {
	winners := make(map[int]bool)
	for _, w := range c.winners {
		winners[w] = true
	}
	sum := 0
	for _, h := range c.have {
		if winners[h] {
			sum++
		}
	}
	return sum
}

func (c card) score() int {
	sum := c.matches()
	if sum == 0 {
		return 0
	}
	// fmt.Printf("card %d: winners=%v have=%v matches=%d score=%d\n", c.num, c.winners, c.have, sum, 1<<(sum-1))
	return 1 << (sum - 1)
}

func parseCard(input string) (card, error) {
	input = strings.ReplaceAll(input, "  ", " ")
	input = strings.ReplaceAll(input, "  ", " ")
	var c card

	data := strings.Split(input, ": ")[1]
	runs := strings.Split(data, " | ")
	var err error
	c.winners, err = util.ParseInts(runs[0], " ")
	if err != nil {
		return c, err
	}
	c.have, err = util.ParseInts(runs[1], " ")
	if err != nil {
		return c, err
	}

	return c, nil
}

func parseCards(inputs []string) ([]card, error) {
	var cards []card
	for i, input := range inputs {
		card, err := parseCard(input)
		if err != nil {
			return nil, err
		}
		card.num = i + 1
		cards = append(cards, card)
	}
	return cards, nil
}

func part1(inputs []string) (int, error) {
	cards, err := parseCards(inputs)
	if err != nil {
		return 0, err
	}
	return util.MappedSum(cards, card.score), nil
}

func part2(inputs []string) (int, error) {
	cards, err := parseCards(inputs)
	if err != nil {
		return 0, err
	}

	counts := make([]int, len(cards))
	// start with one of each card
	for i := range counts {
		counts[i]++
	}

	for i, c := range cards {
		matches := c.matches()
		// fmt.Printf("Card %d says to get %d more cards\n", i+1, matches)
		for j := 0; j < matches; j++ {
			counts[i+1+j] += counts[i]
		}
	}

	return util.MappedSum(counts, func(f int) int { return f }), nil

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
