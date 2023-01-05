package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	values := []int{3, 7, 9, 15, 18, 20}
	found := BinarySearch(values, 9, 0, len(values))

	if found != 2 {
		t.Errorf("expected 2, got %v", found)
	}
}

func TestBinarySearch2(t *testing.T) {
	values := []int{3, 7, 9, 15, 18, 20, 21}
	found := BinarySearch(values, 15, 0, len(values))

	if found != 3 {
		t.Errorf("expected 3, got %v", found)
	}
}
