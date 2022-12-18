package search

type Node interface {
	Win() bool
	Nexts() []Node
}

func Search(n Node) Node {
	if n.Win() {
		return n
	}

	for _, next := range n.Nexts() {
		if result := Search(next); result != nil {
			return result
		}
	}

	return nil
}
