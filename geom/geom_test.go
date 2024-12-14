package geom

import "testing"

func Test90DegreeRotations(t *testing.T) {
	testdata := []struct {
		a, b Vec2
	}{
		{a: N, b: E},
		{a: NE, b: SE},
		{a: E, b: S},
		{a: SE, b: SW},
		{a: S, b: W},
		{a: SW, b: NW},
		{a: W, b: N},
		{a: NW, b: NE},
	}

	for _, tt := range testdata {
		if got := tt.a.Clockwise90(); got != tt.b {
			t.Errorf("want %v.Clockwise90 = %v; got %v", tt.a, tt.b, got)
		}

		if got := tt.b.CounterClockwise90(); got != tt.a {
			t.Errorf("want %v.CounterClockwise90 = %v; got %v", tt.b, tt.a, got)
		}
	}
}
