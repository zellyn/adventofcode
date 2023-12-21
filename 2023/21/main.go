package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = fmt.Printf
var printf = func(string, ...any) {}

func step(m charmap.M) charmap.M {
	res := make(charmap.M, len(m))

	for pos, what := range m {
		switch what {
		case '#':
			res[pos] = '#'
		case '.':
			if res[pos] == 0 {
				res[pos] = '.'
			}
		case 'S', 'O':
			if res[pos] == 0 {
				res[pos] = '.'
			}

			for _, dir := range geom.Compass4 {
				target := pos.Add(dir)
				if m[target] == 0 {
					continue
				}
				if m[target] != '#' {
					res[target] = 'O'
				}
			}
		}
	}

	return res
}

func step2(m charmap.M, size int) {
	for pos, what := range m {
		if what == 'O' {
			m[pos] = '.'

			for _, dir := range geom.Compass4 {
				target := pos.Add(dir)
				source := target
				for source.X < 0 {
					source.X += size
				}
				for source.X >= size {
					source.X -= size
				}
				for source.Y < 0 {
					source.Y += size
				}
				for source.Y >= size {
					source.Y -= size
				}
				if m[source] == '#' {
					m[target] = '#'
				} else {
					m[target] = 'o'
				}
			}
		}
	}
	for pos, what := range m {
		if what == 'o' {
			m[pos] = 'O'
		}
	}
}

func part1(inputs []string, target int) (int, error) {
	m := charmap.Parse(inputs)
	for i := 0; i < target; i++ {
		m = step(m)
		// printf("%v\n\n", m.AsString('/'))
	}
	return m.Count('O'), nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func countAt(m charmap.M, size int, pos geom.Vec2) int {
	offset := pos.Mul(size)
	count := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if m[offset.Add(geom.Vec2{X: x, Y: y})] == 'O' {
				count++
			}
		}
	}
	return count
}

func part2(inputs []string, target int) (int, error) {
	m := charmap.Parse(inputs)
	min, max := m.MinMax()
	size := max.X - min.X + 1

	sPos := geom.V2((size-1)/2, (size-1)/2)

	if m[sPos] != 'S' {
		return 0, fmt.Errorf("Expected S at %s; got %c\n", sPos, m[sPos])
	}
	m[sPos] = 'O'

	printf("target = %d * %d + %d\n", target/size, size, target%size)

	steps := 0
	loops := 2
	printf("Doing %d * %d = %d steps\n", size, loops, size*loops)
	for i := 0; i < loops; i++ {
		printf(" iteration %d\n", i+1)
		for j := 0; j < size; j++ {
			steps++
			step2(m, size)
		}
	}

	printf("Doing %d more steps\n", target%size)
	for i := 0; i < target%size; i++ {
		steps++
		step2(m, size)
	}

	printf("Counting\n")
	counts := make(map[geom.Vec2]int)
	checked := make(map[geom.Vec2]bool)
	queue := []geom.Vec2{{X: 0, Y: 0}}

	var countMinPos, countMaxPos geom.Vec2
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if checked[pos] {
			continue
		}
		checked[pos] = true
		count := countAt(m, size, pos)
		if count > 0 {
			countMinPos = geom.Min2(countMinPos, pos)
			countMaxPos = geom.Max2(countMaxPos, pos)
			counts[pos] = count
			queue = append(queue, pos.N(), pos.S(), pos.E(), pos.W())
		}
	}

	printf("%v\n", counts)

	for y := countMinPos.Y; y <= countMaxPos.Y; y++ {
		if y > countMinPos.Y {
			printf("\n\n\n")
		}
		for x := countMinPos.X; x <= countMaxPos.X; x++ {
			score := counts[geom.Vec2{X: x, Y: y}]
			if score == 0 {
				printf(strings.Repeat(" ", 10))
			} else {
				printf("%10d", score)
			}
		}
	}
	printf("\n")

	printf("In %d steps, we got a total count of %d\n", steps, m.Count('O'))

	evenCount := counts[geom.V2(0, 0)]
	oddCount := counts[geom.V2(1, 0)]
	if loops%2 == 1 {
		evenCount, oddCount = oddCount, evenCount
	}

	northPos := geom.Vec2{X: 0, Y: countMinPos.Y}
	southPos := geom.Vec2{X: 0, Y: countMaxPos.Y}
	westPos := geom.Vec2{X: countMinPos.X, Y: 0}
	eastPos := geom.Vec2{X: countMaxPos.X, Y: 0}

	nCount := counts[northPos]
	sCount := counts[southPos]
	wCount := counts[westPos]
	eCount := counts[eastPos]
	nwCount := counts[northPos.W()]
	neCount := counts[northPos.E()]
	swCount := counts[southPos.W()]
	seCount := counts[southPos.E()]
	nwInnerCount := counts[northPos.SW()]
	neInnerCount := counts[northPos.SE()]
	swInnerCount := counts[southPos.NW()]
	seInnerCount := counts[southPos.NE()]

	printf("\n")
	printf("evenCount:    %d\n", evenCount)
	printf("oddCount:     %d\n", oddCount)
	printf("nCount:       %d\n", nCount)
	printf("sCount:       %d\n", sCount)
	printf("eCount:       %d\n", eCount)
	printf("wCount:       %d\n", wCount)
	printf("nwCount:      %d\n", nwCount)
	printf("neCount:      %d\n", neCount)
	printf("swCount:      %d\n", swCount)
	printf("seCount:      %d\n", seCount)
	printf("nwInnerCount: %d\n", nwInnerCount)
	printf("neInnerCount: %d\n", neInnerCount)
	printf("swInnerCount: %d\n", swInnerCount)
	printf("seInnerCount: %d\n", seInnerCount)
	printf("\n")

	spread := target / size
	printf("spread: %d\n", spread)

	res := nCount + sCount + eCount + wCount
	res += nwCount * spread
	res += neCount * spread
	res += swCount * spread
	res += seCount * spread
	res += nwInnerCount * (spread - 1)
	res += neInnerCount * (spread - 1)
	res += swInnerCount * (spread - 1)
	res += seInnerCount * (spread - 1)

	oddRep := spread * spread
	evenRep := (spread - 1) * (spread - 1)

	res += oddRep * oddCount
	res += evenRep * evenCount

	// (+ 4 (* 4 3)) 16
	// (+ 3 (* 3 2)) 9
	// (+ 2 (* 2 1)) 4
	// (+ 1 (* 1 0)) 1

	return res, nil
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
