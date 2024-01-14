package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var dirs = []string{"U", "D", "L", "R"}

func md5string(s string) string {
	md5Bytes := md5.Sum([]byte(s))
	return hex.EncodeToString(md5Bytes[:])
}

func doors(passcode, path string) string {
	md5Bytes := md5.Sum([]byte(passcode + path))
	dirHex := hex.EncodeToString(md5Bytes[:2])
	res := ""
	for i, c := range dirs {
		if dirHex[i] > 'a' {
			res += c
		}
	}
	return res
}

var area = geom.MakeRectXYs(0, 0, 3, 3)
var target = area.Max

type state struct {
	pos      geom.Vec2
	passcode string
	path     string
}

func (s state) Neighbors() []state {
	res := make([]state, 0, 4)
	ds := doors(s.passcode, s.path)
	for i := range ds {
		d := ds[i : i+1]
		dir := geom.NameToDir[d]
		newPos := s.pos.Add(dir)
		if area.Contains(newPos) {
			res = append(res, state{
				pos:      newPos,
				passcode: s.passcode,
				path:     s.path + d,
			})
		}
	}

	return res
}

func part1(passcode string) (string, error) {
	todo := []state{{passcode: passcode}}

	for len(todo) > 0 {
		this := todo[0]
		todo = todo[1:]
		nn := this.Neighbors()
		for _, n := range nn {
			if n.pos == target {
				return n.path, nil
			}
		}
		todo = append(todo, nn...)
	}

	return "42", nil
}

func part2(passcode string) (int, error) {
	todo := []state{{passcode: passcode}}
	max := 0

	for len(todo) > 0 {
		this := todo[0]
		todo = todo[1:]
		nn := this.Neighbors()
		for _, n := range nn {
			if n.pos == target {
				if len(n.path) > max {
					max = len(n.path)
				}
			} else {
				todo = append(todo, n)
			}
		}
	}

	return max, nil
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
