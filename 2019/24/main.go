package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

type vec2 = geom.Vec2
type vec3 = geom.Vec3

const bug = '#'
const space = '.'
const question = '?'

func next(m map[vec2]rune) map[vec2]rune {
	mm := make(map[vec2]rune, len(m))

	for pos, r := range m {
		neighbors := 0
		for _, nn := range geom.Neighbors4(pos) {
			if m[nn] == bug {
				neighbors++
			}
		}
		if (r == bug && neighbors == 1) || (r == space && (neighbors == 1 || neighbors == 2)) {
			mm[pos] = bug
		} else {
			mm[pos] = space
		}
	}
	return mm
}

func score(m map[vec2]rune) int {
	sum := 0
	score := 1
	min, max := charmap.MinMax(m)

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			if m[geom.Vec2{X: x, Y: y}] == bug {
				sum += score
			}
			score <<= 1
		}
	}
	return sum
}

func neighbors(p vec3) []vec3 {
	if p.X == 2 && p.Y == 2 {
		return nil
	}

	var nn []vec3

	// Up
	if p.Y == 0 {
		nn = append(nn, vec3{2, 1, p.Z - 1})
	} else if p.Y == 3 && p.X == 2 {
		for x := 0; x < 5; x++ {
			nn = append(nn, vec3{x, 4, p.Z + 1})
		}
	} else {
		nn = append(nn, vec3{p.X, p.Y - 1, p.Z})
	}

	// Down
	if p.Y == 4 {
		nn = append(nn, vec3{2, 3, p.Z - 1})
	} else if p.Y == 1 && p.X == 2 {
		for x := 0; x < 5; x++ {
			nn = append(nn, vec3{x, 0, p.Z + 1})
		}
	} else {
		nn = append(nn, vec3{p.X, p.Y + 1, p.Z})
	}

	// Left
	if p.X == 0 {
		nn = append(nn, vec3{1, 2, p.Z - 1})
	} else if p.Y == 2 && p.X == 3 {
		for y := 0; y < 5; y++ {
			nn = append(nn, vec3{4, y, p.Z + 1})
		}
	} else {
		nn = append(nn, vec3{p.X - 1, p.Y, p.Z})
	}

	// Right
	if p.X == 4 {
		nn = append(nn, vec3{3, 2, p.Z - 1})
	} else if p.Y == 2 && p.X == 1 {
		for y := 0; y < 5; y++ {
			nn = append(nn, vec3{0, y, p.Z + 1})
		}
	} else {
		nn = append(nn, vec3{p.X + 1, p.Y, p.Z})
	}

	return nn
}

func makeLevel(m map[vec3]rune, level int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			r := space
			if y == 2 && x == 2 {
				r = question
			}
			if m[vec3{x, y, level}] != 0 {
				panic(vec3{x, y, level})
			}
			m[vec3{X: x, Y: y, Z: level}] = r
		}
	}
}

func next3d(m map[vec3]rune) map[vec3]rune {
	mm := make(map[vec3]rune, len(m))

	// Ensure that neighbors to bugs are built out
	for pos, r := range m {
		if r == bug {
			for _, nn := range neighbors(pos) {
				if m[nn] == 0 {
					makeLevel(m, nn.Z)
				}
			}
		}
	}

	for pos, r := range m {
		if pos.X == 2 && pos.Y == 2 {
			mm[pos] = m[pos]
			continue
		}
		count := 0
		for _, nn := range neighbors(pos) {
			if m[nn] == bug {
				count++
			}
		}
		if (r == bug && count == 1) || (r == space && (count == 1 || count == 2)) {
			mm[pos] = bug
		} else {
			mm[pos] = space
		}
	}
	return mm
}

func printAll(m map[vec3]rune) {
	minLevel, maxLevel := 0, 0
	for pos := range m {
		if pos.Z < minLevel {
			minLevel = pos.Z
		}
		if pos.Z > maxLevel {
			maxLevel = pos.Z
		}
	}

	for level := minLevel; level <= maxLevel; level++ {
		fmt.Printf("Depth: %d:\n", level)
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				fmt.Printf("%c", m[vec3{X: x, Y: y, Z: level}])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func countBugs(m map[vec3]rune) int {
	count := 0
	for _, r := range m {
		if r == bug {
			count++
		}
	}
	return count
}

func run() error {
	ex1 := strings.Split(`....#
#..#.
#..##
..#..
#....`, "\n")

	ex2 := strings.Split(`.....
.....
.....
#....
.#...`, "\n")

	m := charmap.Parse(ex1)
	charmap.Draw(m, ' ')
	fmt.Println()
	charmap.Draw(next(m), ' ')

	fmt.Println()
	m = charmap.Parse(ex2)
	charmap.Draw(m, ' ')
	fmt.Println(score(m))

	scores := map[int]bool{}

	m, err := charmap.Read("input")
	if err != nil {
		return err
	}

	i3 := map[vec3]rune{}
	for pos, r := range m {
		i3[vec3{pos.X, pos.Y, 0}] = r
	}
	i3[vec3{2, 2, 0}] = question

	for {
		s := score(m)
		if scores[s] {
			fmt.Printf("Score: %d\n", s)
			break
		}
		scores[s] = true
		m = next(m)
	}

	m2 := charmap.Parse(ex1)
	m3 := map[vec3]rune{}
	for pos, r := range m2 {
		m3[vec3{pos.X, pos.Y, 0}] = r
	}
	m3[vec3{2, 2, 0}] = question
	printAll(m3)
	fmt.Println()
	for i := 0; i < 10; i++ {
		m3 = next3d(m3)
	}
	printAll(m3)
	fmt.Println(countBugs(m3))

	for i := 0; i < 200; i++ {
		i3 = next3d(i3)
	}
	fmt.Println(countBugs(i3))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
