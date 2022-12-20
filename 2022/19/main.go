package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/lists"
)

type blueprint struct {
	id    int
	costs [4][3]int
}

type robotcost struct {
	index int
	cost  [3]int
}

const template = "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."

func parseBluePrint(input string) blueprint {
	b := blueprint{}
	_, err := fmt.Sscanf(input, template, &b.id, &b.costs[0][0], &b.costs[1][0], &b.costs[2][0], &b.costs[2][1], &b.costs[3][0], &b.costs[3][2])
	if err != nil {
		panic(fmt.Sprintf("Weird input: %q: %v", input, err))
	}
	return b
}

func (b blueprint) String() string {
	return fmt.Sprintf(template, b.id, b.costs[0][0], b.costs[1][0], b.costs[2][0], b.costs[2][1], b.costs[3][0], b.costs[3][2])
}

func parseBlueprints(inputs []string) []blueprint {
	return lists.Map(inputs, parseBluePrint)
}

type state struct {
	blueprints []blueprint
	blueprint  int
	ores       [4]int
	robots     [4]int
	left       int
	best       *int
	// bestPath   *string
	// path       string
}

// End implements graph.Node
func (_ state) End() bool {
	return false
}

// Key implements graph.Node
func (s state) Key() [9]int {
	return [9]int{s.left, s.ores[0], s.ores[1], s.ores[2], s.ores[3], s.robots[0], s.robots[1], s.robots[2], s.robots[3]}
}

// Neighbors implements graph.Node
func (s state) Neighbors() []state {
	max := s.ores[3] + s.left*(s.left-1)/2 + s.left*s.robots[3]
	if max <= *s.best {
		return nil
	}

	if s.left == 0 {
		if s.ores[3] > *s.best {
			*s.best = s.ores[3]
			// *s.bestPath = s.path
			fmt.Println("best:", *s.best)
		}
		return nil
	}

	// debug := s.ores == [4]int{2, 0, 0, 0}
	debug := false
	if debug {
		fmt.Printf("time left: %d\n", s.left)
		fmt.Printf(" ores:%v\n", s.ores)
	}

	// Decrement time remaining
	s.left--

	// One neighbor where we build nothing
	var states []state
	// var summaries []string

	states = append(states, s)
	// summaries = append(summaries, "_")

	// Determine which robots we can build.
	skipCount := 0
OUTER:
	for i := 0; i < 4; i++ {
		costs := s.blueprints[s.blueprint].costs
		if i != 3 {
			max := 0
			for j := i; j < 4; j++ {
				if costs[j][i] > max {
					max = costs[j][i]
				}
			}
			if max <= s.robots[i] {
				skipCount++
				continue OUTER
			}
		}
		cost := costs[i]
		for j, need := range cost {
			if s.ores[j] < need {
				if debug {
					fmt.Printf("  Cannot build recipe %d\n", i)
				}
				continue OUTER
			}
		}
		if debug {
			fmt.Printf(" We can build robot %d. Costs=%v, Ores=%v, Robots=%v\n", i, s.blueprints[s.blueprint].costs, s.ores, s.robots)
		}
		newState := s
		newState.robots[i]++
		for j, need := range cost {
			newState.ores[j] -= need
		}
		if debug {
			fmt.Printf(" We built robot %d. Costs=%v, Ores=%v, Robots=%v\n", i, s.blueprints[s.blueprint].costs, newState.ores, newState.robots)
		}
		states = append(states, newState)
		// summaries = append(summaries, fmt.Sprintf("b%d", i))
	}

	// Don't sit around if we could build anything we need.
	if len(states) == 5 || len(states) == 5-skipCount {
		states = states[1:]
	}

	// Update ores and build result
	for i := range states {

		if debug {
			fmt.Printf(" Before: State:%d Ores=%v, Robots=%v\n", i, states[i].ores, states[i].robots)
			// fmt.Printf("         Path: %q\n", states[i].path)
		}

		for j, num := range s.robots {
			states[i].ores[j] += num
		}

		// states[i].path = states[i].path + fmt.Sprintf(" %d:%s%v%v", states[i].left, summaries[i], states[i].ores, states[i].robots)

		if debug {
			fmt.Printf(" After: State:%d Ores=%v, Robots=%v\n", i, states[i].ores, states[i].robots)
			// fmt.Printf("         Path: %q\n", states[i].path)
		}
	}

	return states
}

func Compare(s1, s2 state) (leq, geq bool) {
	leq, geq = true, true

	if s1.left < s2.left {
		geq = false
	} else if s1.left > s2.left {
		leq = false
	}

	for i := 0; i < 4; i++ {
		if s1.ores[i] < s2.ores[i] || s1.robots[i] < s2.robots[i] {
			geq = false
		}
		if s1.ores[i] > s2.ores[i] || s1.robots[i] > s2.robots[i] {
			leq = false
		}

		if !(leq || geq) {
			return
		}
	}

	return

}

func search(start state) {
	todo := []state{start}

	seen := make(map[[9]int]bool)

	count := 0
	for len(todo) > 0 {
		if count%1000000 == 0 {
			fmt.Printf("count: %d, len(todo): %d, len(seen): %d\n", count, len(todo), len(seen))
		}
		count++
		ns := todo[len(todo)-1].Neighbors()

		for j := 0; j < len(ns); {
			key := ns[j].Key()
			if seen[key] {
				copy(ns[j:], ns[j+1:])
				ns = ns[:len(ns)-1]
			} else {
				seen[key] = true
				j++
			}
		}

		if len(ns) == 0 {
			todo = todo[:len(todo)-1]
			continue
		}

		from := 0
		to := 0

		for from < len(todo)-1 {
			// fmt.Println("  from:", from)
			shadowed := false
			for j := 0; j < len(ns); {
				// fmt.Println("  j:", j, " len(ns):", len(ns))
				leq, geq := Compare(todo[from], ns[j])

				if geq {
					// Existing node completely shadows the neighbor. Get rid of the neighbor.
					copy(ns[j:], ns[j+1:])
					ns = ns[:len(ns)-1]
					fmt.Println("-------------- POP! --------------")
				} else {
					// Our neighbor is good.
					j++
					// Is the existing node shadowed?
					if leq {
						// yep
						shadowed = true
					}
				}
			}
			if !shadowed {
				todo[to] = todo[from]
				to++
			}
			from++
		}
		todo = append(todo[:to], ns...)
	}
}

func part1(inputs []string) (int, error) {
	bps := parseBlueprints(inputs)

	total := 0

	for i := range bps {
		var best int
		// var bestPath string

		start := state{
			blueprints: bps,
			blueprint:  i,
			left:       24,
			best:       &best,
			// bestPath:   &bestPath,
		}
		start.robots[0] = 1

		search(start)

		fmt.Println(bps[i])
		fmt.Printf("Best: %d\n", best)
		// fmt.Printf(" Path: %s\n", bestPath)
		total += bps[i].id * best
	}

	return total, nil
}

func part2(inputs []string) (int, error) {
	bps := parseBlueprints(inputs)
	if len(bps) > 3 {
		bps = bps[:3]
	}

	prod := 1

	for i := range bps {
		var best int
		// var bestPath string

		start := state{
			blueprints: bps,
			blueprint:  i,
			left:       32,
			best:       &best,
			// bestPath:   &bestPath,
		}
		start.robots[0] = 1

		search(start)

		fmt.Println(bps[i])
		fmt.Printf("Best: %d\n", best)
		// fmt.Printf(" Path: %s\n", bestPath)
		prod *= best
	}

	return prod, nil
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
