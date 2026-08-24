package countvowels

import "strings"

// =============================================================================
// COUNT VOWELS
// =============================================================================
// Tingkat   : Mudah
// Konsep    : String iteration, map/set lookup
// Complexity: O(n) time, O(1) space
// =============================================================================
//
// SOAL:
//   Hitung jumlah huruf vokal (a, i, u, e, o) dalam sebuah string.
//   Case insensitive (huruf besar dan kecil dihitung sama).
//
// CONTOH:
//   countVowels("Hello World") -> 3  ('e', 'o', 'o')
//   countVowels("aeiou")       -> 5
//   countVowels("xyz")         -> 0
//   countVowels("")            -> 0
//   countVowels("AEIOU")       -> 5  (case insensitive)
//
// EDGE CASES yang harus disebut:
//   - String kosong        -> 0
//   - Semua konsonan       -> 0
//   - Semua vokal          -> len(s)
//   - Angka/simbol         -> 0 (bukan vokal)
//   - Huruf kapital        -> tetap dihitung (lowercase dulu)
//
// ADA DUA IMPLEMENTASI:
//
//   1. countVowels: hitung total vokal
//      - strings.ToLower dulu, lalu cek tiap karakter
//      - O(n) time, O(1) space
//
//   2. countVowelsDetail: hitung frekuensi tiap vokal
//      - Return map {"a": 2, "i": 1, ...}
//      - Berguna kalau interviewer tanya "breakdown per huruf"
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Lowercase string dulu, iterasi tiap karakter,
//    cek apakah termasuk dalam set vokal {'a','i','u','e','o'}.
//    O(n) time karena satu pass, O(1) space karena set vokal konstan."
//
// JEBAKAN UMUM:
//   - Tidak lowercase -> 'A' tidak ketemu di "aiueo"
//   - Lupa 'u' (sering kelewat: a, e, i, o, u)
// =============================================================================

// countVowels: hitung total vokal — O(n) time, O(1) space
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

// countVowelsDetail: breakdown frekuensi per vokal
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
