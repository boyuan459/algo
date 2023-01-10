package sort

func swap(arr []int, i int, j int) {
	var temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int, left int, right int) int {
	var p = right
	for i := right; i >= left; i-- {
		if arr[i] > arr[p] {
			swap(arr, i, p-1)
			swap(arr, p-1, p)
			p--
		}
	}
	return p
}

func QuickSort(arr []int, left int, right int) {
	if left < right {
		var pivot = partition(arr, left, right)

		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
}

func QuickSelect(arr []int, left int, right int, idxToFind int) int {
	var pivot = partition(arr, left, right)

	if pivot == idxToFind {
		return arr[idxToFind]
	} else if pivot > idxToFind {
		return QuickSelect(arr, left, pivot-1, idxToFind)
	} else {
		return QuickSelect(arr, pivot+1, right, idxToFind)
	}
}
