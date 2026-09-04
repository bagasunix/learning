package stringcompression

import (
	"strconv"
	"strings"
)

// =============================================================================
// STRING COMPRESSION (Run-Length Encoding)
// =============================================================================
// Tingkat   : Sedang
// Konsep    : String iteration, run-length encoding, strings.Builder
// Complexity: O(n) time, O(n) space
// =============================================================================
//
// SOAL:
//   Kompres string dengan menghitung karakter berulang berurutan.
//   Format output: karakter + jumlah kemunculan berturut-turut.
//   BONUS: Kalau hasil kompresi tidak lebih pendek, return string asli.
//   (LeetCode #443, simplified string version)
//
// CONTOH:
//   compress("aabcccccaaa") -> "a2b1c5a3"   (lebih pendek: 8 < 11) ✓
//   compress("abc")         -> "abc"         (BONUS: "a1b1c1"=6 >= "abc"=3)
//   compress("a")           -> "a"           (BONUS: "a1"=2 >= "a"=1)
//   compress("")            -> ""
//   compress("aaaaaaaaaaaa") -> "a12"        (lebih pendek: 3 < 12) ✓
//   compress("aabb")        -> "aabb"        (BONUS: "a2b2"=4 >= "aabb"=4, sama panjang -> asli)
//
// EDGE CASES yang harus disebut:
//   - String kosong              -> ""
//   - Satu karakter              -> return asli (bonus)
//   - Semua karakter berbeda     -> return asli (bonus: "a1b1c1" lebih panjang)
//   - Semua karakter sama        -> selalu lebih pendek (kecuali 1-2 karakter)
//   - Hasil sama panjang         -> return asli (bonus: tidak lebih pendek)
//
// CARA KERJA:
//   Loop dari index 1, bandingkan s[i] dengan s[i-1]:
//   - Sama     -> count++
//   - Berbeda  -> tulis karakter sebelumnya + count, reset count=1
//   Setelah loop -> tulis karakter terakhir + count
//
//   "aabcccccaaa":
//   i=1: s[1]='a'==s[0]='a' -> count=2
//   i=2: s[2]='b'!=s[1]='a' -> write "a2", count=1
//   i=3: s[3]='c'!=s[2]='b' -> write "b1", count=1
//   i=4..7: 'c'=='c' -> count=5
//   i=8: 'a'!='c' -> write "c5", count=1
//   i=9..10: 'a'=='a' -> count=3
//   end: write "a3"
//   result="a2b1c5a3" (8 chars < 11) -> return compressed ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Satu pass, track karakter sebelumnya dan count run-nya.
//    Kalau karakter berubah, tulis char+count, reset count.
//    Setelah loop, tulis karakter terakhir.
//    Bonus: bandingkan panjang, return yang lebih pendek."
//
// JEBAKAN UMUM:
//   - Lupa tulis karakter terakhir setelah loop selesai
//   - Lupa handle string kosong (len(s)==0) sebelum akses s[0]
// =============================================================================

// compress: run-length encoding dengan bonus — O(n) time, O(n) space
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
			sb.WriteByte(s[i-1])
			sb.WriteString(strconv.Itoa(count))
			count = 1
		}
	}
	sb.WriteByte(s[len(s)-1])
	sb.WriteString(strconv.Itoa(count))
	result := sb.String()
	if len(result) >= len(s) {
		return s
	}
	return result
}
