package fun

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
