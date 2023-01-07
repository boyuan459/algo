package doubly_list

import "fmt"

type element[T any] struct {
	val  T
	prev *element[T]
	next *element[T]
}

type DoublyList[T any] struct {
	head, tail *element[T]
}

func New[T any]() *DoublyList[T] {
	return &DoublyList[T]{
		head: nil,
	}
}

func (list *DoublyList[T]) PushBack(val T) {
	if list.head == nil {
		list.head = &element[T]{
			val:  val,
			prev: nil,
			next: nil,
		}
		list.tail = list.head
	} else {
		ele := &element[T]{
			val:  val,
			prev: list.tail,
			next: nil,
		}
		list.tail.next = ele
		list.tail = ele
	}
}

func (list *DoublyList[T]) Print() {
	var current = list.head

	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
	fmt.Println()
}
