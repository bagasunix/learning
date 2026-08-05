package maxsubarray

import "testing"

func TestMaxSubArrayKadane(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if maxSubArray(nums) != 6 {
		t.Errorf("got %d, want 6", maxSubArray(nums))
	}
}

func TestMaxSubArraySingle(t *testing.T) {
	if maxSubArray([]int{1}) != 1 {
		t.Errorf("got %d, want 1", maxSubArray([]int{1}))
	}
}

func TestMaxSubArrayAllPositive(t *testing.T) {
	nums := []int{5, 4, -1, 7, 8}
	if maxSubArray(nums) != 23 {
		t.Errorf("got %d, want 23", maxSubArray(nums))
	}
}

func TestMaxSubArrayAllNegative(t *testing.T) {
	nums := []int{-3, -1, -2}
	if maxSubArray(nums) != -1 {
		t.Errorf("got %d, want -1", maxSubArray(nums))
	}
}

func TestMaxSubArrayTwoElements(t *testing.T) {
	if maxSubArray([]int{-2, -1}) != -1 {
		t.Errorf("got %d, want -1", maxSubArray([]int{-2, -1}))
	}
}
