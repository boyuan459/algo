package doubly_list

import "testing"

func TestDoublyList(t *testing.T) {
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.Print()
}
