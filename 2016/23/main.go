package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/2016/assembunny"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func replaceOps(s *assembunny.State) {
	if len(s.Ops) < 16 {
		return
	}

	/*
			 Replace:
		     4	cpy b c
		     5	inc a
		     6	dec c
		     7	jnz c -2
		     8	dec d
		     9	jnz d -5
			 With:
		     4	mul b d c
		     5	add a c a
		     6	cpy 0 c
		     7	cpy 0 d
		     8	nop
		     9	nop
	*/
	s.Ops[4] = assembunny.Op{Name: "mul", X: "b", Y: "d", Z: "c"}
	s.Ops[5] = assembunny.Op{Name: "add", X: "a", Y: "c", Z: "a"}
	s.Ops[6] = assembunny.Op{Name: "cpy", X: "0", Y: "c"}
	s.Ops[7] = assembunny.Op{Name: "cpy", X: "0", Y: "d"}
	s.Ops[8] = assembunny.Op{Name: "nop"}
	s.Ops[9] = assembunny.Op{Name: "nop"}

	/*
		 Replace:
		    13	dec d
		    14	inc c
		    15	jnz d -2
			With:
		    13	add c d c
		    14	cpy 0 d
		    15	nop
	*/
	s.Ops[13] = assembunny.Op{Name: "add", X: "c", Y: "d", Z: "c"}
	s.Ops[14] = assembunny.Op{Name: "cpy", X: "0", Y: "d"}
	s.Ops[15] = assembunny.Op{Name: "nop"}

}

func part1(inputs []string, initialA int) (int, error) {
	s, err := assembunny.Parse(inputs)
	if err != nil {
		return 0, err
	}
	// s.Debug = true
	replaceOps(s)
	if err := s.SetRegister("a", initialA); err != nil {
		return 0, err
	}
	done := false
	for !done {
		done = s.Step()
	}
	if s.Error != nil {
		return 0, s.Error
	}
	return s.GetRegister("a")
	// return -1, nil
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
