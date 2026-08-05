package reversestring

import "testing"

func TestReverseStringBasic(t *testing.T) {
	if reverseString("hello") != "olleh" {
		t.Errorf("got %s, want olleh", reverseString("hello"))
	}
}

func TestReverseStringSingleChar(t *testing.T) {
	if reverseString("a") != "a" {
		t.Error("single char should return itself")
	}
}

func TestReverseStringEmpty(t *testing.T) {
	if reverseString("") != "" {
		t.Error("empty should return empty")
	}
}

func TestReverseStringRecursive(t *testing.T) {
	if reverseStringRecursive("hello") != "olleh" {
		t.Errorf("got %s, want olleh", reverseStringRecursive("hello"))
	}
}

func TestReverseStringRecursiveSingle(t *testing.T) {
	if reverseStringRecursive("a") != "a" {
		t.Error("single char should return itself")
	}
}
