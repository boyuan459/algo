package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	empty := q.IsEmpty()

	if empty {
		t.Errorf("expected false, got %v", empty)
	}

	val, _ := q.Pop()

	if val != 1 {
		t.Errorf("expected 1, got %v", val)
	}

	q.Print()
}
