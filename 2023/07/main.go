package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = fmt.Printf
var printf = func(string, ...any) {}

type typ = int

const (
	H_UNKNOWN typ = iota
	H_HIGH
	H_PAIR
	H_PAIRS
	H_THREE
	H_HOUSE
	H_FOUR
	H_FIVE
)

var shapetypes = map[string]typ{
	"11111": H_HIGH,
	"22111": H_PAIR,
	"22221": H_PAIRS,
	"33311": H_THREE,
	"33322": H_HOUSE,
	"44441": H_FOUR,
	"55555": H_FIVE,
}

var cardscore map[rune]int
var jokercardscore map[rune]int

func init() {
	cardscore = make(map[rune]int)
	for s, r := range "23456789TJQKA" {
		cardscore[r] = s + 2
	}

	jokercardscore = make(map[rune]int)
	for s, r := range "J23456789TQKA" {
		jokercardscore[r] = s + 2
	}
}

func init() {
}

type hand struct {
	cards       string
	sortedCards string
	typ         typ
	jokerTyp    typ
	bet         int
}

func (h hand) less(other hand, jokers bool) bool {
	scores := cardscore
	if jokers {
		scores = jokercardscore
	}
	printf("less(%v, %v)\n", h, other)

	hTyp := h.typ
	otherTyp := other.typ
	if jokers {
		hTyp = h.jokerTyp
		otherTyp = other.jokerTyp
	}

	if hTyp > otherTyp {
		printf(" easy false based on type\n")
		return false
	}
	if hTyp < otherTyp {
		printf(" easy true based on type\n")
		return true
	}

	hRunes := []rune(h.cards)
	otherRunes := []rune(other.cards)

	for i, r := range hRunes {
		rOther := otherRunes[i]
		printf("  comparing %c to %c: ", r, rOther)
		if scores[r] == scores[rOther] {
			printf("equal; next\n")
			continue
		}
		printf("%v\n", scores[r] < scores[rOther])
		return scores[r] < scores[rOther]
	}

	return false
}

func parseHand(sandi util.StringsAndInts) (hand, error) {
	var h hand

	h.cards = sandi.Strings[0]
	h.bet = sandi.Ints[0]

	if len(h.cards) != 5 {
		return h, fmt.Errorf("weird hand: expected five cards, but got %q", h.cards)
	}
	counts := make(map[rune]int)
	jokerCounts := make(map[rune]int)
	var maxCard rune
	var maxCount int
	var jokerCount int
	for _, r := range h.cards {
		counts[r]++
		if r != 'J' {
			jokerCounts[r]++
			if jokerCounts[r] > maxCount {
				maxCount = jokerCounts[r]
				maxCard = r
			}
		} else {
			jokerCount++
		}
	}
	jokerCounts[maxCard] += jokerCount

	rCards := []rune(h.cards)
	sort.Slice(rCards, func(i, j int) bool {
		ci, cj := rCards[i], rCards[j]
		if counts[ci] != counts[cj] {
			return counts[ci] > counts[cj]
		}
		return cardscore[ci] > cardscore[cj]
	})

	rJokerCards := []rune(h.cards)
	for i, r := range rJokerCards {
		if r == 'J' {
			rJokerCards[i] = maxCard
		}
	}
	sort.Slice(rJokerCards, func(i, j int) bool {
		ci, cj := rJokerCards[i], rJokerCards[j]
		if jokerCounts[ci] != jokerCounts[cj] {
			return jokerCounts[ci] > jokerCounts[cj]
		}
		return jokercardscore[ci] > jokercardscore[cj]
	})

	sCards := string(rCards)
	h.sortedCards = sCards

	shape := strings.Join(util.Map(rCards, func(r rune) string { return strconv.Itoa(counts[r]) }), "")
	h.typ = shapetypes[shape]
	if h.typ == H_UNKNOWN {
		return h, fmt.Errorf("weird shape %q for hand %q", shape, h.cards)
	}

	jokerShape := strings.Join(util.Map(rJokerCards, func(r rune) string { return strconv.Itoa(jokerCounts[r]) }), "")
	h.jokerTyp = shapetypes[jokerShape]
	if h.jokerTyp == H_UNKNOWN {
		return h, fmt.Errorf("weird joker shape %q for hand %q. jokerCounts=%v maxCard=%c maxCount=%d jokerCount=%d jokerCards=%s", jokerShape, h.cards, jokerCounts, maxCard, maxCount, jokerCount, string(rJokerCards))
	}

	return h, nil
}

func parseInput(inputs []string) ([]hand, error) {
	sandi, err := util.ParseStringsAndInts(inputs, 2, []int{0}, []int{1})
	if err != nil {
		return nil, err
	}

	return util.MapE(sandi, parseHand)
}

func part1(inputs []string) (int, error) {
	hands, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	sort.Slice(hands, func(i, j int) bool { return hands[i].less(hands[j], false) })
	sum := 0
	for i, h := range hands {
		printf("%v: %d * %d == %d\n", h, i+1, h.bet, (i+1)*h.bet)
		sum += (i + 1) * h.bet
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	hands, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	sort.Slice(hands, func(i, j int) bool { return hands[i].less(hands[j], true) })
	sum := 0
	for i, h := range hands {
		printf("%v: %d * %d == %d\n", h, i+1, h.bet, (i+1)*h.bet)
		sum += (i + 1) * h.bet
	}
	return sum, nil
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
