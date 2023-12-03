package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type draw struct {
	r, g, b int
}

type game struct {
	id    int
	draws []draw
}

// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
func parseGame(s string) (game, error) {
	gm := game{}
	parts := strings.Split(s, ": ")
	id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])
	if err != nil {
		return gm, err
	}
	gm.id = id

	for _, drawString := range strings.Split(parts[1], "; ") {
		draw, err := parseDraw(drawString)
		if err != nil {
			return gm, err
		}
		gm.draws = append(gm.draws, draw)
	}
	return gm, nil
}

func parseDraw(s string) (draw, error) {
	var d draw

	parts := strings.Split(s, ", ")
	for _, part := range parts {
		countAndColor := strings.Split(part, " ")
		count, err := strconv.Atoi(countAndColor[0])
		if err != nil {
			return d, err
		}
		switch countAndColor[1] {
		case "red":
			d.r = count
		case "green":
			d.g = count
		case "blue":
			d.b = count
		default:
			return d, fmt.Errorf("unknown color: %s", countAndColor[1])
		}
	}

	return d, nil
}

func (gm game) possible(r, g, b int) bool {
	for _, d := range gm.draws {
		if d.r > r || d.g > g || d.b > b {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (gm game) power() int {
	var minr, ming, minb int
	for _, d := range gm.draws {
		minr = max(minr, d.r)
		ming = max(ming, d.g)
		minb = max(minb, d.b)
	}

	return minr * ming * minb
}

func part1(inputs []string, r, g, b int) (int, error) {
	sum := 0
	for _, input := range inputs {
		gm, err := parseGame(input)
		if err != nil {
			return 0, err
		}
		if gm.possible(r, g, b) {
			sum += gm.id
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		gm, err := parseGame(input)
		if err != nil {
			return 0, err
		}
		sum += gm.power()
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
