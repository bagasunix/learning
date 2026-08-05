package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacciBasic(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8}
	result := fibonacci(7)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestFibonacciFirstTwo(t *testing.T) {
	result := fibonacci(2)
	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("got %v, want [0 1]", result)
	}
}

func TestFibonacciOne(t *testing.T) {
	result := fibonacci(1)
	if !reflect.DeepEqual(result, []int{0}) {
		t.Errorf("got %v, want [0]", result)
	}
}

func TestFibonacciZero(t *testing.T) {
	result := fibonacci(0)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
}

func TestFibRecursiveBase(t *testing.T) {
	if fibRecursive(0) != 0 {
		t.Error("fibRecursive(0) should be 0")
	}
	if fibRecursive(1) != 1 {
		t.Error("fibRecursive(1) should be 1")
	}
}

func TestFibRecursiveN10(t *testing.T) {
	if fibRecursive(10) != 55 {
		t.Errorf("fibRecursive(10) = %d, want 55", fibRecursive(10))
	}
}

func TestFibMemoizedMatchesRecursive(t *testing.T) {
	for i := 0; i < 20; i++ {
		if fibMemoized(i) != fibRecursive(i) {
			t.Errorf("mismatch at %d: memoized=%d, recursive=%d",
				i, fibMemoized(i), fibRecursive(i))
		}
	}
}

func TestFibMemoizedLarge(t *testing.T) {
	// fib(50) = 12586269025
	if fibMemoized(50) != 12586269025 {
		t.Errorf("fibMemoized(50) = %d, want 12586269025", fibMemoized(50))
	}
}
