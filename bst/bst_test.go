package bst

import "testing"

func TestBinaryTree(t *testing.T) {
	tree := New()
	// tree.InsertArray([]int{3, 9, 20, 15, 7})
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(3)
	tree.Insert(9)
	t.Log("tree", tree)
}

func TestBSTDepth(t *testing.T) {
	tree := New()
	// tree.InsertArray([]int{3, 9, 20, 15, 7})
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(3)
	tree.Insert(9)
	tree.Insert(18)
	var depth = tree.Depth()
	if depth != 3 {
		t.Errorf("expected 3, got %v", depth)
	}
}

func TestBSTDepthWithArray(t *testing.T) {
	tree := New()
	tree.InsertArray([]int{3, 9, 20, -1, -1, 15, 7})
	var depth = tree.Depth()
	if depth != 3 {
		t.Errorf("expected 3, got %v", depth)
	}
}

func TestBSTLevelOrder(t *testing.T) {
	tree := New()
	// tree.InsertArray([]int{3, 9, 20, 15, 7})
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(3)
	tree.Insert(9)
	tree.Insert(18)
	var levels = tree.LevelOrder()
	if levels[0][0] != 15 {
		t.Errorf("exepcted 15, got %v", levels[0][0])
	}
}
