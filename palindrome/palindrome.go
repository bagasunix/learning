package palindrome

import (
	"strings"
	"unicode"
)

func cleanString(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func isPalindrome(s string) bool {
	cleaned := cleanString(s)
	runes := []rune(cleaned)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeTwoPointer(s string) bool {
	cleaned := cleanString(s)
	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}
	return true
}
