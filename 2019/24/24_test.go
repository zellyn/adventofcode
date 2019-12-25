package main

import (
	"testing"
)

func TestNeighbors(t *testing.T) {
	testdata := []struct {
		pos  vec3
		want []vec3
	}{
		{
			pos:  vec3{0, 0, 17},
			want: []vec3{{1, 2, 16}, {0, 1, 17}, {1, 0, 17}, {2, 1, 16}},
		},
		{
			pos:  vec3{0, 4, 17},
			want: []vec3{{1, 2, 16}, {0, 3, 17}, {1, 4, 17}, {2, 3, 16}},
		},
	}

	for _, tt := range testdata {
		got := neighbors(tt.pos)
		wantMap := map[vec3]bool{}
		for _, w := range tt.want {
			wantMap[w] = true
		}
		for _, g := range got {
			if !wantMap[g] {
				t.Errorf("neighbors(%v) should not include %v", tt.pos, g)
			}
			delete(wantMap, g)
		}
		if len(wantMap) > 0 {
			for w := range wantMap {
				t.Errorf("neighbors(%v) should include %v", tt.pos, w)
			}
		}
	}
}
