package rotatearray

// =============================================================================
// ROTATE ARRAY
// =============================================================================
// Tingkat   : Sedang
// Konsep    : 3-reverse trick, in-place manipulation
// Complexity: O(n) time, O(1) space
// =============================================================================
//
// SOAL:
//   Putar array ke kanan sebanyak k langkah. (LeetCode #189)
//
// CONTOH:
//   rotate([1,2,3,4,5,6,7], 3) -> [5,6,7,1,2,3,4]
//   rotate([-1,-100,3,99], 2)  -> [3,99,-1,-100]
//   rotate([1,2,3], 4)         -> [3,1,2]  (k=4 -> efektif k=1 karena 4%3=1)
//   rotate([1], 5)             -> [1]      (satu elemen tidak berubah)
//
// EDGE CASES yang harus disebut:
//   - k >= n     -> k = k % n (rotasi penuh tidak ada efeknya)
//   - k = 0      -> tidak berubah
//   - Array kosong atau satu elemen -> tidak berubah
//   - k = n      -> tidak berubah
//
// CARA KERJA 3-Reverse Trick:
//   Rotate kanan k langkah = 3 kali reverse:
//   1. Reverse seluruh array
//   2. Reverse bagian kiri [0..k-1]
//   3. Reverse bagian kanan [k..n-1]
//
//   nums=[1,2,3,4,5,6,7], k=3:
//   Step 1 reverse all:    [7,6,5,4,3,2,1]
//   Step 2 reverse [0..2]: [5,6,7,4,3,2,1]
//   Step 3 reverse [3..6]: [5,6,7,1,2,3,4] ✓
//
// KENAPA INI BEKERJA?
//   Rotate kanan k = elemen k terakhir pindah ke depan.
//   Setelah reverse semua, elemen yang harusnya di depan (k terakhir)
//   ada di awal tapi urutannya terbalik -> reverse lagi bagian itu.
//   Demikian juga bagian kanan.
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "3-reverse trick: reverse semua, reverse k pertama, reverse sisanya.
//    O(n) time, O(1) space. Kunci: k = k % n dulu untuk handle k >= n."
//
// APPROACH LAIN (lebih mudah dipikir tapi O(n) space):
//   Buat array baru: nums[(i+k)%n] = nums[i] untuk tiap i
//   Valid tapi butuh extra space O(n)
//
// JEBAKAN UMUM:
//   - Lupa k = k % n -> index out of bounds kalau k >= n
//   - Lupa handle k == 0 setelah modulo (langsung return)
// =============================================================================

// rotate: 3-reverse trick — O(n) time, O(1) space
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k = k % n // handle k >= n
	if k == 0 {
		return
	}
	reverse(nums, 0, n-1) // step 1: reverse semua
	reverse(nums, 0, k-1) // step 2: reverse bagian kiri [0..k-1]
	reverse(nums, k, n-1) // step 3: reverse bagian kanan [k..n-1]
}

func reverse(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
}
