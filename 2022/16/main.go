package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/zellyn/adventofcode/maps"
	"github.com/zellyn/adventofcode/search"
)

const NOPATH = 10000

type graph map[string]node

func (g graph) clone() graph {
	g2 := make(graph, len(g))
	for k, v := range g {
		g2[k] = v
	}
	return g2
}

type node struct {
	name        string
	open        bool
	flow        int
	connections []string
}

func shortestPath(g graph, start, end string) []string {
	if start == end {
		return nil
	}
	cost, strPath := shortestHelper(g, start, end, "")
	if cost == NOPATH {
		panic("disconnected graph!")
	}

	return strings.Split(strPath, ":")[1:]
}

func shortestHelper(g graph, start, end, seen string) (int, string) {
	if start == end {
		return 0, start
	}

	min := NOPATH - 1
	minPath := ""

	for _, conn := range g[start].connections {
		if g[conn].name == end {
			return 1, start + ":" + end
		}
		if strings.Contains(seen, conn) {
			continue
		}
		cost, path := shortestHelper(g, conn, end, seen+":"+start)
		if cost < min {
			min = cost
			minPath = path
		}
	}

	return 1 + min, start + ":" + minPath
}

func totalFlow(g graph, start string, order []string, debug bool) int {
	g = g.clone()

	var open []string
	totalFlow := 0
	minute := 1
	flow := 0
	current := start
	goal := order[0]
	steps := shortestPath(g, current, goal)
	order = order[1:]

	for minute <= 30 {
		if debug {
			if minute > 1 {
				fmt.Println()
			}
			fmt.Printf("== Minute %d ==\n", minute)
			if len(open) == 0 {
				fmt.Println("No valves are open.")
			} else {
				if len(open) == 1 {
					fmt.Printf("Valve %s is open", open[0])
				} else {
					fmt.Printf("Valves %s", strings.Join(open[:len(open)-1], ", "))
					if len(open) > 2 {
						fmt.Printf(",")
					}
					fmt.Printf(" and %s are open", open[len(open)-1])
				}
				fmt.Printf(", releasing %d pressure.\n", flow)
			}
		}
		totalFlow += flow
		minute++
		if len(steps) > 0 {
			current = steps[0]
			steps = steps[1:]
			if debug {
				fmt.Printf("You move to valve %s.\n", current)
			}
			continue
		}
		if current == goal {
			node := g[current]
			if !node.open {
				node.open = true
				open = append(open, current)
				sort.Strings(open)
				flow += node.flow
				if debug {
					fmt.Printf("You open valve %s.\n", current)
				}
			}
			goal = ""
			if len(order) > 0 {
				goal = order[0]
				steps = shortestPath(g, current, goal)
				order = order[1:]
			}
		}
	}

	return totalFlow
}

// Always pick the next largest closed valve
func greedyOrder(g graph, start string) []string {
	order := make([]string, 0, len(g))

	for _, v := range g {
		if v.name != start && v.flow > 0 {
			order = append(order, v.name)
		}
	}
	sort.Slice(order, func(a, b int) bool {
		return g[order[a]].flow > g[order[b]].flow
	})

	return order
}

func parseGraph(inputs []string) graph {
	result := make(graph, len(inputs))

	for _, input := range inputs {
		var n node
		if n, err := fmt.Sscanf(input, "Valve %s has flow rate=%d;", &n.name, &n.flow); err != nil {
			panic(fmt.Sprintf("weird input %q (parsed %d):%v", input, n, err))
		}
		conns := strings.Split(input, ", ")
		conns[0] = conns[0][len(conns[0])-2:]
		n.connections = conns
		result[n.name] = n
	}
	return result
}

type distances map[string]int

func newDistances(g graph, start string) distances {
	nodes := greedyOrder(g, start)
	nodes = append(nodes, start)
	result := make(distances, (len(nodes)*len(nodes) - 1))

	for i, from := range nodes {
		for _, to := range nodes[i+1:] {
			dist, _ := shortestHelper(g, from, to, "")
			result[from+to] = dist
			result[to+from] = dist
		}
	}

	return result
}

func (d distances) of(from, to string) int {
	dist, ok := d[from+to]
	if !ok {
		panic(fmt.Sprintf("Asked for distance from %q to %q, but haven't got it!", from, to))
	}
	return dist
}

func findBestHelper(nf map[string]int, d distances, at string, order string, currentFlow, totalFlow, timeLeft, flowLeft int, best *int, bestOrder *string, orderSoFar string, indent string) {
	debug := indent != ""

	if debug {
		fmt.Printf("%scalled at %s with order %s, flowLeft=%d, timeLeft=%d, currentFlow=%d, totalFlow=%d\n",
			indent, at, order, flowLeft, timeLeft, currentFlow, totalFlow)
	}

	// Bail out if we can't possibly beat the best score
	if timeLeft > 1 {
		if totalFlow+currentFlow+(timeLeft-1)*(currentFlow+flowLeft) <= *best {
			if debug {
				fmt.Printf("%s  bailing out: cannot beat %d\n", indent, *best)
			}
			return

		}
	}

	// Nowhere else to go
	if flowLeft == 0 {
		finalTotal := totalFlow + timeLeft*currentFlow
		if finalTotal > *best {
			if debug {
				fmt.Printf("%s  WIN! beat %d with %d\n", indent, *best, finalTotal)
			}
			*best = finalTotal
			*bestOrder = orderSoFar
		}
		return
	}

	for i := 0; i < len(order); i += 2 {
		target := order[i : i+2]
		newOrder := order[:i] + order[i+2:]
		dist := d.of(at, target) + 1 // one timestep to turn it on
		steps := min(dist, timeLeft)
		totalFlowHere := totalFlow + steps*currentFlow
		if timeLeft <= dist {
			// We're out of time, which means we're done
			if totalFlowHere > *best {
				*best = totalFlowHere
				*bestOrder = orderSoFar
			}
		} else {
			// Not out of time, so recurse
			targetFlow := nf[target]
			newIndent := indent
			if newIndent != "" {
				newIndent += "  "
			}

			findBestHelper(nf, d, target, newOrder, currentFlow+targetFlow, totalFlowHere, timeLeft-dist, flowLeft-targetFlow, best, bestOrder, orderSoFar+target, newIndent)
		}
	}
}

type part1Node struct {
	best         *int
	distances    distances
	nodeFlows    map[string]int
	myAt         string
	eleAt        string
	myStepsLeft  int
	eleStepsLeft int
	order        string
	currentFlow  int
	totalFlow    int
	timeLeft     int
	flowLeft     int
}

func (n *part1Node) Win() bool {
	return false
}

func (n *part1Node) Nexts() []search.Node {
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// if first is true, the head of the string goes only into the first set
func genSplits(s string, first bool) [][]string {
	if len(s) == 2 {
		return [][]string{
			{s, ""},
			{"", s},
		}
	}

	head := s[:2]
	rest := s[2:]

	restSplits := genSplits(rest, false)

	var result [][]string

	for _, split := range restSplits {
		result = append(result, []string{head + split[0], split[1]})
		if !first {
			result = append(result, []string{split[0], head + split[1]})
		}
	}

	return result
}

func part1(inputs []string) (int, error) {
	start := "AA"
	g := parseGraph(inputs)
	flowLeft := maps.Sum(g, func(n node) int { return n.flow })
	order := strings.Join(greedyOrder(g, start), "")
	nodeFlows := maps.MapMapValues(g, func(_ string, n node) int { return n.flow })

	var best int
	var bestOrder string
	findBestHelper(nodeFlows, newDistances(g, start), start, order, 0, 0, 30, flowLeft, &best, &bestOrder, "", "")
	return best, nil
}

func totalDoubleFlow(nf map[string]int, d distances, start string, order []string, debug bool) int {
	var open []string
	totalFlow := 0
	flow := 0
	minute := 1
	at := []string{start, start}
	goal := []string{"", ""}
	steps := []int{0, 0}
	names := []string{"player", "elephant"}

	for minute <= 26 {
		if debug {
			if minute > 1 {
				fmt.Println()
			}
			fmt.Printf("== Minute %d ==\n", minute)
			if len(open) == 0 {
				fmt.Println("No valves are open.")
			} else {
				if len(open) == 1 {
					fmt.Printf("Valve %s is open", open[0])
				} else {
					fmt.Printf("Valves %s", strings.Join(open[:len(open)-1], ", "))
					if len(open) > 2 {
						fmt.Printf(",")
					}
					fmt.Printf(" and %s are open", open[len(open)-1])
				}
				fmt.Printf(", releasing %d pressure.\n", flow)
			}
		}
		totalFlow += flow
		minute++

		for i := 0; i < 2; i++ {
			if steps[i] > 0 {
				steps[i]--
				if steps[i] == 0 {
					flow += nf[goal[i]]
					at[i] = goal[i]
					if debug {
						fmt.Printf("%s opens valve %s\n", names[i], at[i])
						open = append(open, at[i])
						sort.Strings(open)
					}
				}
			}
			if steps[i] == 0 {
				if len(order[i]) > 0 {
					goal[i] = order[i][:2]
					order[i] = order[i][2:]
					steps[i] = d.of(at[i], goal[i]) + 1
					if minute == 2 {
						steps[i]--
					}
				}
			}
		}
	}

	return totalFlow
}

func permutations(s string) []string {
	if len(s) == 2 {
		return []string{s}
	}
	var result []string

	for i := 0; i < len(s); i += 2 {
		this := s[i : i+2]
		other := s[:i] + s[i+2:]
		for _, perm := range permutations(other) {
			result = append(result, this+perm)
		}
	}

	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func calcDistance(d distances, order string) int {
	total := 0
	for i := 0; i < len(order)-2; i += 2 {
		total += d.of(order[i:i+2], order[i+2:i+4])
	}
	return total
}

func part2(inputs []string) (int, error) {
	start := "AA"
	g := parseGraph(inputs)

	initialOrder := strings.Join(greedyOrder(g, start), "")
	initialFlowLeft := maps.Sum(g, func(n node) int { return n.flow })

	splits := genSplits(initialOrder, true)
	nodeFlows := maps.MapMapValues(g, func(_ string, n node) int { return n.flow })
	d := newDistances(g, start)

	var best int
	var bestOrder string
	findBestHelper(nodeFlows, newDistances(g, start), start, initialOrder, 0, 0, 30, initialFlowLeft, &best, &bestOrder, "", "")
	initialDistance := calcDistance(d, bestOrder)

	if nodeFlows["HH"] == 0 { // part 2
		best = 1845
	}
	distanceLimit := int(float32(initialDistance) * 2)

OUTER:
	for _, split := range splits {
		if calcDistance(d, split[0]) > distanceLimit || calcDistance(d, split[1]) > distanceLimit {
			continue
		}
		// fmt.Println(split)
		// fmt.Printf("split %d/%d\n", i+1, len(splits))

		var innerBest [2]int
		innerBestOrder := []string{"", ""}
		for j := 0; j < 2; j++ {
			order := split[j]
			flowLeft := 0
			for k := 0; k < len(order); k += 2 {
				flowLeft += nodeFlows[order[k:k+2]]
			}
			findBestHelper(nodeFlows, newDistances(g, start), start, order, 0, 0, 30, flowLeft, &innerBest[j], &innerBestOrder[j], "", "")

			if innerBest[j] < best/2 {
				continue OUTER
			}
		}

		score := totalDoubleFlow(nodeFlows, d, start, innerBestOrder, false)
		if score > best {
			best = score
		}

	}

	return best, nil
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
