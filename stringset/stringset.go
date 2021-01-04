package stringset

import (
	"bytes"
	"sort"
)

type S map[string]bool

// New returns a new stringset with the given entries.
func New(entries ...string) S {
	result := make(map[string]bool, len(entries))
	for _, entry := range entries {
		result[entry] = true
	}
	return result
}

// OfRunes takes a string, and returns a stringset built by treating each rune
// as a separate entry.
func OfRunes(runes string) S {
	result := make(map[string]bool)
	for _, rune := range runes {
		result[string(rune)] = true
	}
	return result
}

// AddAll adds all the entries in `other` to this set.
func (s S) AddAll(other S) {
	for k := range other {
		s[k] = true
	}
}

func (s S) String() string {
	var buf bytes.Buffer

	buf.WriteRune('{')
	first := true
	for k := range s {
		if first {
			first = false
		} else {
			buf.WriteRune(',')
		}
		buf.WriteString(k)
	}
	buf.WriteRune('}')

	return buf.String()
}

func Intersect(a, b S) S {
	result := make(map[string]bool)

	for k := range a {
		if b[k] {
			result[k] = true
		}
	}

	return result
}

func (s S) Copy() S {
	result := make(map[string]bool, len(s))
	for k, v := range s {
		result[k] = v
	}
	return result
}

func (s S) Keys() []string {
	result := make([]string, 0, len(s))
	for k, v := range s {
		if v {
			result = append(result, k)
		}
	}
	sort.Strings(result)
	return result
}
