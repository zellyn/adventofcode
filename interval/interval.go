package interval

// I represents an (inclusive) interval.
type I [2]int

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
