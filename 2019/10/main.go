package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

func gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if a == 0 {
		return b
	}
	if a == 1 {
		return 1
	}
	m := b % a
	if m == 0 {
		return a
	}
	return gcd(m, a)
}

func ratio(a, b int) [2]int {
	sa, sb := 1, 1
	if a < 0 {
		sa = -1
		a = -a
	}
	if b < 0 {
		sb = -1
		b = -b
	}

	if a == 0 {
		if b == 0 {
			return [2]int{0, 0}
		}
		return [2]int{0, sb}
	} else if b == 0 {
		return [2]int{sa, b}
	}

	g := gcd(a, b)
	return [2]int{sa * a / g, sb * b / g}
}

func parseMap(data string) map[[2]int]bool {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	result := make(map[[2]int]bool)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				result[[2]int{x, y}] = true
			}
		}
	}
	return result
}

func visible(m map[[2]int]bool, from [2]int) int {
	ratios := make(map[[2]int]bool)
	for to := range m {
		if to == from {
			continue
		}
		r := ratio(to[0]-from[0], to[1]-from[1])
		ratios[r] = true
	}
	return len(ratios)
}

type info struct {
	pos    [2]int
	distSq int
}

// byDistSq implements sort.Interface for []info based on the distSq field.
type byDistSq []info

func (a byDistSq) Len() int           { return len(a) }
func (a byDistSq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byDistSq) Less(i, j int) bool { return a[i].distSq < a[j].distSq }

// byAngle implements sort.Interface for ratios.
type byAngle [][2]int

func (a byAngle) Len() int      { return len(a) }
func (a byAngle) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byAngle) Less(i, j int) bool {
	return angle(a[i]) < angle(a[j])
}

func angle(pt [2]int) float64 {
	return math.Mod(math.Atan2(float64(pt[0]), -float64(pt[1]))+2*math.Pi, 2*math.Pi)
}

func stacks(m map[[2]int]bool, from [2]int) map[[2]int][]info {
	result := make(map[[2]int][]info)
	for to := range m {
		if to == from {
			continue
		}
		x := to[0] - from[0]
		y := to[1] - from[1]
		r := ratio(x, y)
		result[r] = append(result[r], info{pos: to, distSq: x*x + y*y})
		sort.Sort(byDistSq(result[r]))
	}
	return result
}

func sequence(m map[[2]int]bool, from [2]int) [][2]int {
	var result [][2]int
	ss := stacks(m, from)
	var ratios [][2]int
	for k := range ss {
		ratios = append(ratios, k)
	}
	sort.Sort(byAngle(ratios))
	done := false
	for !done {
		done = true
		for _, key := range ratios {
			val := ss[key]
			if len(val) > 0 {
				done = false
				result = append(result, val[0].pos)
				ss[key] = val[1:]
			}
		}
	}
	return result
}

func best(m map[[2]int]bool) ([2]int, int) {
	max := -1
	maxPos := [2]int{-1, -1}
	for from := range m {
		v := visible(m, from)
		if v > max {
			max = v
			maxPos = from
		}
	}
	return maxPos, max
}

func run() error {
	input, err := util.ReadFileString("input")
	if err != nil {
		return err
	}
	m := parseMap(input)
	pos, score := best(m)
	fmt.Println(score)
	fmt.Println(sequence(m, pos)[199])
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
