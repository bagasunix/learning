package binarysearch

// =============================================================================
// BINARY SEARCH
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Divide & conquer, rekursi
// Complexity: O(log n) time, O(1) space (iterative) / O(log n) stack (recursive)
// =============================================================================
//
// SOAL:
//   Diberikan sorted array dan target, cari index target dalam array.
//   Return -1 jika tidak ditemukan. (LeetCode #704)
//
// CONTOH:
//   binarySearch([1,3,5,7,9,11], 7)  -> 3   (index ke-3)
//   binarySearch([1,3,5,7,9,11], 4)  -> -1  (tidak ada)
//   binarySearch([1], 1)             -> 0
//   binarySearch([], 5)              -> -1
//
// PRASYARAT: Array harus sudah sorted ascending
//
// EDGE CASES yang harus disebut:
//   - Array kosong           -> -1
//   - Array satu elemen      -> cek langsung
//   - Target di ujung kiri   -> index 0
//   - Target di ujung kanan  -> index n-1
//   - Target tidak ada       -> -1
//
// CARA KERJA:
//   Array: [1, 3, 5, 7, 9, 11], target=7
//   lo=0, hi=5, mid=2 -> arr[2]=5 < 7  -> lo=mid+1=3
//   lo=3, hi=5, mid=4 -> arr[4]=9 > 7  -> hi=mid-1=3
//   lo=3, hi=3, mid=3 -> arr[3]=7 == 7 -> return 3 ✓
//
// ADA DUA IMPLEMENTASI:
//   1. binarySearch (iterative)  : O(log n) time, O(1) space  ← LEBIH BAIK
//   2. binarySearchRecursive     : O(log n) time, O(log n) stack
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "lo=0, hi=n-1. Loop selama lo<=hi.
//    mid = lo + (hi-lo)/2  <-- PENTING: bukan (lo+hi)/2 karena bisa overflow!
//    Kalau arr[mid]==target -> return mid.
//    Kalau arr[mid]<target  -> lo=mid+1 (buang setengah kiri).
//    Kalau arr[mid]>target  -> hi=mid-1 (buang setengah kanan).
//    Keluar loop tanpa ketemu -> return -1."
//
// JEBAKAN UMUM:
//   - mid = (lo+hi)/2 -> OVERFLOW jika lo dan hi sangat besar
//   - mid = lo + (hi-lo)/2 -> AMAN, selalu pakai ini
//   - Loop condition lo < hi (salah) vs lo <= hi (benar)
//     -> lo < hi akan miss elemen terakhir
// =============================================================================

// binarySearch: iterative — O(log n) time, O(1) space
func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // aman dari overflow
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// binarySearchRecursive: O(log n) time, O(log n) stack
func binarySearchRecursive(nums []int, target int) int {
	return bsHelper(nums, target, 0, len(nums)-1)
}

func bsHelper(nums []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return bsHelper(nums, target, mid+1, hi)
	}
	return bsHelper(nums, target, lo, mid-1)
}
