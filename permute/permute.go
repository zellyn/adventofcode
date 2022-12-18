package permute

// FirstsAndRests takes a slice of items. It returns a slice of slices of items,
// where each item in the original list has a turn at being first, and the
// other items are in the original order.
func FirstsAndRests[T any](items []T) [][]T {
	result := make([][]T, 0, len(items))

	for i, item := range items {
		s := make([]T, len(items))
		s[0] = item
		copy(s[1:], items[:i])
		copy(s[i+1:], items[i+1:])
		result = append(result, s)
	}

	return result
}

func PermuteTwoAndRests[T any](items []T) [][]T {
	size := len(items)
	result := make([][]T, 0, size*(size-1))

	for i := 0; i < size; i++ {
		j := 0
		for jj := 0; jj < size-1; jj++ {
			if j == i {
				j++
			}

			row := make([]T, 2, size)
			row[0] = items[i]
			row[1] = items[j]
			first, second := i, j
			if first > second {
				first, second = second, first
			}
			if first > 0 {
				row = append(row, items[:first]...)
			}
			if second-first > 1 {
				row = append(row, items[first+1:second]...)
			}
			if second < size-1 {
				row = append(row, items[second+1:]...)
			}

			result = append(result, row)

			j++
		}
	}
	return result
}
