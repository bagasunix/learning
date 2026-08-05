package mergesorted

import (
	"reflect"
	"testing"
)

func TestMergeSortedBasic(t *testing.T) {
	result := mergeSorted([]int{1, 3, 5}, []int{2, 4, 6})
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestMergeSortedOneEmpty(t *testing.T) {
	result := mergeSorted([]int{1, 2, 3}, []int{})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestMergeSortedBothEmpty(t *testing.T) {
	result := mergeSorted([]int{}, []int{})
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestMergeSortedWithDupes(t *testing.T) {
	result := mergeSorted([]int{1, 3, 3, 5}, []int{2, 3, 4})
	expected := []int{1, 2, 3, 3, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestMergeSortedSecondStartsBefore(t *testing.T) {
	result := mergeSorted([]int{5, 6, 7}, []int{1, 2, 3})
	expected := []int{1, 2, 3, 5, 6, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}
