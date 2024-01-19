package myslices

import (
	"cmp"
	"slices"
)

// Medianish returns the median (for odd-length slices), or the first
// of the two middle entries (for even-length slices).
func Medianish[S ~[]E, E cmp.Ordered](x S) E {
	y := make(S, len(x))
	copy(y, x)
	slices.Sort(y)
	return y[(len(y)-1)/2]
}

// MinMax returns the minimum and maximum of a slice. If the slice has
// length 0, it returns the zero value for both.
func MinMax[S ~[]E, E cmp.Ordered](x S) (E, E) {
	var least, most E
	if len(x) == 0 {
		return least, most
	}
	least, most = x[0], x[0]
	for _, e := range x {
		if e < least {
			least = e
		}
		if e > most {
			most = e
		}
	}

	return least, most
}

// ToSet converts a slice into a map-of-bool where each element in the
// slice maps to true.
func ToSet[K comparable](s []K) map[K]bool {
	set := make(map[K]bool, len(s))
	for _, ss := range s {
		set[ss] = true
	}
	return set
}
