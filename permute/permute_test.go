package permute

import (
	"reflect"
	"testing"
)

func TestFirstsAndRests(t *testing.T) {
	testdata := []struct {
		name  string
		items []int
		want  [][]int
	}{
		{
			name:  "one to five",
			items: []int{1, 2, 3, 4, 5},
			want: [][]int{
				{1, 2, 3, 4, 5},
				{2, 1, 3, 4, 5},
				{3, 1, 2, 4, 5},
				{4, 1, 2, 3, 5},
				{5, 1, 2, 3, 4},
			},
		},
		{
			name:  "one two",
			items: []int{1, 2},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name:  "one",
			items: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			name:  "empty",
			items: []int{},
			want:  [][]int{},
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstsAndRests(tt.items)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want FirstsAndRests(%v) = %v; got %v", tt.items, tt.want, got)
			}
		})
	}
}

func TestPermuteTwoAndRests(t *testing.T) {
	testdata := []struct {
		name  string
		items []int
		want  [][]int
	}{
		{
			name:  "one to five, choose 2",
			items: []int{1, 2, 3, 4, 5},
			want: [][]int{
				{1, 2, 3, 4, 5},
				{1, 3, 2, 4, 5},
				{1, 4, 2, 3, 5},
				{1, 5, 2, 3, 4},

				{2, 1, 3, 4, 5},
				{2, 3, 1, 4, 5},
				{2, 4, 1, 3, 5},
				{2, 5, 1, 3, 4},

				{3, 1, 2, 4, 5},
				{3, 2, 1, 4, 5},
				{3, 4, 1, 2, 5},
				{3, 5, 1, 2, 4},

				{4, 1, 2, 3, 5},
				{4, 2, 1, 3, 5},
				{4, 3, 1, 2, 5},
				{4, 5, 1, 2, 3},

				{5, 1, 2, 3, 4},
				{5, 2, 1, 3, 4},
				{5, 3, 1, 2, 4},
				{5, 4, 1, 2, 3},
			},
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got := PermuteTwoAndRests(tt.items)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want PermuteTwoAndRests(%v) = %v; got %v", tt.items, tt.want, got)
			}
		})
	}
}
