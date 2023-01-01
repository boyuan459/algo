package graph

type Graph struct {
	length  int
	adjList [][]int
}

func New(length int) *Graph {
	return &Graph{
		length:  length,
		adjList: make([][]int, length),
	}
}

func (graph *Graph) Build(values []int) {
	for k, v := range values {
		if v != -1 {
			graph.adjList[v] = append(graph.adjList[v], k)
		}
	}
}
