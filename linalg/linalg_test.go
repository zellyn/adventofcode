package linalg

import (
	"testing"
)

func TestReduction(t *testing.T) {

	testdata := []struct {
		name           string
		rows           [][]int
		wantRows       [][]int
		wantRank       int
		wantImpossible bool
	}{
		{
			name: "2x2 already in lowest form",
			rows: [][]int{
				{1, 0, 1},
				{0, 1, 1},
			},
			wantRows: [][]int{
				{1, 0, 1},
				{0, 1, 1},
			},
			wantRank:       2,
			wantImpossible: false,
		},
		{
			name: "3x3 identity matrix",
			rows: [][]int{
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},
			},
			wantRows: [][]int{
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},
			},
			wantRank:       3,
			wantImpossible: false,
		},
		{
			name: "3x3 out-of-order identity matrix",
			rows: [][]int{
				{0, 1, 0, 1},
				{0, 0, 1, 1},
				{1, 0, 0, 1},
			},
			wantRows: [][]int{
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},
			},
			wantRank:       3,
			wantImpossible: false,
		},
		{
			name: "3x3 out-of-order identity matrix, except with rows multiplied by constants",
			rows: [][]int{
				{0, 5, 0, 5},
				{0, 0, 42, 42},
				{1337, 0, 0, 1337},
			},
			wantRows: [][]int{
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 0, 1, 1},
			},
			wantRank:       3,
			wantImpossible: false,
		},
		{
			name: "3x3 incrementing integers",
			rows: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			wantRows: [][]int{
				{1, 0, -1, -2},
				{0, 1, 2, 3},
				{0, 0, 0, 0},
			},
			wantRank:       2,
			wantImpossible: false,
		},
		{
			name: "3x3 incrementing integers, with extra (non-redundant) row",
			rows: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{9, 8, 6, 7},
			},
			wantRows: [][]int{
				{1, 0, 0, -3},
				{0, 1, 0, 5},
				{0, 0, 1, -1},
				{0, 0, 0, 0},
			},
			wantRank:       3,
			wantImpossible: false,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMatrix(tt.rows)
			// fmt.Println(m.Printable("  ", true))
			m.reduce()
			if gotImpossible := m.Impossible(); gotImpossible != tt.wantImpossible {
				t.Fatalf("want m.Impossible()=%v; got %v", tt.wantImpossible, gotImpossible)
			}
			if tt.wantImpossible {
				return
			}

			if gotRank := m.Rank(); gotRank != tt.wantRank {
				t.Errorf("want rank=%d; got %d", tt.wantRank, gotRank)
			}

			wantM := NewMatrix(tt.wantRows)
			if !m.Equal(wantM) {
				t.Errorf("Want reduced matrix to be:\n%s\n but was:\n%s\n", wantM.Printable("  ", true), m.Printable("  ", true))
			}
		})
	}
}
