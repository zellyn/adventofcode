package interval

import (
	"slices"
	"strconv"
)

// I represents an (inclusive) interval.
type I [2]int

func (i I) String() string {
	return "[" + strconv.Itoa(i[0]) + "," + strconv.Itoa(i[1]) + "]"
}

// New creates a new Interval, given start and end.
func New(start, end int) I {
	return [2]int{start, end}
}

// Contains checks if this interval completely contains the other.
func (i I) Contains(other I) bool {
	return other[0] >= i[0] && other[1] <= i[1]
}

// Overlap checks if this interval overlaps with the other.
func (i I) Overlaps(other I) bool {
	if i[1] < other[0] {
		return false
	}
	if i[0] > other[1] {
		return false
	}
	return true
}

// Adjacent checks if this interval is right next to the other.
func (i I) Adjacent(other I) bool {
	if i[1]+1 == other[0] {
		return true
	}

	if i[0]-1 == other[1] {
		return true
	}
	return false
}

// Merge takes two intervals, and returns one new one that stretches
// from the lowest number in either to the highest number in either.
func (i I) Merge(other I) I {
	res := i

	if other[0] < res[0] {
		res[0] = other[0]
	}

	if other[1] > res[1] {
		res[1] = other[1]
	}

	return res
}

// Sort sorts a slice of intervals by start (and then end, for equal
// starts).
func Sort(intervals []I) {
	slices.SortFunc(intervals, func(a, b I) int {
		aa, bb := a[0], b[0]
		if aa == bb {
			aa, bb = a[1], b[1]
		}
		if aa < bb {
			return -1
		} else if aa > bb {
			return 1
		}
		return 0
	})
}
