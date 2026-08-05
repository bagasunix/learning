package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotateBasic(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, want %v", nums, expected)
	}
}

func TestRotateSmall(t *testing.T) {
	nums := []int{-1, -100, 3, 99}
	rotate(nums, 2)
	expected := []int{3, 99, -1, -100}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, want %v", nums, expected)
	}
}

func TestRotateFullCycle(t *testing.T) {
	nums := []int{1, 2, 3}
	rotate(nums, 3)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, want %v", nums, expected)
	}
}

func TestRotateMoreThanLength(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	rotate(nums, 7) // equivalent to k=2
	expected := []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, want %v", nums, expected)
	}
}

func TestRotateZero(t *testing.T) {
	nums := []int{1, 2, 3}
	rotate(nums, 0)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, want %v", nums, expected)
	}
}

func TestRotateSingle(t *testing.T) {
	nums := []int{42}
	rotate(nums, 5)
	expected := []int{42}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("got %v, want %v", nums, expected)
	}
}
