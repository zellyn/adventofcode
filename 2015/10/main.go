package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

func iterate(s string) string {
	groups := util.GroupString(s)
	result := ""

	for _, g := range groups {
		result += fmt.Sprintf("%d%c", len(g), g[0])
	}
	return result
}

func iterate2(s string) string {
	var buf bytes.Buffer

	last := rune(s[0])
	count := 1
	for _, r := range s[1:] {
		if r == last {
			count++
		} else {
			fmt.Fprintf(&buf, "%d%c", count, last)
			last = r
			count = 1
		}
	}
	fmt.Fprintf(&buf, "%d%c", count, last)

	return buf.String()
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
