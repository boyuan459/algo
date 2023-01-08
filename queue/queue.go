package queue

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	queue []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0),
	}
}

func (q *Queue[T]) Push(val T) int {
	q.queue = append(q.queue, val)
	return len(q.queue)
}

func (q *Queue[T]) Pop() (T, error) {
	var size = len(q.queue)
	if size == 0 {
		var val T
		return val, errors.New("Cannot pop from empty queue")
	}
	var val = q.queue[0]
	q.queue = q.queue[1:]
	return val, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var size = len(q.queue)
	if size == 0 {
		var val T
		return val, errors.New("Cannot peek from empty queue")
	}
	return q.queue[size-1], nil
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *Queue[T]) Print() {
	fmt.Print("[")
	for i := 0; i < len(q.queue); i++ {
		fmt.Print(q.queue[i])
		if i < len(q.queue)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}
