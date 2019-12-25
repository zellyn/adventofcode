package math

import (
	"testing"
)

func TestModExp(t *testing.T) {
	testdata := []struct {
		base     int
		exponent int
		n        int
		want     int
	}{
		{
			base:     2,
			exponent: 5,
			n:        100,
			want:     32,
		},
		{
			base:     -2,
			exponent: 5,
			n:        100,
			want:     -32,
		},
		{
			base:     2,
			exponent: 5,
			n:        10,
			want:     2,
		},
		{
			base:     7,
			exponent: 5,
			n:        100000,
			want:     16807,
		},
		{
			base:     7,
			exponent: 5,
			n:        10000,
			want:     6807,
		},
		{
			base:     7,
			exponent: 5,
			n:        16808,
			want:     16807,
		},
		{
			base:     7,
			exponent: 5,
			n:        16806,
			want:     1,
		},
		{
			base:     7,
			exponent: 5,
			n:        10,
			want:     7,
		},
		{
			base:     17463217478,
			exponent: 5,
			n:        100000000007,
			want:     70601994169,
		},
	}

	for i, tt := range testdata {
		got := ModExp(tt.base, tt.exponent, tt.n)
		if got != tt.want {
			t.Errorf("%d: want ModExp(%d,%d,%d)=%d; got %d", i, tt.base, tt.exponent, tt.n, tt.want, got)
		}
	}
}

func TestModGeometricSum(t *testing.T) {
	testdata := []struct {
		r    int
		n    int
		m    int
		want int
	}{
		{
			r:    2,
			n:    3,
			m:    10,
			want: 7,
		},
		{
			r:    37,
			n:    5,
			m:    10000001,
			want: 1926221,
		},
		{
			r:    37,
			n:    5,
			m:    101,
			want: 50,
		},
	}

	for i, tt := range testdata {
		got, err := ModGeometricSum(tt.r, tt.n, tt.m)
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%d: want ModGeometricSum(%d,%d,%d)=%d; got %d", i, tt.r, tt.n, tt.m, tt.want, got)
		}
	}
}
func TestModInv(t *testing.T) {
	testdata := []struct {
		x    int
		n    int
		want int
		err  bool
	}{
		{
			x:    7,
			n:    10,
			want: 3,
		},
		{
			x:    2018,
			n:    10007,
			want: 8078,
		},
	}

	for i, tt := range testdata {
		got, err := ModInv(tt.x, tt.n)
		if (err != nil) && !tt.err {
			t.Errorf("%d: ModInv(%d,%d): got unexpected error: %v", i, tt.x, tt.n, err)
		}
		if (err == nil) && tt.err {
			t.Errorf("%d: ModInv(%d,%d): error expected, but not found", i, tt.x, tt.n)
		}
		if !tt.err && got != tt.want {
			t.Errorf("%d: want ModInv(%d,%d)=%d; got %d", i, tt.x, tt.n, tt.want, got)
		}
	}
}
