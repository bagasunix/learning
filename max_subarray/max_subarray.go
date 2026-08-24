package maxsubarray

// =============================================================================
// MAXIMUM SUBARRAY (Kadane's Algorithm)
// =============================================================================
// Tingkat   : Sedang
// Konsep    : Dynamic programming, greedy
// Complexity: O(n) time, O(1) space
// =============================================================================
//
// SOAL:
//   Diberikan array integer (bisa negatif), cari jumlah maksimum
//   dari subarray yang berurutan (contiguous subarray).
//   (LeetCode #53)
//
// CONTOH:
//   maxSubArray([-2,1,-3,4,-1,2,1,-5,4]) -> 6   (subarray: [4,-1,2,1])
//   maxSubArray([1])                     -> 1
//   maxSubArray([5,4,-1,7,8])            -> 23   (seluruh array)
//   maxSubArray([-1,-2,-3])              -> -1   (semua negatif, ambil terbesar)
//
// EDGE CASES yang harus disebut:
//   - Semua angka negatif   -> return angka terbesar (kurang negatif)
//   - Satu elemen           -> return elemen itu
//   - Semua positif         -> return sum seluruh array
//
// CARA KERJA Kadane's:
//   Idenya: kalau current sum sudah negatif, lebih baik mulai fresh dari elemen berikutnya
//
//   nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
//   cur=−2, max=−2
//   cur<0  -> cur=1,  max=1
//   cur+(-3)=-2 -> cur=-2, max=1
//   cur<0  -> cur=4,  max=4
//   cur+(-1)=3 -> cur=3, max=4
//   cur+2=5  -> cur=5, max=5
//   cur+1=6  -> cur=6, max=6   ← ini yang terbesar
//   cur+(-5)=1 -> cur=1, max=6
//   cur+4=5  -> cur=5, max=6
//   return 6 ✓
//
// KUNCI JAWABAN saat ditanya interviewer:
//   "Kadane's algorithm: satu pass O(n).
//    Untuk tiap elemen, putuskan: lanjutkan subarray sebelumnya (cur+v)
//    atau mulai subarray baru dari elemen ini (v)?
//    Pilih yang lebih besar. Update max global kalau cur lebih besar.
//    Equivalent: cur = max(v, cur+v)"
//
// JEBAKAN UMUM:
//   - Inisialisasi maxSum = 0 -> salah untuk array semua negatif
//   - Harus inisialisasi maxSum = nums[0]
// =============================================================================

// maxSubArray: Kadane's — O(n) time, O(1) space
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	cur := nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = v // mulai fresh, subarray negatif tidak membantu
		} else {
			cur += v
		}
		if cur > maxSum {
			maxSum = cur
		}
	}
	return maxSum
}
