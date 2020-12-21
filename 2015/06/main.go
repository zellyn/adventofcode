package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

type vec2 = geom.Vec2

func parse(s string) (vec2, error) {
	ints, err := util.ParseInts(s, ",")
	if err != nil {
		return vec2{}, err
	}
	return vec2{X: ints[0], Y: ints[1]}, nil
}

func points(from, to vec2) []vec2 {
	var ps []vec2
	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			ps = append(ps, vec2{X: x, Y: y})
		}
	}
	return ps
}

func run() error {
	lines, err := util.ReadLines("input")
	if err != nil {
		return err
	}

	m := map[vec2]bool{}
	m2 := map[vec2]int{}

	fmt.Println(points(vec2{0, 0}, vec2{3, 3}))
	for _, line := range lines {
		parts := strings.Split(line, " ")
		from, err := parse(parts[len(parts)-3])
		if err != nil {
			return err
		}
		to, err := parse(parts[len(parts)-1])
		if err != nil {
			return err
		}
		pts := points(from, to)
		if parts[0] == "toggle" {
			for _, pt := range pts {
				m[pt] = !m[pt]
				m2[pt] += 2
			}
		} else if parts[1] == "on" {
			for _, pt := range pts {
				m[pt] = true
				m2[pt]++
			}
		} else if parts[1] == "off" {
			for _, pt := range pts {
				m[pt] = false
				m2[pt]--
				if m2[pt] < 0 {
					m2[pt] = 0
				}
			}
		} else {
			panic("weird input")
		}

	}

	count := 0
	for _, v := range m {
		if v {
			count++
		}
	}
	fmt.Println("Lights on:", count)

	brightness := 0
	for _, v := range m2 {
		brightness += v
	}
	fmt.Println("Total brightness:", brightness)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
