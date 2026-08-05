package majorityelement

import "testing"

func TestMajoritySimple(t *testing.T) {
	if majorityElement([]int{3, 2, 3}) != 3 {
		t.Errorf("got %d, want 3", majorityElement([]int{3, 2, 3}))
	}
}

func TestMajorityLonger(t *testing.T) {
	if majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Errorf("got %d, want 2", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	}
}

func TestMajoritySingle(t *testing.T) {
	if majorityElement([]int{1}) != 1 {
		t.Errorf("got %d, want 1", majorityElement([]int{1}))
	}
}

func TestMajorityAllSame(t *testing.T) {
	if majorityElement([]int{5, 5, 5, 5}) != 5 {
		t.Errorf("got %d, want 5", majorityElement([]int{5, 5, 5, 5}))
	}
}

func TestMajorityHalfPlusOne(t *testing.T) {
	if majorityElement([]int{1, 2, 1, 1, 2, 1}) != 1 {
		t.Errorf("got %d, want 1", majorityElement([]int{1, 2, 1, 1, 2, 1}))
	}
}
