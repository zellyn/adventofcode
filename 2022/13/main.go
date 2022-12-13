package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/lexer"
	"github.com/zellyn/adventofcode/lists"
	"github.com/zellyn/adventofcode/util"
)

type listOrInt struct {
	list []*listOrInt
	i    *int
}

func (li *listOrInt) isInt() bool {
	return li.i != nil
}

func (li *listOrInt) String() string {
	if li.isInt() {
		return strconv.Itoa(*li.i)
	}

	return "[" + strings.Join(lists.Map(li.list, func(li *listOrInt) string { return li.String() }), ",") + "]"
}

func ofInt(i int) *listOrInt {
	return &listOrInt{i: &i}
}

func ofList(lis ...*listOrInt) *listOrInt {
	return &listOrInt{list: lis}
}

func parseList(lx *lexer.Lexer) *listOrInt {
	// mt.Printf("   parseList called at %q\n", lx.Rest())

	if ok := lx.AcceptOne('['); !ok {
		panic(fmt.Sprintf("Expecting '['; but found %c at position %d of %q", lx.Peek(), lx.Pos(), lx.Input))
	}

	var result []*listOrInt

	for {
		switch r := lx.Peek(); {
		case r == '[':
			result = append(result, parseList(lx))
		case r >= '0' && r <= '9':
			i, _ := lx.ParseInt()
			result = append(result, ofInt(i))
		case r == ']' && len(result) == 0:
			lx.Next()
			return &listOrInt{}
		default:
			panic(fmt.Sprintf("Weird input '%c' at %d: %q (wanted list or int)", r, lx.Pos(), lx.Input))
		}

		r := lx.Next()
		if r == ']' {
			break
		}
		if r == ',' {
			continue
		}
		panic(fmt.Sprintf("Weird input '%c' at %d: %q (wanted end of list or comma)", r, lx.Pos(), lx.Input))
	}

	return &listOrInt{list: result}
}

func parseNested(s string) *listOrInt {
	lx := lexer.NewLexer(s)

	return parseList(lx)
}

func sgn(i int) int {
	if i < 0 {
		return -1
	}
	if i > 0 {
		return 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cmp(a, b *listOrInt) int {
	if a.isInt() {
		if b.isInt() {
			return sgn(*a.i - *b.i)
		}
		return cmp(ofList(a), b)
	}
	if b.isInt() {
		return cmp(a, ofList(b))
	}

	// both lists
	limit := min(len(a.list), len(b.list))
	for i := 0; i < limit; i++ {
		if c := cmp(a.list[i], b.list[i]); c != 0 {
			return c
		}
	}
	if len(a.list) < len(b.list) {
		return -1
	}
	if len(a.list) > len(b.list) {
		return 1
	}
	return 0
}

func part1(inputs []string) (int, error) {
	pairs := util.SplitAfter(inputs, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})

	sum := 0

	for i, pair := range pairs {
		first := parseNested(pair[0])
		second := parseNested(pair[1])
		c := cmp(first, second)
		if c == 0 {
			panic(fmt.Sprintf("unexpectedly equal lists: %s and %s", first, second))
		}
		if c < 0 {
			sum += i + 1
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	inputs = lists.Filter(inputs, func(s string) bool { return strings.TrimSpace(s) != "" })
	parsed := lists.Map(inputs, parseNested)
	two := parseNested("[[2]]")
	six := parseNested("[[6]]")
	parsed = append(parsed, two, six)
	sort.Slice(parsed, func(a, b int) bool {
		return cmp(parsed[a], parsed[b]) < 0
	})
	var i1, i2 int

	for i, li := range parsed {
		if cmp(two, li) == 0 {
			i1 = i + 1
		} else if cmp(six, li) == 0 {
			i2 = i + 1
		}
	}

	return i1 * i2, nil
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
