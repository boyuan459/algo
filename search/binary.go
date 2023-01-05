package search

func BinarySearch(values []int, target int, left int, right int) int {

	for left <= right {
		mid := (left + right) / 2
		if values[mid] == target {
			return mid
		}
		if values[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
