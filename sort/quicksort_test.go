package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{4, 6, 1, 8, 11, 13, 3, 7}
	QuickSort(arr, 0, len(arr)-1)
	t.Log("arr ", arr)

	res := []int{1, 3, 4, 6, 7, 8, 11, 13}
	var sort = true

	for i := 0; i < len(res); i++ {
		if res[i] != arr[i] {
			sort = false
			break
		}
	}

	if !sort {
		t.Errorf("expected quick sort, but got %v", sort)
	}
}

func TestQuickSelect(t *testing.T) {
	arr := []int{4, 6, 1, 8, 11, 13, 3, 7}

	var sixthLarge = QuickSelect(arr, 0, len(arr)-1, 6)

	t.Log("2th Largest ", sixthLarge)

	if sixthLarge != 11 {
		t.Errorf("expected 11, but got %v", sixthLarge)
	}
}
