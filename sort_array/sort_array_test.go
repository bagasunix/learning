package sortarray

import (
	"reflect"
	"testing"
)

var sortTests = []struct {
	name  string
	input []int
	want  []int
}{
	{"normal", []int{5, 3, 1, 4, 2}, []int{1, 2, 3, 4, 5}},
	{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
	{"single element", []int{42}, []int{42}},
	{"two elements", []int{2, 1}, []int{1, 2}},
	{"empty", []int{}, []int{}},
	{"negatives", []int{-3, 0, -1, 2, -2}, []int{-3, -2, -1, 0, 2}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range sortTests {
		t.Run("bubble/"+tt.name, func(t *testing.T) {
			got := bubbleSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range sortTests {
		t.Run("insertion/"+tt.name, func(t *testing.T) {
			got := insertionSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
