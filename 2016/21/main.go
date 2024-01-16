package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var regexps = []*regexp.Regexp{
	regexp.MustCompile("(swap position) ([0-9]+) with position ([0-9])+"),
	regexp.MustCompile("(swap letter) ([a-z]+) with letter ([a-z])+"),
	regexp.MustCompile("(rotate (?:left|right)) ([0-9]+) steps?"),
	regexp.MustCompile("(rotate based) on position of letter ([a-z])"),
	regexp.MustCompile("(reverse) positions ([0-9]+) through ([0-9]+)"),
	regexp.MustCompile("(move) position ([0-9]+) to position ([0-9]+)"),
}

type instruction struct {
	name    string
	ints    []int
	letters []rune
}

func (i instruction) String() string {
	letters := util.Map(i.letters, func(r rune) string {
		return string(r)
	})

	return fmt.Sprintf("{%s %v %v}", i.name, i.ints, letters)
}

func reverse(runes []rune, start, end int) {
	if start > end {
		start, end = end, start
	}
	for ; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}
}

func rotateLeft(runes []rune, amt int) {
	amt = amt % len(runes)
	if amt == 0 {
		return
	}
	reverse(runes, 0, amt-1)
	reverse(runes, amt, len(runes)-1)
	reverse(runes, 0, len(runes)-1)
}

func rotateRight(runes []rune, amt int) {
	amt = amt % len(runes)
	if amt == 0 {
		return
	}
	rotateLeft(runes, len(runes)-amt)
}

func move(runes []rune, from, to int) {
	if from < to {
		for i := from; i < to; i++ {
			runes[i], runes[i+1] = runes[i+1], runes[i]
		}
		return
	}
	for i := from - 1; i >= to; i-- {
		runes[i], runes[i+1] = runes[i+1], runes[i]
	}
}

func rotationAmount(pos int) int {
	res := pos
	if pos >= 4 {
		res++
	}
	res++
	return res
}

func unrotationAmount(length, pos int) int {
	for oldPos := 0; oldPos < length; oldPos++ {
		amt := rotationAmount(oldPos)
		if (oldPos+amt)%length == pos {
			return amt
		}
	}
	return -1
}

func rotateBasedOnLetter(runes []rune, letter rune) {
	pos := -1
	for i, r := range runes {
		if r == letter {
			pos = i
			break
		}
	}
	if pos == -1 {
		panic(fmt.Sprintf("unable to find letter '%c' in %q", letter, string(runes)))
	}

	amount := rotationAmount(pos)
	rotateRight(runes, amount)
}

func unrotateBasedOnLetter(runes []rune, letter rune) {
	pos := -1
	for i, r := range runes {
		if r == letter {
			pos = i
			break
		}
	}
	if pos == -1 {
		panic(fmt.Sprintf("unable to find letter '%c' in %q", letter, string(runes)))
	}
	amt := unrotationAmount(len(runes), pos)
	if amt == -1 {
		panic(fmt.Sprintf("unable to unrotate to position %d in string of length %d", pos, len(runes)))
	}
	rotateLeft(runes, amt)
}

func swapPosition(runes []rune, pos1, pos2 int) {
	runes[pos1], runes[pos2] = runes[pos2], runes[pos1]
}

func swapLetters(runes []rune, letter1, letter2 rune) {
	for i, r := range runes {
		if r == letter1 {
			runes[i] = letter2
		}
		if r == letter2 {
			runes[i] = letter1
		}
	}
}

func (inst instruction) perform(s string) string {
	res := []rune(s)

	switch inst.name {
	case "move":
		move(res, inst.ints[0], inst.ints[1])
	case "reverse":
		start, end := inst.ints[0], inst.ints[1]
		reverse(res, start, end)
	case "rotate based":
		rotateBasedOnLetter(res, inst.letters[0])
	case "rotate left":
		rotateLeft(res, inst.ints[0])
	case "rotate right":
		rotateRight(res, inst.ints[0])
	case "swap letter":
		swapLetters(res, inst.letters[0], inst.letters[1])
	case "swap position":
		swapPosition(res, inst.ints[0], inst.ints[1])
	default:
		panic(fmt.Sprintf("unknown instruction %q: %v", inst.name, inst))
	}

	return string(res)
}

func (inst instruction) unperform(s string) string {
	res := []rune(s)

	switch inst.name {
	case "move":
		move(res, inst.ints[1], inst.ints[0])
	case "reverse":
		start, end := inst.ints[0], inst.ints[1]
		reverse(res, start, end)
	case "rotate based":
		unrotateBasedOnLetter(res, inst.letters[0])
	case "rotate left":
		rotateRight(res, inst.ints[0])
	case "rotate right":
		rotateLeft(res, inst.ints[0])
	case "swap letter":
		swapLetters(res, inst.letters[0], inst.letters[1])
	case "swap position":
		swapPosition(res, inst.ints[0], inst.ints[1])
	default:
		panic(fmt.Sprintf("unknown instruction %q: %v", inst.name, inst))
	}

	return string(res)
}

func translate(sai util.StringsAndInts) (instruction, error) {
	var res instruction
	if len(sai.Strings) == 0 {
		return res, fmt.Errorf("unknown instruction; no strings matched in line %q", sai.Input)
	}
	res = instruction{
		name: sai.Strings[0],
		ints: sai.Ints,
	}

	for _, s := range sai.Strings[1:] {
		if len(s) != 1 {
			return res, fmt.Errorf("want only 1-char strings, but got %q in line %q", s, sai.Input)
		}
		for _, c := range s {
			res.letters = append(res.letters, c)
		}
	}

	return res, nil
}

func parse(inputs []string) ([]instruction, error) {
	sais, err := util.ParseByRegexps(inputs, regexps)
	if err != nil {
		return nil, err
	}
	return util.MapE(sais, translate)
}

func part1(inputs []string, initial string) (string, error) {
	instructions, err := parse(inputs)
	if err != nil {
		return "", err
	}

	s := initial
	for _, i := range instructions {
		s = i.perform(s)
	}
	return s, nil
}

func part2(inputs []string, initial string) (string, error) {
	instructions, err := parse(inputs)
	if err != nil {
		return "", err
	}

	slices.Reverse(instructions)

	s := initial
	for _, i := range instructions {
		s = i.unperform(s)
	}
	return s, nil
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
