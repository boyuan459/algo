package stack

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	data []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(val T) int {
	s.data = append(s.data, val)
	return len(s.data)
}

func (s *Stack[T]) Pop() (T, error) {
	var size = len(s.data)
	if size == 0 {
		var val T
		return val, errors.New("Cannot pop from empty stack")
	}
	var val = s.data[size-1]
	s.data = s.data[:size-1]
	return val, nil
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Stack[T]) Peek() (T, error) {
	var size = len(s.data)
	if size == 0 {
		var val T
		return val, errors.New("Cannot peek from empty stack")
	}
	return s.data[size-1], nil
}

func (s Stack[T]) Print() {
	fmt.Print("[")
	for i := 0; i < len(s.data); i++ {
		fmt.Print(s.data[i])
		if i < len(s.data)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}
