package bin_tree

import (
	"testing"
)

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

func TestCompleteBinTreeHeight(t *testing.T) {
	tree := New()
	tree.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	height := tree.Height()

	if height != 3 {
		t.Errorf("expected 3, but got %v", height)
	}
}

func TestCompleteBinTreeCountNodes(t *testing.T) {
	tree := New()
	// tree.Insert([]int{1, 2, 3, 4, 5, 6})
	tree.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	count := tree.CountNodes()
	if count != 9 {
		t.Errorf("expected 9, but got %v", count)
	}
}
