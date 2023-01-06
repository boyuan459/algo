package main

import (
	"fmt"

	"github.com/boyuan459/algo/bst"
	"github.com/boyuan459/algo/graph"
	"github.com/boyuan459/algo/list"
)

func main() {
	graph := graph.NewGraph()
	edges := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	graph.Insert(edges, 4)
	var shortest = graph.ShortestPath(2)
	fmt.Println("shortest path", shortest)
	fmt.Println("graph", graph)

	tree := bst.New()
	// tree.InsertArray([]int{3, 9, 20, 15, 7})
	tree.Insert(15)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(3)
	tree.Insert(9)
	// tree.Insert(8)
	tree.Insert(18)
	tree.PreOrderTraverse()
	var depth = tree.Depth()
	fmt.Println("depth", depth)
	var levels = tree.LevelOrder()
	fmt.Println("levels", levels)
	var rightViews = tree.RightSideView()
	fmt.Println("right side views", rightViews)
	var leftViews = tree.LeftSideView()
	fmt.Println("left side views", leftViews)
	var height = tree.Height()
	fmt.Println("height", height)
	var count = tree.CountNodes()
	fmt.Println("count", count)
	var valid = tree.IsValid()
	fmt.Println("valid", valid)

	list := list.New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	fmt.Println("list", list)
	list.Print()
	// list.Reverse()
	list.ReverseBetween(2, 4)
	list.PushBack(6)
	list.Print()
}
