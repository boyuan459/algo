package array

import (
	"fmt"
)

// TwoSum calculates the index of the nums array, and the target
// is the sum
func TwoSum(nums []int, target int) []int {
	sum := make(map[int]int)

	res := []int{}
	for k, v := range nums {
		// fmt.Println(k, v)
		expect := target - v
		if idx, ok := sum[expect]; ok {
			// fmt.Println("found value at index ", idx)
			res = append(res, idx)
			res = append(res, k)
			break
		}
		sum[v] = k
	}

	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxArea(height []int) int {
	var left, right = 0, len(height) - 1
	fmt.Println(left, right)
	maxArea := 0

	for left < right {
		var area = (right - left) * min(height[left], height[right])
		maxArea = max(area, maxArea)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxArea
}

func MoveZeros(nums *[]int) {
	// count the non zero elements
	var count = 0
	for i := 0; i < len(*nums); i++ {
		if (*nums)[i] != 0 {
			// shift the number forward
			(*nums)[count] = (*nums)[i]
			count += 1
		}
	}
	for count < len(*nums) {
		(*nums)[count] = 0
		count += 1
	}
}
