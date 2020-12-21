package intcode

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zellyn/adventofcode/util"
)

// ReadProgram reads a file with a single intCode program.
func ReadProgram(filename string) ([]int64, error) {
	bb, err := util.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseProgram(string(bb))

}

// ParseProgram parses an intcode program to a slice of ints.
func ParseProgram(commaString string) ([]int64, error) {
	input := strings.TrimSpace(commaString)
	entries := strings.Split(input, ",")
	ints := make([]int64, len(entries))
	for i, v := range entries {
		intVal, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ints[i] = intVal
	}
	return ints, nil
}

// OpFunc is the function type for running an opcode.
type OpFunc func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error)

// OpDef defines a single opcode.
type OpDef struct {
	name    string
	code    int
	inputs  int
	outputs int
	reads   int
	writes  int
	fn      OpFunc
}

// Show displays the given OpCode with the given arguments.
func (op OpDef) Show(source []int64) (string, error) {
	if len(source) != 1+op.inputs+op.outputs {
		return "", fmt.Errorf("op %s should be %d instructions long; got %d: %v", op.name, 1+op.inputs+op.outputs, len(source), source)
	}
	result := op.name + " "
	for i, input := range source[1:] {
		if i > 0 {
			result += ","
		}
		mode := getMode(source[0], i)
		switch mode {
		case 0:
			result += fmt.Sprintf("[%d]", input)
		case 1:
			result += fmt.Sprintf("%d", input)
		case 2:
			result += fmt.Sprintf("[BASE+%d]", input)
		default:
			return "", fmt.Errorf("unknown mode (%d) for param %d, op=%s", mode, i, op.name)
		}
	}
	return result, nil
}

var ops = map[int64]OpDef{
	1: {
		name:    "ADD",
		code:    1,
		inputs:  2,
		outputs: 1,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			return pc + 4, 0, []int64{inputs[0] + inputs[1]}, nil, nil
		},
	},
	2: {
		name:    "MUL",
		code:    2,
		inputs:  2,
		outputs: 1,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			return pc + 4, 0, []int64{inputs[0] * inputs[1]}, nil, nil
		},
	},
	3: {
		name:    "IN",
		code:    3,
		reads:   1,
		outputs: 1,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			return pc + 2, 0, []int64{reads[0]}, nil, nil
		},
	},
	4: {
		name:   "OUT",
		code:   4,
		inputs: 1,
		writes: 1,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			// fmt.Printf("writing %d\n", inputs[0])
			return pc + 2, 0, nil, []int64{inputs[0]}, nil
		},
	},
	5: {
		name:   "JNZ",
		code:   5,
		inputs: 2,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			if inputs[0] != 0 {
				return int(inputs[1]), 0, nil, nil, nil
			}
			return pc + 3, 0, nil, nil, nil
		},
	},
	6: {
		name:   "JZ",
		code:   6,
		inputs: 2,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			if inputs[0] == 0 {
				return int(inputs[1]), 0, nil, nil, nil
			}
			return pc + 3, 0, nil, nil, nil
		},
	},
	7: {
		name:    "LT",
		code:    7,
		inputs:  2,
		outputs: 1,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			if inputs[0] < inputs[1] {
				return pc + 4, 0, []int64{1}, nil, nil
			}
			return pc + 4, 0, []int64{0}, nil, nil
		},
	},
	8: {
		name:    "EQ",
		code:    8,
		inputs:  2,
		outputs: 1,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			if inputs[0] == inputs[1] {
				return pc + 4, 0, []int64{1}, nil, nil
			}
			return pc + 4, 0, []int64{0}, nil, nil
		},
	},
	9: {
		name:    "ARB",
		code:    9,
		inputs:  1,
		outputs: 0,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			return pc + 2, inputs[0], nil, nil, nil
		},
	},
	99: {
		name: "HCF",
		code: 99,
		fn: func(pc int, inputs []int64, reads []int64) (newPc int, baseAdj int64, outputs []int64, writes []int64, err error) {
			return pc, 0, nil, nil, nil
		},
	},
}

func getOp(opCodePlusModes int64) (OpDef, error) {
	opCode := opCodePlusModes % 100
	op, ok := ops[opCode]
	if !ok {
		return OpDef{}, fmt.Errorf("unknown opcode %d", opCode)
	}
	return op, nil
}

func getMode(rawOp int64, i int) int {
	for j := 0; j < i+2; j++ {
		rawOp = rawOp / 10
	}
	return int(rawOp % 10)
}

// RunProgram runs an intcode program, modifying the state in place, and using
// slices for reads and writes.
func RunProgram(originalState []int64, reads []int64, debug bool) (state []int64, writes []int64, err error) {
	state = Copy(originalState)

	readsChan := make(chan int64, len(reads))
	writesChan := make(chan int64)
	errChan := make(chan error)
	for _, read := range reads {
		readsChan <- read
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go RunProgramChans(ctx, state, readsChan, writesChan, errChan, nil, debug, "")
	for {
		select {
		case err := <-errChan:
			close(readsChan)
			close(writesChan)
			return state, writes, err
		case write := <-writesChan:
			writes = append(writes, write)
		}
	}
}

func ensureSize(state []int64, index64 int64) []int64 {
	index := int(index64)
	if len(state) > index {
		return state
	}
	zeros := make([]int64, index-len(state)+1)
	return append(state, zeros...)
}

// RunProgramChans runs an intcode program, modifying the state in place, and using
// channels for input and output.
func RunProgramChans(ctx context.Context, originalState []int64, reads <-chan int64, writes chan<- int64, errs chan<- error, updates <-chan [2]int64, debug bool, debugPrefix string) {
	state := make([]int64, len(originalState))
	copy(state, originalState)

	logErr := func(err error) {
		errs <- err
		close(errs)
	}
	pc := 0
	var base int64
	for {
		select {
		case <-ctx.Done():
			logErr(nil)
			return
		case update := <-updates:
			state[update[0]] = update[1]
		default:
			// do nothing
		}

		rawOp := state[pc]
		op, err := getOp(rawOp)
		if err != nil {
			logErr(err)
			return
		}
		// fmt.Printf("Op %s at pc=%d\n", op.name, pc)
		if debug {
			show, err := op.Show(state[pc : pc+1+op.inputs+op.outputs])
			if err != nil {
				logErr(err)
				return
			}
			fmt.Printf("%s: %d: %s\n", debugPrefix, pc, show)
		}
		if rawOp == 99 {
			logErr(nil)
			return
		}
		opInputs := make([]int64, 0, op.inputs)
		for i := 0; i < op.inputs; i++ {
			input := state[pc+1+i]
			switch getMode(state[pc], i) {
			case 0:
				state = ensureSize(state, input)
				input = state[input]
			case 1:
				// do nothing
			case 2:
				state = ensureSize(state, base+input)
				input = state[base+input]
			default:
				logErr(fmt.Errorf("unknown mode (%d) for param %d at pc=%d, op=%s", getMode(state[pc], i), i, pc, op.name))
				return
			}
			opInputs = append(opInputs, input)
			// fmt.Printf("opInputs=%v\n", opInputs)
		}
		for i := 0; i < op.outputs; i++ {
			mode := getMode(state[pc], op.inputs+i)
			if mode != 0 && mode != 2 {
				logErr(fmt.Errorf("op %q at pc=%d declares mode of %d for output %d; want mode 0 or 2", op.name, pc, mode, i))
				return
			}
		}
		var opReads []int64
		if op.reads > 0 {
			opReads = make([]int64, op.reads)
			for i := range opReads {
				select {
				case opReads[i] = <-reads:
					// nothing
				case <-ctx.Done():
					logErr(fmt.Errorf("op %q at pc=%d timed out waiting for read %d/%d", op.name, pc, i+1, op.reads))
					return
				}
			}
		}
		newPc, baseAdj, outputs, opWrites, err := op.fn(pc, opInputs, opReads)
		base += baseAdj
		if err != nil {
			logErr(err)
			return
		}
		if len(outputs) != op.outputs {
			logErr(fmt.Errorf("op %q at pc=%d declares %d outputs; gave %d", op.name, pc, op.outputs, len(outputs)))
			return
		}
		if len(opWrites) != op.writes {
			logErr(fmt.Errorf("op %q at pc=%d declares %d writes; gave %d", op.name, pc, op.writes, len(opWrites)))
			return
		}
		for i, write := range opWrites {
			select {
			case writes <- write:
				// nothing
			case <-ctx.Done():
				logErr(fmt.Errorf("op %q at pc=%d timed out waiting to write %d/%d", op.name, pc, i+1, op.writes))
				return
			}
		}
		// fmt.Printf("outputs=%v\n", outputs)
		for i, output := range outputs {
			mode := getMode(state[pc], op.inputs+i)
			loc := state[pc+1+op.inputs+i]
			if mode == 2 {
				loc = base + loc
			}
			// fmt.Printf("Writing %d to %d\n", output, loc)
			state = ensureSize(state, loc)
			state[loc] = output
		}
		pc = newPc
	}
}

// Copy will copy an intcode program.
func Copy(program []int64) []int64 {
	c := make([]int64, len(program))
	copy(c, program)
	return c
}

// ToAscii converts a slice of int64s to a string.
func ToAscii(ints []int64) string {
	s := ""
	for _, i := range ints {
		s += string(rune(i))
	}
	return s
}

// FromAscii converts a string to a slice of int64s.
func FromAscii(s string) []int64 {
	var ints []int64
	for _, r := range s {
		ints = append(ints, int64(r))
	}
	return ints
}
