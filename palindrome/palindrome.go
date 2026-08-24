package palindrome

import (
	"strings"
	"unicode"
)

// =============================================================================
// PALINDROME
// =============================================================================
// Tingkat   : Mudah
// Konsep    : String manipulation, two-pointer
// Complexity: O(n) time, O(n) space (karena cleanString membuat string baru)
// =============================================================================
//
// SOAL:
//   Cek apakah sebuah string adalah palindrome.
//   Abaikan huruf besar/kecil dan karakter non-alphanumeric.
//   (LeetCode #125)
//
// CONTOH:
//   isPalindrome("racecar")           -> true
//   isPalindrome("A man a plan a canal Panama") -> true
//   isPalindrome("hello")             -> false
//   isPalindrome("")                  -> true  (string kosong = palindrome)
//   isPalindrome("a")                 -> true  (satu karakter = palindrome)
//   isPalindrome("Was it a car or a cat I saw?") -> true
//
// EDGE CASES yang harus disebut:
//   - String kosong         -> true
//   - Satu karakter         -> true
//   - Semua non-alphanumeric ("!@#") -> true (setelah clean jadi "")
//   - Case insensitive: "Racecar" -> true
//
// ADA DUA IMPLEMENTASI:
//   1. isPalindrome       : bandingkan index 0 vs n-1, 1 vs n-2, dst
//   2. isPalindromeTwoPointer: lo/hi pointer bergerak mendekat ke tengah
//      (ini yang lebih sering diminta di interview karena lebih explicit)
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Clean string dulu: lowercase, buang non-alphanumeric.
//    Lalu two-pointer dari kiri dan kanan, jika ada yang tidak cocok -> false.
//    Kalau loop selesai tanpa mismatch -> true."
//
// JEBAKAN UMUM:
//   - Lupa clean string sebelum compare (case sensitive atau ada spasi)
//   - Pakai []byte bukan []rune -> karakter unicode bisa rusak
// =============================================================================

func cleanString(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// isPalindrome: bandingkan index dari luar ke dalam
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

// isPalindromeTwoPointer: eksplisit pakai lo/hi pointer (lebih jelas saat explain)
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
