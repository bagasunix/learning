package factorial

import "testing"

func TestFactorialFive(t *testing.T) {
	if factorial(5) != 120 {
		t.Errorf("factorial(5) = %d, want 120", factorial(5))
	}
}

func TestFactorialZero(t *testing.T) {
	if factorial(0) != 1 {
		t.Errorf("factorial(0) = %d, want 1", factorial(0))
	}
}

func TestFactorialOne(t *testing.T) {
	if factorial(1) != 1 {
		t.Errorf("factorial(1) = %d, want 1", factorial(1))
	}
}

func TestFactorialRecursiveFive(t *testing.T) {
	if factorialRecursive(5) != 120 {
		t.Errorf("factorialRecursive(5) = %d, want 120", factorialRecursive(5))
	}
}

func TestFactorialRecursiveZero(t *testing.T) {
	if factorialRecursive(0) != 1 {
		t.Errorf("factorialRecursive(0) = %d, want 1", factorialRecursive(0))
	}
}

func TestFactorialRecursiveMatchesIterative(t *testing.T) {
	for i := 0; i < 10; i++ {
		if factorial(i) != factorialRecursive(i) {
			t.Errorf("mismatch at %d: iter=%d, rec=%d",
				i, factorial(i), factorialRecursive(i))
		}
	}
}
