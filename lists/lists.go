package lists

// Map applies a function to each value in a slice, yielding a slice with the new values (and type).
func Map[T any, U any](m []T, f func(T) U) []U {
	var result []U = make([]U, len(m))
	for i, v := range m {
		result[i] = f(v)
	}
	return result
}

// Filter applies a function to each value in a slice, yielding a slice with only those elements where the function returned true.
func Filter[T any](m []T, f func(T) bool) []T {
	var result []T = make([]T, 0, len(m))
	for _, v := range m {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
