package graph

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
