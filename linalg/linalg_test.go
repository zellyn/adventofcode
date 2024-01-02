package linalg

import (
	"reflect"
	"testing"
)

func TestReduction(t *testing.T) {

	testdata := []struct {
		name            string
		rows            [][]int
		wantRows        [][]int
		wantRank        int
		wantImpossible  bool
		wantValues      []float64
		wantValuesKnown []bool
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
			wantRank:        2,
			wantImpossible:  false,
			wantValues:      []float64{1, 1},
			wantValuesKnown: []bool{true, true},
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
			wantRank:        3,
			wantImpossible:  false,
			wantValues:      []float64{1, 1, 1},
			wantValuesKnown: []bool{true, true, true},
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
			wantRank:        3,
			wantImpossible:  false,
			wantValues:      []float64{1, 1, 1},
			wantValuesKnown: []bool{true, true, true},
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
			wantRank:        3,
			wantImpossible:  false,
			wantValues:      []float64{1, 1, 1},
			wantValuesKnown: []bool{true, true, true},
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
			wantRank:        2,
			wantImpossible:  false,
			wantValuesKnown: []bool{false, false, false},
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
			wantRank:        3,
			wantImpossible:  false,
			wantValues:      []float64{-3, 5, -1},
			wantValuesKnown: []bool{true, true, true},
		},
		{
			name: "2x2 from parallel, non-intersecting lines",
			rows: [][]int{
				{2, 4, 5},
				{3, 6, 1},
				{5, 10, -4},
			},
			wantImpossible: true,
		},
		{
			name: "2x2 from parallel, colinear lines",
			rows: [][]int{
				{1, -2, 1},
				{2, -4, 2},
				{3, -6, 3},
			},
			wantRows: [][]int{
				{1, -2, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
			wantRank:        1,
			wantImpossible:  false,
			wantValuesKnown: []bool{false, false},
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

			if tt.wantValues != nil || tt.wantValuesKnown != nil {
				gotValues, gotValuesKnown := m.KnownCoefficientFloats()

				if tt.wantValuesKnown != nil {
					if !reflect.DeepEqual(tt.wantValuesKnown, gotValuesKnown) {
						t.Errorf("Want known coefficient flags to be %v; got %v", tt.wantValuesKnown, gotValuesKnown)
					}
				}

				if tt.wantValues != nil {
					if !reflect.DeepEqual(tt.wantValues, gotValues) {
						t.Errorf("Want known coefficient values to be %v; got %v", tt.wantValues, gotValues)
					}
				}
			}
		})
	}
}
