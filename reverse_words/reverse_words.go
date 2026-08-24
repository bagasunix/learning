package reversewords

import "strings"

// =============================================================================
// REVERSE WORDS IN A STRING
// =============================================================================
// Tingkat   : Sedang
// Konsep    : String splitting, two-pointer pada slice
// Complexity: O(n) time, O(n) space
// =============================================================================
//
// SOAL:
//   Balik urutan kata dalam string. Hapus spasi berlebih di awal, akhir,
//   dan di antara kata (normalisasi ke single space). (LeetCode #151)
//
// CONTOH:
//   reverseWords("the sky is blue")   -> "blue is sky the"
//   reverseWords("  hello world  ")   -> "world hello"   (strip leading/trailing)
//   reverseWords("a good   example")  -> "example good a" (multiple spaces -> 1)
//   reverseWords("hello")             -> "hello"
//   reverseWords("   ")               -> ""
//   reverseWords("")                  -> ""
//
// EDGE CASES yang harus disebut:
//   - Spasi di awal/akhir         -> dihapus
//   - Multiple spasi antar kata   -> jadi satu spasi
//   - Satu kata                   -> return kata itu
//   - Semua spasi                 -> return ""
//
// CARA KERJA:
//   1. strings.Fields(s) -> split by whitespace, otomatis handle multiple spaces
//      "  hello   world  " -> ["hello", "world"]
//   2. Reverse slice kata dengan two-pointer (lo/hi)
//      ["hello", "world"] -> ["world", "hello"]
//   3. strings.Join(words, " ") -> "world hello"
//
//   Kenapa strings.Fields?
//   - strings.Split(s, " ") tidak handle multiple spaces
//   - strings.Fields split by ANY whitespace dan buang empty strings
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Pakai strings.Fields untuk split (handle multiple spaces otomatis).
//    Reverse slice dengan two-pointer.
//    Join dengan single space.
//    O(n) time, O(n) space untuk slice kata."
//
// JEBAKAN UMUM:
//   - strings.Split(s, " ") -> menghasilkan "" untuk spasi ganda, perlu filter manual
//   - strings.TrimSpace + strings.Split tidak handle spasi tengah
//   - strings.Fields sudah handle semua kasus spasi -> paling simple
// =============================================================================

// reverseWords: split, reverse, join — O(n) time, O(n) space
func reverseWords(s string) string {
	words := strings.Fields(s) // split & strip whitespace otomatis
	lo, hi := 0, len(words)-1
	for lo < hi {
		words[lo], words[hi] = words[hi], words[lo]
		lo++
		hi--
	}
	return strings.Join(words, " ")
}
