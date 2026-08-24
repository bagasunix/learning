package firstuniquechar

// =============================================================================
// FIRST UNIQUE CHARACTER IN A STRING
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Hash map / frequency array, two-pass
// Complexity: O(n) time, O(1) space (max 26 huruf lowercase)
// =============================================================================
//
// SOAL:
//   Cari index karakter pertama yang tidak berulang dalam string.
//   Return -1 kalau semua karakter berulang. (LeetCode #387)
//
// CONTOH:
//   firstUniqChar("leetcode")     -> 0   ('l' muncul 1x, di index 0)
//   firstUniqChar("loveleetcode") -> 2   ('v' muncul 1x, di index 2)
//   firstUniqChar("aabb")         -> -1  (semua berulang)
//   firstUniqChar("z")            -> 0   (satu karakter pasti unique)
//   firstUniqChar("aab")          -> 2   ('b' di index 2)
//
// EDGE CASES yang harus disebut:
//   - String satu karakter  -> return 0
//   - Semua karakter sama   -> return -1
//   - Semua unik            -> return 0 (karakter pertama)
//
// CARA KERJA (two-pass):
//   Pass 1: hitung frekuensi tiap karakter
//   Pass 2: iterasi dari kiri, return index pertama yang frekuensinya == 1
//
//   "leetcode":
//   freq: l=1, e=3, t=1, c=1, o=1, d=1
//   Pass 2: index 0='l', freq['l']=1 -> return 0 ✓
//
//   "loveleetcode":
//   freq: l=1, o=2, v=1, e=4, t=1, c=1, d=1
//   Pass 2: index 0='l'? freq=1 -> tapi ada 'l' di index 0
//   Wait: "loveleetcode" -> l=1,o=2,v=1,e=4,t=1,c=1,d=1
//   index 0='l' freq=1? -> l muncul 1x -> return 0?
//   Hm, sebenarnya "loveleetcode": l=1,o=2,v=1,e=4... -> index 0='l' return 0
//   Tapi LeetCode bilang return 2 ('v'). Perlu dicek... "loveleetcode":
//   l-o-v-e-l-e-e-t-c-o-d-e -> l muncul 2x (index 0 dan 4)
//   Jadi freq: l=2, o=2, v=1, e=4, t=1, c=1, d=1
//   Pass 2: l=2 skip, o=2 skip, v=1 -> return index 2 ✓
//
// IMPLEMENTASI PAKAI [26]int (lebih efisien dari map):
//   - Hanya berlaku untuk lowercase a-z
//   - Index = c - 'a' (misal 'a'=0, 'b'=1, 'z'=25)
//   - Tidak perlu alokasi map, O(1) space sejati
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Two-pass: pass pertama hitung frekuensi pakai array[26].
//    Pass kedua iterasi string dari kiri, return index pertama
//    yang frekuensinya == 1. O(n) time, O(1) space."
//
// JEBAKAN UMUM:
//   - Single pass tidak bisa: belum tahu frekuensi final saat baru lihat karakter
//   - Pakai map[rune]int juga valid tapi O(n) space
// =============================================================================

// firstUniqChar: two-pass dengan frequency array — O(n) time, O(1) space
func firstUniqChar(s string) int {
	var freq [26]int
	for _, c := range s {
		freq[c-'a']++
	}
	for i, c := range s {
		if freq[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
