package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var dirPipes = map[string]string{
	"N|": "N",
	"S|": "S",

	"E-": "E",
	"W-": "W",

	"SL": "E",
	"WL": "N",

	"SJ": "W",
	"EJ": "N",

	"N7": "W",
	"E7": "S",

	"NF": "E",
	"WF": "S",
}

var valid = map[string]bool{
	"|": true,
	"-": true,
	"L": true,
	"J": true,
	"7": true,
	"F": true,
}

func tracePath(m charmap.M, pos geom.Vec2, dirName string) (int, charmap.M, rune, error) {
	length := 0
	pathMap := make(charmap.M)
	startDirName := dirName
	for {
		pathMap[pos] = m[pos]
		length++
		newPos := pos.Add(geom.Compass4[dirName])
		newRune := string(m[newPos])
		if newRune == "S" {
			for pipe := range valid {
				if dirPipes[dirName+pipe] == startDirName {
					return length, pathMap, rune(pipe[0]), nil
				}
			}

			return 0, pathMap, 0, fmt.Errorf("Cannot find a pipe that goes from %s to %s", dirName, startDirName)
		}
		if !valid[newRune] {
			return 0, nil, 0, fmt.Errorf("ran into non-path rune")
		}
		pipeDirKey := dirName + newRune
		newDirName, ok := dirPipes[pipeDirKey]
		if !ok {
			return 0, nil, 0, fmt.Errorf("tried to enter pipe from invalid direction")
		}
		pos = newPos
		dirName = newDirName
	}
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	start, found := m.Find('S')
	if !found {
		return 0, fmt.Errorf("cannot find S")
	}
	for name := range geom.Compass4 {
		if l, _, _, err := tracePath(m, start, name); err == nil {
			return l / 2, nil
		}
	}
	return 0, fmt.Errorf("cannot find any paths that return")
}

// part2_original_verbose_error_handling is what I actually wrote to solve the problem.
func part2_original_verbose_error_handling(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	start, found := m.Find('S')
	if !found {
		return 0, fmt.Errorf("cannot find S")
	}
	for name := range geom.Compass4 {
		_, pathM, sRune, err := tracePath(m, start, name)
		if err != nil {
			continue
		}
		pathM[start] = sRune
		min, max := pathM.MinMax()
		count := 0
		for y := min.Y; y <= max.Y; y++ {
			inside := false
			lastCorner := 'x'

			for x := min.X; x <= max.X; x++ {
				pos := geom.Vec2{X: x, Y: y}
				c := pathM[pos]
				switch c {
				case '-':
				case '|':
					inside = !inside
				case 'F', 'L':
					if lastCorner != 'x' {
						return 0, fmt.Errorf("Got char corner '%c' at %v when lastCorner='%c' (should be 'x')", c, pos, lastCorner)
					}
					inside = !inside
					lastCorner = c
				case 'J':
					switch lastCorner {
					case 'x':
						return 0, fmt.Errorf("Got char corner 'J' at %v when lastCorner='x'", pos)
					case 'F':
						lastCorner = 'x'
					case 'L':
						lastCorner = 'x'
						inside = !inside
					}
				case '7':
					switch lastCorner {
					case 'x':
						return 0, fmt.Errorf("Got char corner '7' at %v when lastCorner='x'", pos)
					case 'L':
						lastCorner = 'x'
					case 'F':
						lastCorner = 'x'
						inside = !inside
					}
				default:
					if inside {
						count++
					}
				}
			}
		}
		return count, nil
	}
	return 0, fmt.Errorf("cannot find any paths that return")
}

// part2 is the same as part2_original_verbose_error_handling, except
// error checking in the core logic has been removed to make it more
// concise.
func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	start, found := m.Find('S')
	if !found {
		return 0, fmt.Errorf("cannot find S")
	}
	for name := range geom.Compass4 {
		_, pathM, sRune, err := tracePath(m, start, name)
		if err != nil {
			continue
		}
		pathM[start] = sRune
		min, max := pathM.MinMax()
		count := 0
		for y := min.Y; y <= max.Y; y++ {
			inside := false
			lastCorner := 'x'

			for x := min.X; x <= max.X; x++ {
				c := pathM[geom.Vec2{X: x, Y: y}]
				switch c {
				case '-':
				case '|':
					inside = !inside
				case 'F', 'L':
					lastCorner = c
				case 'J':
					if lastCorner == 'F' {
						inside = !inside
					}
				case '7':
					if lastCorner == 'L' {
						inside = !inside
					}
				default:
					if inside {
						count++
					}
				}
			}
		}
		return count, nil
	}
	return 0, fmt.Errorf("cannot find any paths that return")
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
