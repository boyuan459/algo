package bin_tree

import (
	"fmt"
	"math"
)

type BinNode struct {
	data  int
	left  *BinNode
	right *BinNode
}

func NewNode(data int) *BinNode {
	return &BinNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

type BinTree struct {
	root *BinNode
}

func New() *BinTree {
	return &BinTree{
		root: nil,
	}
}

func (tree *BinTree) Insert(values []int) *BinNode {
	var length = len(values)
	if length == 0 {
		tree.root = nil
		return tree.root
	}
	queue := []*BinNode{}
	tree.root = NewNode(values[0])
	queue = append(queue, tree.root)

	var count = 1

	for len(queue) > 0 && count < length {
		var current = queue[0]
		queue = queue[1:]

		if current.left == nil {
			if values[count] != -1 {
				current.left = NewNode(values[count])
			}
			count += 1
			if current.left != nil {
				queue = append(queue, current.left)
			}
			if count >= length {
				break
			}
		}

		if current.right == nil {
			if values[count] != -1 {
				current.right = NewNode(values[count])
			}
			count += 1
			if current.right != nil {
				queue = append(queue, current.right)
			}
			if count >= length {
				break
			}
		}
	}

	return tree.root
}

func (tree *BinTree) print(node *BinNode, depth int) {
	if node.left != nil {
		tree.print(node.left, depth+1)
	}
	spaces := []byte{}
	for i := 0; i < depth; i++ {
		spaces = append(spaces, '-')
	}
	fmt.Printf("%s%d\n", spaces, node.data)
	if node.right != nil {
		tree.print(node.right, depth+1)
	}
}

func (tree *BinTree) Print() {
	tree.print(tree.root, 0)
}

func (tree *BinTree) depth(node *BinNode) int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1
	}
	leftDepth := tree.depth(node.left) + 1
	rightDepth := tree.depth(node.right) + 1

	if leftDepth > rightDepth {
		return leftDepth
	}
	return rightDepth
}

func (tree *BinTree) Depth() int {
	return tree.depth(tree.root)
}

func (tree *BinTree) height(node *BinNode) int {
	if node == nil {
		return 0
	}
	return 1 + tree.height(node.left)
}

func (tree *BinTree) Height() int {
	if tree.root == nil {
		return 0
	}
	return tree.height(tree.root) - 1
}

func (tree *BinTree) nodeExists(idxToFind int, height int) bool {
	var left = 0
	var right = int(math.Pow(2, float64(height))) - 1
	var level = 0
	var node = tree.root

	for level < height && left <= right {
		mid := int(math.Ceil(float64(left+right) / 2))
		if idxToFind >= mid {
			left = mid
			node = node.right
		} else {
			right = mid - 1
			node = node.left
		}
		level += 1
	}
	return node != nil
}

func (tree *BinTree) CountNodes() int {
	if tree.root == nil {
		return 0
	}
	var height = tree.Height()
	if height == 0 {
		return 1
	}
	var upperCount = int(math.Pow(2, float64(height))) - 1
	var left, right = 0, upperCount
	for left < right {
		idxToFind := int(math.Ceil(float64(left+right) / 2))

		if tree.nodeExists(idxToFind, height) {
			left = idxToFind
		} else {
			right = idxToFind - 1
		}
	}
	return upperCount + left + 1
}
