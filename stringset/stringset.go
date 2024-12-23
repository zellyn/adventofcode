package stringset

import (
	"bytes"
	"maps"
	"sort"

	"github.com/zellyn/adventofcode/util"
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

// String returns a clen string representation of the set.
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

// Intersect returns a new set representing strings that are in both a and b.
func Intersect(a, b S) S {
	if len(a) <= len(b) {
		res := make(S, len(a))
		for item := range a {
			if b[item] {
				res[item] = true
			}
		}
		return res
	}

	res := make(S, len(b))
	for item := range b {
		if a[item] {
			res[item] = true
		}
	}
	return res
}

// Clone returns a clone of the set.
func (s S) Clone() S {
	return maps.Clone(s)
}

// Keys returns a sorted slice of all the strings in the set.
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

// ContainsAll returns true if this set has every key in the other one.
func (s S) ContainsAll(other S) bool {
	for item := range other {
		if !s[item] {
			return false
		}
	}
	return true
}

// ContainsAny returns true if this set has any key in the other one.
func (s S) ContainsAny(other S) bool {
	if len(s) <= len(other) {
		for item := range s {
			if other[item] {
				return true
			}
		}
		return false
	}

	for item := range other {
		if s[item] {
			return true
		}
	}
	return false
}

// ClonePlus returns a new set with the given item included.
func (s S) ClonePlus(item string) S {
	return util.SetPlus(s, item)
}

// Union returns a new set containing all items either in this or the other set or both.
func (s S) Union(other S) S {
	res := make(S)

	maps.Copy(res, s)
	maps.Copy(res, other)

	return res
}

// Intersection returns a new set containing all items in both this and the other set.
func (s S) Intersection(other S) S {
	return Intersect(s, other)
}
