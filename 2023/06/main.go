package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type race struct {
	time     int
	distance int
}

func (r race) waysToWin() int {
	// c == (0,time)
	// distance = c * (t-c)
	// ct - c² >= d+1
	// ct - c² - (d+1) >= 0
	// c² -tc +(d+1) <= 0
	// c² -tc +(d+1) == 0
	// b² - 4ac == (-t)²-4d-4

	t := r.time
	d := r.distance

	// fmt.Printf("waysToWin: t=%d d=%d\n", t, d)

	discriminant := t*t - 4*d
	if discriminant < 0 {
		return 0
	}

	plusMinus := math.Sqrt(float64(discriminant))

	epsilon := 1e-6
	minZero := int(math.Ceil((float64(t)-plusMinus)/2 + epsilon))
	maxZero := int(math.Floor((float64(t)+plusMinus)/2 - epsilon))

	// fmt.Printf("      minZero=%d maxZero=%d ways=%d\n", minZero, maxZero, maxZero-minZero+1)

	return maxZero - minZero + 1
}

/*
   t=30
   d=200

   b² - 4ac == (-t)²-4d-4 == 30²-4(200)-4 == 96
   (-b ± √96) / 2
   30 ± √96) / 2
   == (/ (+ 30 (sqrt 96)) 2) 19.898979485566358
   == (/ (- 30 (sqrt 96)) 2) 10.101020514433644
*/

func parseLine(input string) ([]int, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("weird input line: %q", input)
	}

	return util.ParseFieldInts(parts[1])
}

func parseInput(inputs []string) ([]race, error) {
	var result []race
	times, err := parseLine(inputs[0])
	if err != nil {
		return nil, err
	}
	distances, err := parseLine(inputs[1])
	if err != nil {
		return nil, err
	}
	for i, time := range times {
		result = append(result, race{time: time, distance: distances[i]})
	}
	return result, nil
}

func part1(inputs []string) (int, error) {
	races, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	prod := 1
	for _, race := range races {
		waysToWin := race.waysToWin()
		// fmt.Printf("Race=%#v waysToWin=%d\n", race, waysToWin)
		prod *= waysToWin
	}
	return prod, nil
}

func part2(inputs []string) (int, error) {
	newInputs := []string{
		strings.ReplaceAll(inputs[0], " ", ""),
		strings.ReplaceAll(inputs[1], " ", ""),
	}

	return part1(newInputs)
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
