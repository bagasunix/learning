package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSumBasic(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want [0 1]", result)
	}
}

func TestTwoSumNoFirstElement(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("got %v, want [1 2]", result)
	}
}

func TestTwoSumSameValueDifferentIndex(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want [0 1]", result)
	}
}

func TestTwoSumHashmapBasic(t *testing.T) {
	result := twoSumHashmap([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want [0 1]", result)
	}
}

func TestTwoSumHashmapNoFirst(t *testing.T) {
	result := twoSumHashmap([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("got %v, want [1 2]", result)
	}
}

func TestTwoSumHashmapSameValue(t *testing.T) {
	result := twoSumHashmap([]int{3, 3}, 6)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want [0 1]", result)
	}
}
