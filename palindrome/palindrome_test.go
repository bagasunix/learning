package palindrome

import "testing"

func TestIsPalindromeSimpleTrue(t *testing.T) {
	if !isPalindrome("racecar") {
		t.Error("racecar should be palindrome")
	}
}

func TestIsPalindromeSimpleFalse(t *testing.T) {
	if isPalindrome("hello") {
		t.Error("hello should not be palindrome")
	}
}

func TestIsPalindromeWithSpaces(t *testing.T) {
	if !isPalindrome("A man a plan a canal Panama") {
		t.Error("should be palindrome")
	}
}

func TestIsPalindromeSingleChar(t *testing.T) {
	if !isPalindrome("a") {
		t.Error("single char should be palindrome")
	}
}

func TestIsPalindromeEmpty(t *testing.T) {
	if !isPalindrome("") {
		t.Error("empty should be palindrome")
	}
}

func TestIsPalindromeTwoPointerTrue(t *testing.T) {
	if !isPalindromeTwoPointer("racecar") {
		t.Error("racecar should be palindrome")
	}
}

func TestIsPalindromeTwoPointerFalse(t *testing.T) {
	if isPalindromeTwoPointer("hello") {
		t.Error("hello should not be palindrome")
	}
}

func TestIsPalindromeTwoPointerWithSpaces(t *testing.T) {
	if !isPalindromeTwoPointer("A man a plan a canal Panama") {
		t.Error("should be palindrome")
	}
}
