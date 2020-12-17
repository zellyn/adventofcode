package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/charvol"
	"github.com/zellyn/adventofcode/geom"
)

func step3(v charvol.V) charvol.V {
	vv := make(charvol.V)
	min, max := v.MinMax()

	for x := min.X - 1; x <= max.X+1; x++ {
		for y := min.Y - 1; y <= max.Y+1; y++ {
			for z := min.Z - 1; z <= max.Z+1; z++ {
				here := geom.Vec3{X: x, Y: y, Z: z}
				count := 0
				for _, c := range geom.Neighbors26(here) {
					if v[c] == '#' {
						count++
					}
				}

				switch v[here] {
				case 0, '.':
					if count == 3 {
						vv[here] = '#'
					}
				case '#':
					if count == 2 || count == 3 {
						vv[here] = '#'
					}
				default:
					panic(fmt.Sprintf("weird map element: '%c'", v[here]))
				}
			}
		}
	}

	return vv
}

func step4(v charvol.V4) charvol.V4 {
	vv := make(charvol.V4)
	min, max := v.MinMax()

	for w := min.W - 1; w <= max.W+1; w++ {
		for x := min.X - 1; x <= max.X+1; x++ {
			for y := min.Y - 1; y <= max.Y+1; y++ {
				for z := min.Z - 1; z <= max.Z+1; z++ {
					here := geom.Vec4{W: w, X: x, Y: y, Z: z}
					count := 0
					for _, c := range geom.Neighbors80(here) {
						if v[c] == '#' {
							count++
						}
					}

					switch v[here] {
					case 0, '.':
						if count == 3 {
							vv[here] = '#'
						}
					case '#':
						if count == 2 || count == 3 {
							vv[here] = '#'
						}
					default:
						panic(fmt.Sprintf("weird map element: '%c'", v[here]))
					}
				}
			}
		}
	}

	return vv
}

func part1(input charmap.M) (int, error) {
	vol := charvol.FromCharmap(input, 0)
	for i := 0; i < 6; i++ {
		vol = step3(vol)
	}
	return vol.Count('#'), nil
}

func part2(input charmap.M) (int, error) {
	vol := charvol.FromCharmap4(input, 0, 0)
	for i := 0; i < 6; i++ {
		vol = step4(vol)
	}
	return vol.Count('#'), nil
	return 0, nil
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
