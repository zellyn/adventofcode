package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/graph"
	"github.com/zellyn/adventofcode/ioutil"
)

func positions(s, substr string) []int {
	var res []int
	l := len(substr)
	for i := 0; i <= len(s)-l; i++ {
		if s[i:i+l] == substr {
			res = append(res, i)
		}
	}
	return res
}

func replaceAt(s, substr, repl string, pos int) string {
	return s[:pos] + repl + s[pos+len(substr):]
}

func parseReplacements(lines []string) (map[string][]string, error) {
	result := map[string][]string{}

	for _, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("weird input: %q", line)
		}
		result[parts[0]] = append(result[parts[0]], parts[1])
	}

	return result, nil
}

func parseInput(filename string) (map[string][]string, string, error) {
	lines, err := ioutil.ReadLines(filename)
	if err != nil {
		return nil, "", err
	}
	if lines[len(lines)-2] != "" {
		return nil, "", fmt.Errorf(`weird second-last line. Want "", got %q`, lines[len(lines)-2])
	}
	start := lines[len(lines)-1]
	lines = lines[:len(lines)-2]
	repls, err := parseReplacements(lines)
	if err != nil {
		return nil, "", err
	}
	return repls, start, nil
}

func inverseReplacements(repls map[string][]string, target string) map[string][]string {
	res := map[string][]string{}
	unused := setDiff(rightOnly(repls), atomSet(target))

	for k, vv := range repls {
		for _, v := range vv {
			if len(setIntersect(unused, atomSet(v))) > 0 {
				continue
			}
			res[v] = append(res[v], k)
		}
	}

	return res
}

func productions(start string, repls map[string][]string) map[string]bool {
	result := map[string]bool{}

	for k, rs := range repls {
		pps := positions(start, k)
		for _, pp := range pps {
			for _, r := range rs {
				result[replaceAt(start, k, r, pp)] = true
			}
		}
	}
	return result
}

func distinctMolecules(filename string) (int, error) {
	repls, start, err := parseInput(filename)
	if err != nil {
		return 0, err
	}

	return len(productions(start, repls)), nil
}

func count(s string) int {
	count := 0
	for _, r := range s {
		if r <= 'Z' {
			count++
		}
	}
	return count
}

func atomSet(s string) map[string]bool {
	atoms := map[string]bool{}

	for i, r := range s {
		if r <= 'Z' {
			if i < len(s)-1 && s[i+1] > 'Z' {
				atoms[s[i:i+2]] = true
			} else {
				atoms[s[i:i+1]] = true
			}
		}
	}
	return atoms
}

func atomCounts(s string) map[string]int {
	atoms := map[string]int{}

	for i, r := range s {
		if r <= 'Z' {
			if i < len(s)-1 && s[i+1] > 'Z' {
				atoms[s[i:i+2]]++
			} else {
				atoms[s[i:i+1]]++
			}
		}
	}
	return atoms
}

func fewestInStages(repls map[string][]string, startingValue string, goal string, debug bool) (int, error) {
	seen := map[string]bool{}
	todo := map[string]bool{
		startingValue: true,
	}
	steps := 0
	for len(todo) > 0 {
		steps++
		if debug {
			fmt.Printf("iteration %d: len(todo)=%d\n", steps, len(todo))
		}
		next := map[string]bool{}

		maxCount := -1

		i := 0
		for t := range todo {
			if debug {
				i++
				if i%100000 == 0 {
					fmt.Printf(" ...%d: len(next)=%d, len(seen)=%d\n", i, len(next), len(seen))
				}
			}
			for p := range productions(t, repls) {
				if p == goal {
					return steps, nil
				}
				if len(p) > 1 && strings.Contains(p, "e") {
					continue
				}
				if seen[p] {
					continue
				}
				seen[p] = true
				c := count(p)
				if c > maxCount {
					maxCount = c
				}
				next[p] = true
			}
		}

		for k := range seen {
			if count(k) > maxCount {
				delete(seen, k)
			}
		}

		todo = next
	}
	return 0, fmt.Errorf("no solution found")
}

type metadata struct {
	repls map[string][]string
	goal  string
}

type node struct {
	value string
	md    *metadata
}

var _ graph.Node = node{}

func (n node) End() bool {
	return n.value == n.md.goal
}

func (n node) String() string {
	return n.value
}

func (n node) Neighbors() []graph.CostedNode {
	prods := productions(n.value, n.md.repls)
	res := make([]graph.CostedNode, 0, len(prods))
	for prod := range prods {
		if n.md.goal == "e" && strings.Contains(prod, "e") && prod != "e" {
			continue
		}
		res = append(res, graph.CostedNode{
			N: node{
				value: prod,
				md:    n.md,
			},
			Steps: 1,
		})
	}
	return res
}

func fewestSteps(repls map[string][]string, startingValue string, goal string, debug bool) (int, error) {
	start := node{
		value: startingValue,
		md: &metadata{
			repls: repls,
			goal:  goal,
		},
	}
	return graph.DijkstraDebug(start, debug)
}

func rightOnly(repls map[string][]string) map[string]bool {
	result := map[string]bool{}

	for _, prods := range repls {
		for _, prod := range prods {
			atoms := atomSet(prod)
			for atom := range atoms {
				if repls[atom] == nil {
					result[atom] = true
				}
			}
		}
	}

	return result
}

func setDiff(main, subtracted map[string]bool) map[string]bool {
	r := make(map[string]bool, len(main))
	for k := range main {
		if !subtracted[k] {
			r[k] = true
		}
	}
	return r
}

func setIntersect(a, b map[string]bool) map[string]bool {
	r := make(map[string]bool)
	for k := range a {
		if b[k] {
			r[k] = true
		}
	}
	return r
}

func setAdd(a, b map[string]bool) map[string]bool {
	r := make(map[string]bool, len(a))
	for k := range a {
		r[k] = true
	}
	for k := range b {
		r[k] = true
	}
	return r
}

func insideRnAr(inverse map[string][]string) map[string]bool {
	result := map[string]bool{}

	for prod := range inverse {
		i := strings.Index(prod, "Rn")
		if i == -1 {
			continue
		}
		j := strings.Index(prod, "Ar")
		if j < i {
			panic("weird Rn string without Ar")
		}
		result[prod[i+2:j]] = true
	}

	return result
}

func outerRnArs(s string) [][2]int {
	var result [][2]int
	start := -1
	depth := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+2] == "Ar" {
			if depth == 0 {
				start = i
			}
			depth++
		}
		if s[i:i+2] == "Rn" {
			depth--
			if depth == 0 {
				result = append(result, [2]int{start, i})
			}
			if depth < 0 {
				depth = 0
			}
		}
	}
	return result
}

func stepsToReduce(repls map[string][]string, startingValue string) (int, error) {
	steps := count(startingValue) - 1

	cs := atomCounts(startingValue)
	if cs["Rn"] != cs["Ar"] {
		return 0, fmt.Errorf("Got %d 'Rn' atoms != %d 'Ar' atoms", cs["Rn"], cs["Ar"])
	}

	steps -= 2 * cs["Rn"]
	steps -= 2 * cs["Y"]
	return steps, nil
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
