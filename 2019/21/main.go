package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/2019/intcode"
)

func makeInputs() [][]string {
	end := []string{
		"NOT A T", // if there's a hole next…
		"OR T J",  //  …you gotta jump.
		"AND D J", // but not directly into a hole
	}

	bcOptions := [][]string{
		/* 0011 - B */
		/* 0101 - C */

		/* 0000 */ {"Never jump"}, // Never jump
		/* 1000 */ {"^B,^C", "NOT B T", "NOT C J", "AND T J"}, // Jump if ^B,^C
		/* 0100 */ {"^B,C", "NOT B J", "AND C J"}, // Jump if ^B,C
		/* 0010 */ {"B,^C", "NOT C J", "AND B J"}, // Jump if B,^C
		/* 0001 */ {"B,C", "OR B J", "AND C J"}, // Jump if B,C
		/* 1100 */ {"^B", "NOT B J"}, // J=^B
		/* 1010 */ {"^C", "NOT C J"}, // J=^C
		/* 1001 */ {"B=C", "NOT B T", "AND C T", "NOT C J", "AND B J", "AND T J", "NOT J J"}, // B=C
		/* 0110 */ {"B<>C", "NOT B T", "AND C T", "NOT C J", "AND B J", "AND T J"}, // B<>C
		/* 0101 */ {"C", "OR C J"}, // J=C
		/* 0011 */ {"B", "OR B J"}, // J=B
		/* 1110 */ {"! B,C", "OR B J", "AND C J", "NOT J J"}, // Unless B,C
		/* 1101 */ {"! B,^C", "NOT C J", "AND B J", "NOT J J"}, // Unless B,^C
		/* 1011 */ {"! ^B,C", "NOT B J", "AND C J", "NOT J J"}, // Unless ^B,C
		/* 0111 */ {"! ^B,^C", "NOT B T", "NOT C J", "AND T J", "NOT J J"}, // Unless ^B,^C
		/* 1111 */ {"Always jump", "NOT J J"}, // Always jump
	}

	var result [][]string
	for _, bcOption := range bcOptions {
		var res []string
		res = append(res, bcOption...)
		res = append(res, end...)
		result = append(result, res)
	}
	return result
}

/*

  X   X
#####.##.########
   ABCDEFGHI
   OR
   X   X
#####.##.########
    ABCDEFGHI


  X   X
#####.#..########
   ABCDEFGHI

  X   X   X   X
#####.#.#.##..###
   ABCDEFGHI

   X   X
#####..#.########
	ABCDEFGHI

  X   X   X
#####.##.##...###
   ABCDEFGHI
       ABCDEFGHI

   	X   X   X
#####.#.#..##.###
   ABCDEFGHI
	 ABCDEFGHI

   X   X   X
#####..##.##.####
	ABCDEFGHI
	    ABCDEFGHI


Failed ground: "#####.##.########"
Failed ground: "#####.##.##...###"
Failed ground: "#####.#.#..##.###"
Failed ground: "#####.#..########"
Failed ground: "#####..##.##.####"
Failed ground: "#####..#.########"
Missing:        #####.#.#.##..###

*/

func makeInputs2() [][]string {
	end := []string{
		"NOT A T", // if there's a hole next…
		"OR T J",  //  …you gotta jump.
		"AND D J", // but not directly into a hole
	}

	bcOptions := [][]string{
		// {"D,H", "OR D J", "AND H J", "NOT C T", "OR T J"}, // Jump if ^B,^C
		// {"not C", "NOT C J"},
		// {"not C or ^B,^G", "NOT B T", "NOT G J", "AND J T", "NOT C J", "OR T J"},
		// {"not C or ^B,^G, not ^G^H", "NOT B T", "NOT G J", "AND J T", "NOT C J", "OR T J", "OR G T", "AND G T", "OR H T", "AND T J"},
		// {"not C, EorH", "NOT C T", "OR E J", "OR H J", "AND T J"},
		// {"not C, EorH, ^B^E^F^G", "NOT C T", "OR E J", "OR H J", "AND T J", "OR B T", "AND B T", "OR E T", "OR F T", "OR G T", "NOT T T", "OR T J"},
		{"not C, EorH, ^B^E^F^G", "NOT C T", "OR E J", "OR H J", "AND T J", "OR B T", "AND B T", "OR E T", "NOT T T", "OR T J"},
	}

	var result [][]string
	for _, bcOption := range bcOptions {
		var res []string
		res = append(res, bcOption...)
		res = append(res, end...)
		result = append(result, res)
	}
	return result
}

func runProgram(program []int64, boolInputs []string, run bool, debug bool) (int, string, error) {
	startString := "\nWALK\n"

	if run {
		startString = "\nRUN\n"
	}

	name := boolInputs[0]
	boolInputs = boolInputs[1:]
	boolInput := strings.Join(boolInputs, "\n") + startString
	var reads []int64
	for _, r := range boolInput {
		reads = append(reads, int64(r))
	}

	_, writes, err := intcode.RunProgram(program, reads, false)
	if err != nil {
		return 0, "", err
	}
	if writes[len(writes)-1] < 256 {
		if debug {
			fmt.Println("=================================")
			fmt.Println(name)
			fmt.Println(strings.Join(boolInputs, ", "))
			fmt.Println("---------------------------------")
			fmt.Print(boolInput)
			fmt.Println("---------------------------------")
		}
		s := ""
		for _, i := range writes {
			s += string(rune(i))
		}
		if debug {
			fmt.Printf("Output: [%s]\n", s)

		}
		i := strings.Index(s, "Didn't make it across:")
		return 0, s[i+78 : i+95], nil
	}
	if debug {
		fmt.Printf("%s\n%s\n", name, strings.Join(boolInputs, ", "))
	}
	return int(writes[len(writes)-1]), "", nil
}

func run() error {
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return err
	}
	nonZeroScore := 0
	for _, boolInputs := range makeInputs() {
		score, ground, err := runProgram(program, boolInputs, false, false)
		if err != nil {
			return err
		}
		if score > 0 {
			fmt.Println(score)
			nonZeroScore = score
		}
		_ = ground
	}
	fmt.Println("Part 1:", nonZeroScore)

	nonZeroScore = 0
	for _, boolInputs := range makeInputs2() {
		score, ground, err := runProgram(program, boolInputs, true, true)
		if err != nil {
			return err
		}
		if score > 0 {
			fmt.Println(score)
			nonZeroScore = score
		} else {
			fmt.Printf("Failed ground: %q\n", ground)
		}
	}
	fmt.Println("Part 2:", nonZeroScore)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
