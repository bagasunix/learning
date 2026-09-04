package main

import (
	"reflect"
	"testing"
)

func TestQuestion1A(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{-1, 1}, 0},
		{[]int{-1, -7, -5}, -2},
		{[]int{1, 2, 1, 6}, 3},
		{[]int{5}, 1},
		{[]int{-3}, -1},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -2, -3}, -4},
	}
	for _, tt := range tests {
		got := Question1A(tt.arr)
		if got != tt.want {
			t.Errorf("Question1A(%v) = %d, want %d", tt.arr, got, tt.want)
		}
	}
}

func TestQuestion1B(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{-1, 1}, 0},
		{[]int{-1, -7, -5}, -1},
		{[]int{1, 2, 1, 6}, 1},
		{[]int{5}, 5},
		{[]int{-3}, -3},
		{[]int{1, 2, 3}, 1},
		{[]int{-1, -2, -3}, -1},
		{[]int{-7, -5, -3}, -3},
	}
	for _, tt := range tests {
		got := Question1B(tt.arr)
		if got != tt.want {
			t.Errorf("Question1B(%v) = %d, want %d", tt.arr, got, tt.want)
		}
	}
}

func TestQuestion2(t *testing.T) {
	r, err := Question2([]int{0, 1, 0}, 1000, 1010)
	if r != 0 || err == nil || err.Error() != "not found" {
		t.Errorf("got %d %v", r, err)
	}
	r, err = Question2([]int{0, 1, 0}, 1259, 2601)
	if r != 62952 || err != nil {
		t.Errorf("got %d %v", r, err)
	}
	r, err = Question2([]int{}, 12, 34)
	if r != 4321 || err != nil {
		t.Errorf("got %d %v", r, err)
	}
}

func TestQuestion3(t *testing.T) {
	tests := []struct {
		numsA, numsB []int
		want         []int
	}{
		{[]int{1, 2}, []int{1, 3}, []int{1}},
		{[]int{1, 2, 2}, []int{1, 2, 4}, []int{2, 1}},
		{[]int{5, 6}, []int{1, 2}, []int{}},
		{[]int{3, 1, 2}, []int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		got := Question3(tt.numsA, tt.numsB)
		if got == nil {
			got = []int{}
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Question3(%v, %v) = %v, want %v", tt.numsA, tt.numsB, got, tt.want)
		}
	}
}

func TestQuestion4(t *testing.T) {
	got := Question4(1223334)
	want := map[int]int{1: 1, 2: 2, 3: 3, 4: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Question4(1223334) = %v, want %v", got, want)
	}
}

func TestQuestion5(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 14},
		{[]int{1, 2, 3, 4, 5}, 39},
		{[]int{2, 3}, 6},
		{[]int{7}, 49},
		{[]int{}, 0},
	}
	for _, tt := range tests {
		got := Question5(tt.nums...)
		if got != tt.want {
			t.Errorf("Question5(%v) = %d, want %d", tt.nums, got, tt.want)
		}
	}
}
