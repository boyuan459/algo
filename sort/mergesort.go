package sort

func MergeSortOne(arr1 []int, m int, arr2 []int, n int) {
	var p1, p2 = m - 1, n - 1
	for p1 >= 0 && p2 >= 0 {
		var offset = p1 + p2 + 1
		if arr1[p1] < arr2[p2] {
			arr1[offset] = arr2[p2]
			p2--
		} else {
			arr1[offset] = arr1[p1]
			p1--
		}
	}
	for p2 >= 0 {
		arr1[p2] = arr2[p2]
		p2--
	}
}

func MergeSort(arr1 []int, arr2 []int) []int {
	var m = len(arr1)
	var n = len(arr2)
	var arr = make([]int, m+n)

	for m > 0 && n > 0 {
		var offset = m + n - 1
		if arr1[m] < arr2[n] {
			n--
			arr[offset] = arr2[n]
		} else {
			m--
			arr[offset] = arr1[m]
		}
	}

	for m > 0 {
		m--
		arr[m] = arr1[m]
	}

	for n > 0 {
		n--
		arr[n] = arr2[n]
	}

	return arr
}

func Median(arr1 []int, arr2 []int) float32 {
	arr := MergeSort(arr1, arr2)

	var mid = len(arr) / 2
	if len(arr)%2 == 0 {
		return (float32(arr[mid-1]) + float32(arr[mid])) / 2
	} else {
		return float32(arr[mid])
	}
}
