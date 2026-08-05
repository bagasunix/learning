package anagram

import (
	"sort"
	"strings"
)

func cleanAlpha(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func isAnagram(a, b string) bool {
	aClean, bClean := cleanAlpha(a), cleanAlpha(b)
	if len(aClean) != len(bClean) {
		return false
	}
	aRunes, bRunes := []rune(aClean), []rune(bClean)
	sort.Slice(aRunes, func(i, j int) bool { return aRunes[i] < aRunes[j] })
	sort.Slice(bRunes, func(i, j int) bool { return bRunes[i] < bRunes[j] })
	return string(aRunes) == string(bRunes)
}

func isAnagramCounter(a, b string) bool {
	aClean, bClean := cleanAlpha(a), cleanAlpha(b)
	if len(aClean) != len(bClean) {
		return false
	}
	counter := make(map[rune]int)
	for _, c := range aClean {
		counter[c]++
	}
	for _, c := range bClean {
		counter[c]--
		if counter[c] < 0 {
			return false
		}
	}
	return true
}
