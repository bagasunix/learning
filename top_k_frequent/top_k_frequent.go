package topkfrequent

import "sort"

// =============================================================================
// TOP K FREQUENT ELEMENTS
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Hashmap + sorting, frequency counting
// Complexity: O(n log n) time, O(n) space
// =============================================================================
//
// SOAL:
//   Diberikan array integer dan integer k, return k elemen yang paling sering
//   muncul. Urutan output tidak harus terurut. (LeetCode #347)
//
// CONTOH:
//   topKFrequent([1,1,1,2,2,3], 2) -> [1,2]    (1 muncul 3x, 2 muncul 2x)
//   topKFrequent([1], 1)           -> [1]
//   topKFrequent([1,2], 2)         -> [1,2]    (frekuensi sama)
//   topKFrequent([4,4,4,2,2,3], 1) -> [4]      (4 paling sering: 3x)
//
// EDGE CASES yang harus disebut:
//   - k = 1              -> return satu elemen paling sering
//   - k = len(unique)    -> return semua unique element
//   - Semua frekuensi sama -> urutan output tidak penting
//
// PENDEKATAN BERTINGKAT:
//
//   1. Brute force: hitung frekuensi, sort semua, ambil k pertama
//      O(n log n) time, O(n) space -> ini yang dipakai di sini
//
//   2. Min-heap ukuran k:
//      Maintain heap dengan k elemen terbesar
//      O(n log k) time, O(n+k) space -> lebih optimal kalau k << n
//
//   3. Bucket sort (O(n) time):
//      Buat bucket[freq] = []elements
//      Iterasi dari freq tertinggi, ambil k elemen
//      O(n) time, O(n) space -> paling optimal
//
// CARA KERJA (approach yang dipakai):
//   nums=[1,1,1,2,2,3], k=2:
//
//   Step 1 — hitung frekuensi:
//   freq = {1:3, 2:2, 3:1}
//
//   Step 2 — kumpulkan keys:
//   keys = [1, 2, 3]
//
//   Step 3 — sort descending by frequency:
//   sort by freq[key] desc: [1(3x), 2(2x), 3(1x)]
//
//   Step 4 — ambil k pertama:
//   keys[:2] = [1, 2] ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Hashmap untuk hitung frekuensi O(n).
//    Sort keys berdasarkan frekuensi descending O(n log n).
//    Ambil k pertama.
//    Total O(n log n) time, O(n) space.
//    Kalau perlu lebih optimal, bisa pakai bucket sort O(n)."
//
// JEBAKAN UMUM:
//   - Sort array nums langsung (bukan keys by freq) -> tidak memberikan top-k
//   - Lupa bahwa output order tidak penting, test harus sort sebelum compare
// =============================================================================

// topKFrequent: hashmap + sort — O(n log n) time, O(n) space
func topKFrequent(nums []int, k int) []int {
	// step 1: hitung frekuensi tiap elemen
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// step 2: kumpulkan unique keys
	keys := make([]int, 0, len(freq))
	for key := range freq {
		keys = append(keys, key)
	}

	// step 3: sort descending berdasarkan frekuensi
	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	// step 4: return k elemen pertama
	return keys[:k]
}
