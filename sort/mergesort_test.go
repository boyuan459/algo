package sort

import (
	"reflect"
	"testing"
)

func TestMergeSortOne(t *testing.T) {
	var arr1 = []int{4, 7, 9, 10, 0, 0, 0}
	var arr2 = []int{2, 5, 8}

	MergeSortOne(arr1, 4, arr2, 3)
	if reflect.DeepEqual(arr1, []int{2, 4, 5, 7, 8, 9, 10}) {
		t.Log("merged array ok")
	} else {
		t.Errorf("expected merge array: %v", arr1)
	}
}
