package stringcompression

import (
	"fmt"
	"strings"
)

// String Compression
// ------------------
// Kompres string dengan menghitung karakter berulang.
// (LeetCode #443, simplified)
//
// Contoh:
//
//	compress("aabcccccaaa") -> "a2b1c5a3"
//	compress("abc")         -> "a1b1c1"
//	compress("")            -> ""
//
// Bonus: Kembalikan string asli kalau hasil kompresi tidak lebih pendek.
func compress(s string) string {
	if len(s) == 0 {
		return ""
	}
	var sb strings.Builder
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			fmt.Fprintf(&sb, "%c%d", s[i-1], count)
			count = 1
		}
	}
	fmt.Fprintf(&sb, "%c%d", s[len(s)-1], count)
	return sb.String()
}
