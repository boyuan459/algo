package graph

import "testing"

func TestGraph1D(t *testing.T) {
	graph := New(8)
	graph.BuildWith1D([]int{2, 2, 4, 6, -1, 4, 4, 5})
}

func TestGraph2D(t *testing.T) {
	graph := New(6)
	graph.BuildWith2D([][]int{{0, 1}, {1, 2}, {5, 2}, {3, 0}, {3, 4}, {5, 3}, {5, 4}})
	t.Log("graph", graph)
	var acyclic = graph.AcyclicTopologicalSort()
	t.Log("graph", graph)
	if !acyclic {
		t.Errorf("expected true, got %v", acyclic)
	}
}
