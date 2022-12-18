package maps

import "golang.org/x/exp/constraints"

// MapMapValues applies a function to each value in a map, yielding a map with the new values (and type).
func MapMapValues[K comparable, V1, V2 any](m map[K]V1, f func(K, V1) V2) map[K]V2 {
	result := make(map[K]V2, len(m))
	for k, v := range m {
		result[k] = f(k, v)
	}
	return result
}

// MapMax returns the maximum value from a map.
func MapMax[K comparable, V constraints.Ordered](m map[K]V) V {
	var max V
	first := true
	for _, v := range m {
		if first {
			max = v
			first = false
			continue
		}
		if v > max {
			max = v
		}
	}
	return max
}

// Sum returns the sum of values the function returns from a map.
func Sum[K comparable, V1 any, V2 constraints.Integer | constraints.Float | constraints.Complex](m map[K]V1, f func(V1) V2) V2 {
	var sum V2
	for _, v := range m {
		sum += f(v)
	}
	return sum
}

// MapFilter returns a new map, containing only those values for which the filter function
// returns true.
func MapFilter[K comparable, V any](m map[K]V, f func(V) bool) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		if f(v) {
			result[k] = v
		}
	}
	return result
}
