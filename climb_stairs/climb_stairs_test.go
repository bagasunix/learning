package climbstairs

import "testing"

func TestClimbStairs2(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Errorf("got %d, want 2", climbStairs(2))
	}
}

func TestClimbStairs3(t *testing.T) {
	if climbStairs(3) != 3 {
		t.Errorf("got %d, want 3", climbStairs(3))
	}
}

func TestClimbStairs1(t *testing.T) {
	if climbStairs(1) != 1 {
		t.Errorf("got %d, want 1", climbStairs(1))
	}
}

func TestClimbStairs5(t *testing.T) {
	if climbStairs(5) != 8 {
		t.Errorf("got %d, want 8", climbStairs(5))
	}
}

func TestClimbStairs10(t *testing.T) {
	if climbStairs(10) != 89 {
		t.Errorf("got %d, want 89", climbStairs(10))
	}
}

func TestClimbStairs45(t *testing.T) {
	// LeetCode constraint max
	if climbStairs(45) != 1836311903 {
		t.Errorf("got %d, want 1836311903", climbStairs(45))
	}
}
