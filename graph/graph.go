package graph

import "fmt"

type Graph struct {
	length   int
	adjList  [][]int
	indegree []int
}

func New(length int) *Graph {
	return &Graph{
		length:   length,
		adjList:  make([][]int, length),
		indegree: make([]int, length),
	}
}

func (graph *Graph) BuildWith1D(values []int) {
	for k, v := range values {
		if v != -1 {
			graph.adjList[v] = append(graph.adjList[v], k)
		}
	}
}

func (graph *Graph) BuildWith2D(values [][]int) {
	for _, pair := range values {
		var source = pair[0]
		var target = pair[1]
		graph.adjList[source] = append(graph.adjList[source], target)
		graph.indegree[target]++
	}
}

func (graph *Graph) BuildAdjList(values [][]int) {
	for _, pair := range values {
		var source = pair[0]
		var target = pair[1]
		graph.adjList[source] = append(graph.adjList[source], target)
	}
}

func (graph *Graph) PossiblePartition() bool {
	group := make([]int, graph.length+1)

	for i := 1; i < graph.length; i++ {
		if group[i] == 0 {
			group[i] = 1
			queue := make([]int, 0)
			queue = append(queue, i)

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				for _, adj := range graph.adjList[current] {
					// fmt.Println(adj)
					if group[adj] == 0 {
						group[adj] = group[current] * -1
						queue = append(queue, adj)
					} else if group[adj] == group[current] {
						return false
					}
				}
			}
		}
	}
	fmt.Println(group)
	return true
}

func (graph *Graph) AcyclicTopologicalSort() bool {
	stack := make([]int, 0)
	for k, v := range graph.indegree {
		if v == 0 {
			stack = append(stack, k)
		}
	}
	var count = 0
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]
		count++
		for _, conn := range graph.adjList[current] {
			graph.indegree[conn]--
			if graph.indegree[conn] == 0 {
				stack = append(stack, conn)
			}
		}
	}

	return count == graph.length
}

func (graph *Graph) NumOfMinutes(headID int, informTime []int) int {
	if len(graph.adjList[headID]) == 0 {
		return 0
	}
	var max = 0
	for _, conn := range graph.adjList[headID] {
		var time = graph.NumOfMinutes(conn, informTime)
		if time > max {
			max = time
		}
	}
	return max + informTime[headID]
}
