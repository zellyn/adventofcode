package main

import (
	"fmt"
	"math"
	"math/big"
	"math/bits"
	"os"
	"slices"
	"strings"

	"github.com/zellyn/adventofcode/linalg"
	"github.com/zellyn/adventofcode/util"
)

var printf = func(string, ...any) {}

// var printf = fmt.Printf

type config struct {
	target   uint
	opsBits  []uint
	joltages []int
	ops      [][]int
}

type bound struct {
	min *big.Rat
	max *big.Rat
}

func fixedBound(f int) bound {
	return bound{
		min: big.NewRat(int64(f), 1),
		max: big.NewRat(int64(f), 1),
	}
}

func (b bound) Clone() bound {
	return bound{
		min: new(big.Rat).Set(b.min),
		max: new(big.Rat).Set(b.max),
	}
}

func (b bound) String() string {
	return fmt.Sprintf("[%s,%s]", b.min.RatString(), b.max.RatString())
}

func (b bound) Fixed() bool {
	return b.min.Cmp(b.max) == 0
}

func (b bound) MinInt() int {
	return int(b.min.Num().Int64())
}

func (b bound) MaxInt() int {
	return int(b.max.Num().Int64())
}

func (b bound) Size() int {
	return b.MaxInt() - b.MinInt() + 1
}

func parse(inputs []string) ([]config, error) {
	return util.MapE(inputs, func(s string) (config, error) {
		parts := strings.Split(s, " ")
		first := parts[0]
		last := parts[len(parts)-1]
		var ops [][]int
		opsBits, err := util.MapE(parts[1:len(parts)-1], func(s string) (uint, error) {
			ints, err := util.ParseInts(s[1:len(s)-1], ",")
			if err != nil {
				return 0, err
			}

			var opBits uint
			ops = append(ops, ints)
			for _, i := range ints {
				opBits |= (1 << i)
			}
			return opBits, nil
		})
		if err != nil {
			return config{}, err
		}

		// slices.SortFunc(ops, func(a, b []int) int {
		// 	return -cmp.Compare(len(a), len(b))
		// })

		var target uint
		for i, c := range first[1 : len(first)-1] {
			switch c {
			case '.':
			case '#':
				target |= (1 << i)
			default:
				return config{}, fmt.Errorf("weird input char '%c' in %s", c, first)
			}
		}

		joltages, err := util.ParseInts(last[1:len(last)-1], ",")
		if err != nil {
			return config{}, err
		}

		return config{
			target:   target,
			opsBits:  opsBits,
			joltages: joltages,
			ops:      ops,
		}, nil
	})
}

func least1(c config) int {
	printf("%d\n", c)
	count := uint(1 << len(c.opsBits))
	best := len(c.opsBits)

	for mask := range count {
		printf("%b:\n", mask)
		oc := bits.OnesCount(mask)
		if oc > best {
			printf(" too many!\n")
			continue
		}

		target := c.target
		for i, op := range c.opsBits {
			if mask&(1<<i) > 0 {
				target ^= op
				printf(" using %d; target=%d (oc=%d)\n", op, target, oc)
			}
		}
		if target == 0 && oc < best {
			best = oc
			printf("                              best is now %d\n", best)
		}
	}

	return best
}

func least2(joltages []int, ops [][]int) int {
	above := util.Sum(joltages) + 1
	bounds := make([]bound, len(ops))

	for i, op := range ops {
		bounds[i].max = big.NewRat(int64(above), 1)
		bounds[i].min = new(big.Rat)
		for _, idx := range op {
			bounds[i].max = ratMin(bounds[i].max, big.NewRat(int64(joltages[idx]), 1))
		}
	}

	var rows [][]int
	for _, joltage := range joltages {
		row := make([]int, len(ops)+1)
		row[len(row)-1] = joltage
		rows = append(rows, row)
	}

	for i, op := range ops {
		for _, idx := range op {
			rows[idx][i] = 1
		}
	}

	m := linalg.NewMatrix(rows)
	printf("%s\n", m.Printable("", true))
	ks, knowns := m.KnownCoefficients()
	for i, k := range ks {
		if knowns[i] {
			if !k.IsInt() {
				panic(fmt.Sprintf("known coefficient %v is not an int", k))
			}
			bounds[i].min = k
			bounds[i].max = k
		}
	}

	if util.All(knowns) {
		res := 0
		for _, k := range ks {
			res += int(k.Num().Int64())
		}
		return res
	}

	printf("%s\n", m.Printable("", true))
	printf("%v %v\n", ks, knowns)
	printf("%v\n", bounds)

	narrowBounds(bounds, m)

	printf("%v\n", bounds)

	if res, fixed := checkBounds(bounds); fixed {
		return res
	}

	return search("", bounds, m)
}

func narrowBounds(bounds []bound, m *linalg.Matrix) {
	done := false
	for !done {
		done = true
		for _, row := range m.Rows() {
			if row.Empty() {
				continue
			}
			for i := range len(row.Coefficients()) {
				if updateBounds(i, bounds, row) {
					done = false
				}
			}
		}
	}
}

func search(prefix string, bounds []bound, m *linalg.Matrix) int {
	var minBoundIdx int
	minSize := math.MaxInt

	for idx, b := range bounds {
		size := b.Size()
		if size == 1 {
			continue
		}
		if size < minSize {
			minSize = size
			minBoundIdx = idx
		}
	}

	if minSize == math.MaxInt {
		panic(fmt.Sprintf("couldn't find minimum non-zero bound for bounds %v", bounds))
	}

	// printf("%sidx:%d bounds: %v\n", prefix, minBoundIdx, bounds)

	minTotal := math.MaxInt

	minBound := bounds[minBoundIdx]
	lower, upper := minBound.MinInt(), minBound.MaxInt()

	rowInts := make([]int, len(bounds))
	rowInts[minBoundIdx] = 1

	newBounds := make([]bound, len(bounds))

	for val := lower; val <= upper; val++ {
		// printf("%s%d <= %d <= %d\n", prefix, lower, val, upper)
		m2 := m.Clone()
		m2.AddRow(linalg.NewRow(rowInts, val))
		if m2.Impossible() {
			continue
		}

		ks, knowns := m2.KnownCoefficients()
		if util.All(knowns) {
			minTotal = min(minTotal, cSum(ks))
			continue
		}

		for i, b := range bounds {
			newBounds[i] = b.Clone()
		}
		newBounds[minBoundIdx] = fixedBound(val)

		narrowBounds(newBounds, m2)
		if res, fixed := checkBounds(newBounds); fixed {
			minTotal = min(minTotal, res)
			continue
		}

		minTotal = min(minTotal, search(prefix+"  ", newBounds, m2))
	}

	return minTotal
}

func cSum(rats []*big.Rat) int {
	sum := 0
	for _, rat := range rats {
		if !rat.IsInt() {
			return math.MaxInt
		}

		sum += int(rat.Num().Int64())
	}

	return sum
}

func checkBounds(bounds []bound) (int, bool) {
	res := 0

	for _, b := range bounds {
		if !b.Fixed() {
			return 0, false
		}
		res += int(b.min.Num().Int64())
	}

	return res, true
}

var bigIntZero = new(big.Int)
var ratOne = big.NewRat(1, 1)
var ratZero = new(big.Rat)
var bigOne = big.NewInt(1)

func ratIsZero(r *big.Rat) bool {
	return r.Num().Sign() == 0
}

func ratMin(a, b *big.Rat) *big.Rat {
	if a.Cmp(b) < 1 {
		return a
	}
	return b
}

func ratMax(a, b *big.Rat) *big.Rat {
	if a.Cmp(b) < 1 {
		return b
	}
	return a
}

func ratCeil(r *big.Rat) *big.Rat {
	if r.IsInt() {
		return r
	}
	div := new(big.Int)
	div.Div(r.Num(), r.Denom())
	div.Add(div, bigOne)

	return new(big.Rat).SetInt(div)
}

func ratFloor(r *big.Rat) *big.Rat {
	if r.IsInt() {
		return r
	}
	div := new(big.Int)
	div.Div(r.Num(), r.Denom())

	return new(big.Rat).SetInt(div)
}

func ratSort(a, b *big.Rat) (*big.Rat, *big.Rat) {
	if a.Cmp(b) < 1 {
		return a, b
	}
	return b, a
}

func ratsString(rats []*big.Rat) string {
	return "[" + strings.Join(util.Map(rats, func(r *big.Rat) string { return r.RatString() }), ",") + "]"
}

func updateBounds(idx int, bounds []bound, row *linalg.Row) bool {
	if bounds[idx].min.Cmp(bounds[idx].max) == 0 {
		return false
	}
	cs := slices.Clone(row.Coefficients())
	if ratIsZero(cs[idx]) {
		return false
	}
	for i, c := range cs {
		cs[i] = new(big.Rat).Set(c)
	}
	k := new(big.Rat).Set(row.Constant())
	// printf("updateBounds(%d, %v)\n", idx, bounds)
	// printf(" cs=%s, k=%s\n", ratsString(cs), k.RatString())

	if cs[idx].Cmp(ratOne) != 0 {
		inv := new(big.Rat).Inv(cs[idx])
		for _, c := range cs {
			c.Mul(c, inv)
		}
		k.Mul(k, inv)
		// printf(" cs=%s, k=%s\n", ratsString(cs), k.RatString())
	}

	kmin := new(big.Rat).Set(k)
	kmax := new(big.Rat).Set(k)

	for i, c := range cs {
		if i == idx {
			continue
		}
		cmin := new(big.Rat).Mul(c, bounds[i].min)
		cmax := new(big.Rat).Mul(c, bounds[i].max)
		cmin.Neg(cmin)
		cmax.Neg(cmax)
		cmin, cmax = ratSort(cmin, cmax)
		kmin.Add(kmin, cmin)
		kmax.Add(kmax, cmax)
		// printf("  kmin,kmax = %s,%s\n", kmin.RatString(), kmax.RatString())
	}

	kmin = ratMax(kmin, ratZero)
	kmin = ratCeil(kmin)
	kmax = ratFloor(kmax)

	changed := false
	if bounds[idx].min.Cmp(kmin) == -1 {
		changed = true
		bounds[idx].min = kmin
	}
	if bounds[idx].max.Cmp(kmax) == 1 {
		changed = true
		bounds[idx].max = kmax
	}

	return changed
}

func part1(inputs []string) (int, error) {
	fmt.Printf("HERE\n")
	configs, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	sum := 0
	for _, c := range configs {
		sum += least1(c)
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	configs, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	sum := 0
	for _, c := range configs {
		sum += least2(c.joltages, c.ops)
	}
	return sum, nil
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
