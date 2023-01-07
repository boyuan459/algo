package list

import "fmt"

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func New[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
	}
}

func (list *List[T]) PushBack(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) PushFront(v T) {
	if list.head == nil {
		list.head = &element[T]{next: nil, val: v}
		list.tail = list.head
	} else {
		ele := &element[T]{next: nil, val: v}
		ele.next = list.head
		list.head = ele
	}
}

func (list List[T]) Current() *element[T] {
	return list.tail
}

func (list *List[T]) Reverse() {
	var pre *element[T] = nil
	var current = list.head
	var next = current
	list.tail = current

	for current != nil {
		next = current.next
		current.next = pre
		pre = current
		current = next
	}

	list.head = pre
}

func (list *List[T]) ReverseBetween(left, right int) {
	var count = 0
	var pre *element[T] = nil
	var current = list.head
	var next = current
	var prestart = pre
	var start = current

	// find the start node
	for count < left-1 {
		pre = current
		current = current.next
		count++
	}

	prestart = pre
	start = current
	pre = nil

	for count < right && current != nil {
		next = current.next
		current.next = pre
		pre = current
		current = next
		count++
	}
	if current == nil {
		list.tail = start
	}

	if prestart != nil {
		prestart.next = pre
	}
	start.next = current
}

func (list *List[T]) Print() {
	var ele = list.head

	for ele != nil {
		fmt.Printf("%v ", ele.val)
		ele = ele.next
	}
	fmt.Println()
}

func (list *List[T]) IsAcyclic() *element[T] {
	var tortoise = list.head
	var hare = list.head

	for true {
		tortoise = tortoise.next
		hare = hare.next

		if hare != nil {
			hare = hare.next
		}

		// tortoise and hare can reach end, so it's acyclic
		if tortoise == nil || hare == nil {
			return nil
		}

		if tortoise == hare {
			break
		}
	}

	var p1 = list.head
	var p2 = hare

	for p1 != p2 {
		p1 = p1.next
		p2 = p2.next
	}

	return p1
}
