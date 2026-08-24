package majorityelement

// =============================================================================
// MAJORITY ELEMENT (Boyer-Moore Voting Algorithm)
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Boyer-Moore Voting, greedy
// Complexity: O(n) time, O(1) space
// =============================================================================
//
// SOAL:
//   Diberikan array integer, cari elemen yang muncul lebih dari n/2 kali.
//   Dijamin selalu ada majority element. (LeetCode #169)
//
// CONTOH:
//   majorityElement([3,2,3])           -> 3  (3 muncul 2x dari 3 elemen)
//   majorityElement([2,2,1,1,1,2,2])   -> 2  (2 muncul 4x dari 7 elemen)
//   majorityElement([1])               -> 1
//   majorityElement([1,1,2,2,2])       -> 2  (2 muncul 3x dari 5 elemen)
//
// EDGE CASES yang harus disebut:
//   - Satu elemen           -> return elemen itu
//   - Semua elemen sama     -> return elemen itu
//   - Tepat n/2 + 1 muncul  -> masih valid (dijamin ada)
//
// APPROACH BERTINGKAT (tunjukkan yang pertama dulu, lalu optimasi):
//
//   1. Hashmap (O(n) time, O(n) space):
//      Hitung frekuensi tiap elemen, return yang > n/2
//
//   2. Sort (O(n log n) time, O(1) space):
//      Sort array, elemen tengah pasti majority element
//
//   3. Boyer-Moore (O(n) time, O(1) space) ← OPTIMAL:
//      Ini yang dipakai di sini.
//
// CARA KERJA Boyer-Moore:
//   Intuisi: majority element "bertahan" setelah semua "voting" selesai.
//   Tiap elemen lain membatalkan satu kemunculan majority element,
//   tapi majority element lebih banyak jadi tetap menang.
//
//   candidate=3, count=1
//   nums=[3, 2, 3]:
//   v=3: same as candidate -> count=2
//   v=2: different          -> count=1   (saling cancel)
//   v=3: same as candidate  -> count=2
//   Tidak pernah reset -> candidate=3 ✓
//
//   nums=[2,2,1,1,1,2,2]:
//   candidate=2, count=1
//   v=2: count=2
//   v=1: count=1
//   v=1: count=0 -> reset! candidate=1, count=1
//   v=1: count=2
//   v=2: count=1
//   v=2: count=2 -> tapi candidate masih 1??
//   Hmm: setelah reset count=0, candidate=next element
//   v=2(index4): count=0 -> candidate=2, count=1
//   v=2(index5): count=2
//   candidate=2 ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Boyer-Moore: pakai candidate dan count.
//    Kalau count==0, ganti candidate ke elemen sekarang.
//    Kalau elemen sama dengan candidate, count++.
//    Kalau berbeda, count--.
//    Di akhir, candidate pasti majority element karena dia lebih banyak
//    dari semua elemen lain digabung."
//
// JEBAKAN UMUM:
//   - Tidak reset count ke 1 saat ganti candidate (harusnya count=1 bukan count=0+1)
//   - Mengira Boyer-Moore bekerja tanpa jaminan majority element ada -> SALAH
//     (kalau tidak ada jaminan, perlu pass kedua untuk verifikasi)
// =============================================================================

// majorityElement: Boyer-Moore Voting — O(n) time, O(1) space
func majorityElement(nums []int) int {
	candidate, count := nums[0], 1
	for _, v := range nums[1:] {
		if count == 0 {
			candidate = v
			count = 1
		} else if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
