package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/lexer"
)

func debug(l *lexer.Lexer) {
	fmt.Println(l.Input)
	for i := 0; i < l.Pos(); i++ {
		fmt.Printf(" ")
	}
	fmt.Println("^")
}

func expand(s string) (string, error) {
	b := &bytes.Buffer{}
	l := &lexer.Lexer{Input: s}

	for {
		s := l.AcceptNotRun("(")
		b.WriteString(s)
		if !l.AcceptOne('(') {
			break
		}
		length, ok := l.ParseInt()
		if !ok {
			return "", fmt.Errorf("cannot parse length at position %d", l.Pos())
		}
		if !l.AcceptOne('x') {
			return "", fmt.Errorf("cannot find 'x' separator at position %d", l.Pos())
		}
		times, ok := l.ParseInt()
		if !ok {
			return "", fmt.Errorf("cannot parse repeat count at position %d", l.Pos())
		}
		if r := l.Next(); r != ')' {
			l.Backup()
			return "", fmt.Errorf("expected ')' at position %d; got '%c'", l.Pos(), r)
		}
		substr := l.AcceptN(length)
		if len(substr) != length {
			return "", fmt.Errorf("cannot read %d runes at position %d", length, l.Pos())
		}
		for i := 0; i < times; i++ {
			b.WriteString(substr)
		}
	}
	return b.String(), nil
}

func expandedLength(s string) (int, error) {
	l := &lexer.Lexer{Input: s}
	result := 0

	for {
		s := l.AcceptNotRun("(")

		result += len(s)
		if !l.AcceptOne('(') {
			break
		}
		length, ok := l.ParseInt()
		if !ok {
			return 0, fmt.Errorf("cannot parse length at position %d", l.Pos())
		}
		if !l.AcceptOne('x') {
			return 0, fmt.Errorf("cannot find 'x' separator at position %d", l.Pos())
		}
		times, ok := l.ParseInt()
		if !ok {
			return 0, fmt.Errorf("cannot parse repeat count at position %d", l.Pos())
		}
		if r := l.Next(); r != ')' {
			l.Backup()
			return 0, fmt.Errorf("expected ')' at position %d; got '%c'", l.Pos(), r)
		}
		substr := l.AcceptN(length)
		if len(substr) != length {
			return 0, fmt.Errorf("cannot read %d runes at position %d", length, l.Pos())
		}
		expanded, err := expandedLength(substr)
		if err != nil {
			return 0, fmt.Errorf("error expanding substring: %w", err)
		}
		result += times * expanded
	}
	return result, nil
}

func foo(s string) (int, error) {
	return strconv.Atoi(s)
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
