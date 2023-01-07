package list

import "testing"

func TestAcyclic(t *testing.T) {
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	p1 := list.Current()
	list.PushBack(5)
	list.PushBack(6)
	list.PushBack(7)
	list.PushBack(8)
	p2 := list.Current()
	p2.next = p1

	aPoint := list.IsAcyclic()

	t.Log("acyclic", aPoint)
	if aPoint == nil {
		t.Error("expected cyclic, but got acyclic")
	}
}
