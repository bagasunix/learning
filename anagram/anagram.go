package anagram

import (
	"sort"
	"strings"
)

// =============================================================================
// ANAGRAM
// =============================================================================
// Tingkat   : Mudah
// Konsep    : Sorting / hash map / frequency counter
// =============================================================================
//
// SOAL:
//   Cek apakah dua string adalah anagram satu sama lain.
//   Anagram = dua string yang mengandung huruf yang sama, hanya urutannya beda.
//   (LeetCode #242)
//
// CONTOH:
//   isAnagram("listen", "silent")  -> true
//   isAnagram("hello", "world")    -> false
//   isAnagram("Astronomer", "Moon starer") -> true  (kalau ignore spasi & case)
//   isAnagram("", "")              -> true
//   isAnagram("a", "a")            -> true
//   isAnagram("ab", "a")           -> false (panjang beda)
//
// EDGE CASES yang harus disebut:
//   - Panjang berbeda -> langsung false (early return)
//   - String kosong   -> true
//   - Case insensitive & ignore spasi (tergantung soal, tanyakan ke interviewer)
//
// ADA DUA IMPLEMENTASI:
//
//   1. isAnagram (sort-based):
//      - Sort kedua string, compare
//      - O(n log n) time, O(n) space
//      - Lebih mudah dibaca, bagus untuk explain
//
//   2. isAnagramCounter (hashmap-based):
//      - Hitung frekuensi huruf string A, kurangi dengan string B
//      - O(n) time, O(1) space (max 26 huruf)
//      - LEBIH OPTIMAL, ini yang biasanya diharapkan interviewer
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Approach pertama: sort kedua string lalu compare — O(n log n).
//    Approach lebih optimal: frequency counter pakai map/array[26],
//    increment untuk string A, decrement untuk string B.
//    Kalau ada yang negatif -> bukan anagram. O(n) time, O(1) space."
//
// JEBAKAN UMUM:
//   - Tidak cek panjang di awal -> waste waktu proses
//   - Lupa lowercase sebelum compare
// =============================================================================

func cleanAlpha(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// isAnagram: sort-based — O(n log n) time, O(n) space
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

// isAnagramCounter: hashmap-based — O(n) time, O(1) space (lebih optimal)
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
			return false // huruf di B lebih banyak dari A
		}
	}
	return true
}
