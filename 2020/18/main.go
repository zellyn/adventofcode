package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part1(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		i, err := eval(input, false)
		if err != nil {
			return 0, err
		}
		sum += i
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		i, err := eval(input, true)
		if err != nil {
			return 0, err
		}
		sum += i
	}
	return sum, nil
}

var tokenRe = regexp.MustCompile(`\d+|[()+*]`)

type node struct {
	op    string
	val   int
	left  *node
	right *node
}

func (n *node) String() string {
	switch n.op {
	case "+", "*":
		return fmt.Sprintf("(%s %s %s)", n.left, n.op, n.right)
	case "":
		return strconv.Itoa(n.val)
	default:
		panic(fmt.Sprintf("weird node: %v", *n))

	}
}

func (n *node) eval() int {
	if n == nil {
		panic("nil node")
	}
	switch n.op {
	case "":
		return n.val
	case "+":
		return n.left.eval() + n.right.eval()
	case "*":
		return n.left.eval() * n.right.eval()
	}
	panic(fmt.Sprintf("weird op: %q", n.op))
}

func parse(s *source, prec bool) (*node, error) {
	var nodes []*node
	var ops []string

	for {
		if s.peek() == nil || *s.peek() == ")" {
			return nil, fmt.Errorf("unexpected end of input")
		}

		token := *s.next()

		if token == "(" {
			nn, err := parse(s, prec)
			if err != nil {
				return nil, err
			}
			if s.peek() == nil {
				return nil, fmt.Errorf(`want ")"; got nil`)
			}
			if t := s.next(); *t != ")" {
				return nil, fmt.Errorf(`want ")"; got %q`, *t)
			}
			nodes = append(nodes, nn)
		} else {
			i, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}
			nn := &node{
				val: i,
			}
			nodes = append(nodes, nn)
		}

		if s.peek() == nil || *s.peek() == ")" {
			break
		}

		op := *s.next()

		if op != "+" && op != "*" {
			return nil, fmt.Errorf(`want "*" or "+"; got %q`, op)
		}

		ops = append(ops, op)
	}

	result := nodes[0]

	for i, n := range nodes[1:] {
		op := ops[i]

		if prec && i > 0 {
			if op == "*" || result.op != "*" {
				result = &node{
					left:  result,
					right: n,
					op:    op,
				}
			} else {
				result.right = &node{
					op:    "+",
					left:  result.right,
					right: n,
				}
				result.op = "*"
			}
		} else {
			result = &node{
				left:  result,
				right: n,
				op:    op,
			}
		}
	}

	return result, nil
}

func addNode(result **node, n *node, op string) {
	if *result == nil {
		*result = &node{
			op:   op,
			left: n,
		}
		return
	}

	(*result).right = n
	*result = &node{
		op:   op,
		left: *result,
	}
}

type source struct {
	which  int
	tokens []string
}

func (s *source) next() *string {
	if s.which >= len(s.tokens) {
		return nil
	}
	val := s.tokens[s.which]
	s.which++
	return &val
}

func (s *source) peek() *string {
	if s.which >= len(s.tokens) {
		return nil
	}
	val := s.tokens[s.which]
	return &val
}

func eval(expr string, prec bool) (int, error) {
	tokens := &source{
		which:  0,
		tokens: tokenRe.FindAllString(expr, -1),
	}
	tree, err := parse(tokens, prec)
	if tokens.peek() != nil {
		return 0, fmt.Errorf("unexpected unconsumed input: %q", *tokens.next())
	}
	if err != nil {
		return 0, err
	}
	return tree.eval(), nil
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
