package graph

import "testing"

func TestGraph(t *testing.T) {
	graph := New(8)
	graph.Build([]int{2, 2, 4, 6, -1, 4, 4, 5})
}
