package bin_tree

import "testing"

func TestBinTreeDepth(t *testing.T) {
	values := []int{3, 9, 20, -1, -1, 15, 7}
	tree := New()
	tree.Insert(values)
	tree.Print()
	var depth = tree.Depth()
	if depth != 3 {
		t.Errorf("expected 3, but got")
	}
}
