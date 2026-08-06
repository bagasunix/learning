package firstuniquechar

// First Unique Character
// ----------------------
// Cari index karakter pertama yang tidak berulang dalam string.
// Kembalikan -1 kalau tidak ada. (LeetCode #387)
//
// Contoh:
//
//	firstUniqChar("leetcode")     -> 0  ('l')
//	firstUniqChar("loveleetcode") -> 2  ('v')
//	firstUniqChar("aabb")         -> -1
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
