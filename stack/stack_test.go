package stack

import "testing"

func TestStack(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	empty := stack.IsEmpty()

	if empty {
		t.Errorf("expected false, got %v", empty)
	}

	if val, error := stack.Pop(); error != nil {
		if val != 4 {
			t.Errorf("expected v, got %v", val)
		}
	}

	stack.Print()
}
