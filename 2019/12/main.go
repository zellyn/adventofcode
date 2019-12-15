package main

import (
	"github.com/zellyn/adventofcode/geom"
)

type vec3 = geom.Vec3

func step(pos, vel []vec3) ([]vec3, []vec3) {
	for i := range pos {
		for j := i + 1; j < len(pos); j++ {
			diff := pos[i].Add(pos[j].Neg()).Sgn()
			vel[j] = vel[j].Add(diff)
			vel[i] = vel[i].Add(diff.Neg())
		}
	}
	for i := range pos {
		pos[i] = pos[i].Add(vel[i])
	}
	return pos, vel
}

func energy(pos, vel []vec3) int {
	e := 0
	for i := 0; i < len(pos); i++ {
		e += pos[i].AbsSum() * vel[i].AbsSum()
	}
	return e
}

func step1(d *[8]int) {
	if d[0] > d[1] {
		d[4]--
		d[5]++
	} else if d[0] < d[1] {
		d[4]++
		d[5]--
	}
	if d[0] > d[2] {
		d[4]--
		d[6]++
	} else if d[0] < d[2] {
		d[4]++
		d[6]--
	}
	if d[0] > d[3] {
		d[4]--
		d[7]++
	} else if d[0] < d[3] {
		d[4]++
		d[7]--
	}
	if d[1] > d[2] {
		d[5]--
		d[6]++
	} else if d[1] < d[2] {
		d[5]++
		d[6]--
	}
	if d[1] > d[3] {
		d[5]--
		d[7]++
	} else if d[1] < d[3] {
		d[5]++
		d[7]--
	}
	if d[2] > d[3] {
		d[6]--
		d[7]++
	} else if d[2] < d[3] {
		d[6]++
		d[7]--
	}

	d[0] += d[4]
	d[1] += d[5]
	d[2] += d[6]
	d[3] += d[7]

}

func period(pts []int, limit int) int {
	var first [8]int
	var all [8]int
	copy(all[:4], pts)
	copy(first[:], all[:])
	for i := 1; i < limit; i++ {
		step1(&all)
		if all == first {
			return i
		}
	}
	return 0
}

func factors(n int) map[int]int {
	result := make(map[int]int)
	for i := 2; ; i++ {
		if n == 1 {
			return result
		}
		if n%i == 0 {
			count := 0
			for n%i == 0 {
				count++
				n /= i
			}
			result[i] = count
		}
	}
}

func multmaps(a, b, c map[int]int) int {
	ks := make(map[int]int)
	for k, v := range a {
		ks[k] = v
	}
	for k, v := range b {
		if ks[k] < v {
			ks[k] = v
		}
	}
	for k, v := range c {
		if ks[k] < v {
			ks[k] = v
		}
	}
	result := 1
	for k, v := range ks {
		for i := 0; i < v; i++ {
			result *= k
		}
	}
	return result
}
