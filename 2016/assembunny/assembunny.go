package assembunny

import (
	"fmt"
	"strconv"
	"strings"
)

type Op struct {
	Name string
	X    string
	Y    string
	Z    string
}

type State struct {
	Ops     []Op
	IP      int
	Regs    [4]int
	Error   error
	Debug   bool
	Outputs []int
}

// Reset the assembunny machine back to its starting state, but
// leaving the Ops alone (if `tgl` instructions have changed them,
// those changes with NOT be undone).
func (s *State) Reset() {
	s.IP = 0
	s.Regs[0] = 0
	s.Regs[1] = 0
	s.Regs[2] = 0
	s.Regs[3] = 0
	s.Error = nil
	s.Outputs = nil
}

// GetRegister returns the value of the given register, a, b, c, or d.
func (s *State) GetRegister(reg string) (int, error) {
	if reg == "a" || reg == "b" || reg == "c" || reg == "d" {
		return s.Regs[int(reg[0]-'a')], nil
	}

	return 0, fmt.Errorf("unknown register %q", reg)
}

// SetRegister sets the value of the given register, a, b, c, or d.
func (s *State) SetRegister(reg string, value int) error {
	if reg == "a" || reg == "b" || reg == "c" || reg == "d" {
		s.Regs[int(reg[0]-'a')] = value
		return nil
	}

	return fmt.Errorf("unknown register %q", reg)
}

func (s *State) getValue(val string) (int, error) {
	if val == "a" || val == "b" || val == "c" || val == "d" {
		return s.Regs[int(val[0]-'a')], nil
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (s *State) debugf(format string, a ...any) (n int, err error) {
	if !s.Debug {
		return
	}
	return fmt.Printf(format, a...)
}

func (s *State) setValue(reg string, val int) error {
	switch reg {
	case "a", "b", "c", "d":
		s.Regs[int(reg[0]-'a')] = val
		return nil
	}
	return fmt.Errorf("unknown register: %q", reg)
}

func (s *State) doSimple(o Op, f func(i int) (string, int)) bool {
	val, err := s.getValue(o.X)
	if err != nil {
		s.Error = err
		return true
	}
	target, newVal := f(val)
	// The newer `tgl` instruction can create meaningless instructions
	// with numeric (non-register) targets, so only write if we can.
	if target[0] >= 'a' && target[0] <= 'z' {
		if err := s.setValue(target, newVal); err != nil {
			s.Error = err
			return true
		}
	}
	s.IP++
	return s.IP >= len(s.Ops)
}

func (s *State) doAdd(o Op, sub bool) bool {
	val1, err := s.getValue(o.X)
	if err != nil {
		s.Error = err
		return true
	}
	val2, err := s.getValue(o.Y)
	if err != nil {
		s.Error = err
		return true
	}

	val := val1 + val2
	if sub {
		val = val1 - val2
	}

	target := o.Z
	if target[0] >= 'a' && target[0] <= 'z' {
		if err := s.setValue(target, val); err != nil {
			s.Error = err
			return true
		}
	}
	s.IP++
	return s.IP >= len(s.Ops)
}

func (s *State) doMul(o Op) bool {
	val1, err := s.getValue(o.X)
	if err != nil {
		s.Error = err
		return true
	}
	val2, err := s.getValue(o.Y)
	if err != nil {
		s.Error = err
		return true
	}

	val := val1 * val2
	target := o.Z
	if target[0] >= 'a' && target[0] <= 'z' {
		if err := s.setValue(target, val); err != nil {
			s.Error = err
			return true
		}
	}
	s.IP++
	return s.IP >= len(s.Ops)
}

func (s *State) transmit(o Op) bool {
	val, err := s.getValue(o.X)
	if err != nil {
		s.Error = err
		return true
	}
	s.Outputs = append(s.Outputs, val)
	s.IP++
	return s.IP >= len(s.Ops)
}

// Step the assembunny computer one instruction forward. Returns true
// when finished, whether successfully or due to an error.
func (s *State) Step() bool {
	s.debugf("Step: IP=%d, Regs=[%d,%d,%d,%d]\n", s.IP, s.Regs[0], s.Regs[1], s.Regs[2], s.Regs[3])
	if s.Error != nil || s.IP < 0 || s.IP >= len(s.Ops) {
		s.debugf(" already have error: doing nothing\n")
		return true
	}

	o := s.Ops[s.IP]
	s.debugf("op: %v\n", o)
	switch o.Name {
	case "nop":
		s.IP++
		return s.IP >= len(s.Ops)
	case "cpy":
		return s.doSimple(o, func(i int) (string, int) {
			return o.Y, i
		})
	case "inc":
		return s.doSimple(o, func(i int) (string, int) {
			return o.X, i + 1
		})
	case "dec":
		return s.doSimple(o, func(i int) (string, int) {
			return o.X, i - 1
		})
	case "jnz":
		val1, err := s.getValue(o.X)
		if err != nil {
			s.Error = err
			return true
		}
		if val1 == 0 {
			s.IP++
		} else {
			val2, err := s.getValue(o.Y)
			if err != nil {
				s.Error = err
				return true
			}
			s.IP += val2
		}
		return s.IP < 0 || s.IP >= len(s.Ops)
	case "tgl":
		offset, err := s.getValue(o.X)
		if err != nil {
			s.Error = err
			return true
		}
		addr := s.IP + offset
		s.debugf(" toggle %d\n", addr)
		if addr >= 0 && addr < len(s.Ops) {
			before := s.Ops[addr].Name
			switch s.Ops[addr].Name {
			case "inc":
				s.Ops[addr].Name = "dec"
			case "dec", "tgl":
				s.Ops[addr].Name = "inc"
			case "jnz":
				s.Ops[addr].Name = "cpy"
			case "cpy":
				s.Ops[addr].Name = "jnz"
			default:
				s.Error = fmt.Errorf("don't know how to toggle instruction %q", s.Ops[addr].Name)
				return true
			}
			s.debugf(" toggled %q to %q\n", before, s.Ops[addr].Name)
		}
		s.IP++
		return false

	case "add":
		return s.doAdd(o, false)

	case "sub":
		return s.doAdd(o, true)

	case "mul":
		return s.doMul(o)

	case "out":
		return s.transmit(o)

	default:
		s.Error = fmt.Errorf("weird op: %q", o.Name)
		return true
	}
}

// StepUntilOutput clears the Output, then steps until `Step()`
// returns `true`, or an output is emitted. If the boolean is true,
// the machine halted before outputting anything. Otherwise the integer holds the output.
func (s *State) StepUntilOutput() (int, bool) {
	s.Outputs = s.Outputs[:0]

	for len(s.Outputs) == 0 {
		done := s.Step()
		if done {
			return 0, true
		}
	}
	return s.Outputs[0], false
}

// Parse an assembunny program.
func Parse(inputs []string) (*State, error) {
	res := &State{}

	for _, input := range inputs {
		parts := strings.Split(input, " ")
		if len(parts) == 2 {
			parts = append(parts, "")
		}
		res.Ops = append(res.Ops, Op{Name: parts[0], X: parts[1], Y: parts[2]})
	}

	return res, nil
}
