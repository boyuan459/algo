package bst

import (
	"fmt"
)

type BinaryNode struct {
	val   int
	left  *BinaryNode
	right *BinaryNode
}

func NewNode(val int) *BinaryNode {
	return &BinaryNode{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (node *BinaryNode) Insert(val int) bool {
	if val < node.val {
		// fmt.Println("insert left", node.left)
		if node.left != nil {
			return node.left.Insert(val)
		}
		node.left = NewNode(val)
		return true
	} else if val > node.val {
		// fmt.Println("insert right", node.right)
		if node.right != nil {
			return node.right.Insert(val)
		}
		node.right = NewNode(val)
		return true
	}

	return false
}

func (node *BinaryNode) Depth() int {
	if node.left == nil && node.right == nil {
		return 0
	}
	var left, right = 0, 0
	if node.left != nil {
		left = 1 + node.left.Depth()
	}
	if node.right != nil {
		right = 1 + node.right.Depth()
	}
	if left > right {
		return left
	}
	return right
}

func (node *BinaryNode) RightSideView(height int, result *[]int) {
	if node == nil {
		return
	}
	if height >= len(*result) {
		*result = append(*result, node.val)
	}
	if node.right != nil {
		node.right.RightSideView(height+1, result)
	}
	if node.left != nil {
		node.left.RightSideView(height+1, result)
	}
}

func (node *BinaryNode) LeftSideView(height int, result *[]int) {
	if node == nil {
		return
	}
	if height >= len(*result) {
		*result = append(*result, node.val)
	}
	if node.left != nil {
		node.left.LeftSideView(height+1, result)
	}
	if node.right != nil {
		node.right.RightSideView(height+1, result)
	}
}

type BST struct {
	root *BinaryNode
}

func New() *BST {
	return &BST{
		root: nil,
	}
}

func (tree *BST) Insert(val int) {
	if tree.root == nil {
		tree.root = NewNode(val)
	} else {
		tree.root.Insert(val)
	}
}

// insert array not working properly
/*
func (tree *BST) InsertArray(values []int) {
	if len(values) == 0 {
		return
	}
	root := NewNode(values[0])
	var count = 1
	queue := make([]BinaryNode, 0)
	queue = append(queue, *root)
	fmt.Println("queue", queue)

	for len(queue) > 0 {
		var node = queue[0]
		queue = queue[1:]
		fmt.Println("queue", queue)
		fmt.Println("current", node)
		node.left = NewNode(values[count])
		count++
		if count >= len(values) {
			return
		}
		if node.left != nil {
			queue = append(queue, *node.left)
		}

		node.right = NewNode(values[count])
		fmt.Println("current node", node)
		count++
		if count >= len(values) {
			return
		}
		if node.right != nil {
			queue = append(queue, *node.right)
		}
		fmt.Println("queue", queue)
	}
	tree.root = root
}
*/

func (tree *BST) preorder(node *BinaryNode, depth int) {
	if node == nil {
		return
	}
	// fmt.Println("node", node)
	tree.preorder(node.left, depth+1)
	var space = make([]byte, 0)
	for i := 0; i < depth; i++ {
		space = append(space, '.')
	}
	fmt.Printf("%s%d\n", string(space[:]), node.val)
	// fmt.Println(node.val)
	tree.preorder(node.right, depth+1)
}

func (tree *BST) PreOrderTraverse() {
	tree.preorder(tree.root, 0)
}

func (tree *BST) Depth() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.Depth()
}

func (tree *BST) LevelOrder() [][]int {
	if tree.root == nil {
		return [][]int{}
	}
	levels := make([][]int, 0)
	queue := make([]BinaryNode, 0)

	queue = append(queue, *tree.root)

	for len(queue) > 0 {
		// iterate each level once
		level := make([]int, 0)
		total := len(queue)
		count := 0
		// after iteration, the queue will hold children of next level
		for count < total {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.val)
			count++
			if node.left != nil {
				queue = append(queue, *node.left)
			}
			if node.right != nil {
				queue = append(queue, *node.right)
			}
		}
		levels = append(levels, level)
	}
	return levels
}

func (tree *BST) Height() int {
	return tree.Depth() + 1
}

func (tree *BST) RightSideView() []int {
	if tree.root == nil {
		return []int{}
	}
	views := make([]int, 0)
	tree.root.RightSideView(0, &views)
	return views
}

func (tree *BST) LeftSideView() []int {
	if tree.root == nil {
		return []int{}
	}
	views := make([]int, 0)
	tree.root.LeftSideView(0, &views)
	return views
}
