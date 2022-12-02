package main

import (
	"fmt"
	"os"
	"strings"
)

type rps int

const (
	Rock rps = iota
	Paper
	Scissors
)

var outcomes = map[[2]rps]int{
	{Rock, Rock}:         0,
	{Rock, Paper}:        -1,
	{Rock, Scissors}:     1,
	{Paper, Rock}:        1,
	{Paper, Paper}:       0,
	{Paper, Scissors}:    -1,
	{Scissors, Rock}:     -1,
	{Scissors, Paper}:    1,
	{Scissors, Scissors}: 0,
}

var elfMapping = map[string]rps{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var myMapping = map[string]rps{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var outcomeMapping = map[string]int{
	"X": -1,
	"Y": 0,
	"Z": 1,
}

func (r rps) against(other rps) int {
	result, found := outcomes[[2]rps{r, other}]
	if !found {
		panic("weird inputs")
	}
	return result
}

func (r rps) String() string {
	switch r {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return fmt.Sprintf("{weird rps input(%d)}", r)
	}
}

func mapit(s string, m map[string]rps) rps {
	response, found := m[s]
	if !found {
		panic(fmt.Sprintf("weird input: %q for mapping %#v", s, m))
	}

	return response
}

func (other rps) lazyChoiceFor(outcome int) rps {
	for response := Rock; response <= Scissors; response++ {
		if response.against(other) == outcome {
			return response
		}
	}
	panic(fmt.Sprintf("Can't figure out response for %v that gives %d", other, outcome))
}

func score(theirs, ours rps) int {
	outcome := ours.against(theirs)
	return 1 + int(ours) + 3 + 3*outcome
}

func scoreWithMapping(theirCode, ourCode string, theirMapping, ourMapping map[string]rps) int {
	return score(mapit(theirCode, theirMapping), mapit(ourCode, ourMapping))
}

func part1(inputs []string) (int, error) {
	total := 0
	for _, input := range inputs {
		parts := strings.Split(input, " ")
		total += scoreWithMapping(parts[0], parts[1], elfMapping, myMapping)
	}
	return total, nil
}

func part2(inputs []string) (int, error) {
	total := 0
	for _, input := range inputs {
		parts := strings.Split(input, " ")
		theirs := mapit(parts[0], elfMapping)
		desired, ok := outcomeMapping[parts[1]]
		if !ok {
			panic(fmt.Sprintf("weird outcome code %q", parts[1]))
		}
		ours := theirs.lazyChoiceFor(desired)
		total += score(theirs, ours)
	}
	return total, nil
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
