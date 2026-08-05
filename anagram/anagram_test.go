package anagram

import "testing"

func TestIsAnagramTrue(t *testing.T) {
	if !isAnagram("listen", "silent") {
		t.Error("listen/silent should be anagram")
	}
}

func TestIsAnagramFalse(t *testing.T) {
	if isAnagram("hello", "world") {
		t.Error("hello/world should not be anagram")
	}
}

func TestIsAnagramWithSpace(t *testing.T) {
	if !isAnagram("Dormitory", "dirty room") {
		t.Error("should be anagram")
	}
}

func TestIsAnagramCaseInsensitive(t *testing.T) {
	if !isAnagram("Listen", "Silent") {
		t.Error("should be anagram (case insensitive)")
	}
}

func TestIsAnagramCounterTrue(t *testing.T) {
	if !isAnagramCounter("listen", "silent") {
		t.Error("should be anagram")
	}
}

func TestIsAnagramCounterFalse(t *testing.T) {
	if isAnagramCounter("hello", "world") {
		t.Error("should not be anagram")
	}
}

func TestIsAnagramCounterWithSpace(t *testing.T) {
	if !isAnagramCounter("Dormitory", "dirty room") {
		t.Error("should be anagram")
	}
}
