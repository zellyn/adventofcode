package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/zellyn/adventofcode/stringset"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type op int

const (
	OP_NONE op = iota
	OP_XOR
	OP_OR
	OP_AND
)

func (o op) String() string {
	switch o {
	case OP_NONE:
		return "[no operator]"
	case OP_XOR:
		return "xor"
	case OP_OR:
		return "or"
	case OP_AND:
		return "and"
	}
	return fmt.Sprintf("[erroneous op value %d]", o)
}

func (o op) GoString() string {
	return o.String()
}

type gate struct {
	op                     op
	input1, input2, output string
}

func (g gate) String() string {
	return fmt.Sprintf("{%s %s %s -> %s}", g.input1, g.op, g.input2, g.output)
}

func fromBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (g gate) eval(input1, input2 int) int {
	if input1 != 0 && input1 != 1 {
		panic(fmt.Sprintf("eval(%d,%d) failed: %d is not a valid wirevalue", input1, input2, input1))
	}
	if input2 != 0 && input2 != 1 {
		panic(fmt.Sprintf("eval(%d,%d) failed: %d is not a valid wirevalue", input1, input2, input2))
	}

	switch g.op {
	case OP_NONE:
		return -1
	case OP_XOR:
		return input1 ^ input2
	case OP_OR:
		return input1 | input2
	case OP_AND:
		return input1 & input2
	}

	panic(fmt.Sprintf("eval(%d,%d) failed with erroneous operator %s", input1, input2, g.op))
}

func eval(wire string, gates []gate, wireValues map[string]int) int {
	if value, ok := wireValues[wire]; ok {
		return value
	}

	for _, g := range gates {
		if g.output == wire {
			input1 := eval(g.input1, gates, wireValues)
			input2 := eval(g.input2, gates, wireValues)
			output := g.eval(input1, input2)
			wireValues[wire] = output
			return output
		}
	}

	return -1
}

var gateRe = regexp.MustCompile(`^([a-z0-9]{3}) (AND|OR|XOR) ([a-z0-9]{3}) -> ([a-z0-9]{3})$`)

func parse(inputs []string) (map[string]int, []gate, []string, error) {
	wires := stringset.New()
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return nil, nil, nil, fmt.Errorf("want 2 paragraphs; got %d", len(paras))
	}
	wireValues := make(map[string]int, len(paras[0]))
	_, err := util.MapE(paras[0], func(s string) (int, error) {
		wire, value, ok := strings.Cut(s, ": ")
		if !ok {
			return 0, fmt.Errorf("weird wire value input: %q", s)
		}
		wires[wire] = true
		switch value {
		case "0":
			wireValues[wire] = 0
		case "1":
			wireValues[wire] = 1
		default:
			panic(fmt.Sprintf("weird input: %q", s))
		}
		return 0, nil
	})
	if err != nil {
		return nil, nil, nil, err
	}

	gates, err := util.MapE(paras[1], func(s string) (gate, error) {
		match := gateRe.FindStringSubmatch(s)
		g := gate{}
		if match == nil {
			return g, fmt.Errorf("weird gate definition input: %q", s)
		}

		g.input1 = match[1]
		g.input2 = match[3]
		g.output = match[4]
		wires[g.input1] = true
		wires[g.input2] = true
		wires[g.output] = true

		switch match[2] {
		case "AND":
			g.op = OP_AND
		case "OR":
			g.op = OP_OR
		case "XOR":
			g.op = OP_XOR
		default:
			return g, fmt.Errorf("weird gate op %q in gate definition %q", match[2], s)
		}

		return g, nil
	})
	if err != nil {
		return nil, nil, nil, err
	}

	return wireValues, gates, wires.Keys(), nil
}

func part1(inputs []string) (int, error) {
	wireValues, gates, wires, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	zWires := util.Filter(wires, func(s string) bool { return strings.HasPrefix(s, "z") })
	slices.Reverse(zWires)

	result := 0

	for _, zWire := range zWires {
		result <<= 1
		val := eval(zWire, gates, wireValues)
		switch val {
		case 1:
			result += 1
		case 0:
		default:
			panic(fmt.Sprintf("weird value %d for wire %q", val, zWire))
		}
	}

	return result, nil
}

func xy(s string) bool {
	return strings.HasPrefix(s, "x") || strings.HasPrefix(s, "y")
}

func notXY(s string) bool {
	return !xy(s)
}

func z(s string) bool {
	return strings.HasPrefix(s, "z")
}

func notZ(s string) bool {
	return !z(s)
}

func any(string) bool {
	return true
}

func mkOp(wantOp op) func(o op) bool {
	return func(o op) bool {
		return o == wantOp
	}
}

func anyOp(op) bool {
	return true
}

func filter(gates []gate, opPred func(op) bool, inputPred, outputPred func(s string) bool) []gate {
	return util.Filter(gates, func(g gate) bool {
		return inputPred(g.input1) && outputPred(g.output) && opPred(g.op)
	})
}

func part2(inputs []string) (string, error) {
	_, gates, _, err := parse(inputs)
	if err != nil {
		return "", nil
	}

	orInputs := stringset.New()
	for _, g := range filter(gates, mkOp(OP_OR), any, any) {
		orInputs[g.input1] = true
		orInputs[g.input2] = true
	}

	allSubs := stringset.New()
	subForZ := stringset.New()
	for _, g := range filter(gates, mkOp(OP_XOR), notXY, notZ) {
		fmt.Printf("Bad output (should be a z gate): %s\n", g)
		subForZ[g.output] = true
	}
	allSubs.AddAll(subForZ)

	subZ := stringset.New()
	for _, g := range filter(gates, anyOp, any, z) {
		if g.output == "z00" || g.output == "z45" {
			continue
		}

		if g.op != OP_XOR {
			fmt.Printf("Z-output gates should be XOR: %s\n", g)
			subZ[g.output] = true
		}
	}
	allSubs.AddAll(subZ)

	otherSubs := stringset.New()
	for _, g := range filter(gates, mkOp(OP_AND), xy, notZ) {
		if g.input1 == "x00" || g.input2 == "x00" {
			continue
		}
		if !orInputs[g.output] {
			fmt.Printf("x & y output should be an input to an OR gate: %s is not: %s\n", g.output, g)
			otherSubs[g.output] = true
		}
	}
	allSubs.AddAll(otherSubs)

	andOutputs := stringset.New()
	for _, g := range filter(gates, mkOp(OP_AND), any, any) {
		andOutputs[g.output] = true
	}

	for _, g := range filter(gates, mkOp(OP_OR), any, any) {
		if !andOutputs[g.input1] && !allSubs[g.input1] {
			fmt.Printf("or gates should have inputs from and gates; %q isn't: %s\n", g.input1, g)
			otherSubs[g.input1] = true
		}
		if !andOutputs[g.input2] && !allSubs[g.input2] {
			fmt.Printf("or gates should have inputs from and gates; %q isn't: %s\n", g.input2, g)
			otherSubs[g.input2] = true
		}
	}
	allSubs.AddAll(otherSubs)

	fmt.Printf("These shuold be switched with Z wires: %v\n", subForZ)
	fmt.Printf("These Z wires should be switched: %v\n", subZ)
	fmt.Printf("These gates should be switched: %v\n", otherSubs)

	swaps := allSubs.Keys()
	slices.Sort(swaps)

	return strings.Join(swaps, ","), nil
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
