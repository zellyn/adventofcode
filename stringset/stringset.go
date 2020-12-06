package stringset

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

func Intersect(a, b S) S {
	result := make(map[string]bool)

	for k := range a {
		if b[k] {
			result[k] = true
		}
	}

	return result
}
