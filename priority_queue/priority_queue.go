package priority_queue

import "fmt"

type PriorityQueue struct {
	heap       []int
	comparator func(a, b int) bool
}

func New(compare func(a, b int) bool) *PriorityQueue {
	return &PriorityQueue{
		heap:       make([]int, 0),
		comparator: compare,
	}
}

func (pq PriorityQueue) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq PriorityQueue) Size() int {
	return len(pq.heap)
}

func (pq PriorityQueue) parent(i int) int {
	return (i - 1) / 2
}

func (pq PriorityQueue) leftChild(i int) int {
	return 2*i + 1
}

func (pq PriorityQueue) rightChild(i int) int {
	return 2*i + 2
}

func (pq PriorityQueue) compare(i, j int) bool {
	return pq.comparator(pq.heap[i], pq.heap[j])
}

func (pq *PriorityQueue) swap(i, j int) {
	var temp = pq.heap[i]
	pq.heap[i] = pq.heap[j]
	pq.heap[j] = temp
}

func (pq *PriorityQueue) siftUp() {
	var index = pq.Size() - 1

	for index > 0 && pq.compare(index, pq.parent(index)) {
		var parent = pq.parent(index)
		if pq.compare(index, parent) {
			fmt.Println("swap", parent)
			pq.swap(parent, index)
			index = parent
		}
	}
}

func (pq *PriorityQueue) Push(a int) int {
	pq.heap = append(pq.heap, a)
	pq.siftUp()
	return pq.Size()
}

func (pq *PriorityQueue) Peek() int {
	return pq.heap[0]
}

func (pq *PriorityQueue) siftDown() {
	var index = 0
	for index < pq.Size() && (pq.leftChild(index) < pq.Size() && pq.compare(pq.leftChild(index), index)) || (pq.rightChild(index) < pq.Size() && pq.compare(pq.rightChild(index), index)) {
		var child = index
		if pq.rightChild(index) < pq.Size() && pq.compare(pq.rightChild(index), pq.leftChild(index)) {
			child = pq.rightChild(index)
		} else {
			child = pq.leftChild(index)
		}
		pq.swap(child, index)
		index = child
	}
}

func (pq *PriorityQueue) Pop() int {
	var last = pq.Size() - 1
	var temp = pq.heap[0]
	pq.heap[0] = pq.heap[last]
	pq.heap[last] = temp
	pq.heap = pq.heap[0:last]
	pq.siftDown()
	return temp
}
