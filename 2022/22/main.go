package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

type move struct {
	turn  int
	steps int
}

var move_regex = regexp.MustCompile(`([0-9]+|[LR])`)

var dirs = []geom.Vec2{
	{X: 1},
	{Y: 1},
	{X: -1},
	{Y: -1},
}

var chars = []rune{
	'>',
	'v',
	'<',
	'^',
}

var compassChars = []rune{
	'E',
	'S',
	'W',
	'N',
}

func left(dir int) int {
	return (dir + 3) % 4
}

func right(dir int) int {
	return (dir + 1) % 4
}

func opp(dir int) int {
	return (dir + 2) % 4
}

func add(dir int, inc int) int {
	return (dir + 4 + inc) % 4
}

func parseInput(input string) (charmap.M, []move) {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	c := charmap.ParseWithBackground(lines, ' ')
	var moves []move

	pieces := move_regex.FindAllString(parts[1], -1)

	thisMove := move{}

	for _, piece := range pieces {
		if piece[0] == 'L' {
			thisMove.turn = 3
		} else if piece[0] == 'R' {
			thisMove.turn = 1
		} else {
			steps, err := strconv.Atoi(piece)
			if err != nil {
				panic(err)
			}
			thisMove.steps = steps
			moves = append(moves, thisMove)
			thisMove = move{}
		}
	}

	return c, moves
}

func findStart(c charmap.M) geom.Vec2 {
	min, max := c.MinMax()
	pos := geom.Vec2{Y: min.Y}
	for x := min.X; x <= max.X; x++ {
		pos.X = x
		if c.Has(pos) && c[pos] != '#' {
			return pos
		}
	}
	return geom.Vec2{X: -1, Y: -1}
}

func walk(c charmap.M, pos geom.Vec2, dir geom.Vec2, steps int, r rune) geom.Vec2 {
	for i := 0; i < steps; i++ {
		c[pos] = r
		newPos := pos.Add(dir)
		if !c.Has(newPos) {
			for newPos = pos; c.Has(newPos.Sub(dir)); newPos = newPos.Sub(dir) {
			}
		}
		if c[newPos] == '#' {
			return pos
		}
		pos = newPos
	}
	c[pos] = r
	return pos
}

func findSides(c charmap.M) (int, []geom.Rect) {
	var size int
	for size = 1; len(c) > 6*size*size; size++ {
	}
	if len(c) != 6*size*size {
		s1 := size - 1
		panic(fmt.Sprintf("weird map has length %d, which is more than 6×%d×%d=%d and less than 6×%d×%d=%d", len(c), s1, s1, 6*s1*s1, size, size, 6*size*size))
	}

	var result []geom.Rect

	for y := 0; y < 6; y++ {
		pos := geom.Vec2{Y: y}
		for x := 0; x < 6; x++ {
			pos.X = x
			topLeft := pos.Mul(size)
			if c.Has(topLeft) {
				r := geom.Rect{
					Min: topLeft,
					Max: topLeft.Add(geom.Vec2{X: size - 1, Y: size - 1}),
				}
				result = append(result, r)
			}
		}
	}
	return size, result
}

func part1(input string) (int, error) {
	c, moves := parseInput(input)
	pos := findStart(c)
	fmt.Println(pos)
	dir := 0

	for _, m := range moves {
		dir = add(dir, m.turn)
		pos = walk(c, pos, dirs[dir], m.steps, chars[dir])
	}
	return 1000*(pos.Y+1) + 4*(pos.X+1) + dir, nil
}

type cube struct {
	m     charmap.M
	sides []geom.Rect
	size  int
	glue  map[[2]int][2]int
	pos   geom.Vec2
	dir   int
}

func buildCube(m charmap.M) *cube {
	size, sides := findSides(m)
	c := &cube{
		m:     m,
		pos:   findStart(m),
		sides: sides,
		size:  size,
		glue:  make(map[[2]int][2]int),
	}
	c.glueAll()
	return c
}

func (c *cube) glueAll() {
	fmt.Printf("Sides:\n%s\n\n", c.sidePic())
	fmt.Printf("sides: %v\n", c.sides)
	fmt.Printf("size: %d\n", c.size)
	for i, sideA := range c.sides {
		for j := i + 1; j < len(c.sides); j++ {
			sideB := c.sides[j]
			for dir := 0; dir < 4; dir++ {
				otherMin := sideA.Min.Add(dirs[dir].Mul(c.size))
				if sideB.Min == otherMin {
					c.glue[[2]int{i, dir}] = [2]int{j, dir}
					c.glue[[2]int{j, opp(dir)}] = [2]int{i, opp(dir)}
					fmt.Printf("PING! %d\n", len(c.glue))
				}
			}
		}
	}
	if len(c.glue) != 10 {
		panic(fmt.Sprintf("want 10 original glue; got %d", len(c.sides)))
	}

	done := false
	for !done {
		done = true

		for side := 0; side < 6; side++ {
			for dir := 0; dir < 4; dir++ {
				if _, has := c.glue[[2]int{side, left(dir)}]; has {
					continue
				}

				glue, has := c.glue[[2]int{side, dir}]
				if !has {
					continue
				}

				forwardSquare, forwardDir := glue[0], glue[1]
				glue2, ok := c.glue[[2]int{forwardSquare, left(forwardDir)}]
				if !ok {
					continue
				}

				done = false
				diagSquare, diagDir := glue2[0], glue2[1]

				fmt.Printf("went %c from %d to %d going %c, then %c to %d going %c\n",
					compassChars[dir], side,
					forwardSquare, compassChars[forwardDir],
					compassChars[left(forwardDir)],
					diagSquare, compassChars[diagDir])

				existingGlue, exists := c.glue[[2]int{diagSquare, left(diagDir)}]

				if exists {
					panic(fmt.Sprintf("diagonal square's glue already exists! glue[(%d,%c)]=%d, going %d",
						diagSquare, compassChars[left(diagDir)],
						existingGlue[0], compassChars[existingGlue[1]]))
				}
				fmt.Printf(" So %c from %d is %d going %c\n", compassChars[left(dir)], side, diagSquare, compassChars[right(diagDir)])
				fmt.Printf(" And %c from %d is %d going %c\n", compassChars[left(diagDir)], diagSquare, side, compassChars[right(dir)])
				c.glue[[2]int{side, left(dir)}] = [2]int{diagSquare, right(diagDir)}
				c.glue[[2]int{diagSquare, left(diagDir)}] = [2]int{side, right(dir)}
			}
		}
	}
}

func (c *cube) turn(inc int) {
	c.dir = add(c.dir, inc)
}

func (c *cube) side(pos geom.Vec2) (geom.Rect, int) {

	min := pos.Div(c.size).Mul(c.size)
	max := min.Add(geom.Vec2{X: c.size - 1, Y: c.size - 1})
	if pos.X < 0 {
		min.X = -c.size
		max.X = -1
	}
	if pos.Y < 0 {
		min.Y = -c.size
		max.Y = -1
	}
	rect := geom.MakeRect(min, max)

	for i := 0; i < 6; i++ {
		if c.sides[i] == rect {
			return rect, i
		}
	}

	return rect, -1
}

func (c *cube) sidePic() string {
	pic := charmap.Empty()

	for i, side := range c.sides {
		pic[side.Min.Div(c.size)] = '0' + rune(i)
	}

	return pic.AsString(' ')
}

func (c *cube) walk() {
	c.m[c.pos] = chars[c.dir]
	newPos := c.pos.Add(dirs[c.dir])
	newDir := c.dir

	if !c.m.Has(newPos) {
		rect, rectIndex := c.side(c.pos)
		if rectIndex < 0 {
			panic(fmt.Sprintf("BOOM! rect not found for pos %s", c.pos))
		}

		glue, ok := c.glue[[2]int{rectIndex, c.dir}]
		if !ok {
			panic(fmt.Sprintf("BOOM! glue not found for rect %s in direction %d", c.pos, c.dir))
		}

		var newRectIndex int
		newRectIndex, newDir = glue[0], glue[1]
		newRect := c.sides[newRectIndex]

		newTopLeft := newRect.Min

		originalNewPos := newPos
		originalNewRect, _ := c.side(newPos)

		xy := newPos.Sub(originalNewRect.Min)
		offsetX, offsetY := xy.X, xy.Y
		dirChange := (newDir - c.dir + 4) % 4

		switch dirChange {
		case 0:
			// no change
		case 1:
			// turn right
			newTopLeft = geom.Vec2{X: newRect.Max.X, Y: newRect.Min.Y}
		case 2:
			// reverse
			newTopLeft = geom.Vec2{X: newRect.Max.X, Y: newRect.Max.Y}
		case 3:
			// turn left
			newTopLeft = geom.Vec2{X: newRect.Min.X, Y: newRect.Max.Y}
		}

		newXDir := dirs[dirChange]
		newYDir := dirs[right(dirChange)]

		newPos = newTopLeft.Add(newXDir.Mul(offsetX)).Add(newYDir.Mul(offsetY))

		if !newRect.Contains(newPos) {
			fmt.Printf("We were at %s (in %d: %s) going %c, and now we want %s (in %d: %s) going %c\n",
				c.pos, rectIndex, rect, compassChars[c.dir],
				newPos, newRectIndex, newRect, compassChars[newDir])
			fmt.Printf("original newPos: %s\n", originalNewPos)

			fmt.Printf("dirChange: %d\n", dirChange)
			fmt.Printf("xy: %s\n", xy)

			panic(fmt.Sprintf("BOOM!"))
		}

		if !c.m.Has(newPos) {
			panic(fmt.Sprintf("BOOM! going in dir %d from %s, we got position %s, but there's nothing there", c.dir, c.pos, newPos))
		}
	}

	if c.m[newPos] != '#' {
		c.pos = newPos
		c.dir = newDir
		c.m[c.pos] = chars[c.dir]
	}
}

func part2(input string) (int, error) {
	m, moves := parseInput(input)
	c := buildCube(m)

	// fmt.Printf("Sides:\n%s\n\n", c.sidePic())

	for side := range c.sides {
		for dir := range dirs {

			_, ok := c.glue[[2]int{side, dir}]
			if !ok {
				fmt.Printf("%c from side %d is MISSING!\n", compassChars[dir], side)
			} else {
				// fmt.Printf("%c from side %d is side %d, going %c\n", compassChars[dir], side, glue[0], compassChars[glue[1]])
			}
		}
	}

	for _, m := range moves {
		c.turn(m.turn)
		// fmt.Printf("%s\n\n", c.m.AsString(' '))
		for i := 0; i < m.steps; i++ {
			c.walk()
		}
	}
	fmt.Printf("%s\n\n", c.m.AsString(' '))
	return 1000*(c.pos.Y+1) + 4*(c.pos.X+1) + c.dir, nil
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
