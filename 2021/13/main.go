package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type fold struct {
	axis string
	pos  int
}

func parse(inputs []string) (charmap.M, []fold, error) {
	ps := util.LinesByParagraph(inputs)
	coords, err := util.ParseLinesOfInts(ps[0], ",")
	if err != nil {
		return nil, nil, err
	}

	m := charmap.Empty()
	for _, coord := range coords {
		m[geom.V2(coord[0], coord[1])] = '#'
	}

	folds, err := util.MapE(ps[1], func(s string) (fold, error) {
		fields := strings.Fields(s)
		if len(fields) != 3 {
			return fold{}, fmt.Errorf("expected 3 fields in line %q", s)
		}
		axisName, strPos, found := strings.Cut(fields[len(fields)-1], "=")
		if !found {
			return fold{}, fmt.Errorf("expected /[xy]=[0-9]+/ for last field of %q", s)
		}
		pos, err := strconv.Atoi(strPos)
		if err != nil {
			return fold{}, fmt.Errorf("weird position %q in input %q", strPos, s)
		}
		if axisName != "x" && axisName != "y" {
			return fold{}, fmt.Errorf("weird axis %q in input %q", axisName, s)
		}
		return fold{axis: axisName, pos: pos}, nil
	})
	if err != nil {
		return nil, nil, err
	}

	return m, folds, nil
}

func doFold(m charmap.M, f fold) charmap.M {
	m2 := charmap.Empty()

	if f.axis == "x" {
		for pos := range m {
			if pos.X < f.pos {
				m2[pos] = '#'
			} else {
				m2[pos.WithX(f.pos-(pos.X-f.pos))] = '#'
			}
		}
	} else {
		for pos := range m {
			if pos.Y < f.pos {
				m2[pos] = '#'
			} else {
				m2[pos.WithY(f.pos-(pos.Y-f.pos))] = '#'
			}
		}
	}

	return m2
}

func part1(inputs []string) (int, error) {
	m, folds, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	m = doFold(m, folds[0])
	return len(m), nil
}

func run() error {
	m, folds, err := parse(util.MustReadLines("input"))
	if err != nil {
		return err
	}
	for _, f := range folds {
		m = doFold(m, f)
	}
	fmt.Println(charmap.String(m, '.'))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
