package reversestring

// =============================================================================
// REVERSE STRING
// =============================================================================
// Tingkat   : Mudah
// Konsep    : Two-pointer, rekursi, slice manipulation
// Complexity: O(n) time, O(n) space (karena string immutable di Go, perlu []rune)
// =============================================================================
//
// SOAL:
//   Balik urutan karakter dalam sebuah string.
//   (LeetCode #344, versi string)
//
// CONTOH:
//   reverseString("hello")   -> "olleh"
//   reverseString("Go")      -> "oG"
//   reverseString("a")       -> "a"
//   reverseString("")        -> ""
//   reverseString("abcde")   -> "edcba"
//   reverseString("racecar") -> "racecar"  (palindrome)
//
// EDGE CASES yang harus disebut:
//   - String kosong    -> ""
//   - Satu karakter    -> sama
//   - Sudah terbalik   -> return versi terbalik
//   - Unicode/emoji    -> pakai []rune bukan []byte!
//
// ADA DUA IMPLEMENTASI:
//
//   1. reverseString (two-pointer iterative):
//      - Konversi ke []rune, swap i dan j dari luar ke dalam
//      - O(n) time, O(n) space
//      - INI YANG PALING DIHARAPKAN saat interview
//
//   2. reverseStringRecursive:
//      - Base case: panjang <= 1
//      - Rekursi: reverse(s[1:]) + s[0]
//      - O(n) time, O(n) stack space
//      - Demonstrasi rekursi, tapi tidak efisien untuk production
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Konversi ke slice rune dulu biar unicode-safe.
//    Dua pointer: kiri dari index 0, kanan dari index terakhir.
//    Swap, lalu keduanya bergerak ke tengah sampai bertemu."
//
// JEBAKAN UMUM:
//   - Pakai []byte -> karakter multibyte (unicode) akan corrupt
//   - String di Go immutable, harus convert ke []rune dulu
// =============================================================================

// reverseString: two-pointer iterative — O(n) time, O(n) space
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// reverseStringRecursive: O(n) time, O(n) stack — untuk demo rekursi
func reverseStringRecursive(s string) string {
	if len(s) <= 1 {
		return s
	}
	runes := []rune(s)
	return reverseStringRecursive(string(runes[1:])) + string(runes[0])
}
