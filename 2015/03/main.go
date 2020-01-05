package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/ioutil"
)

type vec2 = geom.Vec2

func run() error {
	s, err := ioutil.ReadFileString("input")
	if err != nil {
		return err
	}

	m := map[vec2]int{}
	pos := vec2{X: 0, Y: 0}
	m[pos]++
	for _, r := range s {
		switch r {
		case '^':
			pos.Y--
		case 'v':
			pos.Y++
		case '<':
			pos.X--
		case '>':
			pos.X++
		}
		m[pos]++
	}
	fmt.Println("Houses visited:", len(m))

	var pos2 [2]vec2
	m = map[vec2]int{}
	m[pos2[0]] = 2
	i := 0

	for _, r := range s {
		switch r {
		case '^':
			pos2[i].Y--
		case 'v':
			pos2[i].Y++
		case '<':
			pos2[i].X--
		case '>':
			pos2[i].X++
		}
		m[pos2[i]]++
		i = 1 - i
	}
	fmt.Println("Houses visited:", len(m))

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
