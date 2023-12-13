package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

var printf = func(string, ...any) {}

// var printf = fmt.Printf

const (
	spaces = "                                                        "
	hashes = "########################################################"
	dots   = "........................................................"
)

type row struct {
	input        string
	template     string
	extra        int
	lengths      []int
	positions    [][]int
	previousHash map[int]int
	lastHash     int
	memo         map[int]int
}

func newRow(input string, template string, lengths []int) row {
	r := row{input: input, template: template, lengths: lengths}

	minLength := 0
	for _, length := range lengths {
		minLength += length
		var positions []int
	NEXT_POS:
		for position := 0; position <= len(template)-length; position++ {
			if position > 0 && template[position-1] == '#' {
				continue
			}
			if position+length < len(r.template) && r.template[position+length] == '#' {
				continue
			}
			for j := 0; j < length; j++ {
				if template[position+j] == '.' {
					continue NEXT_POS
				}
			}
			positions = append(positions, position)
		}
		r.positions = append(r.positions, positions)
	}

	r.previousHash = make(map[int]int, len(template))
	r.memo = make(map[int]int)
	r.lastHash = -1
	last := -1
	for i, char := range template {
		r.previousHash[i] = last
		if char == '#' {
			last = i
			r.lastHash = i
		}
	}

	minLength += len(lengths) - 1 // spaces
	r.extra = len(template) - minLength
	return r
}

func (r row) String() string {
	pieces := []string{r.template}
	for i, length := range r.lengths {
		pieces = append(pieces, fmt.Sprintf("%d:%v", length, r.positions[i]))
	}
	return "[" + strings.Join(pieces, " ") + "]"
}

func iprintf(indent int, debug bool, format string, args ...any) {
	if !debug {
		return
	}
	printf(spaces[:indent*2]+format, args...)
}

func (r row) combinationsHelper(minPosition int, run int, counter int, debug bool) int {
	key := minPosition<<16 + run
	if result, ok := r.memo[key]; ok {
		return result
	}

	iprintf(counter, debug, "combinationsHelper(%d, %v)\n", minPosition, r.positions[run:])
	sum := 0

	length := r.lengths[run]
	poses := r.positions[run]

	if minPosition > poses[len(poses)-1] {
		iprintf(counter, debug, "→0  # minPosition(%d) > last position(%d)\n", minPosition, poses[len(poses)-1])
		r.memo[key] = 0
		return 0
	}

	for _, pos := range poses {
		if pos < minPosition {
			iprintf(counter+1, debug, "skipping %d\n", pos)
			continue
		}
		if len(r.positions) == run+1 && pos+length <= r.lastHash {
			iprintf(counter+1, debug, "won't cover last hash(%d) at position %d; continue\n", r.lastHash, pos)
			continue
		}
		if r.previousHash[pos] >= minPosition {
			iprintf(counter+1, debug, "passed # at position %d; break\n", r.previousHash[pos])
			break
		}
		iprintf(counter+1, debug, "trying %d\n", pos)
		c := 1
		if len(r.positions) > run+1 {
			c = r.combinationsHelper(pos+length+1, run+1, counter+2, debug)
		}
		sum += c
	}

	iprintf(counter, debug, "→%d\n", sum)
	r.memo[key] = sum
	return sum
}

func (r row) combinations() int {
	// printf("%s\n", r)
	c := r.combinationsHelper(0, 0, 0, false)
	printf("%q: %d\n", r.input, c)
	return c
}

func (r row) valid(pattern string) bool {
	for i, patternRune := range pattern {
		templateRune := rune(r.template[i])
		if templateRune == '?' {
			continue
		}
		if templateRune != patternRune {
			return false
		}
	}
	return true
}

func allPatterns(extra int, lengths []int) []string {
	var result []string
	for i := 0; i <= extra; i++ {
		prefix := dots[:i] + hashes[:lengths[0]]
		if len(lengths) == 1 {
			result = append(result, prefix+dots[:extra-i])
		} else {
			prefix += "."
			for _, rest := range allPatterns(extra-i, lengths[1:]) {
				result = append(result, prefix+rest)
			}
		}
	}
	return result
}

func (r row) uglyCombinations() int {
	sum := 0
	for _, pattern := range allPatterns(r.extra, r.lengths) {
		if r.valid(pattern) {
			sum++
		}
	}
	return sum
}

func parseRowQuintupled(line string) (row, error) {
	template, lengthsString, found := strings.Cut(line, " ")
	if !found {
		return row{}, fmt.Errorf("weird input line: %q", line)
	}
	var templates, lengthsStrings []string
	for i := 0; i < 5; i++ {
		templates = append(templates, template)
		lengthsStrings = append(lengthsStrings, lengthsString)
	}
	return parseRow(strings.Join(templates, "?") + " " + strings.Join(lengthsStrings, ","))
}

func parseRow(line string) (row, error) {
	template, lengthsString, found := strings.Cut(line, " ")
	if !found {
		return row{}, fmt.Errorf("weird input line: %q", line)
	}

	lengths, err := util.ParseInts(lengthsString, ",")
	if err != nil {
		return row{}, fmt.Errorf("malformed input: %q: %w", line, err)
	}

	return newRow(line, template, lengths), err
}

func part1(inputs []string) (int, error) {
	rows, err := util.MapE(inputs, parseRow)
	if err != nil {
		return 0, err
	}
	return util.MappedSum(rows, row.combinations), nil
	// return util.MappedSum(rows, row.uglyCombinations), nil
}

func part2(inputs []string) (int, error) {
	rows, err := util.MapE(inputs, parseRowQuintupled)
	if err != nil {
		return 0, err
	}
	return util.MappedSum(rows, row.combinations), nil
	// return util.MappedSum(rows, row.uglyCombinations), nil
}

func run() error {
	r, err := parseRow("?###???????? 3,2,1")
	if err != nil {
		return err
	}
	printf("%s\n", r.template)
	for _, pattern := range allPatterns(r.extra, r.lengths) {
		printf("%s\n", pattern)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
