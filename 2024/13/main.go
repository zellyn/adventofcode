package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/math"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type machine struct {
	a     geom.Vec2
	b     geom.Vec2
	prize geom.Vec2
}

var buttonRe = regexp.MustCompile(`^Button [AB]: X[+](\d+), Y[+](\d+)$`)
var prizeRe = regexp.MustCompile(`^Prize: X=(\d+), Y=(\d+)$`)

func parseOne(lines []string) (machine, error) {
	m1 := buttonRe.FindStringSubmatch(lines[0])
	m2 := buttonRe.FindStringSubmatch(lines[1])
	m3 := prizeRe.FindStringSubmatch(lines[2])

	if m1 == nil {
		return machine{}, fmt.Errorf("Weird input %q in\n%s", lines[0], strings.Join(lines, "\n"))
	}
	if m2 == nil {
		return machine{}, fmt.Errorf("Weird input %q in\n%s", lines[1], strings.Join(lines, "\n"))
	}
	if m3 == nil {
		return machine{}, fmt.Errorf("Weird input %q in\n%s", lines[2], strings.Join(lines, "\n"))
	}

	ints, err := util.MapE([]string{m1[1], m1[2], m2[1], m2[2], m3[1], m3[2]}, strconv.Atoi)
	if err != nil {
		return machine{}, err
	}

	return machine{
		a:     geom.Vec2{X: ints[0], Y: ints[1]},
		b:     geom.Vec2{X: ints[2], Y: ints[3]},
		prize: geom.Vec2{X: ints[4], Y: ints[5]},
	}, nil

}

func (m machine) simplify() machine {
	gX := math.MultiGCD(m.a.X, m.b.X, m.prize.X)
	gY := math.MultiGCD(m.a.Y, m.b.Y, m.prize.Y)
	return machine{
		a:     geom.Vec2{X: m.a.X / gX, Y: m.a.Y / gY},
		b:     geom.Vec2{X: m.b.X / gX, Y: m.b.Y / gY},
		prize: geom.Vec2{X: m.prize.X / gX, Y: m.prize.Y / gY},
	}
}

func presses(m machine) int {
	m = m.simplify()

	maxA := m.prize.EachDiv(m.a).Min() + 1
	maxB := m.prize.EachDiv(m.b).Min() + 1

	minTokens := math.MaxInt

	for a := range maxA {
		if a*3 >= minTokens {
			continue
		}
		for b := range maxB {
			tokens := 3*a + b
			if tokens >= minTokens {
				continue
			}
			if m.a.Mul(a).Add(m.b.Mul(b)) == m.prize {
				minTokens = tokens
			}
		}
	}

	if minTokens == math.MaxInt {
		return 0
	}
	return minTokens
}

func presses2(m machine) int {
	fmt.Printf("initial machine: %v\n", m)
	m = m.simplify()
	fmt.Printf("simplified machine: %v\n", m)

	ra := m.a.ReduceDir()
	rb := m.b.ReduceDir()
	if ra == rb {
		if m.prize.ReduceDir() != ra {
			return 0
		}

		panic("Bang")
	}

	// Ax + By = C

	// We have two equations:
	// Ua + Vb = W     // m.a.X * a + m.b.X * b = prize.X
	// Xa + Yb = Z     // m.a.Y * a + m.b.Y * b = prize.Y
	// The lines are either parallel (including co-linear), or there is a unique solution.
	//
	// Vb = W - Ua
	// b = (W - Ua)/V
	// Yb = Z - Xa
	// b = (Z - Xa)/Y

	// b = (-U/V)a + W/V
	// b = (-X/Y)a + Z/Y

	// (W - Ua)/V = (Z - Xa)/Y
	// W/V - Ua/V = Z/Y - Xa/Y
	// Xa/Y - Ua/V = Z/Y - W/V
	// VXa/VY - UYa/VY = VZ/VY - WY/VY
	// VXa - UYa = VZ - WY
	// a(VX - UY) = VZ - WY

	// U = m.a.X
	// V = m.b.X
	// W = m.prize.X
	// X = m.a.Y
	// Y = m.b.Y
	// Z = m.prize.Y
	// a(VX - UY) = VZ - WY
	// a = (VZ - WY) / (VX - UY)
	// a = ((m.b.X)(m.prize.Y) - (m.prize.X)(m.b.Y)) / ((m.b.X)(m.a.Y) - (m.a.X)(m.b.Y))

	numerator_a := m.b.X*m.prize.Y - m.prize.X*m.b.Y
	denominator_a := m.b.X*m.a.Y - m.a.X*m.b.Y

	fmt.Printf("numerator_a: %d * %d - %d * %d == %d\n", m.b.X, m.prize.Y, m.prize.X, m.b.Y, numerator_a)
	fmt.Printf("denominator_a: %d * %d - %d * %d == %d\n", m.b.X, m.a.Y, m.a.X, m.b.Y, denominator_a)

	if numerator_a%denominator_a != 0 {
		return 0
	}

	a := numerator_a / denominator_a
	if a < 0 {
		return 0
	}

	fmt.Printf("a = %d\n", a)

	numerator_b := m.prize.X - m.a.X*a
	if numerator_b < 0 {
		return 0
	}
	denominator_b := m.b.X

	fmt.Printf("numerator_b: %d - %d * %d == %d\n", m.prize.X, m.a.X, a, numerator_b)
	fmt.Printf("denominator_b = %d\n", m.b.X)

	if numerator_b%denominator_b != 0 {
		return 0
	}

	b := numerator_b / denominator_b

	fmt.Printf("b = %d\n", b)

	fmt.Printf("%v * %d = %v\n", m.a, a, m.a.Mul(a))
	fmt.Printf("%v * %d = %v\n", m.b, b, m.b.Mul(b))
	fmt.Printf("sum = %d\n", m.a.Mul(a).Add(m.b.Mul(b)))

	if m.a.Mul(a).Add(m.b.Mul(b)) != m.prize {
		panic(fmt.Sprintf("%d != %d", m.a.Mul(a).Add(m.b.Mul(b)), m.prize))
	}

	return a*3 + b
}

func parse(inputs []string) ([]machine, error) {
	paras := util.LinesByParagraph(inputs)
	return util.MapE(paras, parseOne)
}

func part1(inputs []string) (int, error) {
	machines, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	return util.MappedSum(machines, presses2), nil
}

func part2(inputs []string) (int, error) {
	add := geom.Vec2{X: 1, Y: 1}.Mul(10000000000000)
	machines, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	total := 0

	for i := range machines {
		machines[i].prize = machines[i].prize.Add(add)
		fmt.Printf("%v\n", machines[i])
		tokens := presses2(machines[i])
		fmt.Printf("  %d: %d\n", i, tokens)
		total += tokens
	}

	return total, nil
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
