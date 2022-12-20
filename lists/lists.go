package lists

// Map applies a function to each value in a slice, yielding a slice with the new values (and type).
func Map[T any, U any](s []T, f func(T) U) []U {
	var result []U = make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter applies a function to each value in a slice, yielding a slice with only those elements where the function returned true.
func Filter[T any](s []T, filter func(T) bool) []T {
	var result []T = make([]T, 0, len(s))
	for _, v := range s {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// Split splits slices on elements where the predicate returns true
func Split[T any](s []T, predicate func(T) bool) [][]T {
	var result [][]T

	var current []T
	for _, item := range s {
		if predicate(item) {
			result = append(result, current)
			current = nil
		} else {
			current = append(current, item)
		}
	}
	result = append(result, current)

	return result
}

// Reverse reverses a slice.
func Reverse[T any](s []T) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
