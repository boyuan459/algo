package priority_queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	compare := func(a, b int) bool {
		// t.Log("compare", a < b)
		return a < b
	}
	pq := New(compare)
	pq.Push(4)
	pq.Push(2)
	pq.Push(1)
	pq.Push(3)
	pq.Push(5)

	var v = pq.Peek()
	t.Log("v", v)
	t.Log("pq", pq)

	var p = pq.Pop()
	t.Log("p", p)
	t.Log("pq", pq)

	var p2 = pq.Pop()
	t.Log("p2", p2)
	t.Log("pq", pq)
}
