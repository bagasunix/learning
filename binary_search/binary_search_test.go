package binarysearch

import "testing"

func TestBinarySearchFound(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if binarySearch(nums, 7) != 3 {
		t.Errorf("got %d, want 3", binarySearch(nums, 7))
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if binarySearch(nums, 4) != -1 {
		t.Errorf("got %d, want -1", binarySearch(nums, 4))
	}
}

func TestBinarySearchFirst(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if binarySearch(nums, 1) != 0 {
		t.Errorf("got %d, want 0", binarySearch(nums, 1))
	}
}

func TestBinarySearchLast(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if binarySearch(nums, 11) != 5 {
		t.Errorf("got %d, want 5", binarySearch(nums, 11))
	}
}

func TestBinarySearchEmpty(t *testing.T) {
	if binarySearch([]int{}, 5) != -1 {
		t.Error("empty array should return -1")
	}
}

func TestBinarySearchRecursiveFound(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if binarySearchRecursive(nums, 7) != 3 {
		t.Errorf("got %d, want 3", binarySearchRecursive(nums, 7))
	}
}

func TestBinarySearchRecursiveNotFound(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	if binarySearchRecursive(nums, 4) != -1 {
		t.Errorf("got %d, want -1", binarySearchRecursive(nums, 4))
	}
}

func TestBinarySearchRecursiveEmpty(t *testing.T) {
	if binarySearchRecursive([]int{}, 5) != -1 {
		t.Error("empty array should return -1")
	}
}
