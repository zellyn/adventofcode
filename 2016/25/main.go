package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/2016/assembunny"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func stateToString(s *assembunny.State) string {
	return fmt.Sprintf("%v %d", s.Regs, s.Outputs[0])
}

func part1(inputs []string) (int, error) {
	s, err := assembunny.Parse(inputs)
	if err != nil {
		return 0, err
	}
OUTER:
	for a := 1; ; a++ {
		printf("\n%d: ", a)
		seen := make(map[string]bool)
		s.Reset()
		s.SetRegister("a", a)
		want := 0

		for {
			output, done := s.StepUntilOutput()
			if done {
				continue OUTER
			}
			printf(" %d", output)
			if output != want {
				continue OUTER
			}
			want = 1 - want
			state := stateToString(s)
			if seen[state] {
				return a, nil
			}
			seen[state] = true
		}
	}
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
