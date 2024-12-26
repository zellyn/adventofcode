package pair

type P[T comparable, U comparable] struct {
	A T
	B U
}

// New returns a new pair.
func New[T comparable, U comparable](a T, b U) P[T, U] {
	return P[T, U]{A: a, B: b}
}
