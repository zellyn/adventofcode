package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
	"golang.org/x/exp/maps"
)

// var printf = fmt.Printf
var printf = func(string, ...any) {}

var DOWN = geom.Vec3{X: 0, Y: 0, Z: -1}
var UP = geom.Vec3{X: 0, Y: 0, Z: 1}

type block struct {
	a geom.Vec3
	b geom.Vec3
}

func rng(start, end int) []int {
	res := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		res = append(res, i)
	}
	return res
}

func (blk block) positions() []geom.Vec3 {
	a, b := blk.a, blk.b
	if a.X != b.X {
		return util.Map(rng(a.X, b.X), func(x int) geom.Vec3 { return a.WithX(x) })
	}
	if a.Y != b.Y {
		return util.Map(rng(a.Y, b.Y), func(y int) geom.Vec3 { return a.WithY(y) })
	}
	return util.Map(rng(a.Z, b.Z), func(z int) geom.Vec3 { return a.WithZ(z) })
}

func (blk block) bottomPositions() []geom.Vec3 {
	a, b := blk.a, blk.b
	if a.X != b.X {
		return util.Map(rng(a.X, b.X), func(x int) geom.Vec3 { return a.WithX(x) })
	}
	if a.Y != b.Y {
		return util.Map(rng(a.Y, b.Y), func(y int) geom.Vec3 { return a.WithY(y) })
	}
	return []geom.Vec3{a}
}

func (blk block) topPositions() []geom.Vec3 {
	a, b := blk.a, blk.b
	if a.X != b.X {
		return util.Map(rng(a.X, b.X), func(x int) geom.Vec3 { return a.WithX(x) })
	}
	if a.Y != b.Y {
		return util.Map(rng(a.Y, b.Y), func(y int) geom.Vec3 { return a.WithY(y) })
	}
	return []geom.Vec3{b}
}

func (blk block) Add(delta geom.Vec3) block {
	return block{a: blk.a.Add(delta), b: blk.b.Add(delta)}
}

func parseBlock(input string) (block, error) {
	aStr, bStr, _ := strings.Cut(input, "~")
	aInts, err := util.ParseInts(aStr, ",")
	if err != nil {
		return block{}, err
	}
	bInts, err := util.ParseInts(bStr, ",")
	if err != nil {
		return block{}, err
	}

	a := geom.V3(aInts[0], aInts[1], aInts[2])
	b := geom.V3(bInts[0], bInts[1], bInts[2])

	if a.X > b.X {
		a.X, b.X = b.X, a.X
	}
	if a.Y > b.Y {
		a.Y, b.Y = b.Y, a.Y
	}
	if a.Z > b.Z {
		a.Z, b.Z = b.Z, a.Z
	}

	return block{a: a, b: b}, nil
}

func sortIndicesByBottom(indices []int, blocks []block) {
	sort.Slice(indices, func(i, j int) bool { return blocks[indices[i]].a.Z < blocks[indices[j]].a.Z })
}

func parse(inputs []string) ([]block, error) {
	blocks, err := util.MapE(inputs, parseBlock)
	if err != nil {
		return nil, err
	}

	sort.Slice(blocks, func(i, j int) bool { return blocks[i].a.Z < blocks[j].a.Z })
	return blocks, nil
}

func below(blk block, m map[geom.Vec3]int) []int {
	resMap := make(map[int]bool)

	for _, pos := range blk.bottomPositions() {
		if index, found := m[pos.Add(DOWN)]; found && index > -1 {
			resMap[index] = true
		}
	}

	return maps.Keys(resMap)
}

func above(blk block, m map[geom.Vec3]int) []int {
	resMap := make(map[int]bool)

	for _, pos := range blk.topPositions() {
		if index, found := m[pos.Add(UP)]; found {
			resMap[index] = true
		}
	}

	return maps.Keys(resMap)
}

func setup(blocks []block) (map[geom.Vec3]int, map[int][]int, map[int][]int) {
	var minXY, maxXY geom.Vec2
	for _, blk := range blocks {
		a, b := blk.a.XY(), blk.b.XY()
		minXY = geom.Min2(minXY, a)
		maxXY = geom.Max2(maxXY, b)
	}
	r := geom.MakeRect(minXY, maxXY)

	m := make(map[geom.Vec3]int)
	for _, pos := range r.Positions() {
		m[pos.WithZ(0)] = -1
	}

	// printf("Before falling:\n")
	// for i, blk := range blocks {
	//   printf("%c: %v\n", 'A'+i, blk)
	// }

	for i, blk := range blocks {
		// printf("Dropping block %c:\n", 'A'+i)
		stuck := false
		for !stuck {
			for _, pos := range blk.bottomPositions() {
				if _, found := m[pos.Add(DOWN)]; found {
					stuck = true
					break
				}
			}
			if !stuck {
				// printf(" drop!\n")
				blk = blk.Add(DOWN)
			}
		}
		blocks[i] = blk
		for _, pos := range blk.positions() {
			m[pos] = i
		}
		// printf("m: %v\n", m)
	}

	aboves := make(map[int][]int)
	for i, block := range blocks {
		aboves[i] = above(block, m)
	}
	belows := make(map[int][]int)
	for i, block := range blocks {
		belows[i] = below(block, m)
	}

	return m, aboves, belows
}

func part1(inputs []string) (int, error) {
	blocks, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	_, aboves, belows := setup(blocks)

	count := 0
OUTER:
	for i := range blocks {
		// printf("Block %c:\n", 'A'+i)
		aboveThis := aboves[i]
		if len(aboveThis) == 0 {
			count++
			// printf(" none above: count=%d\n", count)
			continue
		}
		for _, index := range aboveThis {
			belowThis := belows[index]
			if len(belowThis) < 2 {
				continue OUTER
			}
		}
		count++
	}

	return count, nil
}

func fallCount(disintegrated int, blocks []block, aboves map[int][]int, belows map[int][]int) int {
	fallen := map[int]bool{
		disintegrated: true,
	}
	seen := map[int]bool{
		disintegrated: true,
	}

	queue := slices.Clone(aboves[disintegrated])
	printf(" initial queue: %v\n", queue)

OUTER:
	for len(queue) > 0 {
		sortIndicesByBottom(queue, blocks)

		index := queue[0]
		queue = queue[1:]
		if seen[index] {
			continue OUTER
		}
		seen[index] = true

		if blocks[index].a.Z == 1 {
			panic("BOOM")
			// supported by floor: can't drop this one
			continue OUTER
		}
		for _, support := range belows[index] {
			if !fallen[support] {
				// still supported: can't drop this one
				continue OUTER
			}
		}
		// Ok, so we can fall.
		fallen[index] = true
		for _, supported := range aboves[index] {
			if !fallen[supported] {
				queue = append(queue, supported)
			}
		}

	}

	printf(" returning %d\n", len(fallen)-1)
	return len(fallen) - 1
}

func part2(inputs []string) (int, error) {
	blocks, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	_, aboves, belows := setup(blocks)

	count := 0
	for i := range blocks {
		printf("Computing fallcount of block %d/%d:\n", i+1, len(blocks))
		fc := fallCount(i, blocks, aboves, belows)
		printf(" %d\n", fc)
		count += fc
	}

	return count, nil
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
