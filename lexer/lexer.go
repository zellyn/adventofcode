// A simple lexer helper.
// Copied from https://github.com/zellyn/gocool/blob/master/parser/lex.go and modified,
// which is copied from http://golang.org/src/pkg/text/template/parse/lex.go and modified.

package lexer

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// Pos represents a byte position in the original input text.
type Pos int

// Lexer holds the state of the scanner.
type Lexer struct {
	Name  string // the name of the input; used only for error reports
	Input string // the string being scanned
	pos   Pos    // current position in the input
	width Pos    // width of last rune read from input
}

// New creates a new Lexer from a String
func NewLexer(input string) *Lexer {
	return &Lexer{
		Input: input,
	}
}

// EOF is returned at end-of-file.
const EOF = -1

// Next returns the Next rune in the input.
func (l *Lexer) Next() rune {
	if int(l.pos) >= len(l.Input) {
		l.width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.Input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// EOF returns true if the lexer is at the end of its input.
func (l *Lexer) EOF() bool {
	r := l.Next()
	l.Backup()
	return r == EOF
}

// Peek returns but does not consume the next rune in the input.
func (l *Lexer) Peek() rune {
	r := l.Next()
	l.Backup()
	return r
}

// Backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) Backup() {
	l.pos -= l.width
}

// Accept consumes the next rune if it's from the valid set.
func (l *Lexer) Accept(valid string) (rune, bool) {
	r := l.Next()
	if strings.IndexRune(valid, r) >= 0 {
		return r, true
	}
	l.Backup()
	return 0, false
}

func (l *Lexer) AcceptOne(want rune) bool {
	r := l.Next()
	if want == r {
		return true
	}
	l.Backup()
	return false
}

// AcceptRun consumes a run of runes from the valid set.
func (l *Lexer) AcceptRun(valid string) string {
	result := ""
	for r := l.Next(); strings.IndexRune(valid, r) >= 0; r = l.Next() {
		result += string(r)
	}
	l.Backup()
	return result
}

// AcceptN returns up to n runes from the input.
func (l *Lexer) AcceptN(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		r := l.Next()
		if r == EOF {
			return s
		}
		s += string(r)
	}
	return s
}

// AcceptNotRun consumes a run of runes not from the invalid set.
func (l *Lexer) AcceptNotRun(invalid string) string {
	result := ""
	for r := l.Next(); r != EOF && strings.IndexRune(invalid, r) < 0; r = l.Next() {
		result += string(r)
	}
	l.Backup()
	return result
}

func (l *Lexer) mark() [2]Pos {
	return [2]Pos{l.pos, l.width}
}

func (l *Lexer) restore(state [2]Pos) {
	l.pos = state[0]
	l.width = state[1]
}

// Rest shows the rest of the input
func (l *Lexer) Rest() string {
	return l.Input[l.pos:]
}

// ParseInt parses and returns an int. It does not do any bounds checking.
func (l *Lexer) ParseInt() (int, bool) {
	state := l.mark()
	negative := l.AcceptOne('-')
	digits := l.AcceptRun("0123456789")
	if digits == "" {
		l.restore(state)
		return 0, false
	}
	i, err := strconv.Atoi(digits)
	if err != nil {
		l.restore(state)
		return 0, false
	}
	if negative {
		return -i, true
	}
	return i, true
}

// Pos returns the current postion.
func (l *Lexer) Pos() int {
	return int(l.pos)
}
