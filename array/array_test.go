package array

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	var target = 9
	res := TwoSum(nums, target)

	if !reflect.DeepEqual(res, []int{0, 1}) {
		t.Errorf("expect [0 1], but got %v", res)
	}

}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	res := TwoSum(nums, 6)

	if !reflect.DeepEqual(res, []int{1, 2}) {
		t.Errorf("expect [1 2], but got %v", res)
	}
}

func TestMaxArea(t *testing.T) {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := MaxArea(heights)
	t.Log(res)

	t.Errorf("expected 49, but got %v", res)
}
