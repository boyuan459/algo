package graph

import (
	"math"
	"sort"

	"github.com/boyuan459/algo/priority_queue"
)

type Dijkstra struct {
	length  int
	adjList [][][]int
}

func NewGraph() *Dijkstra {
	return &Dijkstra{
		length:  0,
		adjList: make([][][]int, 0),
	}
}

func (graph *Dijkstra) Insert(edges [][]int, nodes int) {
	graph.length = nodes
	graph.adjList = make([][][]int, nodes)
	for _, pairs := range edges {
		var source = pairs[0]
		var target = pairs[1]
		var weight = pairs[2]
		graph.adjList[source-1] = append(graph.adjList[source-1], []int{target - 1, weight})
	}
}

func (graph *Dijkstra) ShortestPath(node int) int {
	distances := make([]int, graph.length)
	for i := 0; i < graph.length; i++ {
		distances[i] = math.MaxInt
	}
	distances[node-1] = 0
	compare := func(a, b int) bool {
		return distances[a] < distances[b]
	}
	pq := priority_queue.New(compare)
	pq.Push(node - 1)

	for !pq.IsEmpty() {
		current := pq.Pop()
		conns := graph.adjList[current]
		for _, pairs := range conns {
			var vertex = pairs[0]
			var weight = pairs[1]
			if distances[current]+weight < distances[vertex] {
				distances[vertex] = distances[current] + weight
				pq.Push(vertex)
			}
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return i > j
	})

	if distances[0] < math.MaxInt {
		return distances[0]
	}
	return -1
}
