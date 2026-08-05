package countvowels

import "strings"

func countVowels(s string) int {
	vowels := "aiueo"
	count := 0
	for _, c := range strings.ToLower(s) {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}
	return count
}

func countVowelsDetail(s string) map[string]int {
	result := map[string]int{"a": 0, "i": 0, "u": 0, "e": 0, "o": 0}
	for _, c := range strings.ToLower(s) {
		key := string(c)
		if _, ok := result[key]; ok {
			result[key]++
		}
	}
	return result
}
