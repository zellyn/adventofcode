package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

func printf(format string, args ...any) {
	// fmt.Printf(format, args...)
}

var moves = [][]geom.Vec2{
	{{X: 0, Y: -1}, {X: -1, Y: -1}, {X: 1, Y: -1}},
	{{X: 0, Y: 1}, {X: -1, Y: 1}, {X: 1, Y: 1}},
	{{X: -1, Y: 0}, {X: -1, Y: -1}, {X: -1, Y: 1}},
	{{X: 1, Y: 0}, {X: 1, Y: -1}, {X: 1, Y: 1}},
}

func moveAll(m charmap.M, start int) (charmap.M, bool) {
	proposed := make(map[geom.Vec2]int, len(m))
	want := make(map[geom.Vec2]geom.Vec2, len(m))

OUTER:
	for pos := range m {
		printf("elf %c at %s:\n", m[pos], pos)
		alone := true
		for _, target := range pos.Neighbors8() {
			if m.Has(target) {
				alone = false
				break
			}
		}
		if alone {
			proposed[pos] = 1
			printf(" alone! staying put. proposed[%s]=%d\n", pos, proposed[pos])
			want[pos] = pos
			continue OUTER
		}

		for i := 0; i < 4; i++ {
			move := moves[(start+i)%4]
			target := pos.Add(move[0])
			printf(" considering %s\n", target)
			if !m.Has(target) && !m.Has(pos.Add(move[1])) && !m.Has(pos.Add(move[2])) {
				proposed[target]++
				printf("  yep! proposed[%s] is now %d\n", target, proposed[target])
				want[pos] = target
				continue OUTER
			}
		}
		proposed[pos] = 1
		want[pos] = pos

	}

	result := charmap.Empty()
	moved := false

	for pos := range m {
		target, ok := want[pos]
		if !ok {
			panic(fmt.Sprintf("boom! want[%s] is missing!", pos))
		}
		if proposed[target] == 1 {
			result[target] = m[pos]
			if target != pos {
				moved = true
			}
		} else {
			if proposed[target] == 0 {
				panic(fmt.Sprintf("boom! proposed[%s]==0", target))
			}
			result[pos] = m[pos]
		}
	}

	return result, moved
}

func part1(inputs []string) (int, error) {
	m := charmap.ParseWithBackground(inputs, '.')
	// name := 'A'
	// for pos := range m {
	// 	m[pos] = name
	// 	name++
	// }
	printf("== Initial State ==\n")
	printf("%s\n", m.AsString('.'))
	for i := 0; i < 10; i++ {
		m, _ = moveAll(m, i%4)
		printf("== End of Round %d ==\n", i+1)
		printf("%s\n", m.AsString('.'))
	}

	min, max := m.MinMax()
	return (max.Y-min.Y+1)*(max.X-min.X+1) - len(m), nil
}

func part2(inputs []string) (int, error) {
	m := charmap.ParseWithBackground(inputs, '.')
	// name := 'A'
	// for pos := range m {
	// 	m[pos] = name
	// 	name++
	// }
	printf("== Initial State ==\n")
	printf("%s\n", m.AsString('.'))
	for i := 0; ; i++ {
		if i > 0 && i%10000 == 0 {
			fmt.Println(i)
		}
		var moved bool
		m, moved = moveAll(m, i%4)
		if !moved {
			return i + 1, nil
		}
		printf("== End of Round %d ==\n", i+1)
		printf("%s\n", m.AsString('.'))
	}
	return 42, nil
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
