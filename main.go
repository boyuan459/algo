package main

import (
	"fmt"

	"github.com/boyuan459/algo/graph"
)

func main() {
	graph := graph.NewGraph()
	edges := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	graph.Insert(edges, 4)
	var shortest = graph.ShortestPath(2)
	fmt.Println("shortest path", shortest)
	fmt.Println("graph", graph)
}
